package shared

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const notifPropertyKey = "notif_property"

// NotifProperty merepresentasikan satu record pada tabel notif_property.
type NotifProperty struct {
	// ── identitas ────────────────────────────────────────
	ID int `gorm:"column:id;primaryKey" json:"id"`

	// ── email / SMTP ──────────────────────────────────────
	MailProtocol *string `gorm:"column:mail_protocol" json:"mail_protocol"`
	SmtpHost     *string `gorm:"column:smtp_host"     json:"smtp_host"`
	SmtpUser     *string `gorm:"column:smtp_user"     json:"smtp_user"`
	SmtpPass     *string `gorm:"column:smtp_pass"     json:"smtp_pass"`
	SmtpPort     *string `gorm:"column:smtp_port"     json:"smtp_port"`

	// ── SMS ───────────────────────────────────────────────
	SmsApiKey    *string `gorm:"column:sms_api_key"    json:"sms_api_key"`
	SmsUrlServer *string `gorm:"column:sms_url_server" json:"sms_url_server"`

	// ── WhatsApp ──────────────────────────────────────────
	WaApiKey    *string `gorm:"column:wa_api_key"    json:"wa_api_key"`
	WaUrlServer *string `gorm:"column:wa_url_server" json:"wa_url_server"`

	// ── Mail Chimp ────────────────────────────────────────
	McApiKey    *string `gorm:"column:mc_api_key"    json:"mc_api_key"`
	McUrlServer *string `gorm:"column:mc_url_server" json:"mc_url_server"`

	// ── Push Notification ─────────────────────────────────
	PushAppId  *string `gorm:"column:push_app_id" json:"push_app_id"`
	PushApiKey *string `gorm:"column:push_apikey"  json:"push_apikey"`

	// ── timestamp ─────────────────────────────────────────
	Created *time.Time `gorm:"column:created" json:"created_at"`
	Updated *time.Time `gorm:"column:updated" json:"updated_at"`
	Deleted *time.Time `gorm:"column:deleted" json:"deleted_at"`
}

// TableName sets the name of the database table.
func (NotifProperty) TableName() string {
	return "notif_property"
}

// SetToRedis menyimpan data notif_property dari DB ke Redis.
func (n *NotifProperty) SetToRedis(db *gorm.DB, rdb *redis.Client) error {
	if err := db.First(n).Error; err != nil {
		return err
	}

	data, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return rdb.Set(context.Background(), notifPropertyKey, data, 0).Err()
}

// GetFromRedis mengambil data notif_property dari Redis.
func (n *NotifProperty) GetFromRedis(rdb *redis.Client) error {
	val, err := rdb.Get(context.Background(), notifPropertyKey).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), n)
}
