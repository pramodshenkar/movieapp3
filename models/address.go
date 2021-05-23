package Models

type Address struct {
	Houseno string `bson:"houseno"`
	Street  string `bson:"street"`
	City    string `bson:"city"`
	State   string `bson:"state"`
}
