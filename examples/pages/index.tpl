{{template "layout" .}}
{{define "head"}}<title>Home</title>{{end}}

{{define "body"}}
    <h2>Welcome to home page</h2>
    {{block "form" .}}{{end}}
{{end}}