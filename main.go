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

var nothings = []struct {
	prompt    string
	result    string
	linebreak bool
}{
	{"", "無", false},                           // Japanese
	{"GET / HTTP/1.1", "404 Not Found", true},  // HTTP
	{"GET / HTTP/1.1", "204 No Content", true}, // HTTP
	{">>>", "None", false},                     // Python
	{"irb(main):001:0>", "nil", false},         // Ruby
	{">", "undefined", false},                  // Javascript
	{"$ cat /dev/null", "", true},              // POSIX and shell
	{"$ nslookup nanimona.in", "** server can't find  nanimona.in: NXDOMAIN", true},
	// DNS
	{"GET / HTTP/1.1", "ERR_EMPTY_RESPONSE", true}, // HTTP, ELB, and Chrome
	{"", "∅", false},                               // maths
	{"", "void", false},                            // C
	{"", "NULL", false},                            // C
}

const template = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title>nanimona.in</title>
<style type="text/css">
	body {
		background-color: #101020;
		color: #808080;
	}
	div.nothing {
		text-align: center;
	}
	p.nothing {
		text-align: left;
		font-size: 300%%;
		font-family: monospace;
		margin: 3em 0em 3em 0em;
		display: inline-block;
	}
	span.input {
		font-weight: bold;
	}
	p.footer {
		text-align: right;
		margin: 1em;
	}
	a {
		text-decoration: none;
	}
	a:link, a:visited {
		color: #409040;
	}
	a:hover, a:active {
		color: #60B060;
	}
</style>
</head>
<body>
<div class="nothing"><p class="nothing">
<span class="prompt">%s</span>%s
<span class="input">%s</span>
</p></div>
<p class="footer">Fork me at <a href="https://github.com/zunda/nanimona.in">GitHub</a></p>
</body>
</html>
`

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
