package main

import (
	"fmt"
	"go_vue_template/cmd/handlers"
    "go_vue_template/ui"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Default values
const DEFAULT_LOGGING_PATH = "./logs"
const DEFAULT_PORT = "8080"

// Control whether stdout console output should be suppressed (only works if logging is deactivated)
const SUPPRESS_CONSOLE_OUTPUT = false

// API Paths (embed trailing slashes to retain all URL control here)

const VERSION = "v1/"
const API_PATH = "api/" + VERSION

/*
Main entry point for web version of IG Parser.
*/
func main() {
	// Load .env file
    err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

    hc := handlers.NewHandlerContext(/*DB connections etc...*/)

	// Check for custom port
	port := os.Getenv("port")
	if port == "" {
		port = DEFAULT_PORT
	}

    log.Println("Server running...")

	// Launch web server
    err = http.ListenAndServeTLS(":"+port, "cert.pem", "key.pem", router(hc))
    //http.ListenAndServe(":"+port, router(hc)) // Use this if you don't want HTTPS
	if err != nil {
		log.Fatal("Web service stopped. Error:", err)
	}

}

type rou struct {
    mux *http.ServeMux
    hc *handlers.HandlerContext
}

func newRouter(hc *handlers.HandlerContext) rou {
    var r rou
    r.hc = hc
    r.mux = http.NewServeMux()
    return r
}

func (rou *rou) handle(path string, f func(hc *handlers.HandlerContext, w http.ResponseWriter, r *http.Request)){
    rou.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
        f(rou.hc, w, r)
    })
}

// router creates a http handler for all of IGParsers endoints given a handler context.
// Endpoints such as hc.Datasets and hc.Statements are themselves routers for CRUD within their scope.
func router(hc *handlers.HandlerContext) http.Handler {
	mux := http.NewServeMux()

	// index page
	mux.HandleFunc("/", indexHandler)

	// static files
	staticFS, err := fs.Sub(ui.StaticFiles, "dist")
	if err != nil {
		log.Panic("Could not load static files from dist.")
	}
	httpFS := http.FileServer(http.FS(staticFS))
	mux.Handle("/static/", httpFS)

	// Information API
	mux.HandleFunc("/"+API_PATH, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,"TEMPLATE VALUE: REPLACE WITH DOCS OR REMOVE ENTIRELY")
	})

	// Action APIs
    /*
	    mux.HandleFunc("/"+SOME_PATH, hc.SOME_HANDLER)
    */

	return mux
}

// indexHandler takes care of any base url calls and sends user to our Single Page Application (SPA) frontend.
func indexHandler(w http.ResponseWriter, r *http.Request) {

	// Prevent any other methods than get to frontend
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Prevent any failed API calls being directed to frontend
	if strings.HasPrefix(r.URL.Path, "/api") {
		http.NotFound(w, r)
	}

	// Handle favicon edge case
	if r.URL.Path == "/favicon.ico" {
		rawfile, _ := ui.StaticFiles.ReadFile("dist/favicon.ico")
		w.Write(rawfile)
		return
	}

	// Read static files from production ready frontend.
	rawfile, err := ui.StaticFiles.ReadFile("dist/index.html")
	if err != nil {
		http.Error(w, "Could not reach index.html", http.StatusInternalServerError)
		return
	}

	w.Write(rawfile)
}
