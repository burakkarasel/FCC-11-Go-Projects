package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3712393710581-3700753206471-HLOg4nM7hEKaYsf7wVK5l3zm")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03LUME2DP0-3708730999254-2c0848966abf985183f2f735c26d9477786c62da2960dbfd3d8fbe9acd85a33d")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "age calculator (yob = year of birth)",
		Example:     "my yob is 2022",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)

			if err != nil {
				fmt.Println(err)
			}
			age := 2022 - yob
			r := fmt.Sprintf("You are %d year old", age)

			err = response.Reply(r)

			if err != nil {
				fmt.Println("error while replying", err)
			}
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
	}
}
