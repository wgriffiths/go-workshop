package main

import (
	"log"
	"net/http"
	"time"
)

//OClock Gives you the time of day
func OClock(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().Format("2006-01-02T15:04")))
}

func main() {
	http.HandleFunc("/oclock", OClock)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
