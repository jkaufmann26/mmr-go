package frontend

import (
	"net/http"
)


func IntitializeServer() *http.ServeMux{
    router := http.NewServeMux()
    router.HandleFunc("/", Index)

    return router

}

func Index(w http.ResponseWriter, r *http.Request){
   if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    
    MustRenderHtml(w, "homepage.html", nil)

    switch r.Method {
    case http.MethodGet:
        // Handle the GET request...

    case http.MethodPost:
        // Handle the POST request...

    case http.MethodOptions:
        w.Header().Set("Allow", "GET, POST, OPTIONS")
        w.WriteHeader(http.StatusNoContent)

    default:
        w.Header().Set("Allow", "GET, POST, OPTIONS")
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }    
}
