package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

// printCommandEvents prints information about each command event received on the analytics channel
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	// Set Slack bot and app tokens as environment variables
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6196336599091-6219482593120-XH1Q3N0aj0O9JO8MHYqSsCGv")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A065PF5SQMB-6219433482976-a2f68aa7625bfd64a719b4a90f9751f9ef3eeee179aa6dbbdea377c9e0489cce")

	// Create a new Slack bot client
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// Start a goroutine to print command events
	go printCommandEvents(bot.CommandEvents())

	// Define a command for the bot
	bot.Command("My yob is <year>", &slacker.CommandDefinition{
		Description: "Birthyear calculator",
		Examples:    []string{"My yob is 2020"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			// Extract the year parameter from the command
			year := request.Param("year")

			// Convert the year to an integer
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}

			// Calculate the age
			age := 2023 - yob

			// Prepare the response message
			r := fmt.Sprintf("Your age is %d", age)

			// Reply to the user
			response.Reply(r)
		},
	})

	// Create a context with cancellation for managing the lifecycle of the bot
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start listening for events
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
