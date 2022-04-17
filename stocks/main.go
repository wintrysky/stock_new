package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ww/stocks/global"
	"ww/stocks/orm"
	"ww/stocks/router"
	"ww/stocks/template"
	"ww/stocks/utils"
	"ww/stocks/xlog"
)
func init() {
	utils.InitLog("./logs/", "service", "gb18030")
	initLog() // DB Log
}

func initLog() error {
	defer xlog.Sync()

	err := xlog.Init(&xlog.LogSettings{
		Level:    xlog.DefaultLevel,
		Path:     xlog.DefaultPath,
		FileName: xlog.DefaultFileName,
		CataLog:  xlog.DefaultCataLog,
		Caller:   xlog.DefaultCaller,
	})
	return err
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func startServer() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	r.Use(cors())

	router.UseRouters(r)

	r.Run(":8020")
}

func initTemplate(){
	template.InitFutuColumnMap()
}


func initORM() {
	dbSettings := &orm.DbSettings{}
	global.ReadKey("db", dbSettings)
	err :=orm.InitDB(dbSettings)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	initORM()

	initTemplate()
	// start
	startServer()
}