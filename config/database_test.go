package config

import (
	"testing"
)

func TestConnection(t *testing.T) {
	ConnectDB()
	if DB.Error != nil {
		t.Errorf("Gagal koneksi")
	}
}
