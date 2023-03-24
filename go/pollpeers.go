package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const rpcCommand string = `{"jsonrpc":"2.0","method":"admin_peers","params":[],"id":1}`

type PeerResponse struct {
	Peers []Peer `json:"result"`
}

type Peer struct {
	Id         string
	Name       string
	Enode      string
}

func main() {
	pollRate := flag.Duration("pollrate", 3, "poll rate in seconds")
	addr := flag.String("addr", "localhost", "ip address of RPC API")
	port := flag.String("port", "8545", "RPC port")
	
	flag.Parse()

	client := &http.Client{}

	timer := time.NewTicker(*pollRate * time.Second)

	for range timer.C {
		rpcPayload := strings.NewReader(rpcCommand)

		resp, err := client.Post(fmt.Sprintf("http://%s:%s", *addr, *port), "application/json", rpcPayload)
		if err != nil {
			log.Println("Error making JSON-RPC request:", err)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Failed to ready resp body:", err)
			continue
		}

		var response PeerResponse
		json.Unmarshal(body, &response)

		log.Println("Connected peers:", len(response.Peers))
		for i, peer := range response.Peers {
			log.Printf("%5d. %s\n", i+1, peer.Id)
		}

		resp.Body.Close()
	}
}
