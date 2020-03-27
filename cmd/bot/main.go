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

	topOfThePops := flag.Bool("topOfThePops", false, "toggle for TopOfThePops")
	topOfTheSimps := flag.Bool("topOfTheSimps", false, "toggle for TopOfTheSimps")
	topOfTheNarcissists := flag.Bool("topOfTheNarcissists", false, "toggle for TopOfTheNarcissists")
	topPoster := flag.Bool("topPoster", false, "toggle for TopPoster")
	textFrequencyAnalysis := flag.Bool("textFrequencyAnalysis", false, "toggle for TextFrequencyAnalysis")
	topMessages := flag.Bool("topMessages", false, "toggle for TopMessages")

	flag.Parse()

	if *accessToken == "" || *botID == "" || *groupID == "" {
		flag.Usage()
		os.Exit(1)
	}

	if !*topOfThePops && !*topOfTheSimps && !*topOfTheNarcissists && !*topPoster && !*textFrequencyAnalysis && !*topMessages {
		fmt.Print("Must toggle at least one of: ")
		fmt.Println("topOfThePops, topOfTheSimps, topOfTheNarcissists, topPoster, textFrequencyAnalysis, topMessages")
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

	if *textFrequencyAnalysis {
		err = bot.Post(stats.SprintTextFrequencyAnalysis(*limit), nil)
		if err != nil {
			log.Panic(err)
		}
	}

	if *topMessages {
		for _, part := range botBufferedMessage(stats.SprintTopMessages(*limit), "\n\n") {
			err = bot.Post(part, nil)
			if err != nil {
				log.Panic(err)
			}
		}
	}
}

// botBufferedMessage returns a list of strings no bigger than
// what is allowed to be sent as a GroupMe Bot (length: 1000)
func botBufferedMessage(s, sep string) []string {
	list := []string{}
	var strBuilder string

	split := strings.Split(s, sep)
	for _, part := range split {
		if len(strBuilder)+len(part)+len(sep) <= 1000 {
			strBuilder += part + sep
		} else {
			list = append(list, strings.TrimSpace(strBuilder))
			strBuilder = part + sep
		}
	}

	if len(strBuilder) > 0 {
		list = append(list, strings.TrimSpace(strBuilder))
	}

	return list
}
