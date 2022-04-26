package crisps

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/crispcam/crispapi/auth"
	"github.com/crispcam/crispapi/catalog"
	crispcamSave "github.com/crispcam/crispapi/crispcam-save"
	"github.com/crispcam/crispapi/reviews"
	"github.com/crispcam/crispapi/search"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

func CatalogAll(ctx context.Context) (results catalog.Results, err error) {
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return results, errors.New("unable to read config")
	}
	// Process items from catalog
	u := config.CrispCam.Services.Catalog + config.CrispCam.Paths.Catalog.All
	bodyBytes, err := Request(ctx, u, "GET", nil)
	if err != nil {
		return results, err
	}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return results, err
	}

	return results, nil
}

func CatalogItem(ctx context.Context, id string) (item catalog.Item, err error) {
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return item, errors.New("unable to read config")
	}
	// Process items from catalog
	u := config.CrispCam.Services.Catalog + config.CrispCam.Paths.Catalog.Single + "/" + id
	bodyBytes, err := Request(ctx, u, "GET", nil)
	if err != nil {
		return item, err
	}
	err = json.Unmarshal(bodyBytes, &item)
	if err != nil {
		return item, err
	}

	return item, nil
}

func ReviewAll(ctx context.Context) (results reviews.Ratings, err error) {
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return results, errors.New("unable to read config")
	}
	// Process items from catalog
	u := config.CrispCam.Services.Reviews + config.CrispCam.Paths.Reviews.Ratings
	bodyBytes, err := Request(ctx, u, "GET", nil)
	if err != nil {
		return results, err
	}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return results, err
	}

	return results, nil
}

func ReviewItem(ctx context.Context, id string) (result reviews.Rating, err error) {
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return result, errors.New("unable to read config")
	}
	// Process items from catalog
	u := config.CrispCam.Services.Reviews + config.CrispCam.Paths.Reviews.Rating + "/" + id
	bodyBytes, err := Request(ctx, u, "GET", nil)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func CategoryAll(ctx context.Context) (results catalog.Categories, err error) {
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return results, errors.New("unable to read config")
	}
	// Process items from catalog
	u := config.CrispCam.Services.Catalog + config.CrispCam.Paths.Catalog.Categories
	bodyBytes, err := Request(ctx, u, "GET", nil)
	if err != nil {
		return results, err
	}
	err = json.Unmarshal(bodyBytes, &results)
	if err != nil {
		return results, err
	}

	return results, nil
}

func SearchResults(ctx context.Context, q string) (results []search.Result, err error) {
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return results, errors.New("unable to read config")
	}
	// Process items from catalog
	u := config.CrispCam.Services.Search + config.CrispCam.Paths.Search.Search + "/" + q
	bodyBytes, err := Request(ctx, u, "GET", nil)
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

func UserInfo(ctx context.Context, uid string) error {
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return errors.New("unable to read config")
	}
	u := config.CrispCam.Services.Auth + config.CrispCam.Paths.Auth.BasicUser + "/" + uid
	_, err := Request(ctx, u, "GET", nil)
	if err != nil {
		return err
	}
	return nil
}

func Assets(ctx context.Context) error {
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return errors.New("unable to read config")
	}
	u := config.CrispCam.Services.Assets + "/v1/simulate"
	_, err := Request(ctx, u, "GET", nil)
	if err != nil {
		return err
	}
	return nil
}

func AuthZ(ctx context.Context, token string) (user auth.User, err error) {
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return user, errors.New("unable to read config")
	}
	u := config.CrispCam.Services.Auth + config.CrispCam.Paths.Auth.User
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return user, err
	}

	if !strings.HasPrefix(token, "Bearer") {
		token = fmt.Sprintf("Bearer %v", token)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("authorization", token)

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

func SavedItems(ctx context.Context, saveType SaveType) (results crispcamSave.Results, err error) {
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return results, errors.New("unable to read config")
	}
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
	bodyBytes, err := Request(ctx, u, "GET", nil)
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

func SavedItem(ctx context.Context, transaction string) (result crispcamSave.Image, err error) {
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return result, errors.New("unable to read config")
	}
	// Process items from catalog
	u := config.CrispCam.Services.CrispCamSave + config.CrispCam.Paths.CrispCamSave.Single + "/" + transaction
	bodyBytes, err := Request(ctx, u, http.MethodGet, nil)
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

func UpdateItem(ctx context.Context, transaction string, match string, reviewed bool) (err error) {
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return errors.New("unable to read config")
	}
	// Process items from catalog
	u := config.CrispCam.Services.CrispCamSave + config.CrispCam.Paths.CrispCamSave.Update + "/" + transaction + "/" + match + "/" + fmt.Sprintf("%v", reviewed)
	_, err = Request(ctx, u, http.MethodPut, nil)
	if err != nil {
		return err
	}
	return nil
}

func Flavours(ctx context.Context, images crispcamSave.Results) (flavours crispcamSave.Flavours, err error) {
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return flavours, errors.New("unable to read config")
	}
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
	req = req.WithContext(ctx)
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

func DeleteItem(ctx context.Context, transaction string) (err error) {
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return errors.New("unable to read config")
	}
	u := config.CrispCam.Services.CrispCamSave + config.CrispCam.Paths.CrispCamSave.Delete + "/" + transaction
	req, err := http.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// Persist http headers
	req = req.WithContext(ctx)
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

func CSV(ctx context.Context) (bodyBytes []byte, err error) {
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return bodyBytes, errors.New("unable to read config")
	}
	u := config.CrispCam.Services.CrispCamSave + config.CrispCam.Paths.CrispCamSave.CSV
	bodyBytes, err = Request(ctx, u, "GET", nil)
	if err != nil {
		log.Println("Request to " + u + " failed: " + err.Error())
		return bodyBytes, err
	}
	if len(bodyBytes) == 0 {
		return bodyBytes, nil
	}

	return bodyBytes, nil
}

func UpdateCatalogItem(ctx context.Context, token string, item catalog.Item) (err error) {
	println("Updating catalog item")
	config, ok := ctx.Value(CtxKey).(*Config)
	if !ok {
		return errors.New("unable to read config")
	}
	// Process items from catalog
	u := config.CrispCam.Services.Catalog + config.CrispCam.Paths.Catalog.Update
	println(u)
	jsonBody, err := json.Marshal(item)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, u, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("authorization", "Bearer "+token)
	// Persist http headers
	req = req.WithContext(ctx)
	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	resp, err := client.Do(req)
	fmt.Printf("%v", resp)
	return nil
}
