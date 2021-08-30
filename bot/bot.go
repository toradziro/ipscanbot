package main

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"regexp"
)

const (
	accessKey = "819fd023013dee1687749eeab7478640"
	accessBot = "1990350601:AAFVplgLwyOM7ZjPP5Fw1WHQKv9W_ey4eMI"
	endl = "\n"
	pattern = "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
)

func main() {
	var jsonRead JSONfile

	bot, err := tgbotapi.NewBotAPI(accessBot)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	upds, err := bot.GetUpdatesChan(u)

	db := openDb()

	for upd := range upds {
		if upd.Message == nil {
			continue
		}
		log.Printf("[%s]%s", upd.Message.From.UserName, upd.Message)
		//isAdm := isAdmin(db, upd.Message.From.UserName)
		//if isAdm == true {
			// admin menu
		//} else {
			// user menu
		//}
		msg := tgbotapi.NewMessage(upd.Message.Chat.ID, upd.Message.Text)

		ipRequest := string(upd.Message.Text)

		re, _ := regexp.Compile(pattern)
		matched := re.MatchString(ipRequest)

		if matched == true {
			ipInfoRespond, err := getJSON(ipRequest)
			if err != nil {
				log.Printf("Error during reading from API: %v\n", err)
			}

			err = json.Unmarshal([]byte(ipInfoRespond), &jsonRead)
			if err != nil {
				log.Printf("Error during parse JSON: %v\n", err)
			}

			msg.ReplyToMessageID = upd.Message.MessageID
			msg.Text = fillMsg(jsonRead)
			addInDatabase(upd.Message.From.UserName, upd.Message.From.ID, msg.Text, db)
		} else if ipRequest == "/get_last" {
			msg.Text = getLastReq(upd.Message.From.UserName, db)
		} else {
			msg.Text = fillInfo()
		}
		_, err = bot.Send(msg)
		if err != nil {
			log.Printf("Message wasn't recieved")
		}
	}
	db.Close()
}


