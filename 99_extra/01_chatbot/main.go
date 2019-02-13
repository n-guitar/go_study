package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
)

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("Hello Event")

			case *slack.MessageEvent:
				log.Printf("Message: %v\n", ev.Msg)
				rtm.SendMessage(rtm.NewOutgoingMessage("こんさんまじかよ・・・", ev.Channel))
				// var stev string
				// stev = ev.string[]
				// if strings.Contains(stev, "こん") {
				// 	rtm.SendMessage(rtm.NewOutgoingMessage("こんさんまじかよ・・・", ev.Channel))
				// } else {
				// 	rtm.SendMessage(rtm.NewOutgoingMessage("何を言っているのかわかりません。", ev.Channel))
				// }

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1

			}
		}
	}
}

func main() {
	api := slack.New("tokenをセット")
	os.Exit(run(api))
}
