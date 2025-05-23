package bot

import (
	"rem/db"
	"rem/models"
	"time"
)

func CreateReminder(guildID, channelID, userID, message string, remindAt time.Time) {
	r := models.Reminder{
		GuildID:   guildID,
		ChannelID: channelID,
		UserID:    userID,
		Message:   message,
		RemindAt:  remindAt,
	}
	db.DB.Create(&r)

	// ゴルーチンで通知予約
	go func(rem models.Reminder) {
		sleepDuration := time.Until(rem.RemindAt)
		if sleepDuration > 0 {
			time.Sleep(sleepDuration)
		}
		sendReminder(rem)
		db.DB.Delete(&rem)
	}(r)
}

func sendReminder(r models.Reminder) {
	if Session != nil {
		Session.ChannelMessageSend(r.ChannelID, "<@"+r.UserID+"> お知らせわよ！: "+r.Message)
	}
}
