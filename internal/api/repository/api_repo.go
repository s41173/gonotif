package repository

import (
	"go-notif/config"
	"go-notif/internal/api/model"
)

func Save(notif *model.Notif) error {
	return config.DB.Save(notif).Error
}

func GetLatestID() (int, error) {
	var id int
	err := config.DB.Model(&model.Notif{}).
		Select("id").
		Order("id DESC").
		Limit(1).
		Scan(&id).Error
	if err != nil {
		return 0, err
	}
	return id, nil
}

func Set_status(notifID int) error {

	// expired := time.Now().Add(24 * time.Hour)
	return config.DB.Model(&model.Notif{}).
		Where("id = ?", notifID).
		Updates(map[string]interface{}{
			"status": 1,
		}).Error
}
