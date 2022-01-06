package main

import (
	docs "TMA/docs"
	router "TMA/internal/tma_management/rest"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func main(){
	//init env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	//route
	r := router.Setup()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		u := v1.Group("/user")
		{
			u.GET("/get_user", router.GetUser())
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

}

