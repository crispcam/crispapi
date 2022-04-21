package auth

import "time"

type ID struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
}

type User struct {
	Id         string    `json:"id" firestore:"id"`
	Name       string    `json:"name" firestore:"name"`
	GivenName  string    `json:"given_name" firestore:"givenName"`
	FamilyName string    `json:"family_name" firestore:"familyName"`
	Picture    string    `json:"picture" firestore:"picture"`
	Locale     string    `json:"locale" firestore:"locale"`
	Email      string    `json:"email" firestore:"email"`
	Admin      bool      `json:"admin" firestore:"admin"`
	LastLogin  time.Time `json:"last_login" firestore:"lastLogin"`
	FirstLogin time.Time `json:"first_login" firestore:"firstLogin"`
	LoggedIn   bool      `json:"logged_in" firestore:"loggedIn"`
}
type BasicUser struct {
	Id         string `json:"id" firestore:"id"`
	Name       string `json:"name" firestore:"name"`
	GivenName  string `json:"given_name" firestore:"givenName"`
	FamilyName string `json:"family_name" firestore:"familyName"`
	Picture    string `json:"picture" firestore:"picture"`
	Locale     string `json:"locale" firestore:"locale"`
}
