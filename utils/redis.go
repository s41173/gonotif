package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"go-notif/config"
	authdto "go-notif/internal/api/dto"
	"time"
)

// SetSession simpan session user ke Redis
func Set_User_Session(user authdto.CustomerSession, duration time.Duration) error {
	key := fmt.Sprintf("session:%d", user.UserID)
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return config.RDB.Set(config.Ctx, key, data, duration).Err()
}

// GetSession ambil session dari Redis
func Get_User_Session(userID int64) (*authdto.CustomerSession, error) {
	key := fmt.Sprintf("session:%d", userID)
	val, err := config.RDB.Get(config.Ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var session authdto.CustomerSession
	if err := json.Unmarshal([]byte(val), &session); err != nil {
		return nil, err
	}

	return &session, nil
}

// DeleteSession hapus session dari Redis
func Delete_User_Session(userID int64) error {
	key := fmt.Sprintf("session:%d", userID)
	return config.RDB.Del(config.Ctx, key).Err()
}

// ================== Product Cache =====================
func Set_Product_Cache(key string, data interface{}, duration time.Duration) error {
	json_data, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return config.RDB.Set(context.Background(), key, json_data, duration).Err()
}

func Get_Product_Cache(key string, dest interface{}) error {
	val, err := config.RDB.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

func Delete_Product_Cache(key string) error {
	return config.RDB.Del(context.Background(), key).Err()
}

func Product_Cache_Key(suffix string) string {
	return fmt.Sprintf("product:%s", suffix)
}

// voucher
func SetVoucherSelected(customerID int, voucherID int, duration time.Duration) error {
	key := fmt.Sprintf("voucher:selected:%d", customerID)
	return config.RDB.Set(config.Ctx, key, voucherID, duration).Err()
}

func GetVoucherSelected(customerID int) (int, error) {
	key := fmt.Sprintf("voucher:selected:%d", customerID)
	val, err := config.RDB.Get(config.Ctx, key).Int()
	if err != nil {
		return 0, err
	}
	return val, nil
}

func DelVoucherSelected(customerID int) error {
	key := fmt.Sprintf("voucher:selected:%d", customerID)
	return config.RDB.Del(config.Ctx, key).Err()
}
