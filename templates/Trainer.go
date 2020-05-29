package templates

type Trainer struct {
	Username   string   `bson:"username" json:"username"`
	Gender     string   `bson:"gender" json:"gender"`
	Expertise  []string `bson:"expertise" json:"expertise"`
	Experience string   `bson:"experience" json:"experience"`
	Programs   string   `bson:"programs" json:"programs"`
	Website    string   `bson:"website" json:"website"`
}
