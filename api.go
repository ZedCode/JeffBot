package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// to reduce duplicate lookups, we should find the user ID in the API and store that.
// This calls out to the following URL:
//    https://api.playbattlegrounds.com/shards/pc-na/players?filter[playerNames]=PLAYER_NAME
// using the appropriate headers and auth
func getPlayerData(playerName string) PlayersAPIResponse {
	r := PlayersAPIResponse{}
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.playbattlegrounds.com/shards/pc-na/players?filter[playerNames]="+playerName, bytes.NewBuffer([]byte("")))
	if err != nil {
		log.Println("Failed to prepare API Request: ", err.Error())
		return r
	}
	req.Header.Add("Accept", "application/vnd.api+json")
	req.Header.Add("Authorization", "Bearer "+pubgToken)
	res, err := client.Do(req)
	if err != nil {
		log.Println("Failed to make request: ", err.Error())
		return r
	}
	defer res.Body.Close()
	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Failed to read body: ", err.Error())
		return r
	}
	err = json.Unmarshal(payload, &r)
	if err != nil {
		log.Println("Failed to parse JSON: ", err.Error())
		return r
	}
	return r
}
