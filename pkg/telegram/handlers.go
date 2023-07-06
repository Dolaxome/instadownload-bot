package telegram

import (
	"database/sql"
	"fmt"
	"github.com/Dolaxome/instadownload-bot/pkg/storage"
	"github.com/Dolaxome/instadownload-bot/pkg/storydw"
	"github.com/Dolaxome/instadownload-bot/pkg/storylink"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
	"log"
	"os"
	"path"
	"strings"
)

const (
	commandStart    = "start"
	commandDownload = "download"

	replyDW    = "YES"
	replyStart = "Hi! Before we get started you need to send me a cookie values to access the Instagram story API\nTo pass values, write (value name): (value) (without brackets)\nRead more: https://telegra.ph/How-to-get-Cookie-value-11-22"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandDownload:
		return b.handleDWCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	//db open
	connStr := "user=postgres password=admin dbname=postgres sslmode=disable"
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	if strings.HasPrefix(msg.Text, "ds_user_id:") {
		msg1 := tgbotapi.NewMessage(message.Chat.ID, "ds_user_id saved!")

		//storylink.SetUserId(strings.TrimPrefix(message.Text, "ds_user_id:"))
		err := storage.DbAddDsID(sqlDB, message.Chat.ID, strings.TrimPrefix(message.Text, "ds_user_id:"))
		if err != nil {
			return err
		}

		_, err = b.bot.Send(msg1)
		return err
	}
	if strings.HasPrefix(msg.Text, "sessionid:") {
		msg1 := tgbotapi.NewMessage(message.Chat.ID, "sessionid saved!")

		//storylink.SetSessionId(strings.TrimPrefix(message.Text, "sessionid:"))
		err := storage.DbAddSession(sqlDB, message.Chat.ID, strings.TrimPrefix(message.Text, "sessionid:"))
		if err != nil {
			return err
		}

		_, err = b.bot.Send(msg1)
		return err
	}
	if strings.HasPrefix(msg.Text, "csrftoken:") {
		msg1 := tgbotapi.NewMessage(message.Chat.ID, "csrftoken saved!")

		//storylink.SetCsrfToken(strings.TrimPrefix(message.Text, "csrftoken:"))
		err := storage.DbAddCsrf(sqlDB, message.Chat.ID, strings.TrimPrefix(message.Text, "csrftoken:"))
		if err != nil {
			return err
		}

		_, err = b.bot.Send(msg1)
		return err
	}

	if msg.Text == "get ds" {
		dsuserid, err := storage.GetDS(sqlDB, message.Chat.ID)

		fmt.Println(dsuserid)
		msg1 := tgbotapi.NewMessage(message.Chat.ID, dsuserid)
		//_, err = b.bot.Send(msg2)
		_, err = b.bot.Send(msg1)
		return err
	}
	if msg.Text == "get sess" {
		sess, err := storage.GetSessID(sqlDB, message.Chat.ID)

		fmt.Println(sess)
		msg1 := tgbotapi.NewMessage(message.Chat.ID, sess)
		_, err = b.bot.Send(msg1)
		return err
	}
	if msg.Text == "get csrf" {
		csrf, err := storage.GetCSRF(sqlDB, message.Chat.ID)

		fmt.Println(csrf)
		msg1 := tgbotapi.NewMessage(message.Chat.ID, csrf)
		_, err = b.bot.Send(msg1)
		return err
	}
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, replyStart)

	//db open
	connStr := "user=postgres password=admin dbname=postgres sslmode=disable"
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	err = storage.InitUserRow(sqlDB, message.Chat.ID)

	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleDWCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, replyDW)

	//setup
	connStr := "user=postgres password=admin dbname=postgres sslmode=disable"
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer sqlDB.Close()
	var ds string
	var sess string
	var csrf string
	ds, err = storage.GetDS(sqlDB, message.Chat.ID)
	if err != nil {
		return err
	}
	if ds == "ds_user_id is missing" {
		return err
	}
	sess, err = storage.GetSessID(sqlDB, message.Chat.ID)
	if err != nil {
		return err
	}
	if ds == "sessionid is missing" {
		return err
	}
	csrf, err = storage.GetCSRF(sqlDB, message.Chat.ID)
	if err != nil {
		return err
	}
	if ds == "csrftoken is missing" {
		return err
	}
	storylink.SetDS(ds)
	storylink.SetSESS(sess)
	storylink.SetCSRF(csrf)

	//downloading
	err = storydw.DownloadAll()
	if err != nil {
		return err
	}
	//defer os.RemoveAll("./stories")
	items, _ := os.ReadDir("./stories")
	for _, item := range items {
		subitems, _ := os.ReadDir("./stories/" + item.Name())
		for _, subitem := range subitems {
			fileext := path.Ext(subitem.Name())
			if fileext == ".jpg" {
				msg1 := tgbotapi.NewPhotoUpload(message.Chat.ID, "./stories/"+item.Name()+"/"+subitem.Name())
				_, err = b.bot.Send(msg1)
				if err != nil {
					return err
				}
			}
			if fileext == ".mp4" {
				msg1 := tgbotapi.NewVideoUpload(message.Chat.ID, "./stories/"+item.Name()+"/"+subitem.Name())
				_, err = b.bot.Send(msg1)
				if err != nil {
					return err
				}
			}
		}
	}
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "unknown command")

	_, err := b.bot.Send(msg)
	return err
}
