package images

import (
	"fmt"
	"os"
)

func loadPuzzle(path string) ([][]int, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(file))

	return [][]int{}, nil
}
