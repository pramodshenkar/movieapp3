package handler

import (
	"strconv"

	function "github.com/pramodshenkar/movieapp3/functions"
	model "github.com/pramodshenkar/movieapp3/models"

	"github.com/gin-gonic/gin"
)

//--------------- Add One Director -------------

/*
http://localhost:8080/director

{
			"ID":       1,
			"Name":     "Rajkumar Hirani",
			"Address":   "10C",
}
*/
func AddDirectorHandler(c *gin.Context) {

	var director model.Director
	c.Bind(&director)
	res, err := function.AddDirector(director)

	if err != nil {
		c.JSON(409, err)
	} else {
		c.JSON(200, res)
	}
}

//--------------- Add Multiple Director -------------
/*
http://localhost:8080/directors

[
	{
		"ID": 2,
		"Name": "PK",
		"Budget": "10C",
		"Director": "Rajkumar Hirani",
	},
	{
		"ID": 3,
		"Name": "Happy new year",
		"Budget": "10C",
		"Director": "Rajkumar Hirani",
	},
	{
		"ID": 4,
		"Name": "Bahubali",
		"Budget": "10C",
		"Director": "Rajkumar Hirani",
	},
	{
		"ID": 5,
		"Name": "Saho",
		"Budget": "10C",
		"Director": "Rajkumar Hirani",
	}
]

*/
func AddManyDirectorHandler(c *gin.Context) {

	var directors []model.Director
	c.Bind(&directors)
	res, err := function.AddManyDirectors(directors)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

//--------------- Get Director By Id -------------
/*
http://localhost:8080/director?id=2
*/

func GetDirectorsByIdHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	director, err := function.GetDirectorsById(id)

	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, director)
	}
}

//--------------- Get All Directors -------------
/*
http://localhost:8080/directors
*/

func GetAllDirectorsHandler(c *gin.Context) {
	Directors, err := function.GetAllDirectors()
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, Directors)
	}
}

//--------------- Get Director By Id -------------
/*
http://localhost:8080/director

{
			"ID":       1,
			"Name":     "Joker",
			"Budget":   "10C",
			"Director": "Ashutosh Gowarikar"
}
*/

func UpdateDirectorHandler(c *gin.Context) {
	var director model.Director
	c.Bind(&director)

	res, err := function.UpdateDirector(director)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

//-------------- Get Director By Id -------------
/*
http://localhost:8080/director?id=2
*/

func DeleteOneDirectorHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	res, err := function.DeleteOneDirector(id)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

//-------------- Get Director By Id -------------
/*
http://localhost:8080/directors
*/
func DeleteAllDirectorsHandler(c *gin.Context) {
	res, err := function.DeleteAllDirectors()
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

func AddDemoDirectorsHandler(c *gin.Context) {

	var directors = []model.Director{
		{
			ID:   1,
			Name: "Rajkumar Hirani",
			Address: model.Address{
				Houseno: "2B",
				Street:  "ABC",
				City:    "Mumbai",
				State:   "Maharashtra",
			},
		},
		{
			ID:   2,
			Name: "SS Rajmauli",
			Address: model.Address{
				Houseno: "3C",
				Street:  "ABC",
				City:    "Benglore",
				State:   "Tamilnadu",
			},
		},
		{
			ID:   3,
			Name: "Ashutosh Govarikar",
			Address: model.Address{
				Houseno: "4D",
				Street:  "ABC",
				City:    "Panvel",
				State:   "Maharashtra",
			},
		},
		{
			ID:   4,
			Name: "Rohit Shetty",
			Address: model.Address{
				Houseno: "5E",
				Street:  "ABC",
				City:    "Mumbai",
				State:   "Maharashtra",
			},
		},
	}
	c.Bind(&directors)
	res, err := function.AddManyDirectors(directors)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}
