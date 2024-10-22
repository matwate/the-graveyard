package main

type link_schema struct {
	link           string
	expectedOutput bool
	onError        string
}

var link_schemas = []link_schema{
	{link: "https://github.com/user/repo", expectedOutput: true, onError: "invalid link"},
	{link: "http://github.com/user/repo", expectedOutput: false, onError: "invalid link"},
	{link: "https://github.com/", expectedOutput: true, onError: "invalid link"},
	{link: "https://github.com/user", expectedOutput: true, onError: "invalid link"},
	{link: "https://example.com/user/repo", expectedOutput: false, onError: "invalid link"},
	{link: "https://github.com/user/repo/issues", expectedOutput: true, onError: "invalid link"},
	{link: "ftp://github.com/user/repo", expectedOutput: false, onError: "invalid link"},
	{link: "https://github.com/user/repo/pull/1", expectedOutput: true, onError: "invalid link"},
	{link: "https://gitlab.com/user/repo", expectedOutput: false, onError: "invalid link"},
	{link: "https://github.com/user/repo/wiki", expectedOutput: true, onError: "invalid link"},
}

func assertLink_parsing(schema link_schema) (output bool) {
	if parse_link(schema.link) == true {
		output = true
	} else {
		output = false
	}
	return
}

/*
func TestLink_parsing(t *testing.T) {
	for _, schema := range link_schemas {
		if assertLink_parsing(schema) != schema.expectedOutput {
			t.Error(schema.onError)
		}
	}
}
*/
