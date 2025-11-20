package lensa

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"io/fs"
	"path/filepath"
)

type Render struct {
	pagesDir string // directory for pages
	partsDir string // directory for layouts, components etc.
	ext      string // template file extension.

	cache map[string]*template.Template // the store for UseCache
	//TODO: FuncMap
}

// Creates Render with default values. Throws panic if there's path error.
//
// Directory for pages (index, dashboard etc)  must be "pages"
//
// Directory for parts (components, layout etc) must be "partials"
//
// File extension for templates must be ".tpl"
func Default() *Render { return New("pages", "partials", ".tpl") }

// Creates Render with custom values. Throws panic if there's path error.
func New(pagesDir, partsDir, ext string) *Render {
	// assert
	if _, err := glob(pagesDir, ext); err != nil {
		panic(err)
	}
	if _, err := glob(partsDir, ext); err != nil {
		panic(err)
	}
	return &Render{pagesDir, partsDir, ext, nil}
}

// Parse all pages into cache
func (r *Render) UseCache() {
	r.cache = make(map[string]*template.Template)

	pages, _ := glob(r.pagesDir, r.ext) // assert already done in New()
	parts, _ := glob(r.partsDir, r.ext)

	for _, p := range pages {
		files := append([]string{p}, parts...)

		t, err := template.ParseFiles(files...)
		if err != nil {
			panic(err)
		}

		r.cache[p] = t
	}
}

// Parses the page file and all the needed parts in it.
// if UseCache, immediately writes the template.
func (r *Render) View(w io.Writer, page string, data any) error {
	page = filepath.Join(r.pagesDir, page+r.ext)

	if r.cache != nil {
		t, ok := r.cache[page]
		if !ok {
			return errors.New("no template found")
		}

		return write(w, t, data)
	}

	parts, err := glob(r.partsDir, r.ext)
	if err != nil {
		return err
	}

	files := append([]string{page}, parts...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}

	return write(w, t, data)
}

func glob(root, ext string) ([]string, error) {
	var files []string

	walk := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		n := info.Name()
		// this is just HasSuffix
		if len(n) >= len(ext) && n[len(n)-len(ext):] == ext {
			files = append(files, path)
		}

		return nil
	}

	return files, filepath.Walk(root, walk)
}

func write(w io.Writer, t *template.Template, data any) error {
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return err
	}
	_, err := buf.WriteTo(w)
	return err
}
