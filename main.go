package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stevenhowes/PakGo"
)

var pak PakGo.PakFile

func thing(w http.ResponseWriter, req *http.Request) {
	out, err := pak.ReadFile(req.URL.String()[1:])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
	} else {
		w.Write(out)
	}
}

func main() {

	pak0, err := PakGo.PakLoad("example.pak")
	if err != nil {
		panic(err)
	}
	defer pak0.PakClose()

	pak = pak0

	fmt.Println("--------")

	out, err := pak0.ReadFile("folder1/file1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	fmt.Println("--------")

	out, err = pak0.ReadFile("file2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	fmt.Println("--------")

	rtr := mux.NewRouter()
	rtr.HandleFunc("/file2.txt", thing).Methods("GET")
	rtr.HandleFunc("/folder1/file1.txt", thing).Methods("GET")
	rtr.HandleFunc("/invalid.txt", thing).Methods("GET")

	//rtr.PathPrefix("/").Handler(http.FileServer( **TODO: :)** ))

	http.ListenAndServe("127.0.0.1:8881", rtr)

	select {}
}
