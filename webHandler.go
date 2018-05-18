package main
import (
        "fmt"
        "net/http"
    )
func homeHandeler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"i love her")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"this our about page")
}
func main() {
	http.HandleFunc("/",homeHandeler);
	http.HandleFunc("/about",aboutHandler);
	http.ListenAndServe(":8000",nil)
}
