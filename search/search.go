package search

type Result struct {
	ID     string `json:"id,omitempty"`
	Weight int64  `json:"weight,omitempty"`
	Name   string `json:"name,omitempty"`
	Image  string `json:"img,omitempty"`
}

type Item struct {
	Prefix  string
	Matches []Result
}

type MatchPair struct {
	Name   string
	Weight int64
}
