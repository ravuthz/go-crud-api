package router

import (
	"github.com/gin-gonic/gin"
)
 
func InitRouter() *gin.Engine {
   r := gin.Default()
   r.GET("/movies", getMovies)
   r.GET("/movies/:id", getMovie)
   r.POST("/movies", postMovie)
   r.PUT("/movies/:id", putMovie)
   r.DELETE("/movies/:id", deleteMovie)
   return r
}
