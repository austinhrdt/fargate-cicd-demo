// demo api
package main

import (
    "fmt"
    "log"
    "net/http"
)

// home serves home page
func home(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello, World!")
}

// main application entrypoint
func main() {
    http.HandleFunc("/", home)
    log.Fatal(http.ListenAndServe(":80", nil))
}
