package web

import (
	"fmt"
	"github.com/Tondorf/tppdr/common"
	"net/http"
	"strings"
)

type Webserver struct {
	ch chan common.BrowserEvent
}

func NewWebserver(ch chan common.BrowserEvent) Webserver {
	web := Webserver{ch}
	http.Handle("/", http.FileServer(http.Dir("./wwwdata/")))
	http.HandleFunc("/press", web.handlePress)
	http.HandleFunc("/release", web.handleRelease)
	return web
}

func (web *Webserver) Listen() {
	if err := http.ListenAndServe(":1234", nil); err != http.ErrServerClosed {
		fmt.Printf("Ouch, server closed: %v\n", err)
	}
}

func (web *Webserver) handlePress(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		fmt.Println("Url Param 'key' is missing")
		return
	}
	key := keys[0]
	fmt.Println("press:", key, "from client", r.RemoteAddr)

	ip := strings.Split(r.RemoteAddr, ":")[0]
	bke := common.BrowserEvent{Origin: ip, Key: key, Typ: common.Press}
	web.ch <- bke // send key to the channel
}

func (web *Webserver) handleRelease(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		fmt.Println("Url Param 'key' is missing")
		return
	}
	key := keys[0]
	fmt.Println("release:", key, "from client", r.RemoteAddr)

	ip := strings.Split(r.RemoteAddr, ":")[0]
	bke := common.BrowserEvent{Origin: ip, Key: key, Typ: common.Release}
	web.ch <- bke // send key to the channel
}
