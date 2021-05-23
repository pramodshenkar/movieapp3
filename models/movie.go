package Models

//Issue - struct to map with mongodb documents
type Movie struct {
	ID        int    `bson:"_id"`
	Name      string `bson:"name"`
	Budget    string `bson:"budget"`
	Director  int    `bson:"director"`
	Producers []int  `bson:"producers"`
	Actors    []int  `bson:"actors"`
}
