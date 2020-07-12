package daemon

import (
	"net/http"
)

func renderFile(w http.ResponseWriter, r *http.Request, fs http.FileSystem) {
	http.FileServer(fs).ServeHTTP(w, r)
}

func renderNotFound(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte(message))
}
