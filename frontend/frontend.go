package frontend

import (
	"github.com/crispcam/crispapi/catalog"
	"github.com/crispcam/crispapi/reviews"
	"github.com/crispcam/crispapi/search"
)

type CatalogPage struct {
	Title           string
	AutocompleteURL string
	Catalog         catalog.Results
	CatalogError    bool
	Reviews         reviews.Ratings
	ReviewsError    bool
	Stars           [10]int
	Query           string
}

type ItemPage struct {
	Title           string
	AutocompleteURL string
	Item            catalog.Item
	ItemEmpty       bool
	Review          reviews.Rating
	ReviewError     bool
	Stars           [10]int
	Query           string
}

type CategoryPage struct {
	Title           string
	AutocompleteURL string
	Categories      catalog.Categories
	CategoryError   bool
	Catalog         catalog.Results
	CatalogError    bool
	Query           string
}

type SearchPage struct {
	Title           string
	AutocompleteURL string
	ResultsError    bool
	ResultsEmpty    bool
	Query           string
	Results         []search.Result
}

type ErrorPage struct {
	Title           string
	AutocompleteURL string
	StatusCode      int
	ErrorMessage    string
	Query           string
}

type GenericPage struct {
	Title           string
	AutocompleteURL string
	Query           string
}
