package templates

type User struct {
	Username       string         `bson:"username" json:"username"`
	Gender         string         `bson:"gender" json:"gender"`
	WeightKg       int            `bson:"weightKg" json:"weightKg"`
	HeightCm       int            `bson:"heightCm" json:"heightCm"`
	Build          string         `bson:"build" json:"build"`
	Goals          string         `bson:"goals" json:"goals"`
	Equipment      []string       `bson:"equipment" json:"equipment"`
	Statistics     map[string]int `bson:"statistics" json:"statistics"`
	Allergies      []string       `bson:"allergies" json:"allergies"`
	ProfileImage   string         `bson:"profileImage" json:"profileImage"`
	ProgressImages []string       `bson:"progressImages" json:"progressImages"`
	PayAcctID      string         `bson:"payAcctID" json:"payAcctID"`
	Password       string         `bson:"password" json:"password"`
	Tokens         []string       `bson:"tokens" json:"tokens"`
}
