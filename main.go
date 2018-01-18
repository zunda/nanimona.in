package main

import (
	"fmt"
	"html"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const template = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
<html>
<head>
<title>nanimona.in</title>
<style type="text/css">
	body {
		background-color: #101020;
		color: #808080;
	}
	p.nothing {
		text-align: center;
		font-size: 300%%;
		font-family: monospace;
		margin: 3em;
	}
	span.input {
		font-weight: bold;
	}
</style>
</head>
<body>
<p class="nothing">
<span class="prompt">%s</span>%s
<span class="input">%s</span>
</p>
</body>
</html>
`

var nothings = []struct {
	prompt    string
	result    string
	linebreak bool
}{
	{"", "無", false},                          // Japanese
	{"GET / HTTP/1.1", "404 Not Found", true}, // HTTP
	{">>>", "None", false},                    // Python
	{"irb(main):001:0>", "nil", false},        // Ruby
	{">", "undefined", false},                 // Javascript
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	rand.Seed(time.Now().UnixNano())

	h := http.NewServeMux()
	h.HandleFunc("/favicon.ico", http.NotFound)
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		x := nothings[rand.Intn(len(nothings))]
		b := ""
		if x.linebreak {
			b = "<br>"
		}
		fmt.Fprintf(w, template, html.EscapeString(x.prompt), b, html.EscapeString(x.result))
	})

	log.Println("Listening at port " + port)
	err := http.ListenAndServe(":"+port, h)
	log.Fatal(err)
}
