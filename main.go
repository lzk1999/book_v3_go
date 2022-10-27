package main

import (
	"book_v3_go/mysql"
	"book_v3_go/redis"
	"book_v3_go/routes"
	"book_v3_go/setting"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//	1.加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Println("配置文件初始化失败", err)
	}
	//  2.初始化Mysql
	if err := mysql.Init(); err != nil {
		fmt.Println("Mysql初始化失败", err)
	}
	//defer mysql.DB.C
	//	3.初始化Redis
	if err := redis.Init(); err != nil {
		fmt.Println("Redis初始化失败", err)
	}
	defer redis.Rdb.Close()
	//	4.注册路由
	r := routes.Setup()
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server listen err:%s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 在此阻塞
	<-quit

	ctx, channel := context.WithTimeout(context.Background(), 5*time.Second)

	defer channel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown error")
	}
	log.Println("server exiting...")

}
