package templates

//User struct stores user data
type User struct {
	Username       string         `bson:"username" json:"username"`
	Type           string         `bson:"type" json:"type"`
	Gender         string         `bson:"gender" json:"gender"`
	Email          string         `bson:"email" json:"email"`
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
	Plans          []string       `bson:"plans" json:"plans"`
}
