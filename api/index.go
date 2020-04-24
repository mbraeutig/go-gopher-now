package api

import (
	"io"
	"net/http"
)

const index = "" +
	`
	<head>
	<html>
		<link rel="shortcut icon" type="image/x-icon" href="favicon.ico">
		<h1>Want to see a gopher, try this!</h1>
		<h1><a href="gopher">Gopher</a></h1>
	</html>
	</head>
	`

// Index - write the index
func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, index)
}
