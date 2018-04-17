package responder

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type textHTMLResponder struct{}

func (responder textHTMLResponder) ContentType() string {
	return "text/html"
}

type htmlResponse struct {
	UUID    string
	Version int
}

func (responder textHTMLResponder) WriteResponse(uuid uuid.UUID, w *http.ResponseWriter) error {
	tmpl, err := template.New("gimmeanuuid").Parse(htmlTemplate)
	if err != nil {
		return err
	}

	tmpl.Execute(*w, htmlResponse{UUID: uuid.String(), Version: int(uuid.Version())})
	return nil
}

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>gimme an uuid {{ .Version }}</title>
    <style>
        * {
            box-sizing: border-box;
        }

        html,
        body,
        main {
            margin: 0;
            width: 100%;
            height: 100%;
        }

        main {
            display: flex;
            justify-content: center;
            align-items: center;
            text-align: center;
        }

        h1 {
            margin: 0;
        }
    </style>
</head>

<body>
    <main>
        <h1>{{ .UUID }}</h1>
    </main>
</body>

</html>`
