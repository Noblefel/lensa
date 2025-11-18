{{template "admin.layout" .}}
{{define "head"}}<title>Dashboard</title>{{end}}

{{define "body"}}
    <h2>Welcome to dashboard page</h2>
    {{block "table" .}}{{end}}
{{end}}
