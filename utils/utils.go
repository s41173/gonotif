// utils/password.go

package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// GetLocalTime mengembalikan waktu saat ini dalam zona waktu lokal
func GetLocalTime() time.Time {
	// Ganti dengan nama zona waktu yang sesuai
	loc, err := time.LoadLocation("Asia/Jakarta") // Misalnya untuk WIB
	if err != nil {
		panic(err) // Atau tangani kesalahan sesuai kebutuhan Anda
	}
	return time.Now().In(loc) // Kembalikan waktu saat ini dalam zona waktu lokal
}

func SplitSpace(s string) string {

	// lowercase
	s = strings.ToLower(s)

	// replace non alphanumeric dengan "-"
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	s = reg.ReplaceAllString(s, "-")

	// trim "-"
	s = strings.Trim(s, "-")

	return s
}

func BaseURL(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + c.Request.Host
}

// format rupiah (simple version)
func FormatRupiah(amount int) string {
	str := fmt.Sprintf("%d", amount)
	n := len(str)

	if n <= 3 {
		return str
	}

	var result []byte
	mod := n % 3
	if mod > 0 {
		result = append(result, str[:mod]...)
		if n > mod {
			result = append(result, '.')
		}
	}

	for i := mod; i < n; i += 3 {
		result = append(result, str[i:i+3]...)
		if i+3 < n {
			result = append(result, '.')
		}
	}

	return string(result)
}

// format tanggal
func FormatDate(t time.Time) string {
	return t.Format("02 Jan 2006 15:04")
}

func FormatDateString(s string) string {
	formats := []string{
		time.RFC3339,          // 2006-01-02T15:04:05Z
		"2006-01-02 15:04:05", // MySQL datetime
		"2006-01-02",          // date only
	}

	for _, layout := range formats {
		t, err := time.Parse(layout, s)
		if err == nil {
			return t.Format("02 Jan 2006 15:04")
		}
	}

	return s // return as-is kalau semua gagal
}

// helper untuk ambil string dari *string pointer di map
func ConvertToString(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(*string); ok && s != nil {
		return *s
	}
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
