package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
/* check file url cointaining vuln RCE for interact with agent-shellter
/*/
func DoCheckFileBackdoor(url string) {
	fmt.Println("[+] Checking shellter file ..")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	if resp.StatusCode == 200 {
		fmt.Println("[+] Interactive shellter found !!")
		return
	}
	fmt.Println("[!] Failed. Exploit Maybe no use")
	return
}

/* Send Payload in QueryParams
 */
func DoRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	fmt.Print(string(body))
	fmt.Println("\n")
}

/* Send Payload in body form */
func DoPostRequestPayloadInBody(url, payload string) {

	req, err := http.NewRequest("GET", url, strings.NewReader(payload))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
