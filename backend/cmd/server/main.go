package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gitwb-c/crm.saas/backend/internal/db"
	"github.com/gitwb-c/crm.saas/backend/internal/db/cache"
	_ "github.com/gitwb-c/crm.saas/backend/internal/ent/runtime"
	"github.com/gitwb-c/crm.saas/backend/internal/http"
	"github.com/gitwb-c/crm.saas/backend/pkg/env"
)

func main() {
	if err := env.InitEnv(); err != nil {
		return
	}
	client, e := db.NewClient()
	if e != nil {
		log.Fatalf("failed to init Ent client: %v", e)
		return
	}
	_, er := cache.NewCacheClient()
	if er != nil {
		log.Fatalf("failed to init redis client: %v", er)
		return
	}
	srv, err := db.Init(client)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	defer db.Close(client)

	r := http.GlobalRouter(client, srv)
	r.Run(fmt.Sprintf(":%v", os.Getenv("SERVER_PORT")))
}
