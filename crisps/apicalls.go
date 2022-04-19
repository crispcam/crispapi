package crisps

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/crispcam/crispapi/auth"
	"github.com/crispcam/crispapi/catalog"
	"github.com/crispcam/crispapi/reviews"
	"github.com/crispcam/crispapi/search"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"io/ioutil"
	"log"
	"net/http"
)

func CatalogAll(r *http.Request, config Config, ctx context.Context) (catalog.Results, error) {
	var results catalog.Results
	// Process items from catalog
	u := config.CrispCam.Services.Catalog + config.CrispCam.Paths.Catalog.All
	bodyBytes, err := Request(r, u, "GET", nil, ctx)
	if err != nil {
		return results, err
	}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return results, err
	}

	return results, nil
}

func CatalogItem(r *http.Request, config Config, id string, ctx context.Context) (catalog.Item, error) {
	var item catalog.Item
	// Process items from catalog
	u := config.CrispCam.Services.Catalog + config.CrispCam.Paths.Catalog.Single + "/" + id
	bodyBytes, err := Request(r, u, "GET", nil, ctx)
	if err != nil {
		return item, err
	}
	err = json.Unmarshal(bodyBytes, &item)
	if err != nil {
		return item, err
	}

	return item, nil
}

func ReviewAll(r *http.Request, config Config, ctx context.Context) (reviews.Ratings, error) {
	var results reviews.Ratings
	// Process items from catalog
	u := config.CrispCam.Services.Reviews + config.CrispCam.Paths.Reviews.Ratings
	bodyBytes, err := Request(r, u, "GET", nil, ctx)
	if err != nil {
		return results, err
	}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return results, err
	}

	return results, nil
}

func ReviewItem(r *http.Request, config Config, id string, ctx context.Context) (reviews.Rating, error) {
	var result reviews.Rating
	// Process items from catalog
	u := config.CrispCam.Services.Reviews + config.CrispCam.Paths.Reviews.Rating + "/" + id
	bodyBytes, err := Request(r, u, "GET", nil, ctx)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func CategoryAll(r *http.Request, config Config, ctx context.Context) (catalog.Categories, error) {
	var results catalog.Categories
	// Process items from catalog
	u := config.CrispCam.Services.Catalog + config.CrispCam.Paths.Catalog.Categories
	bodyBytes, err := Request(r, u, "GET", nil, ctx)
	if err != nil {
		return results, err
	}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return results, err
	}

	return results, nil
}

func SearchResults(r *http.Request, config Config, q string, ctx context.Context) ([]search.Result, error) {
	var results []search.Result
	// Process items from catalog
	u := config.CrispCam.Services.Search + config.CrispCam.Paths.Search.Search + "/" + q
	bodyBytes, err := Request(r, u, "GET", nil, ctx)
	if err != nil {
		log.Println("Request to " + u + " failed: " + err.Error())
		return results, err
	}
	if len(bodyBytes) == 0 {
		return results, nil
	}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		log.Println("Marshall response from " + u + " failed: " + err.Error())
		return results, err
	}

	return results, nil
}

func UserInfo(r *http.Request, config Config, uid string, ctx context.Context) error {
	u := config.CrispCam.Services.Auth + config.CrispCam.Paths.Auth.User + "/" + uid
	_, err := Request(r, u, "GET", nil, ctx)
	if err != nil {
		return err
	}
	return nil
}

func Assets(r *http.Request, config Config, ctx context.Context) error {
	u := config.CrispCam.Services.Assets + "/v1/simulate"
	_, err := Request(r, u, "GET", nil, ctx)
	if err != nil {
		return err
	}
	return nil
}

func AuthZ(r *http.Request, config Config, token string) (user auth.User, err error) {
	u := config.CrispCam.Services.Auth + config.CrispCam.Paths.Auth.User
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, u, nil)
	if err != nil {
		return user, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("authorization", "Bearer "+token)

	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	resp, err := client.Do(req)
	if err != nil {
		return user, err
	}
	if resp.StatusCode != http.StatusOK {
		return user, errors.New(fmt.Sprintf("upstream status code %d (request URI: %v)", resp.StatusCode, u))
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return user, err
	}
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		log.Println("Marshall response from " + u + " failed: " + err.Error())
		return user, err
	}

	return user, nil
}
