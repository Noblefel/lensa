package lensa

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var templateHome = []byte(`{{template "layout" .}}{{define "body"}}hello {{index . "name"}}{{end}}`)

var templateLayout = []byte(`{{define "layout"}}title-{{block "body" .}}{{end}}-footer{{end}}`)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func join(s ...string) string { return filepath.Join(s...) }

func TestView(t *testing.T) {
	dir := t.TempDir()
	pagesDir := join(dir, "pages")
	partsDir := join(dir, "parts")

	must(os.Mkdir(pagesDir, os.ModePerm))
	must(os.Mkdir(partsDir, os.ModePerm))

	must(os.WriteFile(join(pagesDir, "index.tem"), templateHome, os.ModePerm))
	must(os.WriteFile(join(partsDir, "layout.tem"), templateLayout, os.ModePerm))

	var (
		buf    = &strings.Builder{}
		render = New(pagesDir, partsDir, ".tem")
		data   = map[string]any{"name": "JEB"}
	)

	if err := render.View(buf, "index", data); err != nil {
		t.Fatal("expecting no error, got:", err)
	}

	want := "title-hello JEB-footer"

	if buf.String() != want {
		t.Fatalf("\nwant:%s\ngot:%s", want, buf.String())
	}

	if err := render.View(buf, "x", data); err == nil {
		t.Fatal("expecting error for no template")
	}

	buf.Reset()

	nestDir := join(partsDir, "nested")
	must(os.Mkdir(nestDir, os.ModePerm))
	must(os.WriteFile(
		join(nestDir, "nested.tem"),
		[]byte(`{{define "nested"}}nested{{end}}`),
		os.ModePerm,
	))

	must(os.Remove(join(pagesDir, "index.tem")))
	must(os.WriteFile(
		join(pagesDir, "index.tem"),
		[]byte(`hello {{block "nested" .}}{{end}}`),
		os.ModePerm,
	))

	if err := render.View(buf, "index", nil); err != nil {
		t.Fatal("expecting no error for nested partials, got:", err)
	}

	want = "hello nested"

	if buf.String() != want {
		t.Fatalf("\nwant:%s\ngot:%s", want, buf.String())
	}

	render.UseCache()

	if err := render.View(buf, "index", nil); err != nil {
		t.Fatal("expecting no error for cached template, got:", err)
	}

	buf.Reset()

	must(os.Remove(join(pagesDir, "index.tem")))
	must(os.WriteFile(
		join(pagesDir, "index.tem"),
		[]byte(`from function: {{ greet }}`),
		os.ModePerm,
	))

	render = New(pagesDir, partsDir, ".tem")
	render.UseFuncs(map[string]any{
		"greet": func() string { return "greetings" },
	})

	if err := render.View(buf, "index", nil); err != nil {
		t.Fatal("expecting no error for template func, got:", err)
	}

	want = "from function: greetings"

	if buf.String() != want {
		t.Fatalf("\nwant:%s\ngot:%s", want, buf.String())
	}
}
