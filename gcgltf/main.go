package main

// #include <gcgltf.h>
import "C"

func loadModel(path string) {

	aiScene := C.loadFile(C.CString(path))
	if aiScene != nil {
		defer C.freeScene(aiScene)
	} else {
		return

	}
}

func main() {
	loadModel("model.gltf")
}
