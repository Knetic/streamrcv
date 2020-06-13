package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"flag"
	"net/http"
	"time"
)

var passkeyPath string

var passkeys map[string]string
var lastRetrieved time.Time

func main() {	
	var host string

	flag.StringVar(&passkeyPath, "f", "/etc/sauth/passkeys.conf", "Path of the passkey file")
	flag.StringVar(&host, "h", "0.0.0.0:1930", "IP/Port to listen on")

	pk, err := loadPasskeys(passkeyPath)
	if err != nil {
		fmt.Printf("unable to load passkeys from '%s': %v\n", passkeyPath, err)
		os.Exit(1)
	}
	passkeys = pk
	lastRetrieved = time.Now()

	http.HandleFunc("/auth", handleAuth)
	
	fmt.Printf("Listening on '%s'\n", host)
	err = http.ListenAndServe(host, nil)
	if err != nil {
		fmt.Printf("unable to serve http: %v\n", err)
	}
}

func handleAuth(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form body", 400)
		return
	}

	names, ok := r.Form["name"]
	if !ok || len(names) <= 0 {
		http.Error(w, "no 'name' provided in form body", 400)
		return
	}
	streamKey := names[0] 

	pk, ok := r.Form["pk"]
	if !ok || len(pk) <= 0 {
		http.Error(w, "no 'pk' provided in form body", 400)
		return
	}
	passkey := strings.Join(pk, "")

	err = refreshPasskeys()
	if err != nil {
		http.Error(w, "unable to refresh passkeys", 500)
		return
	}

	actual, ok := passkeys[streamKey]
	if !ok {
		http.Error(w, "no such stream key registered", 401)
		return
	}

	if actual != passkey {
		http.Error(w, "given pk doesn't match registered", 403)
		return
	}

	// if everything went through, no problem.
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

func refreshPasskeys() error {

	diff := time.Since(lastRetrieved)
	if diff.Seconds() >= 5 {
		pk, err := loadPasskeys(passkeyPath)
		if err != nil {
			return err
		}
		passkeys = pk
	}

	return nil
}