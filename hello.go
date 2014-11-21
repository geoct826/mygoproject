package hello

import (
    "io"
    "fmt"
    "html"
    "net/http"
    "appengine"
    "appengine/urlfetch"
    "github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r* http.Request) {
    fmt.Fprintf(w, "Hello, World! %q", html.EscapeString(r.URL.Path))
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
    
    r.HandleFunc("/", handler)
    r.HandleFunc("/{.path:.*}", cloudAdminHandler)
    
    http.Handle("/", r)
    //http.HandleFunc("/", handler)
    //http.HandleFunc("/cloudadmin", cloudAdminHandler)
    //http.HandleFunc("/datastore", cloudAdminHandler)
}
