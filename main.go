package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	handler "github.com/pramodshenkar/movieapp3/handler"
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

	//------------Producer---------------------------------

	r.POST("/producer", handler.AddProducerHandler)
	r.POST("/producers", handler.AddManyProducerHandler)

	r.GET("/producer", handler.GetProducersByIdHandler)
	r.GET("/producers", handler.GetAllProducersHandler)

	r.PUT("/producer", handler.UpdateProducerHandler)
	r.DELETE("/producer", handler.DeleteOneProducerHandler)
	r.DELETE("/producers", handler.DeleteAllProducersHandler)

	r.POST("/demoproducers", handler.AddDemoProducersHandler)


	//------------Actor---------------------------------

	r.POST("/actor", handler.AddActorHandler)
	r.POST("/actors", handler.AddManyActorHandler)

	r.GET("/actor", handler.GetActorsByIdHandler)
	r.GET("/actors", handler.GetAllActorsHandler)

	r.PUT("/actor", handler.UpdateActorHandler)
	r.DELETE("/actor", handler.DeleteOneActorHandler)
	r.DELETE("/actors", handler.DeleteAllActorsHandler)

	r.POST("/demoactors", handler.AddDemoActorsHandler)


	//------------MOVIE---------------------------------

	r.PUT("/add-producer-to-movie", handler.AddProducerToMovieHandler)
	r.PUT("/remove-producer-frome-movie", handler.RemoveProducerFromMovieHandler)

	r.PUT("/add-actor-to-movie", handler.AddActorToMovieHandler)
	r.PUT("/remove-actor-frome-movie", handler.RemoveActorFromMovieHandler)


	r.Run()

}
