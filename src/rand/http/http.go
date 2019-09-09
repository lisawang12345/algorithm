package main

import (
	"net/http"
	"time"
)

/*func handleTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	w.Write([]byte(currentTime.String()))
}
func whoAmI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(("I am lisa")))
}

func main() {
	http.HandleFunc("/time", handleTime)
	http.HandleFunc("/whoAreYou", whoAmI)

	http.ListenAndServe(":8081", nil)
}*/
type MyHandler struct {
}

//实现接口的ServeHTTP方法
func (this *MyHandler) handleTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	w.Write([]byte(currentTime.String()))
}
func main() {
	myHandler := &MyHandler{}
	http.Handle("/", myHandler)
	http.ListenAndServe(":8080", nil)
}
