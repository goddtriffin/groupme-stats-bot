package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/MagnusFrater/groupme"
	groupmestatsbot "github.com/MagnusFrater/groupme-stats-bot"
)

func main() {
	accessToken := flag.String("accessToken", "", "GroupMe API client access token")
	botID := flag.String("botID", "", "GroupMe Bot ID")
	groupID := flag.String("groupID", "", "GroupMe Group ID")
	limit := flag.Int("limit", 5, "number of items to list")

	blacklist := flag.String("blacklist", "", "blacklist of comma-delimited User IDs")

	topOfThePops := flag.Bool("topOfThePops", false, "toggle for TopOfThePops")
	topOfTheSimps := flag.Bool("topOfTheSimps", false, "toggle for TopOfTheSimps")
	topOfTheNarcissists := flag.Bool("topOfTheNarcissists", false, "toggle for TopOfTheNarcissists")
	topPoster := flag.Bool("topPoster", false, "toggle for TopPoster")
	mostCharismatic := flag.Bool("mostCharismatic", false, "toggle for MostCharismatic")
	topLurker := flag.Bool("topLurker", false, "toggle for TopLurker")
	topRambler := flag.Bool("topRambler", false, "toggle for TopRambler")
	textFrequencyAnalysis := flag.Bool("textFrequencyAnalysis", false, "toggle for TextFrequencyAnalysis")
	topMessages := flag.Bool("topMessages", false, "toggle for TopMessages")
	topReposts := flag.Bool("topReposts", false, "toggle for TopReposts")

	flag.Parse()

	if *accessToken == "" || *botID == "" || *groupID == "" {
		flag.Usage()
		os.Exit(1)
	}

	if !*topOfThePops && !*topOfTheSimps && !*topOfTheNarcissists && !*topPoster &&
		!*mostCharismatic && !*topLurker && !*topRambler && !*textFrequencyAnalysis &&
		!*topMessages && !*topReposts {
		fmt.Print("Must toggle at least one of: ")
		fmt.Print("topOfThePops, topOfTheSimps, topOfTheNarcissists, topPoster, ")
		fmt.Print("mostCharismatic, topLurker, topRambler, textFrequencyAnalysis, ")
		fmt.Println("topMessages, topReposts")

		flag.Usage()
		os.Exit(1)
	}

	client := groupme.NewClient(groupme.V3BaseURL, *accessToken)
	bot := groupme.NewBot(groupme.V3BaseURL, *botID, *groupID, "", "")

	messages, err := client.AllMessages(bot.GroupID)
	if err != nil {
		log.Panic(err)
	}

	stats := groupmestatsbot.NewStats(messages)

	// blacklist Bot User ID if it exists
	if *blacklist != "" {
		for _, userID := range strings.Split(*blacklist, ",") {
			stats.Blacklist(userID)
		}
	}

	stats.Analyze()

	if *topOfThePops {
		err = bot.Post(stats.SprintTopOfThePops(*limit), nil)
		if err != nil {
			log.Panic(err)
		}
	}

	if *topOfTheSimps {
		err = bot.Post(stats.SprintTopOfTheSimps(*limit), nil)
		if err != nil {
			log.Panic(err)
		}
	}

	if *topOfTheNarcissists {
		err = bot.Post(stats.SprintTopOfTheNarcissists(*limit), nil)
		if err != nil {
			log.Panic(err)
		}
	}

	if *topPoster {
		err = bot.Post(stats.SprintTopPoster(*limit), nil)
		if err != nil {
			log.Panic(err)
		}
	}

	if *mostCharismatic {
		err = bot.Post(stats.SprintMostCharismatic(*limit), nil)
		if err != nil {
			log.Panic(err)
		}
	}

	if *topLurker {
		err = bot.Post(stats.SprintTopLurker(*limit), nil)
		if err != nil {
			log.Panic(err)
		}
	}

	if *topRambler {
		err = bot.Post(stats.SprintTopRambler(*limit), nil)
		if err != nil {
			log.Panic(err)
		}
	}

	if *textFrequencyAnalysis {
		err = bot.Post(stats.SprintTextFrequencyAnalysis(*limit), nil)
		if err != nil {
			log.Panic(err)
		}
	}

	if *topMessages {
		err = bot.Post(stats.SprintTopMessages(*limit), nil)
		if err != nil {
			log.Panic(err)
		}
	}

	if *topReposts {
		err = bot.Post(stats.SprintTopReposts(*limit), nil)
		if err != nil {
			log.Panic(err)
		}
	}
}
