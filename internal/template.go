package internal

import (
	"bytes"
	"fmt"
	"text/template"
)

var tpl = template.Must(template.New("response").Parse(`<!DOCTYPE html>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
        <meta name="go-import" content="{{ .AliasModPath }} {{ .Protocol }} {{ .ActualUrl }}">
    </head>
</html>
`))

type templateData struct {
	AliasModPath string
	Protocol     string
	ActualUrl    string
}

func CreateVanityServerResponse(vanityHost string, aliasPath string, l ActualLocation) string {
	td := templateData{
		AliasModPath: fmt.Sprintf("%s/%s", vanityHost, aliasPath),
		Protocol:     l.Protocol,
		ActualUrl:    l.Uri,
	}
	b := bytes.Buffer{}
	tpl.Execute(&b, td)
	return b.String()
}
