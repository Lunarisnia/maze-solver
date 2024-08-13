package main

import (
	"github.com/gin-gonic/gin"
	"maze-solver/internal/api"
)

func main() {
	r := gin.Default()
	api.NewSolveMazeController(r.Group("api/v1"))

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
