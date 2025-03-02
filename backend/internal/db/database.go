package db

import (
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent"
	"github.com/tuoitrevohoc/gofw/backend/internal/config"

	_ "github.com/lib/pq"
)

func NewDB(config *config.AppConfig) (*ent.Client, error) {
	return ent.Open("postgres", config.Db.Url)
}
