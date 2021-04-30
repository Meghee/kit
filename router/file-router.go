package router

import (
	"net/http"
	"os"
	"path/filepath"
)

// FileRouter is the route handler for routing files|assets.
//
// usage: mux.HandleFunc("/assets/", router.FileRouter("./static/assets", "/assets/"))
func FileRouter(directory, routePath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		workDir, _ := os.Getwd()
		files := http.FileServer(http.Dir(filepath.Join(workDir, directory)))
		fileServer := http.StripPrefix(routePath, files)
		fileServer.ServeHTTP(w, r)
	}
}
