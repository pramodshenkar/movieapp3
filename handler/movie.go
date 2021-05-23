package handler

import (
	"strconv"

	function "github.com/pramodshenkar/movieapp3/functions"
	model "github.com/pramodshenkar/movieapp3/models"

	"github.com/gin-gonic/gin"
)

//--------------- Add One Movie -------------

/*
http://localhost:8080/movie

{
			"ID":       1,"Budget"
			"Name":     "Munnabhai MBBS",
			"Budget":   "10C","Name"
			"Director": "Rajkumar Hirani"
}
*/
func AddMovieHandler(c *gin.Context) {

	var movie model.Movie
	c.Bind(&movie)
	res, err := function.AddMovie(movie)

	if err != nil {
		c.JSON(409, err)
	} else {
		c.JSON(200, res)
	}
}

//--------------- Add Multiple Movie -------------
/*
http://localhost:8080/movies

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
func AddManyMovieHandler(c *gin.Context) {

	var movies []model.Movie
	c.Bind(&movies)
	res, err := function.AddManyMovies(movies)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

//--------------- Get Movie By Id -------------
/*
http://localhost:8080/movie?id=2
*/

func GetMoviesByIdHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	movie, err := function.GetMoviesById(id)

	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, movie)
	}
}

//--------------- Get All Movies -------------
/*
http://localhost:8080/movies
*/

func GetAllMoviesHandler(c *gin.Context) {
	Movies, err := function.GetAllMovies()
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, Movies)
	}
}

//--------------- Get Movie By Id -------------
/*
http://localhost:8080/movie

{
			"ID":       1,
			"Name":     "Joker",
			"Budget":   "10C",
			"Director": "Ashutosh Gowarikar"
}
*/

func UpdateMovieHandler(c *gin.Context) {
	var movie model.Movie
	c.Bind(&movie)

	res, err := function.UpdateMovie(movie)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

//-------------- Get Movie By Id -------------
/*
http://localhost:8080/movie?id=2
*/

func DeleteOneMovieHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	res, err := function.DeleteOneMovie(id)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

//-------------- Get Movie By Id -------------
/*
http://localhost:8080/movies
*/
func DeleteAllMoviesHandler(c *gin.Context) {
	res, err := function.DeleteAllMovies()
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

func AddDemoMoviesHandler(c *gin.Context) {

	var movies = []model.Movie{
		{
			ID:        1,
			Name:      "PK",
			Budget:    "10C",
			Director:  1,
			Producers: []int{},
			Actors:    []model.ActorProfile{},
		},
		{
			ID:        2,
			Name:      "Happy new year",
			Budget:    "10C",
			Director:  2,
			Producers: []int{},
			Actors:    []model.ActorProfile{},
		},
		{
			ID:        3,
			Name:      "Bahubali",
			Budget:    "10C",
			Director:  3,
			Producers: []int{},
			Actors:    []model.ActorProfile{},
		},
		{
			ID:        4,
			Name:      "Saho",
			Budget:    "10C",
			Director:  4,
			Producers: []int{},
			Actors:    []model.ActorProfile{},
		}}

	c.Bind(&movies)
	res, err := function.AddManyMovies(movies)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}
