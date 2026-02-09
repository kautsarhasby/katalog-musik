package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kautsarhasby/katalog-musik/internal/configs"
	membershipHandler "github.com/kautsarhasby/katalog-musik/internal/handlers/memberships"
	"github.com/kautsarhasby/katalog-musik/internal/models/memberships"
	membershipRepo "github.com/kautsarhasby/katalog-musik/internal/repository/memberships"
	membershipSvc "github.com/kautsarhasby/katalog-musik/internal/service/memberships"
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
		log.Fatal("Failed to initialize database")
	}
	db.AutoMigrate(&memberships.User{})

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Users
	membershipsRepository := membershipRepo.NewRepository(db)
	membershipsService := membershipSvc.NewService(cfg, membershipsRepository)
	membershipsHandler := membershipHandler.NewHandler(r, membershipsService)
	membershipsHandler.RegisterRoute()

	if err := r.Run(cfg.Service.Port); err != nil {
		log.Fatalf("Failed to run server %v", err)
	}

}
