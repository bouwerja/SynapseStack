package testing

import (
	db "backend/database/operations"
	"testing"
)

func TestDBConnection(t *testing.T) {
	dbConn := db.InsertScraper()
	t.Log(dbConn)
}
