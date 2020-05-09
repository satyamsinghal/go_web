package main

import (
	"net/http"
	"text/template"
)

type Context struct {
	FirstName string
	Message   string
	URL       string
	beers     []string
	Title     string
}

// type MyHandler struct {
// }

// func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	path := req.URL.Path[1:]

// 	log.Println(path)

// 	data, err := ioutil.ReadFile(string(path))

// 	if err == nil {

// 		var contentType string

// 		if strings.HasSuffix(path, ".jpg") {
// 			contentType = "image/jpg"
// 		}
// 		w.Header().Add("Content Type", contentType)
// 		w.Write(data)
// 	} else {
// 		w.WriteHeader(404)
// 		w.Write([]byte("404 my friend " + http.StatusText(404)))
// 	}

// }

func main() {
	//http.Handle("/", new(MyHandler))
	//http.HandleFunc("/", myTempFunc)
	//http.ListenAndServe(":8080", http.FileServer(http.Dir("")))
	http.HandleFunc("/", myFunc1)
	http.ListenAndServe(":8080", nil)

}

func myFunc1(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content Type", "text/html")
	temp, err := template.New("mytemplate").Parse(doc)
	if err == nil {
		context := Context{
			"Vimal",
			"likes beer",
			req.URL.Path,
			[]string{"A", "B", "C"},
			"favourite",
		}
		temp.Execute(w, context)
	}
}

// func myTempFunc(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Add("Content Type", "text/html")
// 	temp, err := template.New("mytemplate").Parse(doc)
// 	if err == nil {
// 		temp.Execute(w, req.URL.Path[1:])
// 	}
// }

const doc = `
<!DOCTYPE html>
<html>
<head>
<title>Page Title</title>
</head>
<body>

	<h1>{{.FirstName}} likes "{{.Message}}"</h1>
	{{if eq .URL "/nobeer"}}
		<h2> Srry we are out of beer {{.FirstName}}</h2> 
	{{else}}
		<h2> welcome {{.FirstName}}</h2>
		<ul>
			{{range.beers}}
			<li>{{.}}</li>
			{{end}}
		</ul>
	{{end}}

	<hr>

	<h2>total information</h2>
	<p> {{.}}<\p>
	
</body>
</html>
`
