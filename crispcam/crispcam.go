package crispcam

import "github.com/crispcam/crispapi/catalog"

type Request struct {
	Image     string  `json:"image,omitempty"`
	Threshold float64 `json:"threshold,omitempty"`
	Save      bool    `json:"save,omitempty"`
}

type Response struct {
	Results []catalog.Item `json:"results,omitempty"`
}

type Save struct {
	Image   string   `json:"image,omitempty"`
	Results Response `json:"crisps_results"`
	Save    bool     `json:"save"`
}
