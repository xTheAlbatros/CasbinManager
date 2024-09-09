package main

import (
	"CasbinManager/config"
	"CasbinManager/controllers"
	"CasbinManager/middleware"
	"CasbinManager/services"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

const CasdoorCert = `cert`

func InitCasdoorConfig() {
	casdoorsdk.InitConfig(
		"http://localhost:8000",
		"clientID",
		"clientSecret",
		CasdoorCert,
		"Org",
		"App",
	)
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	db := config.InitDB()

	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatalf("failed to initialize Casbin adapter: %v", err)
	}

	enforcer, err := casbin.NewEnforcer("resources/casbin_model.conf", adapter)
	if err != nil {
		log.Fatalf("failed to initialize Casbin enforcer: %v", err)
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		log.Fatalf("failed to load Casbin policy: %v", err)
	}

	InitCasdoorConfig()

	casbinService := services.NewCasbinService(db)

	r.Use(middleware.CasbinMiddleware(enforcer))

	casbinController := controllers.NewCasbinController(casbinService)
	r.GET("/api/rules", casbinController.GetRules)
	r.POST("/api/rules", casbinController.AddRule)
	r.PUT("/api/rules/:id", casbinController.UpdateRule)
	r.DELETE("/api/rules/:id", casbinController.DeleteRule)

	if err := r.Run(":8083"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
