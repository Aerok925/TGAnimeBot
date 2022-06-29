package modules

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io/ioutil"
	"log"
	"net/http"
)

type AnimeData struct {
	Name      string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func (anime *AnimeData) ConvectAnimeToMsg(id int64) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(id, fmt.Sprintf("Anime: %s\nCharacter: %s\nQuote: %s\n", anime.Name, anime.Character, anime.Quote))
	return msg
}

type Specialbot struct {
	Bot     *tgbotapi.BotAPI
	U       tgbotapi.UpdateConfig
	Updates tgbotapi.UpdatesChannel
	Msg     tgbotapi.MessageConfig
}

func (bot *Specialbot) RandomAnime() ([]AnimeData, error) {
	url := "https://animechan.vercel.app/api/random"
	anime := make([]AnimeData, 1)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Can`t connect to server")
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Can read from body")
		return nil, err
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &anime[0])
	fmt.Println(anime)
	if err != nil {
		fmt.Println(err)
		log.Printf("Can`t pars from body")
		return nil, err
	}
	fmt.Println(anime)
	return anime, nil
}

func (bot *Specialbot) FoundName(name string) ([]AnimeData, error) {
	url := "https://animechan.vercel.app/api/quotes/anime?title=" + name
	var anime []AnimeData
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Can`t connect to server")
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Can read from body")
		return nil, err
	}
	err = json.Unmarshal(body, &anime)
	if err != nil {
		return nil, err
	}
	return anime, nil
}

func (bot *Specialbot) InitBot(token string, debug bool, timeUpdate int) error {
	var err error
	bot.Bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	bot.Bot.Debug = debug

	log.Printf("Authorized on account %s", bot.Bot.Self.UserName)

	bot.U = tgbotapi.NewUpdate(0)
	bot.U.Timeout = timeUpdate

	bot.Updates = bot.Bot.GetUpdatesChan(bot.U)
	return nil
}
