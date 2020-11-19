package db

import (
	"fmt"

	"github.com/sharmarajdaksh/go-pwd/config"

	"github.com/sharmarajdaksh/go-pwd/secrets"

	"gorm.io/gorm"
)

// Password represents the database representation of a Password record
type Password struct {
	gorm.Model
	App      string
	URL      string
	Username string
	Email    string
	Password string
}

// NewPassword returns a new Password object
func NewPassword(app, url, username, email, password string) Password {
	return Password{
		App:      app,
		URL:      url,
		Username: username,
		Email:    email,
		Password: password,
	}
}

// Save persists a Password instance to database
func (p *Password) Save() error {
	var err error
	p.Password, err = secrets.EncryptString(
		config.GetMasterSecret(),
		p.Password,
	)
	if err != nil {
		return fmt.Errorf("failed to save password: %w", err)
	}

	if tx := db.Create(p); tx.Error != nil {
		return fmt.Errorf("failed to save Password: %w", tx.Error)
	}
	return nil
}

// GetRecordsByApp returns Password records for a given app
func GetRecordsByApp(app string) ([]Password, error) {
	return fetchRecordsForQuery(&Password{App: app})
}

// GetRecordsByEmail returns Password records for a given username
func GetRecordsByEmail(email string) ([]Password, error) {
	return fetchRecordsForQuery(&Password{Email: email})
}

// GetRecordsByUsername returns Password records for a given username
func GetRecordsByUsername(username string) ([]Password, error) {
	return fetchRecordsForQuery(&Password{Username: username})
}

// GetAllRecords returns all password records in the database
func GetAllRecords() ([]Password, error) {
	return fetchRecordsForQuery(&Password{})
}

func fetchRecordsForQuery(p *Password) ([]Password, error) {
	rows, err := db.Model(&Password{}).Where(p).Rows()
	defer rows.Close()

	if err != nil {
		return []Password{}, fmt.Errorf("failed to query database: %w", err)
	}

	var ps []Password

	decKey := config.GetMasterSecret()

	for rows.Next() {
		var p Password
		err = db.ScanRows(rows, &p)
		if err != nil {
			return []Password{}, fmt.Errorf("failed to parse data: %w", err)
		}

		p.Password, err = secrets.DecryptString(decKey, p.Password)
		if err != nil {
			return []Password{}, fmt.Errorf("failed to encrypt data: %w", err)
		}

		ps = append(ps, p)
	}

	return ps, nil
}
