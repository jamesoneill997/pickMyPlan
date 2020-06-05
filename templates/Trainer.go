package templates

//Trainer struct stores trainer account data
type Trainer struct {
	Username   string   `bson:"username" json:"username"`
	Type       string   `bson:"type" json:"type"`
	Gender     string   `bson:"gender" json:"gender"`
	Email      string   `bson:"email" json:"email"`
	Password   string   `bson:"password" json:"password"`
	Expertise  []string `bson:"expertise" json:"expertise"`
	Experience string   `bson:"experience" json:"experience"`
	Plans      string   `bson:"plans" json:"plans"`
	Website    string   `bson:"website" json:"website"`
}
