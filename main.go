package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type requestData struct {
	Proto            string
	RequestURI       string
	RemoteAddr       string
	Host             string
	ContentLength    string
	TransferEncoding []string
	Header           map[string][]string
}

var version = "development-build"
var responseFormat string

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	d := getRequestData(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(d)
}

func textHandler(w http.ResponseWriter, r *http.Request) {
	d := getRequestData(r)
	fmt.Fprintln(w, d)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	d := getRequestData(r)

	t, _ := template.New("debug_page").Parse(`<!doctype html>
<html lang=en>
	<head>
		<meta charset=utf-8>
		<title>DummyServer</title>
		<link href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAB9klEQVR4AY2SA6wdQRiFt25YBbXNa9uubdu2bdu2bdu2+2zzdHYf912e5Avn+5WhCiZ6r7BmxD7xBsLzuHuGlMTX1tj4x+ZHMZdUdSh/IVJPQhwBNKQAkr44GRLf2zPi7xrHe5f3S2VEyri7XokJg/WIemhG/BMzI+fyyZGZ+EAj8db94+9dUhgNOqhUOsS8pyV3El5Ygtz33iWqQY98b6sGU8eYGM7vtTJ8uGlnFUj85ACua0qyu+8VtaULRB6T5T7U67ImeXbR7jZF/ENzN1aB8P1Cfc7hYm/okPjBDpddjxWzTSzxxHYr3l23I/mVvR+rQNR2TmkiZxLAcECCf3slSHhtZRU4uNGCl1fs+HJQOtX9iHvFB2g5H24FcnixkhN3c1StKqwCsTtF5X5tFyUEUuDNBsGVm+PqNHCboqWBu/XyPGFugZi7WkTdVLkRekbci/KUxk05K5o25WBgBzU2TSK7bm6W+X1+VRTkz9raS9zk5cuXl541e2H0rFkLsGXzLty4fgeP9890k3/Mr5r2Y261CpSnyJWazRKpAg5HCwweNBz9e/WK+r6g6tN8Bf5+n1d1mJuYM8H8+ctSFi1aiV27DuDe3Ue4c/sBDh06xPk5vyr/+/zqmuszqKKUrzRp1qx/k6Yc5ge2b9cFNpvrGeUn/wGWvMPf5TkecgAAAABJRU5ErkJggg==" rel="icon" type="image/png">
	</head>
	<body>
		<h1>DummyServer Debug Info</h1>

		<h2>Request Info</h2>
		<ul>
			<li><strong>Protocol:</strong> {{.Proto}}</li>
			<li><strong>Request URI:</strong> {{.RequestURI}}</li>
			<li><strong>Remote Address URI:</strong> {{.RemoteAddr}}</li>
			<li><strong>Host:</strong> {{.Host}}</li>
			<li><strong>Content Length:</strong> {{.ContentLength}}</li>
			<li><strong>Transfer Encoding:</strong> {{.TransferEncoding}}</li>
		</ul>

		<h2>Request Headers</h1>
		<ul>
			{{ range $key, $value := .Header }}
			   <li><strong>{{ $key }}</strong>: {{ $value }}</li>
			{{ end }}
		</ul>
	</body>
</html>`)

	t.Execute(w, d)
}

func getRequestData(r *http.Request) requestData {
	return requestData{
		r.Proto,
		r.RequestURI,
		r.RemoteAddr,
		r.Host,
		strconv.Itoa(int(r.ContentLength)),
		r.TransferEncoding,
		r.Header}
}

func isJSONRequest(r *http.Request) bool {
	c := r.Header.Get("Content-Type")
	return strings.Contains(c, "/json") || strings.Contains(c, "+json")
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URI: \"%s\", Remote Address: \"%s\", Host: \"%s\"", r.RequestURI, r.RemoteAddr, r.Host)

	switch responseFormat {
	case "html":
		htmlHandler(w, r)
	case "json":
		jsonHandler(w, r)
	case "text":
		textHandler(w, r)
	default:
		if isJSONRequest(r) {
			jsonHandler(w, r)
		} else {
			htmlHandler(w, r)
		}
	}
}

func main() {
	showVersion := flag.Bool("version", false, "Print version information.")
	listenAddress := flag.String("address", ":8080", "The host (optional) and port to listen on (e.g. 127.0.0.1:8080 or :4000)")
	flag.StringVar(&responseFormat, "format", "auto", "The format of the response \"auto\", \"html\", \"json\", or \"text\"")
	flag.Parse()

	fmt.Printf("dummyserver: %s\n", version)
	if *showVersion {
		os.Exit(0)
	}

	fmt.Printf("Response formatter: %s\n", responseFormat)
	fmt.Printf("Listening for requests on %s ...\n\n", *listenAddress)

	http.HandleFunc("/", pageHandler)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
