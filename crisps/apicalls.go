package crisps

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/crispcam/crispapi/auth"
	"github.com/crispcam/crispapi/catalog"
	crispcam_save "github.com/crispcam/crispapi/crispcam-save"
	"github.com/crispcam/crispapi/reviews"
	"github.com/crispcam/crispapi/search"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

func CatalogAll(r *http.Request, config Config) (catalog.Results, error) {
	var results catalog.Results
	// Process items from catalog
	u := config.CrispCam.Services.Catalog + config.CrispCam.Paths.Catalog.All
	bodyBytes, err := Request(r, u, "GET", nil, r.Context())
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
	bodyBytes, err := Request(r, u, "GET", nil, r.Context())
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
	bodyBytes, err := Request(r, u, "GET", nil, r.Context())
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
	bodyBytes, err := Request(r, u, "GET", nil, r.Context())
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
	bodyBytes, err := Request(r, u, "GET", nil, r.Context())
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
	bodyBytes, err := Request(r, u, "GET", nil, r.Context())
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

func UserInfo(r *http.Request, config Config, uid string) error {
	u := config.CrispCam.Services.Auth + config.CrispCam.Paths.Auth.BasicUser + "/" + uid
	_, err := Request(r, u, "GET", nil, r.Context())
	if err != nil {
		return err
	}
	return nil
}

func Assets(r *http.Request, config Config) error {
	u := config.CrispCam.Services.Assets + "/v1/simulate"
	_, err := Request(r, u, "GET", nil, r.Context())
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

type SaveType int

const (
	All SaveType = iota
	Reviewed
	NotReviewed
)

func SavedItems(r *http.Request, config Config, saveType SaveType) (results crispcam_save.Results, err error) {
	var u string
	if saveType == All {
		u = config.CrispCam.Services.CrispCamSave + config.CrispCam.Paths.CrispCamSave.All
	} else {
		u = config.CrispCam.Services.CrispCamSave + config.CrispCam.Paths.CrispCamSave.Reviewed
		if saveType == Reviewed {
			u = config.CrispCam.Services.CrispCamSave + config.CrispCam.Paths.CrispCamSave.Reviewed + "?reviewed=true"
		} else if saveType == NotReviewed {
			u = config.CrispCam.Services.CrispCamSave + config.CrispCam.Paths.CrispCamSave.Reviewed + "?reviewed=false"
		} else {
			return results, errors.New(fmt.Sprintf("unknown type %v", saveType))
		}
	}
	bodyBytes, err := Request(r, u, "GET", nil, r.Context())
	if err != nil {
		return results, err
	}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		log.Println("Marshall response from " + u + " failed: " + err.Error())
		return results, err
	}
	// Sort the result by date ascending
	sort.Slice(results.Results, func(i, j int) bool {
		return results.Results[i].Modified.Before(results.Results[j].Modified)
	})
	return results, nil
}

func SavedItem(r *http.Request, config Config, transaction string) (result crispcam_save.Image, err error) {
	// Process items from catalog
	u := config.CrispCam.Services.CrispCamSave + config.CrispCam.Paths.CrispCamSave.Single + "/" + transaction
	bodyBytes, err := Request(r, u, http.MethodGet, nil, r.Context())
	if err != nil {
		log.Println("Request to " + u + " failed: " + err.Error())
		return result, err
	}
	if len(bodyBytes) == 0 {
		return result, nil
	}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		log.Println("Marshall response from " + u + " failed: " + err.Error())
		return result, err
	}
	return result, nil
}

func UpdateItem(r *http.Request, config Config, transaction string, match string, reviewed bool) (err error) {
	// Process items from catalog
	u := config.CrispCam.Services.CrispCamSave + config.CrispCam.Paths.CrispCamSave.Update + "/" + transaction + "/" + match + "/" + fmt.Sprintf("%v", reviewed)
	_, err = Request(r, u, http.MethodPut, nil, r.Context())
	if err != nil {
		return err
	}
	return nil
}

func Flavours(r *http.Request, config Config, images crispcam_save.Results) (flavours crispcam_save.Flavours, err error) {
	u := config.CrispCam.Services.CrispCamSave + config.CrispCam.Paths.CrispCamSave.Flavours
	jsonBody, err := json.Marshal(images.Results)
	if err != nil {
		return flavours, err
	}
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(jsonBody))
	if err != nil {
		return flavours, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// Persist http headers
	req = req.WithContext(r.Context())
	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return flavours, err
	}
	if resp.StatusCode != http.StatusOK {
		return flavours, errors.New(fmt.Sprintf("invalid status code: %v", resp.StatusCode))
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return flavours, err
	}
	err = json.Unmarshal(bodyBytes, &flavours)
	if err != nil {
		log.Println("Marshall response from " + u + " failed: " + err.Error())
		return flavours, err
	}
	return flavours, nil
}

func DeleteItem(r *http.Request, config Config, transaction string) (err error) {
	u := config.CrispCam.Services.CrispCamSave + config.CrispCam.Paths.CrispCamSave.Delete + "/" + transaction
	req, err := http.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// Persist http headers
	req = req.WithContext(r.Context())
	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("invalid status code: %v", resp.StatusCode))
	}
	return nil
}

func CSV(r *http.Request, config Config) (bodyBytes []byte, err error) {
	u := config.CrispCam.Services.CrispCamSave + config.CrispCam.Paths.CrispCamSave.CSV
	bodyBytes, err = Request(r, u, "GET", nil, r.Context())
	if err != nil {
		log.Println("Request to " + u + " failed: " + err.Error())
		return bodyBytes, err
	}
	if len(bodyBytes) == 0 {
		return bodyBytes, nil
	}

	return bodyBytes, nil
}
