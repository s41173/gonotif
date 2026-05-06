package service

import (
	"bytes"
	"fmt"
	"go-notif/config"
	"go-notif/internal/api/repository"
	"go-notif/internal/shared"
	"go-notif/utils"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	notifmodel "go-notif/internal/api/model"

	"go.uber.org/zap"
	// "gorm.io/gorm/utils"
)

func Send_wa(number string, message string) bool {

	var notifproperty shared.NotifProperty
	if err := notifproperty.GetFromRedis(config.RDB); err != nil {
		config.Log.Error("Error get notif property", zap.Error(err))
		return false
	}

	if notifproperty.WaUrlServer == nil || notifproperty.WaApiKey == nil {
		config.Log.Error("WA config is not set",
			zap.String("wa_url_server", fmt.Sprintf("%v", notifproperty.WaUrlServer)),
			zap.String("wa_api_key", fmt.Sprintf("%v", notifproperty.WaApiKey)),
		)
		return false
	}

	apiKey := utils.ConvertToString(notifproperty.WaApiKey)
	urlServer := utils.ConvertToString(notifproperty.WaUrlServer)

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// sama seperti POSTFIELDS di PHP
	_ = writer.WriteField("target", number)
	_ = writer.WriteField("message", message)
	_ = writer.WriteField("countryCode", "62")

	writer.Close()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("POST", urlServer, &b)
	if err != nil {
		fmt.Println("error request:", err)
		return false
	}

	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error send:", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		fmt.Println("failed status:", resp.Status)
		return false
	}

	return true
}

func Add_notif(custID int, custName string, sentTo string, subject string, content string, notifType int, modul string, campaign int) (int, error) {

	// ── validasi required ─────────────────────────────────
	if custID <= 0 {
		return 400, fmt.Errorf("custID wajib diisi")
	}
	if strings.TrimSpace(custName) == "" {
		return 400, fmt.Errorf("custName wajib diisi")
	}
	if strings.TrimSpace(sentTo) == "" {
		return 400, fmt.Errorf("sentTo wajib diisi")
	}
	if strings.TrimSpace(subject) == "" {
		return 400, fmt.Errorf("subject wajib diisi")
	}
	if strings.TrimSpace(content) == "" {
		return 400, fmt.Errorf("content wajib diisi")
	}
	if notifType <= 0 {
		return 400, fmt.Errorf("notifType wajib diisi")
	}
	if strings.TrimSpace(modul) == "" {
		return 400, fmt.Errorf("modul wajib diisi")
	}
	if campaign < 0 {
		return 400, fmt.Errorf("campaign tidak valid")
	}

	// ── simpan ke DB ──────────────────────────────────────
	joined := time.Now()
	newData := notifmodel.Notif{
		Customer: custID,
		CustName: &custName,
		SentTo:   sentTo,
		Subject:  subject,
		Content:  content,
		Type:     int16(notifType),
		Modul:    modul,
		Campaign: int16(campaign),
		Created:  &joined,
	}

	if err := repository.Save(&newData); err != nil {
		config.Log.Error("Error Add notif", zap.Error(err))
		return 500, fmt.Errorf("Error Add Notif")
	}

	if notifType == 7 {
		stts := Send_wa(sentTo, subject)
		if stts != true {
			return 500, fmt.Errorf("Error Sent Notif")
		}
	}

	// Get latest ID
	latestID, err := repository.GetLatestID()
	if err != nil {
		config.Log.Error("Error Get Latest", zap.Error(err))
		return 500, fmt.Errorf("Error Latest")
	}

	err = repository.Set_status(latestID)
	if err != nil {
		config.Log.Error("Error Set Status", zap.Error(err))
		return 500, fmt.Errorf("Error Set Status")
	}

	return 200, nil
}
