package main

//hi
import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// git
// templ is a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP for HTTP request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	globalRoom := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})

	http.Handle("/room", globalRoom)
	// start room
	go globalRoom.run()

	// starting the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
