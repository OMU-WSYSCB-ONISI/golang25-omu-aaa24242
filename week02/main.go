package main
import (
	"fmt"
	"net/http"
)
func main() {
http.HandleFunc("/hello", hellohandler)
}
func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Glitch !")
}
