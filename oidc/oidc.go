package oidc

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	oidclib "github.com/coreos/go-oidc/v3/oidc"
	"github.com/crispcam/crispapi/auth"
	"github.com/crispcam/crispapi/crisps"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	Expiry         = "expire"
	Token          = "token"
	Redirect       = "redirect"
	Authorization  = "authorization"
	IDToken        = "id-token"
	AppConfigKey   = "crispcam-oidc-config"
	UserSessionKey = "crispcam-session"
)

type SessionConfig struct {
	SessionStore *redisstore.RedisStore
	Config       *crisps.Config
	OidcProvider *oidclib.Provider
	Oauth2Config *oauth2.Config
	Verifier     *oidclib.IDTokenVerifier
}

type UserSession struct {
	User  auth.User
	Valid bool
}

func SessionActive(ctx context.Context, session *sessions.Session) (userSession UserSession, err error) {
	sessionConfig, ok := ctx.Value(AppConfigKey).(*SessionConfig)
	if !ok {
		println("Cannot get session config!")
	}
	userSession.Valid = false
	expiry, err := time.Parse(time.RFC3339, fmt.Sprintf("%v", session.Values[Expiry]))
	if err != nil {
		// This means there's no expiry time - redirect to login
		return userSession, nil
	}
	if expiry.Sub(time.Now()).Seconds() <= 0 {
		return userSession, nil
	}

	// Verify token
	idToken := session.Values[IDToken]

	// Parse and verify ID Token payload.
	_, err = sessionConfig.Verifier.Verify(ctx, fmt.Sprintf("%v", idToken))
	if err != nil {
		log.Println(err.Error())
		return userSession, err
	}

	// Obtain AuthZ
	user, err := crisps.AuthZ(ctx, fmt.Sprintf("%v", session.Values[Token]))
	if err != nil {
		log.Println(err.Error())
		return userSession, err
	}
	return UserSession{
		User:  user,
		Valid: true,
	}, nil

}

func CheckSession(w http.ResponseWriter, r *http.Request) (userSession UserSession, err error) {

	sessionConfig, ok := r.Context().Value(AppConfigKey).(*SessionConfig)
	if !ok {
		println("Cannot get session config!")
	}

	session, err := sessionConfig.SessionStore.Get(r, sessionConfig.Config.Session.Key)
	if err != nil {
		return UserSession{}, err
	}
	userSession, err = SessionActive(r.Context(), session)
	if err != nil {
		return userSession, err
	}
	if !userSession.Valid {
		state, err := randString(16)
		if err != nil {
			return userSession, err
		}
		// Save current page
		redir := fmt.Sprintf("%v", r.URL.Path)
		session.Values[Redirect] = redir
		err = session.Save(r, w)
		if err != nil {
			log.Println("Warning: Could not save session:", err.Error())
		}

		setCallbackCookie(w, r, "state", state)
		http.Redirect(w, r, sessionConfig.Oauth2Config.AuthCodeURL(state), http.StatusFound)
	}

	return userSession, nil
}

func LoginCallback(w http.ResponseWriter, r *http.Request, sessionConfig SessionConfig) error {
	ctx := r.Context()
	state, err := r.Cookie("state")
	if err != nil {
		return err
	}
	if r.URL.Query().Get("state") != state.Value {
		return errors.New("state did not match")
	}
	oauth2Token, err := sessionConfig.Oauth2Config.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		return err
	}
	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		log.Println("Could not extract raw ID token")
		return errors.New("could not extract raw ID token")
	}

	session, err := sessionConfig.SessionStore.Get(r, sessionConfig.Config.Session.Key)
	if err != nil {
		return err
	}
	session.Values[Expiry] = oauth2Token.Expiry.Format(time.RFC3339)
	session.Values[Token] = oauth2Token.AccessToken
	session.Values[IDToken] = rawIDToken

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	http.Redirect(w, r, fmt.Sprintf("%v", session.Values[Redirect]), http.StatusFound)

	return nil

}

func randString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func setCallbackCookie(w http.ResponseWriter, r *http.Request, name, value string) {
	c := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   r.TLS != nil,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
}
