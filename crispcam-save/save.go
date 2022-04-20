package crispcam_save

import (
	"github.com/crispcam/crispapi/catalog"
	"time"
)

type Results struct {
	Results []Image `json:"results,omitempty" firestore:"results,omitempty"`
}

type ResultsWithSource struct {
	Results []Image `json:"results,omitempty" firestore:"results,omitempty"`
	Source  string  `json:"source,omitempty" firestore:"source,omitempty"`
}

type Image struct {
	Accurate      string         `json:"accurate,omitempty" firestore:"accurate"`
	BestMatch     catalog.Item   `json:"best_match,omitempty" firestore:"bestMatch,omitempty"`
	AllMatches    []catalog.Item `json:"all_matches,omitempty" firestore:"allMatches"`
	Created       time.Time      `json:"created,omitempty" firestore:"created"`
	Modified      time.Time      `json:"modified,omitempty" firestore:"modified"`
	Reviewed      bool           `json:"reviewed,omitempty" firestore:"reviewed"`
	Reviewer      string         `json:"reviewer,omitempty" firestore:"reviewer"`
	Source        string         `json:"source,omitempty" firestore:"source"`
	Transaction   string         `json:"transaction,omitempty" firestore:"transaction"`
	ImageLocation string         `json:"image_location,omitempty" firestore:"imageLocation"`
	RealMatch     Match          `json:"real_match,omitempty" firestore:"realMatch"`
}

type Match struct {
	Matched   bool   `json:"matched,omitempty" firestore:"matched"`
	RealMatch string `json:"real_match,omitempty" firestore:"realMatch"`
	Valid     bool   `json:"valid,omitempty" firestore:"valid"`
}

type Flavours struct {
	Results map[string][]Flavour `json:"results,omitempty" firestore:"results,omitempty"`
}
type Flavour struct {
	ID       string  `json:"id,omitempty" firestore:"id"`
	Score    float64 `json:"score,omitempty" firestore:"score"`
	FullName string  `json:"full_name,omitempty" firestore:"fullName"`
}

/**
 * Not crisps (a selfie etc)
 */
const (
	NoResult      = "no_result"
	NotListed     = "not_listed"
	ReviewedField = "reviewed"
	NotReviewed   = "not_reviewed"
	All           = "all"
	NoChange      = "NO_CHANGE"
)
