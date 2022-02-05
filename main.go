package main

import (
	"GoBackend/pkg/check"
	"GoBackend/pkg/db"
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func checkHandler(w http.ResponseWriter, r *http.Request) {
	s := r.PostFormValue("text")
	hashs := check.StringHashDefault(s)
	hits := db.GetHitsByHashs(hashs)
	res, _ := json.Marshal(hits)
	fmt.Fprintf(w, string(res))
}

func main() {
	db.LoadDbDes()
	http.HandleFunc("/check", checkHandler)
	http.ListenAndServe(":8090", nil)
}
