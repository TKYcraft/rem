package bot

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var Session *discordgo.Session

func Start() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Discordセッション作成エラー:", err)
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages | discordgo.IntentsMessageContent

	Session = dg
	dg.AddHandler(handleMessage)

	err = dg.Open()
	if err != nil {
		log.Fatal("Discord接続失敗:", err)
	}

	log.Println("Botが起動しました。")
	select {} // 無限待機
}
