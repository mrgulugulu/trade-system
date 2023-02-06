// package server 与http服务相关
package server

import (
	"fmt"
	"net/http"
	"trade-system/internal/log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Addr string
	Port string
}

func (s *Server) Run() {
	r := gin.Default()

	r.GET("/kLine1Min", queryKLineIn1Min)
	r.GET("/kLine5Min", queryKLineIn5Min)
	// r.GET("/filmInfo/top10", top10)
	// r.DELETE("/filmInfo", delete)
	// r.POST("/filmInfo", update)

	ser := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.Addr, s.Port),
		Handler: r,
	}
	// 搞个signal来监听，实现优雅关闭
	log.Sugar.Infof("listening ", ser.Addr)
	ser.ListenAndServe()
	// gracefulExitServer(ser)
}

// gracefulExitServer 服务器优雅退出
// func gracefulExitServer(ser *http.Server) {
// 	ch := make(chan os.Signal, 1)
// 	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
// 	<-ch
// 	nowTime := time.Now()
// 	movieTitleList, err := dao.D.GetMovieSetMembers(config.QueryMovieSet)
// 	if err != nil {
// 		log.Printf("get movie set member error: %v", err)
// 	}
// 	for _, title := range movieTitleList {
// 		viewNum, err := dao.D.GetMovieViewNumber(title)
// 		if err != nil {
// 			log.Printf("get movie view number error: %v", err)
// 		}
// 		err = dao.D.UpdateMovieInfo(title, "view_number", viewNum)
// 		if err != nil {
// 			log.Printf("update movie viewNumber error: %v", err)
// 		}
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	err = ser.Shutdown(ctx)
// 	if err != nil {
// 		log.Printf("shutdown error: %v", err)
// 	}
// 	fmt.Println("-----exited-----", time.Since(nowTime))
// }
