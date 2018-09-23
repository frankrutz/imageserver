package main
 
import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "strconv"
    "unicode/utf8"
    "unicode"
    "strings"
)
func onlyLetters(s string) bool {
    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}
func toLetters(s string) string {
    var ret string
    for _, r := range s {
        if unicode.IsLetter(r) {
           ret=ret+string(r);
        }
    }
    return ret
}
func makeimage(vorname string) int {
    var rc int = 1
    var imglength int
    inputname := vorname
    printtext := inputname
    outtext   := printtext+ "Life"
    outtextlength := utf8.RuneCountInString(outtext)
    imglength = 77 * outtextlength
    cmnd := exec.Command("./printimage.sh",inputname, printtext, strconv.Itoa(imglength))
    cmnd.Start()
    log.Println("log")
    return rc
}

func imageMakerHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "hello: 404 not found.", http.StatusNotFound)
        return
    }
    switch r.Method {
    case "GET":
         http.ServeFile(w, r, "form.html")
    case "POST":
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }
        name := strings.TrimSpace(r.FormValue("name"))
        name=toLetters(name)
        if len(name) == 0 {
          name="Noname"
        }
        if name[len(name)-1:] == "s" || name[len(name)-1:] == "x" {
           name=name+""
        } else {
           name=name+"s"
        }
        if onlyLetters(name) {
           makeimage(name)
           fmt.Fprintf(w,"<!DOCTYPE html><HTML>")
           fmt.Fprintf(w,"<CENTER><a href=\"./tmp/%s.png\">click to download %s.png</n></CENTER></HTML>",name,name)
           fmt.Fprintf(w,"<meta http-equiv=\"refresh\" content=\"1; url=./tmp/%s.png\" />",name)
        } else {
           fmt.Fprintf(w,"invalid input.\nNo spaces.\nNo spacial characters.\nPlease go back and try again\n")
        }
    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}

func main() {
    http.HandleFunc("/", imageMakerHandler)
    http.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("./tmp/"))))
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatal(err)
    }
}
