package main

import (
	"log"

	"github.com/raedatoui/assimp"
)

func main() {
	postProcessing := uint(assimp.Process_Triangulate) | uint(assimp.Process_FlipUVs)
	log.Println(postProcessing)
}
