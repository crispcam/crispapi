package catalog

type Results struct {
	Results map[string]Item `json:"results" firestore:"results,omitempty"`
}

type Item struct {
	ID            string    `json:"id" firestore:"id,omitempty"`
	Brand         string    `json:"brand" firestore:"brand,omitempty"`
	Name          string    `json:"name" firestore:"name,omitempty"`
	Images        Images    `json:"images" firestore:"images,omitempty"`
	Categories    []string  `json:"categories" firestore:"categories"`
	FlavourGroups []string  `json:"flavour_groups" firestore:"flavourGroups"`
	Ingredients   []string  `json:"ingredients,omitempty" firestore:"ingredients"`
	Nutrition     Nutrition `json:"nutrition" firestore:"nutrition,omitempty"`
	Score         Accuracy  `json:"accuracy,omitempty"`
}

type Accuracy struct {
	Score float64 `json:"score,omitempty"`
}

type Images struct {
	HqImage string `json:"hq_image" firestore:"hqImage,omitempty"`
	LqImage string `json:"lq_image" firestore:"lqImage,omitempty"`
}
type Nutrition struct {
	EnergyKj     int     `json:"energy_kj" firestore:"energyKj,omitempty"`
	EnergyCal    int     `json:"energy_cal" firestore:"energyCal,omitempty"`
	Fat          float64 `json:"fat" firestore:"fat,omitempty"`
	SaturatedFat float64 `json:"saturated_fat" firestore:"saturated_fat,omitempty"`
	Carbs        float64 `json:"carbs" firestore:"carbs,omitempty"`
	Sugar        float64 `json:"sugar" firestore:"sugar,omitempty"`
	Fibre        float64 `json:"fibre" firestore:"fibre,omitempty"`
	Protein      float64 `json:"protein" firestore:"protein,omitempty"`
	Salt         float64 `json:"salt" firestore:"salt,omitempty"`
}

type Categories struct {
	Categories map[string]Category `json:"categories"`
}

type Category struct {
	Category string   `json:"category"`
	Members  []string `json:"members"`
}
