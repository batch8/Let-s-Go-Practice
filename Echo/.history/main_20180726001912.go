package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"./handler"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", handler.GreetingPage())
	e.GET("/hello/:username", handler.GreetingPage())

	// サーバー起動
	e.Start(":8888")
}
