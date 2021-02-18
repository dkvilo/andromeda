package object

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// AndromedaObject Struct
type AndromedaObject struct {
	X, Y, Z, U, V float64
}

// Load and Parse Object
func Load(path string) []float32 {
	
	var vertexIndices []float32
	
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("Object %q not found on disk: %v", path, err))
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if string(scanner.Text()[0]) == "a"{
			x, _ := strconv.ParseFloat(strings.Split(scanner.Text(), " ")[1], 64)
			y, _ := strconv.ParseFloat(strings.Split(scanner.Text(), " ")[2], 64)
			z, _ := strconv.ParseFloat(strings.Split(scanner.Text(), " ")[3], 64)
			u, _ := strconv.ParseFloat(strings.Split(scanner.Text(), " ")[4], 64)
			v, _ := strconv.ParseFloat(strings.Split(scanner.Text(), " ")[5], 64)
			vertexIndices = append(vertexIndices, float32(x), float32(y), float32(z), float32(u), float32(v))
		}
	}
	return vertexIndices
}


