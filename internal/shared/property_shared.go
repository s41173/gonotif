package shared

import (
	"time"

	"gorm.io/gorm"
)

// Property merepresentasikan satu record pada tabel property.
type Property struct {
	// ── identitas ────────────────────────────────────────
	IDPrimary  int     `db:"idPrimary" json:"id"`
	Name       string  `db:"name"             json:"name"`
	SiteName   string  `db:"site_name"        json:"site_name"`
	Logo       *string `db:"logo"             json:"logo"`
	Manager    *string `db:"manager"          json:"manager"`
	Accounting *string `db:"accounting"       json:"accounting"`

	// ── alamat & lokasi ──────────────────────────────────
	Address    string `db:"address"          json:"address"`
	Coordinate string `db:"coordinate"       json:"coordinate"`
	Zip        int    `db:"zip"              json:"zip"`
	City       string `db:"city"             json:"city"`

	// ── kontak ───────────────────────────────────────────
	Phone1         string  `db:"phone1"           json:"phone1"`
	Phone2         string  `db:"phone2"           json:"phone2"`
	Fax            string  `db:"fax"              json:"fax"`
	Email          string  `db:"email"            json:"email"`
	BillingEmail   string  `db:"billing_email"    json:"billing_email"`
	TechnicalEmail *string `db:"technical_email"  json:"technical_email"`
	CcEmail        string  `db:"cc_email"         json:"cc_email"`
	EmailLink      *string `db:"email_link"       json:"email_link"`

	// ── bank ─────────────────────────────────────────────
	AccountName string `db:"account_name"     json:"account_name"`
	AccountNo   string `db:"account_no"       json:"account_no"`
	Bank        string `db:"bank"             json:"bank"`

	// ── SEO / media ──────────────────────────────────────
	MetaDescription string  `db:"meta_description" json:"meta_description"`
	MetaKeyword     string  `db:"meta_keyword"     json:"meta_keyword"`
	UrlUpload       *string `db:"url_upload"       json:"url_upload"`
	ImageUrl        *string `db:"image_url"        json:"image_url"`

	// ── timestamp ────────────────────────────────────────
	Created *time.Time `db:"created"          json:"created_at"`
	Updated *time.Time `db:"updated"          json:"updated_at"`
	Deleted *time.Time `db:"deleted"          json:"deleted_at"`
}

// GetProperty retrieves property data from the database.
func (p *Property) GetProperty(db *gorm.DB) (map[string]interface{}, error) {
	var property Property

	if err := db.First(&property).Error; err != nil {
		return nil, err
	}

	// Handle url_upload — fallback ke "./" jika nil
	urlUpload := "./"
	if property.UrlUpload != nil {
		urlUpload = *property.UrlUpload
	}

	// Handle image_url — fallback ke default path jika nil
	imageUrl := "base_url/images/"
	if property.ImageUrl != nil {
		imageUrl = *property.ImageUrl
	}

	data := map[string]interface{}{
		"name":            property.Name,
		"address":         property.Address,
		"coordinate":      property.Coordinate,
		"phone1":          property.Phone1,
		"phone2":          property.Phone2,
		"fax":             property.Fax,
		"email":           property.Email,
		"email_link":      property.EmailLink,
		"billing_email":   property.BillingEmail,
		"technical_email": property.TechnicalEmail,
		"cc_email":        property.CcEmail,
		"zip":             property.Zip,
		"city":            property.City,
		"account":         property.AccountName,
		"acc_no":          property.AccountNo,
		"bank":            property.Bank,
		"manager":         property.Manager,
		"accounting":      property.Accounting,
		"site_name":       property.SiteName,
		"logo":            property.Logo,
		"url_upload":      urlUpload,
		"image_url":       imageUrl,
		"meta_desc":       property.MetaDescription,
		"meta_key":        property.MetaKeyword,
	}

	return data, nil
}

// TableName sets the name of the database table.
func (Property) TableName() string {
	return "property"
}
