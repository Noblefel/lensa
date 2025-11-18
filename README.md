A wrapper around go html/template to skip the tedious parsing part. Not recommended

**THIS IS MOSTLY UNTESTED, SO DONT USE IT 😁**

```
go get github.com/Noblefel/lensa
```

Quick example writes to standard output (see examples):

```go
func main() {
	render := lensa.Default()
	render.View(os.Stdout, "index", nil)
}
```

Basic example for web (see examples):

```go
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
```

For caching, simply use UseCache function:

```go
	render := lensa.Default()
    render.UseCache() // <-- should be done once after creation
```

**Warning** - make sure you separate the templates for pages (like index, dashboard etc) and templates for parts (layout, components) in their own directory.
