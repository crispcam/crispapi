package reviews

type Ratings struct {
	Ratings map[string]Rating `json:"ratings,omitempty"`
}

type Rating struct {
	ID     string `json:"id,omitempty"`
	Score  int    `json:"score,omitempty"`
	Colour string `json:"colour,omitempty"`
}
