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

const MaxRetries = 3

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

func GenerateBestOfInfoFile(outputDir string, year string) (err error) {
	log.Print("Getting Coub best of data")
	coubs, err := RetrieveBestOf(year)
	if err != nil {
		return err
	}

	outputFile, _ := json.MarshalIndent(coubs, "", " ")
	_ = ioutil.WriteFile(outputDir+"bestof.json", outputFile, 0644)

	return nil
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

func DownloadFile(filepath string, url string) (err error) {
	if url == "" {
		return nil
	}
	// Check if a file already exists at filepath
	if _, err := os.Stat(filepath); err == nil {
		//log.Print("File already exists at ", filepath)
		return nil
	}

	// Get the data and retry if it fails
	var resp *http.Response
	for i := 0; i < MaxRetries; i++ {
		resp, err = http.Get(url)
		if err != nil {
			log.Print("Failed to download file: ", err, "retrying...")
			if i == MaxRetries-1 {
				return err
			}
			continue
		}
		break
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

// RetrieveBestOf a complete array of CoubInfo structs for the best of category
func RetrieveBestOf(year string) (coubs []Coub, err error) {
	// Get first page of profile info with an http.Get request
	oldAPI := false
	log.Print("Retrieving first page to parse best of info")
	var yearID string
	switch year {
	case "2021":
		yearID = "62"
	case "2020":
		yearID = "53"
	case "2019":
		yearID = "14"
	case "2018":
		yearID = "2018"
		oldAPI = true
	case "2017":
		yearID = "2017"
		oldAPI = true
	case "2016":
		yearID = "2016"
		oldAPI = true
	case "2015":
		yearID = "2015"
		oldAPI = true
	case "2014":
		yearID = "2014"
		oldAPI = true
	case "2013":
		yearID = "2013"
		oldAPI = true
	case "2012":
		yearID = "2012"
		oldAPI = true
	}
	var timelineFirstPage string
	if !oldAPI {
		timelineFirstPage = "https://coub.com/api/v2/best/" + yearID + "/coubs?type=coubs&page=1"
	} else {
		timelineFirstPage = "https://coub.com/api/v2/best/" + yearID + "/likes?page=1"
	}
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

		var timelinePage string
		if !oldAPI {
			timelinePage = "https://coub.com/api/v2/best/" + yearID + "/coubs?type=coubs&page=" + strconv.Itoa(i)
		} else {
			timelinePage = "https://coub.com/api/v2/best/" + yearID + "/likes?page=" + strconv.Itoa(i)
		}
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
