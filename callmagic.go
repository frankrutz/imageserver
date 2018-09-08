package main
import (
    "os"
    "log"
    "os/exec"
    "strconv"
    "unicode/utf8"
)
func main() {
    var imglength int
    inputname := os.Args[1]
    printtext := inputname+"s"
    outtext   := printtext+ "Life"
    outtextlength := utf8.RuneCountInString(outtext)
    imglength = 74 * outtextlength
    cmnd := exec.Command("/root/src/src/imageserver/tmp/printimage.sh",inputname, printtext, strconv.Itoa(imglength))
    //cmnd.Run() // and wait
    cmnd.Start()
    log.Println("log")
}
