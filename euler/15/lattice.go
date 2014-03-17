package main
import "fmt"

// Note: this can also be solved "more easily" by just
// 	calculating 40 choose 20. I kind of enjoyed writing
//	this dynamic programming solution though :)

var paths [21][21]int

func countPaths(width int, height int) int {
	if width > 20 || height > 20 {
		panic ("Can't go larger than 20, yo")
	}

	for i := 0; i < len(paths); i++ {
		paths[i][0] = 1
		paths[0][i] = 1
	}

	for w := 1; w <= width; w++ {
		for h := 1; h <= height; h++ {
			paths[w][h] = paths[w-1][h] + paths[w][h-1]
		}
	}
	
	return paths[width][height]
}

func main() {
	fmt.Println(countPaths(20, 20))
}