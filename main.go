package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var movies = []Movie{}

func CreateMovie(c *gin.Context) {
	var newMovie Movie
	if err := c.ShouldBindJSON(&newMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newMovie.ID = len(movies) + 1
	movies = append(movies, newMovie)
	c.JSON(http.StatusOK, newMovie)
}

func main() {
	r := gin.Default()
	r.POST("/movie", CreateMovie)
	r.Run(":8080")
}
