package hello

import (
    "io"
    "html"
    "net/http"
    "appengine"
    "appengine/urlfetch"
    "github.com/gorilla/mux"
    "html/template"
)

func handler(w http.ResponseWriter, r* http.Request) {
    t, _ := template.ParseFiles("view/index.html")
    t.Execute(w, nil)
}

func copyHeader(dst, src http.Header) {
    for k, w := range src {
        for _, v := range w {
            dst.Add(k, v)
        }
    }
}
    
func copyResponse(r *http.Response, w http.ResponseWriter) {
    copyHeader(w.Header(), r.Header)
    w.WriteHeader(r.StatusCode)
    io.Copy(w, r.Body)
}
    
func cloudAdminHandler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    client := urlfetch.Client(c)
    
    view := html.EscapeString(r.URL.Path)
    if view == "/cloudadmin" {
        view = ""
    }
    
    resp, err := client.Get("http://0.0.0.0:8000" + view)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    copyResponse(resp, w)
}

func init() {
    r := mux.NewRouter()

    s1 := r.PathPrefix("/api").Methods("GET").Subrouter()
    s2 := r.PathPrefix("/api").Methods("POST").Subrouter()
    
    s1.HandleFunc("/users", userGetHandler)
    s2.HandleFunc("/users", userPostHandler)

    r.HandleFunc("/", handler)
    r.HandleFunc("/{.path:.*}", cloudAdminHandler)
    
    http.Handle("/", r)
}
