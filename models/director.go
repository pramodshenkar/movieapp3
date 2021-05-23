package Models

type Director struct {
	ID      int    `bson:"_id"`
	Name    string `bson:"name"`
	Address Address
}
