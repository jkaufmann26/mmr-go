package main

import (
	"fmt"
	"go-gaming/init/frontend"
	"net/http"
)

func main() {
	server := frontend.IntitializeServer()
    err :=	http.ListenAndServe("localhost:3000", server)
    if err != nil {
        fmt.Print("An error occurred starting the server", err)
    }
}
