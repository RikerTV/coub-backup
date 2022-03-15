package main

import (
	"flag"
	"log"
	"time"
)

func main() {
	// Parse the command line flags.
	// They are set if the short or long option is given, e.g.:
	//   `go run main.go -v` or `go run main.go --verbose`
	//   sets the verbose flag.
	directory := flag.String("directory", "./", "select output directory")
	flag.Parse()

	// user string is the first non-flag argument
	user := flag.Arg(0) // get the first non-flag argument

	log.Println("Attempting to backup Coub videos for user:", user)
	log.Println("Output directory:", *directory)

	err := Run(user, *directory)
	if err != nil {
		log.Fatal(err)
	}
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
