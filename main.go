package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// Parse the command line flags.
	// They are set if the short or long option is given, e.g.:
	//   `go run main.go -v` or `go run main.go --verbose`
	//   sets the verbose flag.
	directory := flag.String("directory", "./backups/", "select output directory")
	bestof := flag.String("bestof", "", "select best of to download")
	community := flag.String("community", "", "select community to download")
	pages := flag.Int("pages", 10, "select number of COMMUNITY pages to download")
	featured := flag.Bool("featured", false, "select featured to download")
	day := flag.Bool("day", false, "select day to download")
	flag.Parse()

	if *bestof != "" {
		year, err := strconv.Atoi(*bestof)
		if err != nil {
			log.Fatal("Best of must be a number")
		}

		// if the year is less than 2012 then quit
		if year < 2012 {
			log.Fatal("Year must be greater than 2012")
		}

		log.Println("Attempting to backup Coub videos for best of year: ", year)
		log.Println("Output directory:", *directory)
		// append a / to the end of the directory if it doesn't exist
		if (*directory)[len(*directory)-1] != '/' {
			*directory += "/"
		}

		err = RunBestOf(*directory, strconv.Itoa(year))
		if err != nil {
			log.Fatal(err)
		}
	} else if *community != "" {
		log.Println("Attempting to backup Coub videos for community: ", *community)
		log.Println("Note: There is a maximum of 500 pages that can be downloaded, of 10 coubs per page.")
		log.Println("Output directory:", *directory)

		if *pages > 500 {
			*pages = 500
		} else if *pages < 1 {
			*pages = 1
		}

		err := RunCommunity(*directory, *community, *pages)
		if err != nil {
			log.Fatal(err)
		}
	} else if *featured || *day {
		log.Println("Attempting to backup Coub videos for featured or day")
		log.Println("Output directory:", *directory)

		if *featured {
			log.Println("Attempting to backup Coub videos for featured")
			err := RunFeatured(*directory)
			if err != nil {
				log.Fatal(err)
			}
		} else if *day {
			log.Println("Attempting to backup Coub videos for day")
			err := RunDay(*directory)
			if err != nil {
				log.Fatal(err)
			}
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

func RunCommunity(dir string, community string, pages int) (err error) {
	switch community {
	case "animals-pets":
		break
	case "mashup":
		break
	case "music":
		break
	case "blogging":
		break
	case "standup-jokes":
		break
	case "movies":
		break
	case "anime":
		break
	case "gaming":
		break
	case "cartoons":
		break
	case "art":
		break
	case "live-pictures":
		break
	case "news":
		break
	case "sports":
		break
	case "science-technology":
		break
	case "food-kitchen":
		break
	case "celebrity":
		break
	case "nature-travel":
		break
	case "fashion":
		break
	case "dance":
		break
	case "cars":
		break
	case "memes":
		break
	case "nsfw":
		break
	default:
		log.Println("Unrecognized community, please chose from one of the following:")
		log.Println("animals-pets, mashup, music, blogging, standup-jokes, movies, anime, gaming, cartoons, art, live-pictures, news, sports, science-technology, food-kitchen, celebrity, nature-travel, fashion, dance, cars, memes, nsfw")
	}

	log.Println("Creating directory if not exists:", dir)
	err = DirectorySetup(community, dir)
	if err != nil {
		return err
	}

	outputDir := dir + community + "/"

	// if file doesn't exist
	if !FileExists(outputDir + community + ".json") {
		err = GenerateCommunityInfoFile(outputDir, community, pages)
		if err != nil {
			return err
		}
	} else {
		log.Println("Coub info file already exists for community: ", community, ", skipping generation.")
		log.Println("Proceeding to downloads for ", community)
	}

	log.Println("Beginning downloads in 5 seconds...")
	time.Sleep(5 * time.Second)

	err = ReadCoub(outputDir, community)
	if err != nil {
		return err
	}
	return nil
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
		log.Println("Coub info file already exists for bestof ", year, ", skipping generation.")
		log.Println("Proceeding to downloads for bestof ", year)
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

func RunFeatured(dir string) (err error) {
	log.Println("Creating directory if not exists:", dir+"featured")
	err = DirectorySetup("featured", dir)
	if err != nil {
		return err
	}

	outputDir := dir + "featured" + "/"

	// if file exists, delete it
	if FileExists(outputDir + "featured.json") {
		err = os.Remove(outputDir + "featured.json")
		if err != nil {
			return err
		}
	}
	err = GenerateFeaturedInfoFile(outputDir)
	if err != nil {
		return err
	}

	log.Println("Beginning downloads in 5 seconds...")
	time.Sleep(5 * time.Second)

	err = ReadCoub(outputDir, "featured")
	if err != nil {
		return err
	}

	return nil
}

func RunDay(dir string) (err error) {
	log.Println("Creating directory if not exists:", dir+"coub-of-the-day")
	err = DirectorySetup("coub-of-the-day", dir)
	if err != nil {
		return err
	}

	outputDir := dir + "coub-of-the-day" + "/"

	// if file exists, delete it
	if FileExists(outputDir + "coub-of-the-day.json") {
		err = os.Remove(outputDir + "coub-of-the-day.json")
		if err != nil {
			return err
		}
	}
	err = GenerateCoubOfDayInfoFile(outputDir)
	if err != nil {
		return err
	}

	log.Println("Beginning downloads in 5 seconds...")
	time.Sleep(5 * time.Second)

	err = ReadCoub(outputDir, "coub-of-the-day")
	if err != nil {
		return err
	}

	return nil
}
