package localml

type Request struct {
	Instances []Instance `json:"instances,omitempty"`
}
type Instance struct {
	Image Payload `json:"image_bytes"`
	Key   string  `json:"key"`
}
type Payload struct {
	Image string `json:"b64,omitempty"`
}

type Response struct {
	Predictions []Predictions `json:"predictions,omitempty"`
}

type Predictions struct {
	Labels []string  `json:"labels,omitempty"`
	Scores []float64 `json:"scores,omitempty"`
}
