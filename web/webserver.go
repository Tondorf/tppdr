package web

import (
	"fmt"
	"net/http"
)

func Webserver() {
	http.Handle("/", http.FileServer(http.Dir("./wwwdata/")))
	http.HandleFunc("/press", handleDown)
	http.HandleFunc("/release", handleUp)
	if err := http.ListenAndServe(":8000", nil); err != http.ErrServerClosed {
		fmt.Printf("Ouch, server closed: %v\n", err)
	}
}

func handleDown(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		fmt.Println("Url Param 'key' is missing")
		return
	}
	key := keys[0]
	fmt.Println("press:", key, "from client", r.RemoteAddr)
}

func handleUp(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		fmt.Println("Url Param 'key' is missing")
		return
	}
	key := keys[0]
	fmt.Println("release:", key, "from client", r.RemoteAddr)
}
