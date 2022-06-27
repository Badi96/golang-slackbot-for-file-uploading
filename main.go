package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	//Never share token with someone else!
	os.Setenv("SLACK_BOT_TOKEN", "Replace_with_Bot_token_here") // token from Oauth & Permissions
	os.Setenv("CHANNEL_ID", "Repace_with_channel_id")           // the channel ID in the slack channel of the workspace
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	ChannelArr := []string{os.Getenv("CHANNEL_ID")}
	FileArr := []string{"cat.jpeg"}
	for i := 0; i < len(FileArr); i++ {

		params := slack.FileUploadParameters{
			Channels: ChannelArr,
			File:     FileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
