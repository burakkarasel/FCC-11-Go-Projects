package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	// You need to put your bot and channel tokens
	os.Setenv("SLACK_BOT_TOKEN", "YOUR_TOKEN_HERE")
	os.Setenv("CHANNEL_ID", "YOUR_TOKEN_HERE")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"YOUR_FILE_NAME"}

	// used for loop so we can upload more than 1 file
	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
