package main

import (
	. "TGAnimeBot/config"
	"TGAnimeBot/modules"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	var bot modules.Specialbot

	err := bot.InitBot(GetToken(), true, 60)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Authorized on account %s", bot.Bot.Self.UserName)
	for update := range bot.Updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				bot.Msg = tgbotapi.NewMessage(update.Message.From.ID, "Привет, здесь ты можешь получить как случайное аниме, так и произвести поиск аниме по названию!")
			case "random":
				anime, err := bot.RandomAnime()
				if err != nil {
					bot.Msg = tgbotapi.NewMessage(update.Message.From.ID, "Привет, здесь ты можешь получить как случайное аниме, так и произвести поиск аниме по названию!")
				} else {
					bot.Msg = anime[0].ConvectAnimeToMsg(update.Message.From.ID)
				}
			default:
				anime, err := bot.FoundName(update.Message.Text)
				if err != nil {
					bot.Msg = tgbotapi.NewMessage(update.Message.From.ID, "Не удалось найти аниме по вашему запросу")
				} else {
					bot.Msg = anime[0].ConvectAnimeToMsg(update.Message.From.ID)
				}
			}
			bot.Bot.Send(bot.Msg)
			bot.Msg.Text = ""
		}
	}
}
