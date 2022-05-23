package sql

import (
	"fmt"
	"github.com/diploma/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
	"testing"
)

func TestDB(t *testing.T, dsn string) (*gorm.DB, func(...string)) {
	t.Helper()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}

	if err = sqlDB.Ping(); err != nil {
		t.Fatal(err)
	}

	if err = db.AutoMigrate(&domain.MovieEntity{}); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s", strings.Join(tables, ", ")))
		}

		sqlDB.Close()
	}
}
