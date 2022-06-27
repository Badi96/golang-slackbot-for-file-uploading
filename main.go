package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	//Never share token with someone else!
	os.Setenv("SLACK_BOT_TOKEN", "REPLACE_WITH_SLACK_BOT_TOKEN") // token from Oauth & Permissions
	os.Setenv("CHANNEL_ID", "REPLACE_WITH_CANNEL_ID")            // the channel ID in the slack channel of the workspace
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
