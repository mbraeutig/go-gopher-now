package api

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var gopher []byte

func init() {
	if err := readFile("gophercolor.png", &gopher); err != nil {
		log.Fatal(err)
	}
}

func readFile(filename string, data *[]byte) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	*data = b
	return nil
}

// Gopher prints a gopher.
func Gopher(w http.ResponseWriter, r *http.Request) {
	// Write the gopher image to the response writer.
	if _, err := io.Copy(w, bytes.NewReader(gopher)); err != nil {
		http.Error(w, fmt.Sprintf("Error writing response: %v", err), http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "image/png")
	// Read the gopher image file.
	// f, err := os.Open("gophercolor.png")
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
	// 	return
	// }
	// defer f.Close()

	// Write the gopher image to the response writer.
	// if _, err := io.Copy(w, f); err != nil {
	// 	http.Error(w, fmt.Sprintf("Error writing response: %v", err), http.StatusInternalServerError)
	// }
	// w.Header().Add("Content-Type", "image/png")
}
