package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type HandsOn struct {
	Time     time.Time `json:"time"`
	Hostname string    `json:"hostname"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// resp := fmt.Sprintf("la hora es %v y hostname es %v", time.Now(), os.Getenv("HOSTNAME"))
		// resp = fmt.Sprintf(`{"message": "%s"}`, resp)

		resp := HandsOn{
			Time:     time.Now(),
			Hostname: os.Getenv("HOSTNAME"),
		}
		jsonResp, err := json.Marshal(&resp)

		if err != nil {
			panic(err)
		}

		w.Write([]byte(jsonResp))
	})

	fmt.Println("server running at port 9090")
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		panic(err)
	}

}
