package main

import (
	"clean-architecture-beego/internal/domain"
	productHandler "clean-architecture-beego/internal/product/delivery/http"
	productRepo "clean-architecture-beego/internal/product/repository"
	productUcase "clean-architecture-beego/internal/product/usecase"
	userHandler "clean-architecture-beego/internal/user/delivery/http"
	userRepo "clean-architecture-beego/internal/user/repository"
	userUcase "clean-architecture-beego/internal/user/usecase"
	"clean-architecture-beego/pkg/database"
	"clean-architecture-beego/pkg/helpers/response"
	"clean-architecture-beego/pkg/jwt"
	"clean-architecture-beego/pkg/middleware"
	"github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/client/cache/redis"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"net/http"
	"time"
)

// @title BEE SAMPLE API
// @version v1
// @contact.name Kredit Plus
// @contact.url https://kreditplus.com
// @contact.email support@kreditplus.com
// @description api "sample using beego framework"
// @termsOfService https://dev-kreditmu.kreditplus.com/terms

// @BasePath /api
// @query.collection.format multi
func main() {

	// default vars
	var (
		requestTimeout = 3
	)

	if err := beego.LoadAppConfig("ini", "./conf/app.conf"); err != nil {
		panic(err)
	}

	// initialization database
	db := database.DB()

	// migrate database
	if err := db.AutoMigrate(&domain.Product{}, &domain.Customer{}, &domain.Order{}, &domain.User{}); err != nil {
		panic(err)
	}

	// global timeout
	if timeout, err := beego.AppConfig.Int("timeout"); err == nil {
		requestTimeout = timeout
	}

	timeoutContext := time.Duration(requestTimeout) * time.Second

	// swagger config
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// middleware cors
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowMethods:    []string{http.MethodGet, http.MethodPost},
		AllowAllOrigins: true,
	}))

	// jwt middleware
	if auth, err := jwt.NewJwt(&jwt.Options{
		Issuer:      "backend",
		SignMethod:  jwt.HS256,
		SecretKey:   "secret",
		ExpiredTime: 1500,
		Locations:   "header:Authorization",
		IdentityKey: "uid",
	}); err == nil {
		if bm, err := cache.NewCache("redis", `{"conn":"127.0.0.1:6379"}`); err != nil {
			panic(err)
		} else {
			auth.SetAdapter(bm)

			// user handler
			userRepository := userRepo.NewUserRepository(db)
			userUseCase := userUcase.NewUserUseCase(timeoutContext, userRepository)
			userHandler.NewUserHandler(userUseCase, auth)

			beego.InsertFilterChain("/api/*", middleware.NewJwtMiddleware().JwtMiddleware(auth))
		}
	} else {
		panic(err)
	}

	// default error handler
	beego.ErrorController(&response.ErrorController{})

	// product handler
	productRepository := productRepo.NewProductRepository(db)
	productUseCase := productUcase.NewProductUseCase(timeoutContext, productRepository)
	productHandler.NewProductHandler(productUseCase)

	beego.Run()
}