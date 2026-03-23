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

	t.Log(tv.ServerName)
}
