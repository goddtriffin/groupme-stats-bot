package main

import (
	"flag"
	"log"
	"os"

	"github.com/MagnusFrater/groupme"
	groupmestatsbot "github.com/MagnusFrater/groupme-stats-bot"
)

func main() {
	accessToken := flag.String("accessToken", "", "GroupMe API client access token")
	botID := flag.String("botID", "", "GroupMe Bot ID")
	groupID := flag.String("groupID", "", "Group ID")
	flag.Parse()

	if *accessToken == "" || *botID == "" || *groupID == "" {
		flag.Usage()
		os.Exit(1)
	}

	client := groupme.NewClient(groupme.V3BaseURL, *accessToken)
	bot := groupme.NewBot(groupme.V3BaseURL, *botID, *groupID, "", "")

	messages, err := client.AllMessages(bot.GroupID)
	if err != nil {
		log.Fatal(err)
	}

	stats := groupmestatsbot.NewStats(messages)
	stats.Analyze()

	err = bot.Post(stats.SprintTopOfThePops(5))
	if err != nil {
		log.Fatal(err)
	}

	err = bot.Post(stats.SprintTopOfTheSimps(5))
	if err != nil {
		log.Fatal(err)
	}

	err = bot.Post(stats.SprintTopOfTheNarcissists(5))
	if err != nil {
		log.Fatal(err)
	}

	err = bot.Post(stats.SprintTopPoster(5))
	if err != nil {
		log.Fatal(err)
	}

	err = bot.Post(stats.SprintTextFrequencyAnalysis(5))
	if err != nil {
		log.Fatal(err)
	}
}
