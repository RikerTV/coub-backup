package main

import (
	"flag"
	"log"
	"strconv"
	"time"
)

func main() {
	// Parse the command line flags.
	// They are set if the short or long option is given, e.g.:
	//   `go run main.go -v` or `go run main.go --verbose`
	//   sets the verbose flag.
	directory := flag.String("directory", "./backups/", "select output directory")
	bestof := flag.Bool("bestof", false, "select best of to download")
	year := flag.Int("year", 2022, "select year to download")
	flag.Parse()

	if *bestof {
		// if the year is less than 2012 then quit
		if *year < 2012 {
			log.Fatal("Year must be greater than 2012")
		}

		log.Println("Attempting to backup Coub videos for best of year: ", *year)
		log.Println("Output directory:", *directory)
		// append a / to the end of the directory if it doesn't exist
		if (*directory)[len(*directory)-1] != '/' {
			*directory += "/"
		}

		err := RunBestOf(*directory, strconv.Itoa(*year))
		if err != nil {
			log.Fatal(err)
		}
	} else {

		// user string is the first non-flag argument
		user := flag.Arg(0) // get the first non-flag argument

		log.Println("Attempting to backup Coub videos for user:", user)
		log.Println("Output directory:", *directory)

		err := Run(user, *directory)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func RunBestOf(dir string, year string) (err error) {
	log.Println("Creating directory if not exists:", dir)
	err = DirectorySetup("bestof-"+year, dir)
	if err != nil {
		return err
	}

	outputDir := dir + "bestof-" + year + "/"

	// if file doesn't exist
	if !FileExists(outputDir + "bestof.json") {
		err = GenerateBestOfInfoFile(outputDir, year)
		if err != nil {
			return err
		}
	} else {
		log.Println("Coub info file already exists for user:", "bestof", ", skipping generation.")
		log.Println("Proceeding to downloads for ", "bestof")
	}

	log.Println("Beginning downloads in 5 seconds...")
	time.Sleep(5 * time.Second)

	err = ReadCoub(outputDir, "bestof")
	if err != nil {
		return err
	}
	return nil
}

func Run(user string, dir string) (err error) {
	log.Println("Creating directory if not exists:", dir+user)
	err = DirectorySetup(user, dir)
	if err != nil {
		return err
	}

	outputDir := dir + user + "/"

	// if file doesn't exist
	if !FileExists(outputDir + user + ".json") {
		err = GenerateInfoFile(outputDir, user)
		if err != nil {
			return err
		}
	} else {
		log.Println("Coub info file already exists for user:", user, ", skipping generation.")
		log.Println("Proceeding to downloads for ", user)
	}

	log.Println("Beginning downloads in 5 seconds...")
	time.Sleep(5 * time.Second)

	err = ReadCoub(outputDir, user)
	if err != nil {
		return err
	}

	return nil
}
