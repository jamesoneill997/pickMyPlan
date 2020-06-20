package templates

//Program struct stores program template
type Program struct {
	//name does not need to be unique
	Name     string `bson:"name" json:"name"`
	Category string `bson:"category" json:"category"`
	Workout  struct {
		Exercises  []string `bson:"exercises" json:"exercises"`
		Equipment  []string `bson:"equipment" json:"equipment"`
		Duration   int      `bson:"duration" json:"duration"`
		TargetArea string   `bson:"targetArea" json:"targetArea"`
	} `bson:"workout" json:"workout"`
	//Diet contains meals
	Diet struct {
		Breakfast []string `bson:"breakfast" json:"breakfast"`
		Lunch     []string `bson:"luch" json:"lunch"`
		Dinner    []string `bson:"dinner" json:"dinner"`
		Snacks    []string `bson:"snacks" json:"snacks"`
	}

	//username of trainer
	Creator string `bson:"creator" json:"creator"`
}
