package main

import (
    "log"
    "net/http"

    "github.com/rs/cors"
    "github.com/yanran666/calculator/server/calculatorv1connect"

)

func main() {
    mux := http.NewServeMux()

    // 实例化我们自己实现的服务
    calcSvc := &CalculatorServiceServer{}

    // 利用 connect 包装
    path, handler := calculatorv1connect.NewCalculatorServiceHandler(calcSvc)
    mux.Handle(path, handler)

    // 配置 CORS 中间件
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // 允许前端地址
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Connect-Protocol-Version"},
		AllowCredentials: true,
	})
	handlerWithCors := c.Handler(mux)

	log.Println("Calculator server listening on :8080...")
	if err := http.ListenAndServe(":8080", handlerWithCors); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
