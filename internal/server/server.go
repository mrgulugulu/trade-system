// package server 与http服务相关
package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"trade-system/internal/log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Addr string
	Port string
}

func (s *Server) Run() {

	r := gin.Default()
	kLine1Min := r.Group("/kLine1Min")
	{
		kLine1Min.GET("", queryKLineIn1Min)
		kLine1Min.GET("/:key", queryKLineIn1MinWithKey)
	}
	kLine5Min := r.Group("/kLine5Min")
	{
		kLine5Min.GET("", queryKLineIn5Min)
		kLine5Min.GET("/:key", queryKLineIn5MinWithKey)
	}

	ser := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.Addr, s.Port),
		Handler: r,
	}
	// 搞个signal来监听，实现优雅关闭
	log.Sugar.Infof("listening %v", ser.Addr)

	go ser.ListenAndServe()
	gracefulExitServer(ser)
}

// gracefulExitServer 服务器优雅退出
func gracefulExitServer(ser *http.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-ch
	nowTime := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := ser.Shutdown(ctx)
	if err != nil {
		log.Sugar.Errorf("shutdown error: %v", err)
	}
	log.Sugar.Info("-----exited-----", time.Since(nowTime))
}
