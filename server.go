package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "pong")
}

func hook(w http.ResponseWriter, r *http.Request) {
	os.Chdir(Config.Path)

	cmd := exec.Command(Config.Cmd, Config.Args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "error: "+err.Error()+"\n")
	} else {
		w.WriteHeader(http.StatusOK)
	}

	res := string(out)
	log.Println("out: ", res)
	io.WriteString(w, res)
}
