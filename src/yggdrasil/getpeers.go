package yggdrasil

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetPeers() map[string][]string {
	peersBody := loadPeers()
	peersMap := encodeBytesPeers(peersBody)

	peers := make(map[string][]string)
	for country, value := range peersMap {

		country = strings.Replace(country, ".md", "", -1)
		peers[country] = []string{}

		for peer := range value {
			peers[country] = append(peers[country], peer)
		}
	}

	defer peersBody.Close()
	return peers
}

func loadPeers() io.ReadCloser {
	res, err := http.Get("https://publicpeers.neilalexander.dev/publicnodes.json")
	if err != nil {
		log.Fatal(err)
	}
	return res.Body
}

func encodeBytesPeers(body io.ReadCloser) map[string]map[string]interface{} {
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}

	var result map[string]map[string]interface{}
	json.Unmarshal(bytes, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
