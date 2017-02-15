package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type clock struct {
	Time jsonTime `json:"time"`
}

type jsonTime time.Time

func (t jsonTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).Format(`"2006-01-02T15:04"`)), nil
}

//OClock Gives you the time of day
func OClock(w http.ResponseWriter, r *http.Request) {

	c := clock{Time: jsonTime(time.Now())}

	cJSON, err := json.Marshal(c)
	if err != nil {
		fmt.Errorf("Can't marshal time - %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write([]byte(cJSON))
}

func main() {
	http.HandleFunc("/oclock", OClock)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
