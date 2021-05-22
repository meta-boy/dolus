package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type author struct {
	Name string `json:"name"`
}

type embed struct {
	Description string `json:"description"`
	Author      author `json:"author"`
}

type embeds struct {
	Embeds []embed `json:"embeds"`
}

func send(hook string, content string, user string) {
	embeds := embeds{
		Embeds: []embed{
			{
				Description: content,
				Author: author{
					Name: user,
				},
			},
		},
	}

	postBody, _ := json.Marshal(embeds)

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(hook, "application/json", responseBody)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Println(sb)

}

func main() {

	webhookPtr := flag.String("webhook", "foo", "Discord webhook")
	userPtr := flag.String("user", "Anon", "an int")

	flag.Parse()

	oldContent, err := clipboard.ReadAll()

	if err != nil {
		log.Fatalln(err)
	}

	for {
		content, err := clipboard.ReadAll()
		if err != nil {
			log.Fatalln(err)
		}

		if oldContent != content {
			send(*webhookPtr, content, *userPtr)
			log.Println("Send content")
			oldContent = content

		}
		time.Sleep(500 * time.Millisecond)
	}
}
