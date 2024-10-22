package main

import (
	"log"

	"github.com/bloeys/assimp-go/asig"
)

func main() {
	// Create a new importer
	postProcessingEffects := asig.PostProcessTriangulate | asig.PostProcessFlipUVs
	importer, release, err := asig.ImportFile("path/to/model.obj", postProcessingEffects)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(importer)
	release()
}
