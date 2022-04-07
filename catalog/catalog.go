package catalog

type Results struct {
	Results map[string]Item `json:"results" firestore:"results,omitempty"`
}

type Item struct {
	ID            string    `json:"id,omitempty" firestore:"id,omitempty"`
	Brand         string    `json:"brand,omitempty" firestore:"brand,omitempty"`
	Name          string    `json:"name,omitempty" firestore:"name,omitempty"`
	Images        Images    `json:"images,omitempty" firestore:"images,omitempty"`
	Categories    []string  `json:"categories,omitempty" firestore:"categories"`
	FlavourGroups []string  `json:"flavour_groups,omitempty" firestore:"flavourGroups"`
	Ingredients   []string  `json:"ingredients,omitempty,omitempty" firestore:"ingredients"`
	Nutrition     Nutrition `json:"nutrition,omitempty" firestore:"nutrition,omitempty"`
	Score         Accuracy  `json:"accuracy,omitempty"  firestore:"accuracy,omitempty"`
}

type Accuracy struct {
	Score float64 `json:"score,omitempty"`
}

type Images struct {
	HqImage string `json:"hq_image,omitempty" firestore:"hqImage,omitempty"`
	LqImage string `json:"lq_image,omitempty" firestore:"lqImage,omitempty"`
}
type Nutrition struct {
	EnergyKj     int     `json:"energy_kj,omitempty" firestore:"energyKj,omitempty"`
	EnergyCal    int     `json:"energy_cal,omitempty" firestore:"energyCal,omitempty"`
	Fat          float64 `json:"fat,omitempty" firestore:"fat,omitempty"`
	SaturatedFat float64 `json:"saturated_fat,omitempty" firestore:"saturated_fat,omitempty"`
	Carbs        float64 `json:"carbs,omitempty" firestore:"carbs,omitempty"`
	Sugar        float64 `json:"sugar,omitempty" firestore:"sugar,omitempty"`
	Fibre        float64 `json:"fibre,omitempty" firestore:"fibre,omitempty"`
	Protein      float64 `json:"protein,omitempty" firestore:"protein,omitempty"`
	Salt         float64 `json:"salt,omitempty" firestore:"salt,omitempty"`
}

type Categories struct {
	Categories map[string]Category `json:"categories"`
}

type Category struct {
	Category string   `json:"category"`
	Members  []string `json:"members"`
}
