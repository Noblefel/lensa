{{define "admin.layout"}}
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    {{block "head" .}}{{end}}
  </head>
  <body>
    <h1>Admin Layout</h1>
    {{block "body" .}}{{end}}
  </body>
</html>
{{end}}