package Models

type Actor struct {
	ID      int    `bson:"_id"`
	Name    string `bson:"name"`
	Address Address
}
