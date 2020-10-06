package handler

import (
	"net/http"
	"io"
	"io/ioutil"
	"mime"
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
)

func parseBody(reader io.ReadCloser, header http.Header, data interface{}) error {
	defer reader.Close()

	mediatype, _, _ := mime.ParseMediaType(header.Get("Content-Type"))

	if body, err := ioutil.ReadAll(reader); err != nil {
		return err
	} else if mediatype == "application/json" {
		return json.Unmarshal(body, data)
	} else {
		return errors.New(string(body))
	}
}

func renderResponse(w http.ResponseWriter, r *http.Request, statusCode int, v interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	render.JSON(w, r, v)

	return
}
