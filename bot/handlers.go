package bot

import (
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Println("なんか来た:", m.Content)

	if m.Author.Bot {
		return
	}

	content := strings.TrimSpace(m.Content)

	switch {
	case strings.HasPrefix(content, "!reminite "):
		args := strings.SplitN(content, " ", 3)
		if len(args) < 3 {
			s.ChannelMessageSend(m.ChannelID, "ex: `!reminite <分> <メッセージ>`")
			return
		}

		minutes, err := time.ParseDuration(args[1] + "m")
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "ちゃんと指定するわよ")
			return
		}
		msg := args[2]
		remindAt := time.Now().Add(minutes)

		CreateReminder(m.GuildID, m.ChannelID, m.Author.ID, msg, remindAt)
		s.ChannelMessageSend(m.ChannelID, "覚えたわよ！ "+remindAt.Format("2006-01-02_15:04"))

	case strings.HasPrefix(content, "!rem "):
		args := strings.SplitN(content, " ", 3)
		if len(args) < 3 {
			s.ChannelMessageSend(m.ChannelID, "使い方: `!rem <YYYY-MM-DD_HH:MM> <メッセージ>`")
			return
		}

		layout := "2006-01-02_15:04"
		remindAt, err := time.ParseInLocation(layout, args[1], time.Local)
		if err != nil || remindAt.Before(time.Now()) {
			s.ChannelMessageSend(m.ChannelID, "ちゃんと指定するわよ！ ex: `2025-06-01_14:30`")
			return
		}

		msg := args[2]
		CreateReminder(m.GuildID, m.ChannelID, m.Author.ID, msg, remindAt)
		s.ChannelMessageSend(m.ChannelID, "覚えたわよ！ "+remindAt.Format("2006-01-02_15:04"))
	}
}
