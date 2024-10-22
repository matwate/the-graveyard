package core

import (
	"matwa/blogger/data"
	"net/http"
	"strconv"
)

var GetBlog HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	// Endpoint: GET /blog/{title}/{index}

	// Get the title and index from the URL
	t := r.PathValue("title")
	i := r.PathValue("index")

	// Find the blog by title
	blog, _ := data.FindBlogByTitle(t)

	if blog != (data.Blog{}) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(blog.Title))
		return
	}

	// Find the blog by index

	// Convert the index to an integer
	Idx, err := strconv.Atoi(i)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	blog, err = data.FindBlogByIndex(Idx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Blog not found"))
		return
	}

	if blog != (data.Blog{}) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(blog.Title))
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Blog not found"))

}
