# Overview

Holiday telegram bot can handle one command and 4 country option. You can use it to get today's holiday in selected country

## Usage

Clone the repository and run the program using a Go compiler:

```
git clone https://github.com/c1kzy/Telegram-Holiday-Bot.git
cd telegram-holiday-bot
go run main.go 
```

Click the link "**https://t.me/FMWeatherBot**" and run the bot

### Deploy a server

You can deploy your server any way you want, but I find it really quick and easy to use ngrok. Ngrok allows you to expose applications running on your local machine to the public internet.
### How to install it?
You can download it from the website directly
```
https://ngrok.com/download
```
or install ngrok via Chocolatey

```
$ choco install ngrok
```

### Running a server using ngrok

Once you install ngrok, you can run this command on another terminal on your system:

```
ngrok http <port>

```
Example:
```
ngrok http port 8080
```

[Server example]()

Here, https://ed40-178-150-143-55.ngrok.io is the public temporary URL for the server running on port 8080 of my machine.

Now, all we need to do is let telegram know that our bot has to talk to this url whenever it receives any message. We do this through the telegram API. Enter this in your terminal :
```
curl -F "url=https://ed40-178-150-143-55.ngrok.io"  https://api.telegram.org/bot<your_api_token>/setWebhook

```
