package db

import (
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent"

	_ "github.com/lib/pq"
)

func NewEntClient(dbUrl string) (*ent.Client, error) {
	return ent.Open("postgres", dbUrl)
}
