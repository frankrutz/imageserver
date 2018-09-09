package main
import (
    "log"
    "net/http"
)
func main() {
    http.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("./tmp/"))))
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
