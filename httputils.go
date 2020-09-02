package goutils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

func RenderError(w http.ResponseWriter, err error, status int) {
	errMessage := "Unknown error"
	if err != nil {
		errMessage = err.Error()
	}
	http.Error(w, errMessage, status)
}

func RenderJSON(w http.ResponseWriter, object interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(object); err != nil {
		log.Error().Caller().Err(err).Msgf("rendering json failed")
	}
}

func RenderText(w http.ResponseWriter, object interface{}, status int) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(status)
	w.Write([]byte("pong"))
}

// ReadBody -- do not use to read body unless it is reading the error, or in debug mode
// for serious things, use io.Copy
func ReadBody(b io.ReadCloser) string {
	body, err := ioutil.ReadAll(b)
	if err != nil {
		return fmt.Sprintf("read from body failed: %s", err)
	}
	defer b.Close()
	return string(body)
}
