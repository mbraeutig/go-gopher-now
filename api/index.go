package api

import (
	"io"
	"net/http"
)

const index = "" +
	`
<html>
    <h1>Testing now.sh/zeit.co</h1>
</html>
`

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, index)
}
