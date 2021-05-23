package main

import (
	"fmt"

	handler "github.com/pramodshenkar/movieapp3/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi")

	r := gin.Default()

	//------------MOVIE---------------------------------

	r.POST("/movie", handler.AddMovieHandler)
	r.POST("/movies", handler.AddManyMovieHandler)

	r.GET("/movie", handler.GetMoviesByIdHandler)
	r.GET("/movies", handler.GetAllMoviesHandler)

	r.PUT("/movie", handler.UpdateMovieHandler)
	r.DELETE("/movie", handler.DeleteOneMovieHandler)
	r.DELETE("/movies", handler.DeleteAllMoviesHandler)

	r.POST("/demomovies", handler.AddDemoMoviesHandler)

	//------------MOVIE---------------------------------

	r.POST("/producer", handler.AddProducerHandler)
	r.POST("/producers", handler.AddManyProducerHandler)

	r.GET("/producer", handler.GetProducersByIdHandler)
	r.GET("/producers", handler.GetAllProducersHandler)

	r.PUT("/producer", handler.UpdateProducerHandler)
	r.DELETE("/producer", handler.DeleteOneProducerHandler)
	r.DELETE("/producers", handler.DeleteAllProducersHandler)

	r.POST("/demoproducers", handler.AddDemoProducersHandler)

	//------------MOVIE---------------------------------
	r.POST("/add-producer-to-movie", handler.AddProducerToMovieHandler)

	r.Run()

}
