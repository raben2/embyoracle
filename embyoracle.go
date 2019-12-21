package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/raben2/embyoracle/config"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func GetEmbyUserId(embyPath string) string {
	userMatch := ""
	conf := config.New()
	conf.Emby.EmbyUrl += embyPath
	fmt.Println("talking to server at: " + conf.Emby.EmbyUrl)
	r, err := http.NewRequest("GET", conf.Emby.EmbyUrl, nil)
	if err != nil {
		panic(err.Error())
	}
	r.Header.Add("User-Agent", "embyoracle")
	r.Header.Add("X-Emby-Token", conf.Emby.EmbyToken)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	var users []Users
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(data, &users)

	for u := range users {
		if strings.Compare(conf.Emby.EmbyUserName, users[u].Name) == 0 {
			userMatch := users[u].Id
			return userMatch
		}

	}
	return userMatch

}

func getUserMovieIds(embyUserId string) map[string]string {

	conf := config.New()
	embyUrl := conf.Emby.EmbyUrl
	embyUrl += "/emby/Users/"
	embyUrl += embyUserId
	embyUrl += "/Items"
	embyUrl += "?Recursive=true&IncludeItemTypes=Movie"
	r, err := http.NewRequest("GET", embyUrl, nil)
	if err != nil {
		panic(err.Error())
	}
	r.Header.Add("User-Agent", "embyoracle")
	r.Header.Add("X-Emby-Token", conf.Emby.EmbyToken)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	movies := &Movies{}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	movieList := make(map[string]string)
	json.Unmarshal(data, &movies)
	for i := range movies.Items {

		if !movies.Items[i].UserData.Played {
			movieList[movies.Items[i].Name] = movies.Items[i].ID
			continue

		}
	}
	return movieList

}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Searching random movie to watch...")
	embyUserId := GetEmbyUserId("/Users/")
	embyMovies := getUserMovieIds(embyUserId)

	numMovies := len(embyMovies)
	randMovie := rand.Intn(numMovies)
	for n := 0; n < numMovies; n++ {
		if int(n) == randMovie {
			for name, id := range embyMovies {
				fmt.Printf("Found random movie %v with ID %s\n", name, id)
				break
			}

		}
	}
}
