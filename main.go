package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("Starting the server at 8080!")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(cors.Default().Handler)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	FileServer(r, "/", filesDir)

	s := http.Server{Addr: ":8080", Handler: r}

	// For testing purposes
	go test(&s)

	s.ListenAndServe()

}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func test(s *http.Server) {
	for _, arg := range os.Args[1:] {
		if arg == "--test" {

			resp, err := http.Get("http://localhost:8080/temp.txt")
			if err != nil {
				log.Fatalln(err)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			sb := string(body)

			if sb != "notessss" {
				fmt.Println("FAIL!!")
			} else {
				fmt.Println("PASS!!")
			}

			fmt.Println("Stopping http server!")
			if err := s.Shutdown(context.Background()); err != nil {
				panic(err) // shutting down the server
			}
		}
	}
}
