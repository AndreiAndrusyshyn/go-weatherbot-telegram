package main

import (
	"awesomeProject1/Picture_generate"
	"awesomeProject1/weather_api"
	"fmt"
	"gopkg.in/syfaro/telegram-bot-api.v4"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	var chanel tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	chanel.Timeout = 60

	upd, _ := bot.GetUpdatesChan(chanel)

	for {
		select {
		case update := <-upd:

			ChatID := update.Message.Chat.ID

			Text := update.Message.Text

			parsed_weather := weather_api.Get_weather(Text)

			temp := string(fmt.Sprintf("%.1f", parsed_weather.Main.Temp-273.15))
			Picture_generate.Text = []string{temp + "°C", Text}
			Picture_generate.Picture_generatd()

			msg := tgbotapi.NewPhotoUpload(ChatID, "/root/go/src/awesomeProject1/result.jpg")

			bot.Send(msg)

		}
	}

}
