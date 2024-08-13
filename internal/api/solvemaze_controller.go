package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"maze-solver/internal/api/models"
	"maze-solver/internal/mazes"
	"net/http"
)

type SolveMazeController struct {
}

func NewSolveMazeController(group *gin.RouterGroup) *SolveMazeController {
	controller := SolveMazeController{}
	g := group.Group("/maze")
	g.POST("/", controller.SubmitMaze)
	return &controller
}

func (s *SolveMazeController) SubmitMaze(c *gin.Context) {
	var userMaze models.UserMaze
	if err := c.ShouldBindJSON(&userMaze); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
	}
	maze := userMaze.ConvertToMaze()
	var algo mazes.Algorithm
	switch userMaze.Algorithm {
	case models.FloodFill:
		algo = mazes.NewFloodFill(maze)
	case models.AStar:
		algo = mazes.NewAStar(maze)
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unknown algorithm"})
		return
	}
	solvable := true
	solved, err := mazes.Solve(algo)
	if err != nil && !errors.Is(err, mazes.UnsolvableError{}) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error solving maze"})
	}
	if err != nil && errors.Is(err, mazes.UnsolvableError{}) {
		solvable = false
	}
	c.JSON(http.StatusOK, gin.H{
		"solvable": solvable,
		"solution": solved,
	})
}
