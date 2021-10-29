package crisps

import (
	"encoding/json"
	"github.com/crispcam/crispapi/catalog"
	"github.com/crispcam/crispapi/reviews"
	"github.com/crispcam/crispapi/search"
	"net/http"
)

func CatalogAll(r *http.Request, config Config) (catalog.Results, error) {
	var results catalog.Results
	// Process items from catalog
	u := config.CrispCam.Services.Catalog + config.CrispCam.Paths.Catalog.All
	bodyBytes, err := Request(r, u, "GET", nil)
	if err != nil {
		return results, err
	}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return results, err
	}

	return results, nil
}

func CatalogItem(r *http.Request, config Config, id string) (catalog.Item, error) {
	var item catalog.Item
	// Process items from catalog
	u := config.CrispCam.Services.Catalog + config.CrispCam.Paths.Catalog.Single + "/" + id
	bodyBytes, err := Request(r, u, "GET", nil)
	if err != nil {
		return item, err
	}
	err = json.Unmarshal(bodyBytes, &item)
	if err != nil {
		return item, err
	}

	return item, nil
}

func ReviewAll(r *http.Request, config Config) (reviews.Ratings, error) {
	var results reviews.Ratings
	// Process items from catalog
	u := config.CrispCam.Services.Reviews + config.CrispCam.Paths.Reviews.Ratings
	bodyBytes, err := Request(r, u, "GET", nil)
	if err != nil {
		return results, err
	}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return results, err
	}

	return results, nil
}

func ReviewItem(r *http.Request, config Config, id string) (reviews.Rating, error) {
	var result reviews.Rating
	// Process items from catalog
	u := config.CrispCam.Services.Reviews + config.CrispCam.Paths.Reviews.Rating + "/" + id
	bodyBytes, err := Request(r, u, "GET", nil)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func CategoryAll(r *http.Request, config Config) (catalog.Categories, error) {
	var results catalog.Categories
	// Process items from catalog
	u := config.CrispCam.Services.Catalog + config.CrispCam.Paths.Catalog.Categories
	bodyBytes, err := Request(r, u, "GET", nil)
	if err != nil {
		return results, err
	}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return results, err
	}

	return results, nil
}

func SearchResults(r *http.Request, config Config, q string) ([]search.Result, error) {
	var results []search.Result
	// Process items from catalog
	u := config.CrispCam.Services.Search + config.CrispCam.Paths.Search.Search + "/" + q
	bodyBytes, err := Request(r, u, "GET", nil)
	if err != nil {
		return results, err
	}
	if len(bodyBytes) == 0 {
		return results, nil
	}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return results, err
	}

	return results, nil
}

func UserInfo(r *http.Request, config Config, uid string) error {
	u := config.CrispCam.Services.Auth + config.CrispCam.Paths.Auth.User + "/" + uid
	_, err := Request(r, u, "GET", nil)
	if err != nil {
		return err
	}
	return nil
}

func Assets(r *http.Request, config Config) error {
	u := config.CrispCam.Services.Assets + "/v1/simulate"
	_, err := Request(r, u, "GET", nil)
	if err != nil {
		return err
	}
	return nil
}
