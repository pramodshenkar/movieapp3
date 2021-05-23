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

	// c.JSON(200, movieid)
	// c.JSON(200, producerid)

	function.AddProducerToMovie(movieid, producerid)

}
