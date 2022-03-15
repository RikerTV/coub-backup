package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// RetrieveProfile a complete array of CoubInfo structs for a given profile
func RetrieveProfile(username string) (coubs []Coub, err error) {
	// Get first page of profile info with an http.Get request
	log.Print("Retrieving first page to parse channel info")

	timelineFirstPage := "https://coub.com/api/v2/timeline/channel/" + username + "?page=1&per_page=25"
	res, err := http.Get(timelineFirstPage)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var Coubs []Coub

	var CoubInfoPage CoubInfo
	err = json.Unmarshal(b, &CoubInfoPage)
	if err != nil {
		return nil, err
	}

	Coubs = append(Coubs, CoubInfoPage.Coubs...)

	log.Print("Pages to parse: ", CoubInfoPage.TotalPages)

	// Get all pages of profile info with an http.Get request
	for i := 2; i <= CoubInfoPage.TotalPages; i++ {
		log.Print("Retrieving page ", i)

		timelinePage := "https://coub.com/api/v2/timeline/channel/" + username + "?page=" + strconv.Itoa(i) + "&per_page=25"
		res, err := http.Get(timelinePage)
		if err != nil {
			return nil, err
		}

		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var CoubInfoPage CoubInfo
		err = json.Unmarshal(b, &CoubInfoPage)
		if err != nil {
			return nil, err
		}

		Coubs = append(Coubs, CoubInfoPage.Coubs...)
	}

	return Coubs, nil
}

func GenerateInfoFile(outputDir string, user string) (err error) {
	log.Print("Getting Coub user data for: ", user)
	coubs, err := RetrieveProfile(user)
	if err != nil {
		return err
	}

	outputFile, _ := json.MarshalIndent(coubs, "", " ")
	_ = ioutil.WriteFile(outputDir+user+".json", outputFile, 0644)

	return nil
}

func DownloadFile(filepath string, url string) error {
	// Check if a file already exists at filepath
	if _, err := os.Stat(filepath); err == nil {
		log.Print("File already exists at ", filepath)
		return nil
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func FileNameFromURL(url string) (filename string) {
	filename = url[strings.LastIndex(url, "/")+1:]
	return
}
