package core

import (
	"encoding/json"
	"io"
	"matwa/blogger/data"
	"net/http"
)

var NewBlog HandlerFunc = func(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request body is required"))
		return
	}

	if s, err := io.ReadAll(r.Body); err == nil {

		var blog data.Blog
		if err := json.Unmarshal(s, &blog); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to parse JSON"))
			return
		}

		// Validate required fields
		if blog.Title == "" || blog.Content == "" || blog.Author == "" || blog.Time == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing required fields"))
			return
		}

		// Additional validation or processing logic here

		// Save the blog to the database or perform other actions
		data.AddBlog(blog)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Blog created successfully"))

	}

}
