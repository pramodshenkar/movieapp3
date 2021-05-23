package handler

import (
	"strconv"

	function "github.com/pramodshenkar/movieapp3/functions"
	model "github.com/pramodshenkar/movieapp3/models"

	"github.com/gin-gonic/gin"
)

//--------------- Add One Producer -------------

/*
http://localhost:8080/producer

{
			"ID":       1,
			"Name":     "Rajkumar Hirani",
			"Address":   "10C",
}
*/
func AddProducerHandler(c *gin.Context) {

	var producer model.Producer
	c.Bind(&producer)
	res, err := function.AddProducer(producer)

	if err != nil {
		c.JSON(409, err)
	} else {
		c.JSON(200, res)
	}
}

//--------------- Add Multiple Producer -------------
/*
http://localhost:8080/producers

[
	{
		"ID": 2,
		"Name": "PK",
		"Budget": "10C",
		"Producer": "Rajkumar Hirani",
	},
	{
		"ID": 3,
		"Name": "Happy new year",
		"Budget": "10C",
		"Producer": "Rajkumar Hirani",
	},
	{
		"ID": 4,
		"Name": "Bahubali",
		"Budget": "10C",
		"Producer": "Rajkumar Hirani",
	},
	{
		"ID": 5,
		"Name": "Saho",
		"Budget": "10C",
		"Producer": "Rajkumar Hirani",
	}
]

*/
func AddManyProducerHandler(c *gin.Context) {

	var producers []model.Producer
	c.Bind(&producers)
	res, err := function.AddManyProducers(producers)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

//--------------- Get Producer By Id -------------
/*
http://localhost:8080/producer?id=2
*/

func GetProducersByIdHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	producer, err := function.GetProducersById(id)

	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, producer)
	}
}

//--------------- Get All Producers -------------
/*
http://localhost:8080/producers
*/

func GetAllProducersHandler(c *gin.Context) {
	Producers, err := function.GetAllProducers()
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, Producers)
	}
}

//--------------- Get Producer By Id -------------
/*
http://localhost:8080/producer

{
			"ID":       1,
			"Name":     "Joker",
			"Budget":   "10C",
			"Producer": "Ashutosh Gowarikar"
}
*/

func UpdateProducerHandler(c *gin.Context) {
	var producer model.Producer
	c.Bind(&producer)

	res, err := function.UpdateProducer(producer)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

//-------------- Get Producer By Id -------------
/*
http://localhost:8080/producer?id=2
*/

func DeleteOneProducerHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	res, err := function.DeleteOneProducer(id)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

//-------------- Get Producer By Id -------------
/*
http://localhost:8080/producers
*/
func DeleteAllProducersHandler(c *gin.Context) {
	res, err := function.DeleteAllProducers()
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

func AddDemoProducersHandler(c *gin.Context) {

	var producers = []model.Producer{
		{
			ID:   2,
			Name: "Ashutosh Gowarikar",
			Address: model.Address{
				Houseno: "2B",
				Street:  "ABC",
				City:    "Mumbai",
				State:   "Maharashtra",
			},
		},
		{
			ID:   3,
			Name: "SS RajMauli",
			Address: model.Address{
				Houseno: "3C",
				Street:  "ABC",
				City:    "Benglore",
				State:   "Tamilnadu",
			},
		},
		{
			ID:   4,
			Name: "Rohit Shetty",
			Address: model.Address{
				Houseno: "4D",
				Street:  "ABC",
				City:    "Panvel",
				State:   "Maharashtra",
			},
		},
		{
			ID:   5,
			Name: "Farah Khan",
			Address: model.Address{
				Houseno: "5E",
				Street:  "ABC",
				City:    "Mumbai",
				State:   "Maharashtra",
			},
		},
	}
	c.Bind(&producers)
	res, err := function.AddManyProducers(producers)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}
