package worker

import (
	"context"
	"encoding/json"

	"go-notif/config"
	"go-notif/internal/api/service"

	"go.uber.org/zap"
)

func StartNotifWorker(ctx context.Context) {
	config.Log.Info("Notif worker started, waiting for messages...")

	for {
		// BRPop standby sampai ada data, timeout 0 = selamanya
		result, err := config.RDB.BRPop(ctx, 0, "notif_channel").Result()
		if err != nil {
			// kalau context cancelled (app shutdown), stop worker
			if ctx.Err() != nil {
				config.Log.Info("Notif worker stopped")
				return
			}
			config.Log.Error("Error BRPop", zap.Error(err))
			continue
		}

		// result[0] = channel name, result[1] = data
		// fmt.Println("Isi Channel : ", result[1])

		var data map[string]interface{}
		if err := json.Unmarshal([]byte(result[1]), &data); err != nil {
			config.Log.Error("Failed unmarshal", zap.Error(err))
			continue
		}

		// ambil nilai dari map
		custID := int(data["customer"].(float64))
		custName, _ := data["custname"].(string)
		sentTo, _ := data["sentto"].(string)
		subject, _ := data["subject"].(string)
		content, _ := data["content"].(string)
		notifType := int(data["type"].(float64))
		modul, _ := data["modul"].(string)

		// simpan ke DB
		// code, err := service.Add_notif(custID, custName, sentTo, subject, content, notifType, modul, 0)
		// if err != nil {
		// 	config.Log.Error("Failed Add_notif",
		// 		zap.Int("code", code),
		// 		zap.Error(err),
		// 	)
		// 	continue
		// }

		go func(custID int, custName, sentTo, subject, content, modul string, notifType int) {
			code, err := service.Add_notif(custID, custName, sentTo, subject, content, notifType, modul, 0)
			if err != nil {
				config.Log.Error("Failed Add_notif",
					zap.Int("code", code),
					zap.Error(err),
				)
				return
			}
		}(custID, custName, sentTo, subject, content, modul, notifType)

		// go processNotif(payload)

	}
}

// func processNotif(payload utils.NotifPayload) {
// 	switch payload.Type {
// 	case 7: // WhatsApp
// 		stts := service.Send_wa(payload.SentTo, payload.Content)
// 		if !stts {
// 			config.Log.Error("Failed send WA",
// 				zap.String("sentto", payload.SentTo),
// 			)
// 		}
// 	// tambah case lain sesuai kebutuhan
// 	// case 1: // Email
// 	// case 2: // SMS
// 	default:
// 		config.Log.Warn("Unknown notif type", zap.Int("type", payload.Type))
// 	}
// }
