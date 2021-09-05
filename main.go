package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ShunyaNagashige/golang-jwt-sample/config"
	handler "github.com/ShunyaNagashige/golang-jwt-sample/interface/handler/rest"
	"github.com/ShunyaNagashige/golang-jwt-sample/utils"
	"github.com/julienschmidt/httprouter"
)

func main() {
	if err := config.SetEnv("./config/env"); err != nil {
		if err != nil {
			log.Fatalf("Error in main function : %+v\n", err)
		}
	}

	if err := utils.LoggingSettings(os.Getenv("LOG_FILE")); err != nil {
		log.Fatalf("Error in main function : %+v\n", err)
	}

	// db, err := config.ConnectDB()
	// if err != nil {
	// 	log.Fatalf("Error in main function : %+v\n", err)
	// }

	// // 依存関係を定義（依存性の注入）
	// userPersistence := persistence.NewUserPersistence(db)
	// userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler()

	// ルーティングの設定
	router := httprouter.New()
	router.POST("/users", userHandler.Create)
	router.GET("/auth", userHandler.Auth)

	if err := http.ListenAndServe(":10000", router); err != nil {
		log.Fatalf("Error in main function : %+v\n", err)
	}

	// router := gin.Default()

	// router.POST("/users", userHandler.Create)

	// router.Run(":10000")

	// // ルーティングの設定
	// router := httprouter.New()
	// router.GET("/api/users", userHandler.Index)

	// // サーバ起動
	// http.ListenAndServe(":8080", router)
}

/*
type Handler interface{
	ServeHTTP(ResponseWriter,*Request)
}
である．
つまり，上記のServeHTTPさえ持っていれば
Handlerを実装したことになる．
*/
