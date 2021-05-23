package handler

import (
	"strconv"

	function "github.com/pramodshenkar/movieapp3/functions"
	model "github.com/pramodshenkar/movieapp3/models"

	"github.com/gin-gonic/gin"
)

//--------------- Add One Actor -------------

/*
http://localhost:8080/actor

{
			"ID":       1,
			"Name":     "Rajkumar Hirani",
			"Address":   "10C",
}
*/
func AddActorHandler(c *gin.Context) {

	var actor model.Actor
	c.Bind(&actor)
	res, err := function.AddActor(actor)

	if err != nil {
		c.JSON(409, err)
	} else {
		c.JSON(200, res)
	}
}

//--------------- Add Multiple Actor -------------
/*
http://localhost:8080/actors

[
	{
		"ID": 2,
		"Name": "PK",
		"Budget": "10C",
		"Actor": "Rajkumar Hirani",
	},
	{
		"ID": 3,
		"Name": "Happy new year",
		"Budget": "10C",
		"Actor": "Rajkumar Hirani",
	},
	{
		"ID": 4,
		"Name": "Bahubali",
		"Budget": "10C",
		"Actor": "Rajkumar Hirani",
	},
	{
		"ID": 5,
		"Name": "Saho",
		"Budget": "10C",
		"Actor": "Rajkumar Hirani",
	}
]

*/
func AddManyActorHandler(c *gin.Context) {

	var actors []model.Actor
	c.Bind(&actors)
	res, err := function.AddManyActors(actors)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

//--------------- Get Actor By Id -------------
/*
http://localhost:8080/actor?id=2
*/

func GetActorsByIdHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	actor, err := function.GetActorsById(id)

	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, actor)
	}
}

//--------------- Get All Actors -------------
/*
http://localhost:8080/actors
*/

func GetAllActorsHandler(c *gin.Context) {
	Actors, err := function.GetAllActors()
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, Actors)
	}
}

//--------------- Get Actor By Id -------------
/*
http://localhost:8080/actor

{
			"ID":       1,
			"Name":     "Joker",
			"Budget":   "10C",
			"Actor": "Ashutosh Gowarikar"
}
*/

func UpdateActorHandler(c *gin.Context) {
	var actor model.Actor
	c.Bind(&actor)

	res, err := function.UpdateActor(actor)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

//-------------- Get Actor By Id -------------
/*
http://localhost:8080/actor?id=2
*/

func DeleteOneActorHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	res, err := function.DeleteOneActor(id)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

//-------------- Get Actor By Id -------------
/*
http://localhost:8080/actors
*/
func DeleteAllActorsHandler(c *gin.Context) {
	res, err := function.DeleteAllActors()
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

func AddDemoActorsHandler(c *gin.Context) {

	var actors = []model.Actor{
		{
			ID:   1,
			Name: "Prabhas",
			Address: model.Address{
				Houseno: "2B",
				Street:  "ABC",
				City:    "Mumbai",
				State:   "Maharashtra",
			},
		},
		{
			ID:   2,
			Name: "Rajkumar Rao",
			Address: model.Address{
				Houseno: "3C",
				Street:  "ABC",
				City:    "Benglore",
				State:   "Tamilnadu",
			},
		},
		{
			ID:   3,
			Name: "Athulya Shetty",
			Address: model.Address{
				Houseno: "4D",
				Street:  "ABC",
				City:    "Panvel",
				State:   "Maharashtra",
			},
		},
		{
			ID:   4,
			Name: "Shahid Kappor",
			Address: model.Address{
				Houseno: "5E",
				Street:  "ABC",
				City:    "Mumbai",
				State:   "Maharashtra",
			},
		},
	}
	c.Bind(&actors)
	res, err := function.AddManyActors(actors)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}
