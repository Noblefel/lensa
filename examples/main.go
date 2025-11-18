package main

import (
	"net/http"

	"github.com/Noblefel/lensa"
)

// func main() {
// 	render := lensa.Default()
// 	render.View(os.Stdout, "index", nil)
// }

func main() {
	render := lensa.Default()
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render.View(w, "index", nil)
	})

	mux.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		render.View(w, "dashboard", nil)
	})

	http.ListenAndServe(":8080", mux)
}
