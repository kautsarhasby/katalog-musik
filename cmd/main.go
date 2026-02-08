package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kautsarhasby/katalog-musik/internal/configs"
	"github.com/kautsarhasby/katalog-musik/pkg/internalsql"
)

func main() {
	r := gin.Default()
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolders(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Failed initiated config", err)
	}

	cfg = configs.Get()
	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Faile to initialize database")
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if err := r.Run(cfg.Service.Ports); err != nil {
		log.Fatalf("Failed to run server %v", err)
	}

}
