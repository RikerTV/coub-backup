package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const DownloadInterval = time.Millisecond * 250

// ReadCoub Accepts a Coub struct
// It generates a directory for the coub, creates the info file for it
// And finally downloads all data for it
func ReadCoub(rootdir string, user string) (err error) {

	// rootdir should be the path to the user directory
	// From there we will create our sub directories

	coubs, err := GetNonRecoubs(rootdir, user)
	if err != nil {
		return err
	}
	log.Println("Total Coubs to process: " + strconv.Itoa(len(coubs)))

	for i, coub := range coubs {
		coub.Title = strings.TrimSpace(coub.Title)
		log.Println("Processing Coub: " + coub.Title)
		// Create the directory for the coub
		outdir, err := CreateCoubDir(rootdir, coub)
		if err != nil {
			return err
		}
		if i > 2 {
			break
		}

		// Create the info file for the coub
		err = CreateCoubInfoFiles(outdir, coub)
		if err != nil {
			return err
		}

		// Download all data for the coub
		err = DownloadCoubData(outdir, coub)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetNonRecoubs(dir string, user string) (coubs []Coub, err error) {
	// Open the json file for the user
	jsonFile, err := os.Open(dir + "/" + user + ".json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	var tmpCoubs []Coub

	// Unmarshal the json file into a Coubs struct
	err = json.NewDecoder(jsonFile).Decode(&tmpCoubs)
	//log.Print(len(tmpCoubs))
	for _, coub := range tmpCoubs {
		if coub.Type != "Coub::Recoub" {
			coubs = append(coubs, coub)
		}
	}
	return coubs, nil
}

func CreateCoubDir(dir string, coub Coub) (outdir string, err error) {

	dir = strings.TrimRight(dir, "/")
	CoubYear := coub.CreatedAt.Year()
	CoubMonth := coub.CreatedAt.Month()

	err = CreateDirectory(dir + "/" + strconv.Itoa(CoubYear))
	if err != nil {
		return "", err
	}

	err = CreateDirectory(dir + "/" + strconv.Itoa(CoubYear) + "/" + strconv.Itoa(int(CoubMonth)))
	if err != nil {
		return "", err
	}

	outdir = dir + "/" + strconv.Itoa(CoubYear) + "/" + strconv.Itoa(int(CoubMonth)) + "/" + coub.Title
	err = CreateDirectory(outdir)
	if err != nil {
		return "", err
	}

	return outdir, nil
}

func CreateCoubInfoFiles(dir string, coub Coub) (err error) {
	// First we dump the coub struct into a json file
	outputFile, _ := json.MarshalIndent(coub, "", " ")
	err = ioutil.WriteFile(dir+"/metadata.json", outputFile, 0644)
	if err != nil {
		return err
	}

	infoFile, err := os.Create(dir + "/info.txt")
	if err != nil {
		fmt.Println("Unable to open file: %s", err)
	}

	_, err = infoFile.WriteString("Title: " + coub.Title + "\n")
	if err != nil {
		return err
	}
	_, err = infoFile.WriteString("Created At: " + coub.CreatedAt.String() + "\n")
	if err != nil {
		return err
	}

	_, err = infoFile.WriteString("Duration: " + fmt.Sprintf("%.2f", coub.Duration) + "\n")
	if err != nil {
		return err
	}

	_, err = infoFile.WriteString("Views: " + strconv.Itoa(coub.ViewsCount) + "\n")
	if err != nil {
		return err
	}

	_, err = infoFile.WriteString("Recoubs: " + strconv.Itoa(coub.RecoubsCount) + "\n")
	if err != nil {
		return err
	}

	_, err = infoFile.WriteString("Source: " + fmt.Sprintf("%v", coub.ExternalDownload) + "\n")

	_, err = infoFile.WriteString("Tags: ")

	for i, tag := range coub.Tags {
		if i == len(coub.Tags)-1 {
			_, err = infoFile.WriteString(tag.Title + "\n")
		} else {
			_, err = infoFile.WriteString(tag.Title + ", ")
		}
		if err != nil {
			return err
		}
	}

	err = infoFile.Close()
	if err != nil {
		return err
	}

	return nil
}

func DownloadCoubData(rootdir string, coub Coub) (err error) {

	log.Print("Downloading Frames for Coub: " + coub.Title)
	err = DownloadFirstFrameVersions(rootdir, coub)
	if err != nil {
		return err
	}

	log.Print("Downloading Images for Coub: " + coub.Title)
	err = DownloadImageVersions(rootdir, coub)
	if err != nil {
		return err
	}

	log.Print("Downloading Media Files for Coub: " + coub.Title)
	err = DownloadFileVersions(rootdir, coub)
	if err != nil {
		return err
	}

	return nil
}

func DownloadFileVersions(filepath string, coub Coub) (err error) {

	url := coub.FileVersions.HTML5.Video.Med.URL
	err = DownloadFile(filepath+"/"+FileNameFromURL(url), url)
	if err != nil {
		log.Println("Error downloading Medium Quality HTML5 Video for: " + coub.Title + ": " + err.Error())
	}
	time.Sleep(DownloadInterval)

	url = coub.FileVersions.HTML5.Video.High.URL
	err = DownloadFile(filepath+"/"+FileNameFromURL(url), url)
	if err != nil {
		log.Println("Error downloading High Quality HTML5 Video for: " + coub.Title + ": " + err.Error())
	}
	time.Sleep(DownloadInterval)

	url = coub.FileVersions.HTML5.Video.Higher.URL
	err = DownloadFile(filepath+"/"+FileNameFromURL(url), url)
	if err != nil {
		log.Println("Error downloading Higher Quality HTML5 Video for: " + coub.Title + ": " + err.Error())
	}
	time.Sleep(DownloadInterval)

	url = coub.FileVersions.HTML5.Audio.High.URL
	err = DownloadFile(filepath+"/"+FileNameFromURL(url), url)
	if err != nil {
		log.Println("Error downloading Higher Quality HTML5 Audio for: " + coub.Title + ": " + err.Error())
	}
	time.Sleep(DownloadInterval)

	url = coub.FileVersions.HTML5.Audio.Med.URL
	err = DownloadFile(filepath+"/"+FileNameFromURL(url), url)
	if err != nil {
		log.Println("Error downloading Medium Quality HTML5 Audio for: " + coub.Title + ": " + err.Error())
	}
	time.Sleep(DownloadInterval)

	// We do not download mobile versions, because they are the same as the medium quality HTML5 versions
	/*
		url = coub.FileVersions.Mobile.Video
		err = DownloadFile(filepath+"/"+FileNameFromURL(url), url)
		if err != nil {
			log.Println("Error downloading Mobile Video for" + coub.Title + ": " + err.Error())
		}
		time.Sleep(DownloadInterval)

		url = coub.FileVersions.Mobile.Audio[0]
		err = DownloadFile(filepath+"/"+FileNameFromURL(url), url)
		if err != nil {
			log.Println("Error downloading Mobile Audio for" + coub.Title + ": " + err.Error())
		}
		time.Sleep(DownloadInterval)
	*/

	url = coub.FileVersions.Share.Default
	err = DownloadFile(filepath+"/"+FileNameFromURL(url), url)
	if err != nil {
		log.Println("Error downloading Default Share File for: " + coub.Title + ": " + err.Error())
	}

	url = coub.FileVersions.Share.Default
	err = DownloadFile(filepath+"/"+coub.Title+".mp4", url)
	if err != nil {
		log.Println("Error downloading (renamed) Default Share File for: " + coub.Title + ": " + err.Error())
	}

	return nil
}

func DownloadImageVersions(filepath string, coub Coub) (err error) {
	template := coub.ImageVersions.Template
	for _, version := range coub.ImageVersions.Versions {
		url := strings.Replace(template, "%{version}", version, -1)
		err = DownloadFile(filepath+"/"+FileNameFromURL(url), url)
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func DownloadFirstFrameVersions(filepath string, coub Coub) (err error) {
	template := coub.FirstFrameVersions.Template
	for _, version := range coub.FirstFrameVersions.Versions {
		url := strings.Replace(template, "%{version}", version, -1)
		err = DownloadFile(filepath+"/"+FileNameFromURL(url), url)
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}
