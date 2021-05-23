package Models

type ActorProfile struct {
	Actorid int    `bson:"actorid"`
	Role    string `bson:"role"`
	Salary  string `bson:"salary"`
}
