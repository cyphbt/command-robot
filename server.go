package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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
	if Config.Secret != "" {
		body := make([]byte, 0)
		_, err := r.Body.Read(body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "read body error: "+err.Error())
			return
		}
		defer r.Body.Close()
		if !verifySignature(Config.Secret, r.Header.Get("X-Hub-Signature-256"), body) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

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

func verifySignature(secret, signature string, payload []byte) bool {
	if len(signature) < 8 {
		return false
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expectedSignature), []byte(signature)[7:])
}
