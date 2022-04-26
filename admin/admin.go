package admin

import (
	"github.com/crispcam/crispapi/auth"
	"github.com/crispcam/crispapi/catalog"
)

type CatalogPage struct {
	Title string
	Items catalog.Results
	User  auth.User
}

type CatalogItemPage struct {
	Title string
	Item  catalog.Item
	User  auth.User
}
