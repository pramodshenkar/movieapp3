package Models

type Producer struct {
	ID      int    `bson:"_id"`
	Name    string `bson:"name"`
	Address Address
}
