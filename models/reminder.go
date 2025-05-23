package models

import "time"

type Reminder struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	GuildID   string    `gorm:"type:varchar(32);index"` // サーバーID
	ChannelID string    `gorm:"type:varchar(32);index"` // チャンネルID
	UserID    string    `gorm:"type:varchar(32);index"` // ユーザーID
	Message   string    `gorm:"type:text;not null"`     // リマインドメッセージ
	RemindAt  time.Time `gorm:"not null;index"`         // 通知時間
	CreatedAt time.Time `gorm:"autoCreateTime"`         // 登録日時
}
