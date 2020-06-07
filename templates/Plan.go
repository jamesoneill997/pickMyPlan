package templates

//Program struct stores program template
type Program struct {
	Category string `bson:"category" json:"category"`
	Exercies []struct {
		Equipment  []string `bson:"equipment" json:"equipment"`
		Duration   int      `bson:"duration" json:"duration"`
		TargetArea string   `bson:"targetArea" json:"targetArea"`
		WeightKg   int      `bson:"weightKg" json:"weightKg"`
	} `bson:"exercies" json:"exercies"`
	Diet struct {
		Breakfast []struct {
			Name        string   `bson:"name" json:"name"`
			Ingredients []string `bson:"breakfast" json:"breakfast"`
		} `bson:"breakfast" json:"breakfast"`

		Lunch []struct {
			Name        string   `bson:"name" json:"name"`
			Ingredients []string `bson:"ingredients" json:"ingredients"`
		} `bson:"luch" json:"lunch"`

		Dinner []struct {
			Name        string   `bson:"name" json:"name"`
			Ingredients []string `bson:"ingredients" json:"ingredients"`
		} `bson:"dinner" json:"dinner"`

		Snacks []struct {
			Name        string   `bson:"name" json:"name"`
			Ingredients []string `bson:"snacks" json:"snacks"`
		}
	}
}
