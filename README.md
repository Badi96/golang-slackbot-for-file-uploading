# Golang slackbot for file uploading

## Setup Enviroment 
```go mod init github.com/badi96/file-bot```
```go get "github.com/slack-go/slack"```
```go mod tidy```
```go get github.com/slack-go/slack```
```go build```

## Prerequisites
- Download and install Go (https://golang.org/dl/)
- Slack Installed with a slack account and worspace setup that you are the admin to.

## Setting up the slack bot environment 
In main.go, replace the strings ```"Replace_with_Bot_token_here"``` and ```"Replace_with_socket_token_here"``` below with your specific tokens.
```go 
os.Setenv("SLACK_BOT_TOKEN", "Replace_with_Bot_token_here") // Token from Oauth & Permissions
```

```go
os.Setenv("CHANNEL_ID", "Repace_with_channel_id") //Channel ID in the slack channel of the workspace
```

1. To find your token, you need to got to api.slack.com/apps. Press ```create new app``` and choose ```From scratch```, name you app with your worspace you are admin to.
2. Choose ```Socket Mode``` in the left menu and enable socket mode. Give a name for the socket/app token. This tutorial does not require this socket token. 
3. Go to ```Event subscriptions```, enable event subscription. Go to Subscribe to bot events, and add all the event the bot should subscribe to. For this example, add: 
    - ```app_mention```
    - ```channels:history``` 
    - ```channels:read``` 
    - ```chat:write``` 
    - ```chat:write.public``` 
    - ```im:history``` 
    - ```mpim:history``` 
    - ```files:read```
    - ```files:write```
    - ```im:read```
    - ```im:write```
    - ```remote_files:read```
    - ```remote_files:share```
    - ```remote_files:write```


     
4. Finally, press on ```Install to Workspace``` and accept. Take the bot User Oauth Token and replace with ```"Replace_with_socket_token_here"``` above in the code. 
## Commands 
In this project folder, run in terminal:
```go mod init github.com/badi96/slack-age-bot```
```go get "github.com/shomali11/slacker"```
```go mod tidy```
```go get github.com/shomali11/slacker```
```go build```
```go run main.go```

## Runing the bot in Slack
First add the bot in the slack channel. TO do this, go to the slack application and write in the channel ```@file-bot``` to add the bot. Then in this source directory, write in the terminal: ```go run main.go```.
