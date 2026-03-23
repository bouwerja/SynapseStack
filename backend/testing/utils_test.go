package testing

import (
	"backend/utils"
	"testing"
)

func TestDotEnv(t *testing.T) {
	tv, err := utils.GetDatabaseDetails()
	if err != nil {
		t.Fatal(err)
	}

	if tv.ServerName != "" {
		t.Log("Test dot env passed")
	}
}
