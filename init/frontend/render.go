package frontend

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

func MustRenderHtml(w http.ResponseWriter, path string, data map[string]any) {
    workingDir, err  := os.Getwd()
    if err != nil {
        panic(fmt.Errorf("An error occurred getting working dir %w", err))
    }
    fmt.Println(workingDir)
	template, err := template.ParseFiles(workingDir + "/init/frontend/html/"+path)
	if err != nil {
		panic(fmt.Errorf("Something went wrong parsing path: %s, %w", path, err))
	}

	err = template.Execute(w, data)
	if err != nil {
		panic(fmt.Errorf("Something went wrong parsing data from map %w", err))
	}
}
