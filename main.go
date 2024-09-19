package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

// Home represents the home data
type Home struct {
}

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Join internally call path.Clean to prevent directory traversal
	path := filepath.Join(h.staticPath, r.URL.Path)

	// check whether a file exists or is a directory at the given path
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		// file does not exist or path is a directory, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static file
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	// templates:
	tmpl := make(map[string]*template.Template)
	tmpl["index.html"] = template.Must(template.ParseFiles("./web/templates/index.html", "./web/templates/base.html"))
	tmpl["register.html"] = template.Must(template.ParseFiles("./web/templates/register.html", "./web/templates/base.html"))
	tmpl["recover.html"] = template.Must(template.ParseFiles("./web/templates/recover.html", "./web/templates/base.html"))
	tmpl["network.html"] = template.Must(template.ParseFiles("./web/templates/network.html", "./web/templates/base.html"))

	// router:
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl["index.html"].ExecuteTemplate(w, "base", &Home{})
		if err != nil {
			log.Printf("error: %s", err.Error())
		}
	})

	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl["register.html"].ExecuteTemplate(w, "base", &Home{})
		if err != nil {
			log.Printf("error: %s", err.Error())
		}
	})

	router.HandleFunc("/recover", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl["recover.html"].ExecuteTemplate(w, "base", &Home{})
		if err != nil {
			log.Printf("error: %s", err.Error())
		}
	})

	router.HandleFunc("/network", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl["network.html"].ExecuteTemplate(w, "base", &Home{})
		if err != nil {
			log.Printf("error: %s", err.Error())
		}
	})

	spa := spaHandler{staticPath: "web/assets", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	// server:
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
