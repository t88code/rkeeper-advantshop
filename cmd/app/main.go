package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"rkeeper2advantshop/internal/handler"
	"rkeeper2advantshop/pkg/advantshop"
	"rkeeper2advantshop/pkg/config"
	check "rkeeper2advantshop/pkg/license"
	"rkeeper2advantshop/pkg/logging"
	_ "rkeeper2advantshop/pkg/telegram"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Start service Advantshop Integrations v1.0") // todo version
	defer logger.Info("End Main")

	check.Check()
	cfg := config.GetConfig()

	//go telegram.BotStart()

	_, err := advantshop.NewClient(cfg) // todo contex
	if err != nil {
		logger.Fatal(err)
	}
	router := httprouter.New()

	router.GET("/GetCardInfoEx", handler.GetCardInfoEx)
	router.GET("/FindByEmail", handler.FindByEmail)
	router.POST("/TransactionsEx", handler.TransactionsEx)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.SERVICE.PORT), RequestLogger{h: router, l: logger}))
}

type RequestLogger struct {
	h http.Handler
	l *logging.Logger
}

func (rl RequestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	rl.l.Debugf("Started %s %s", r.Method, r.URL.Path)
	rl.l.Debug("Request: ", r)
	rl.l.Debug("Method: ", r.Method)
	rl.l.Debug("Host: ", r.Host)
	rl.l.Debug("URL: ", r.URL)
	rl.l.Debug("RequestURI: ", r.RequestURI)
	rl.l.Debug("path: ", r.URL.Path)
	rl.l.Debug("Form: ", r.Form)
	rl.l.Debug("MultipartForm: ", r.MultipartForm)
	rl.l.Debug("ContentLength: ", r.ContentLength)
	rl.l.Debug("Header: ", r.Header)
	rl.h.ServeHTTP(w, r)
	rl.l.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
}

// todo delete
//func init() {
//	logger := logging.GetLogger()
//
//	logger.Println("Start main init...")
//	defer logger.Println("End main init.")
//	cfg := config.GetConfig()
//	var err error
//
//	_ = wooapi.NewAPI(cfg.WOOCOMMERCE.URL, cfg.WOOCOMMERCE.Key, cfg.WOOCOMMERCE.Secret)
//
//	_, err = rk7api.NewAPI(cfg.RK7.URL, cfg.RK7.User, cfg.RK7.Pass, "REF")
//	if err != nil {
//		logger.Fatal("failed main init; rk7api.NewAPI; ", err)
//	}
//
//	_, err = rk7api.NewAPI(cfg.RK7MID.URL, cfg.RK7MID.User, cfg.RK7MID.Pass, "MID")
//	if err != nil {
//		logger.Fatal("failed main init; rk7api.NewAPI; ", err)
//	}
//
//	_, err = cache.NewCacheMenu()
//	if err != nil {
//		logger.Error("failed in cache.NewCacheMenu()")
//	}
//
//	if database.Exists(cfg.DBSQLITE.DB) != true {
//		logger.Info(cfg.DBSQLITE.DB, " not exist")
//		err := database.CreateDB(cfg.DBSQLITE.DB)
//		if err != nil {
//			logger.Fatalf("%s, %v", cfg.DBSQLITE.DB, err)
//		}
//	} else {
//		logger.Info(cfg.DBSQLITE.DB, " exist")
//	}
//}
