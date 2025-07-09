package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Fungsi untuk menerjemahkan pesan error validasi dan database ke dalam map
func TranslateErrorMessage(err error) map[string]string {
	pesanError := make(map[string]string)

	// Validasi dari validator.v10
	if validasi, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validasi {
			field := fieldErr.Field()
			switch fieldErr.Tag() {
			case "required":
				pesanError[field] = fmt.Sprintf("%s wajib diisi", field)
			case "email":
				pesanError[field] = "Format email tidak valid"
			case "unique":
				pesanError[field] = fmt.Sprintf("%s sudah terdaftar", field)
			case "min":
				pesanError[field] = fmt.Sprintf("%s minimal %s karakter", field, fieldErr.Param())
			case "max":
				pesanError[field] = fmt.Sprintf("%s maksimal %s karakter", field, fieldErr.Param())
			case "numeric":
				pesanError[field] = fmt.Sprintf("%s harus berupa angka", field)
			default:
				pesanError[field] = "Input tidak valid"
			}
		}
	}

	// Error dari database (GORM)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "username") {
				pesanError["Username"] = "Username sudah terdaftar"
			}
			if strings.Contains(err.Error(), "email") {
				pesanError["Email"] = "Email sudah terdaftar"
			}
		} else if err == gorm.ErrRecordNotFound {
			pesanError["Error"] = "Data tidak ditemukan"
		}
	}

	return pesanError
}

// Fungsi untuk mendeteksi error duplikasi entri pada database
func IsDuplicateEntryError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "Duplicate entry") ||
		strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}
