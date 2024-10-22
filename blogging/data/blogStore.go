package data

import "errors"

type Blog struct {
	Title   string
	Content string
	Author  string
	Time    string
	Index   int
}

func NewBlog(title, content, author, time string) Blog {
	return Blog{
		Title:   title,
		Content: content,
		Author:  author,
		Time:    time,
		Index:   len(BlogStore),
	}
}

func AddBlog(blog Blog) Blog {
	BlogStore[blog.Title] = blog
	return blog
}

func FindBlogByIndex(index int) (Blog, error) {
	for _, blog := range BlogStore {
		if blog.Index == index {
			return blog, nil
		}
	}
	return Blog{}, errors.New("Blog not found")
}

func FindBlogByTitle(title string) (Blog, error) {
	b, ok := BlogStore[title]
	if !ok {
		return Blog{}, errors.New("Blog not found")
	} else {
		return b, nil
	}

}

var BlogStore = map[string]Blog{
	"Matwa": {
		Title:   "Matwa's Blog",
		Content: "This is Matwa's blog",
		Author:  "Matwa",
		Time:    "2024-06-04",
		Index:   0,
	},
}
