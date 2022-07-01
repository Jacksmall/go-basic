package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":8081", "连接的地址")

var templ = template.Must(template.New("qr").Parse(templateStr))

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
    <input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`

func main() {
	flag.Parse()

	server := &http.Server{
		Addr:           *addr,
		Handler:        http.HandlerFunc(QR),
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 10 << 20, // 10M
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func QR(w http.ResponseWriter, r *http.Request) {
	templ.Execute(w, r.FormValue("s"))
}
