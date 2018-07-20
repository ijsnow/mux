package examples

import (
	"net/http"

	"github.com/ijsnow/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ProductsHandler)
	http.Handle("/", r)
}

// HomeHandler handles the home route.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

// ProductsHandler handles the products route.
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Products"))
}
