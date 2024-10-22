package core

import (
	"matwa/blogger/data"
	"net/http"
)

var UpdateBlog HandlerFunc = func(w http.ResponseWriter, r *http.Request) {

	blogTitle := r.PathValue("title")

	blog, err := data.FindBlogByTitle(blogTitle)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Blog not found"))
		return
	}

	blog.Title = r.FormValue("title")
	blog.Content = r.FormValue("content")

	data.BlogStore[blogTitle] = blog

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Blog updated successfully"))

}
