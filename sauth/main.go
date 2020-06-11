package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"flag"
	"net/http"
)

var passkeys map[string]string

func main() {

	var passkeyPath string
	var host string

	flag.StringVar(&passkeyPath, "-f", "/etc/sauth/passkeys.conf", "Path of the passkey file")
	flag.StringVar(&host, "-h", "127.0.0.1:1930", "IP/Port to listen on")

	pk, err := loadPasskeys(passkeyPath)
	if err != nil {
		fmt.Printf("unable to load passkeys from '%s': %v\n", passkeyPath, err)
		os.Exit(1)
	}
	passkeys = pk
	http.HandleFunc("/auth", handleAuth)
	
	fmt.Printf("Listening on '%s'\n", host)
	err = http.ListenAndServe(host, nil)
	if err != nil {
		fmt.Printf("unable to serve http: %v\n", err)
	}
}

func handleAuth(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("handling auth\n")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form body", 400)
	}

	fmt.Printf("form values: %v\n", r.Form)
}

func loadPasskeys(path string) (map[string]string, error) {

	ret := make(map[string]string)

	fd, err := os.Open(path)
	if err != nil {
		return ret, err
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		
		line := scanner.Text()
		idx := strings.Index(line, " ")
		
		if idx <= 0 {
			continue
		}

		ret[line[0:idx]] = line[idx+1:]
	}

	return ret, nil
}