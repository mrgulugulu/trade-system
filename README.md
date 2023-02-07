# trade-system

## 软件功能
从redis的publish消息中获取消息对的最新成交价格，定时生成1min和5min的k线，k线数据包括高，开，收，低，收，量。生成后将存入mysql，并提供api接口，可获取k线信息。

## 代码目录结构
├── cmd  
│   ├── main  
│   │   └── main.go               // 主程序  
│   ├── trade-price-gen  
│   │   └── trade-price-gen.go    // 生成随机的交易对信息，并存入mysql  
│   └── trade-price-read  
│       └── trade-price-read.go   // 从mysql读取交易对信息，并通过redis的publish发布  
├── config  
│   ├── config.go                 // 配置以及变量的定义  
│   └── config.yaml               // 配置文件  
├── go.mod  
├── go.sum  
├── internal  
│   ├── cache  
│   │   └── cache.go              // go-cache的初始化  
│   ├── dao  
│   │   ├── dao.go                  
│   │   ├── mysql.go              // 操作mysql  
│   │   └── redis.go              // 操作redis  
│   ├── kline  
│   │   └── kline.go              // k线信息的生成函数  
│   ├── log  
│   │   └── log.go                // 日志文件的初始化  
│   ├── model  
│   │   ├── model.go              // 程序使用的数据库模型定义  
│   │   ├── model.sql             // 用于mysql建表  
│   │   └── trade-pair.sql        // 用于mysql建表  
│   └── server  
│       ├── middleware.go         // 中间件（暂没使用）  
│       ├── router.go             // 路由  
│       └── server.go             // 服务器  
└── README.md  

## 快速启动
1. 初次使用需要先在config.yaml配置好信息，并在mysql建立对应的表格。  
2. 启动cmd/trade-price-gen/trade-price-gen.go， 生成随机交易对并写入mysql（建议生成多一点，默认是10w条）。  
3. 启动cmd/main/main.go，即可订阅redis的对应频道读取信息，通过redis的publish功能发布最新k线信息并启动http服务  
4. 启动cmd/trade-price-read/trade-price-read.go，从mysql中读取交易对信息，并以每秒10笔的速度publish出去  

## API实例 (如无amount，则默认为十条信息)
`GET /kLine1Min?amount=k` ：返回最新k条1分钟k线信息  
`GET /kLine5Min?amount=k` ：返回最新k条5分钟k线信息  
`GET /kLine1Min/key?amount=k` ：返回倒序以key字段排列的k条数据；如`/kLine1Min/open?amount=k`，则返回最大的k个开盘价1分钟k线信息并以倒序排列  
`GET /kLine5Min/key?amount=k` ：返回倒序以key字段排列的k条数据；如`/kLine5Min/open?amount=k`，则返回最大的k个开盘价5分钟k线信息并以倒序排列  


