package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/beevik/ntp"
)

type clock struct {
	Time   jsonTime `json:"time"`
	Server string
}

type jsonTime time.Time

func (t jsonTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).Format(`"2006-01-02T15:04"`)), nil
}

var chann = make(chan clock)

func ntpWorker(host string) {
	t, err := ntp.Time(host)
	if err != nil {
		fmt.Printf("Can't get NTP time %v - %v", host, err.Error())
		return
	}
	// write to a channel
	chann <- clock{Server: host, Time: jsonTime(t)}
}

//OClock Gives you the time of day
func OClock(w http.ResponseWriter, r *http.Request) {
	go ntpWorker("europe.pool.ntp.org")
	go ntpWorker("africa.pool.ntp.org")
	go ntpWorker("north-america.pool.ntp.org")

	// read from a channel
	c := <-chann

	cJSON, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("Can't marshal time - %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write([]byte(cJSON))
}

func main() {
	http.HandleFunc("/oclock", OClock)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
