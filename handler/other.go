package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	function "github.com/pramodshenkar/movieapp3/functions"
)

/*
localhost:8080/add-producer-to-movie?movieid=2&producerid=2
*/
func AddProducerToMovieHandler(c *gin.Context) {
	movieid, _ := strconv.Atoi(c.Query("movieid"))
	producerid, _ := strconv.Atoi(c.Query("producerid"))

	res, err := function.AddProducerToMovie(movieid, producerid)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

func RemoveProducerFromMovieHandler(c *gin.Context) {
	movieid, _ := strconv.Atoi(c.Query("movieid"))
	producerid, _ := strconv.Atoi(c.Query("producerid"))

	res, err := function.RemoveProducerFromMovie(movieid, producerid)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}


func AddActorToMovieHandler(c *gin.Context) {
	movieid, _ := strconv.Atoi(c.Query("movieid"))
	actorid, _ := strconv.Atoi(c.Query("actorid"))

	res, err := function.AddActorToMovie(movieid, actorid)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}

func RemoveActorFromMovieHandler(c *gin.Context) {
	movieid, _ := strconv.Atoi(c.Query("movieid"))
	actorid, _ := strconv.Atoi(c.Query("actorid"))

	res, err := function.RemoveActorFromMovie(movieid, actorid)
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, res)
	}
}
