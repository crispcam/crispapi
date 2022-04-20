package crispcam_review

import (
	"github.com/crispcam/crispapi/auth"
	save "github.com/crispcam/crispapi/crispcam-save"
)

type ReviewPage struct {
	SaveOK   bool
	Items    save.Results
	User     auth.User
	Title    string
	Flavours save.Flavours
	Source   string
}

type LoginPage struct {
	User  auth.User
	Title string
}
