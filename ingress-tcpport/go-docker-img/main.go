package main

import (
    "fmt"
    "os"
    "log"
    "net/http"
)

func main() {

    host, _ := os.Hostname()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "##################\n")
        fmt.Fprintf(w, "Hello From Host %q\n", host)
        fmt.Fprintf(w, "Hello From Host %q\n", host)
        fmt.Fprintf(w, "Hello From Host %q\n", host)
        fmt.Fprintf(w, "##################\n")
    })

    log.Fatal(http.ListenAndServe(":80", nil))

}
