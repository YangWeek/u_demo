package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"u_demo/conf"
	"u_demo/controller"
	"u_demo/dao/mysql"
	"u_demo/dao/redis"
	ginzap "u_demo/gin_zap"
	"u_demo/router"
)

// @title u_demo项目接口文档
// @version 1.0
// @description 社区项目
// @termsOfService http://swagger.io/terms/

// @contact.name yang
// @contact.email 2033231795@qq.com

// 暂时无
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:9099
// @BasePath /api/v1
func main() {

	//检查命令行参数
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: bluebell config.yaml")
		return
	}

	if err := conf.Init_Viper(os.Args[1]); err != nil {
		fmt.Printf("conf init error is %v\n", err)
		return
	}

	if err := ginzap.InitGin_Zap(conf.Conf.LogConfig, conf.Conf.Mode); err != nil {
		fmt.Printf("fin_zap err is %v\n", err)
		return
	}

	if err := mysql.Init(conf.Conf.MySQLConfig); err != nil {
		fmt.Printf("mysql err is %v\n", err)
		return
	}
	defer mysql.Close()

	if err := redis.Init(conf.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	// 初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}

	r := router.SetRouter(conf.Conf.Mode)

	srv := &http.Server{
		Addr:    ":9099",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")

}
