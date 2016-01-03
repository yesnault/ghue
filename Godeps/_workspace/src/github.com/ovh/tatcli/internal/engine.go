package internal

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// GetSkipLimit gets skip and limit in args array
// default skip to 0 and limit to 10
func GetSkipLimit(args []string) (string, string) {
	skip := "0"
	limit := "10"
	if len(args) == 3 {
		skip = args[1]
		limit = args[2]
	} else if len(args) == 2 {
		skip = args[0]
		limit = args[1]
	}
	return skip, limit
}

func initRequest(req *http.Request) {
	req.Header.Set("Tat_username", viper.GetString("username"))
	req.Header.Set("Tat_password", viper.GetString("password"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")
}

// GetWantReturn GET on path and check is http return code is 200 and
// returns results
func GetWantReturn(path string) string {
	return ReqWant("GET", http.StatusOK, path, nil)
}

// PutWantReturn PUT on path, check is http return code is 200 and
// returns results
func PutWantReturn(path string) string {
	return ReqWant("PUT", http.StatusOK, path, nil)
}

// PutWant PUT on path and check is http return code is 200
func PutWant(path string, jsonStr []byte) string {
	return ReqWant("PUT", http.StatusCreated, path, jsonStr)
}

// PostWant POST on path and check is http return code is 201
func PostWant(path string, jsonStr []byte) string {
	return ReqWant("POST", http.StatusCreated, path, jsonStr)
}

// DeleteWant DELETE on path and check is http return code is 200
func DeleteWant(path string, jsonStr []byte) {
	ReqWant("DELETE", http.StatusOK, path, jsonStr)
}

// IsHTTPS returns true if url begins with https
func IsHTTPS() bool {
	return strings.HasPrefix(viper.GetString("url"), "https")
}

func getHTTPClient() *http.Client {
	var tr *http.Transport
	if IsHTTPS() {
		tlsConfig := GetTLSConfig()
		tr = &http.Transport{TLSClientConfig: tlsConfig}
	} else {
		tr = &http.Transport{}
	}

	return &http.Client{Transport: tr}
}

// GetTLSConfig returns tls.config with flag InsecureSkipVerify
func GetTLSConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: SSLInsecureSkipVerify,
	}
}

// ReqWant execute a request of method (POST, PUT, DELETE) on path, checks
// if return HTTP Code is equals to wantCode
func ReqWant(method string, wantCode int, path string, jsonStr []byte) string {
	ReadConfig()

	if viper.GetString("url") == "" {
		fmt.Fprintf(os.Stderr, "Invalid Configuration : invalid URL. See tatcli config --help\n")
		os.Exit(1)
	}

	var req *http.Request
	if jsonStr != nil {
		req, _ = http.NewRequest(method, viper.GetString("url")+path, bytes.NewReader(jsonStr))
	} else {
		req, _ = http.NewRequest(method, viper.GetString("url")+path, nil)
	}

	initRequest(req)
	resp, err := getHTTPClient().Do(req)
	Check(err)
	defer resp.Body.Close()

	if resp.StatusCode != wantCode || Verbose {
		fmt.Printf("Response Status: %s and we want %d\n", resp.Status, wantCode)
		fmt.Printf("Request path: %s on %s\n", method, viper.GetString("url")+path)
		fmt.Printf("Request: %s\n", string(jsonStr))
		fmt.Printf("Response Headers: s%s\n", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Response Body: %s\n", string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	Check(err)
	return getJSON(body)
}

func getJSON(s []byte) string {
	if Pretty {
		var out bytes.Buffer
		json.Indent(&out, s, "", "  ")
		return out.String()
	}
	return string(s)
}

// Check checks error, if != nil, throw panic
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
