package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rkeeper-advantshop/internal/handler"
	"rkeeper-advantshop/pkg/config"
	"rkeeper-advantshop/pkg/crm"
	"rkeeper-advantshop/pkg/crm/options/api"
	check "rkeeper-advantshop/pkg/license"
	"rkeeper-advantshop/pkg/logging"
	"rkeeper-advantshop/pkg/rk7api"
	"rkeeper-advantshop/pkg/telegram"
	"time"
)

// TODO сделать зашифрованный канал
// сделать файнд майл
// TODO сделать офлайн
// TODO сделать получение кодов из окружения и дату окончания
// TODO сделать лицензию из файла
// TODO несколько транзакций одновременно
// TODO обработку отмены транзакции
// TODO веб морду для облака
// TODO хранение логов
// TODO удалить лишние модули
// TODO проверь лицензии перед отправкой

func main() {
	loggerMain, err := logging.NewLogger(
		true,
		"main.log",
		"main",
		"main")
	if err != nil {
		fmt.Println(err)
	}

	loggerMain.Info("Start service RestoCRM")
	defer loggerMain.Info("End Main")

	loggerTelegram, err := logging.NewLogger(
		true,
		"telegram.log",
		"telegram",
		"telegram")
	if err != nil {
		loggerMain.Fatal(err)
	}

	loggerConfig, err := logging.NewLogger(
		true,
		"config.log",
		"config",
		"config")
	if err != nil {
		loggerMain.Fatal(err)
	}

	check.Check()
	cfg, err := config.GetConfig(loggerConfig)
	if err != nil {
		loggerMain.Fatal(err)
	}

	_, err = rk7api.NewAPI(cfg.RK7MID.URL,
		cfg.RK7MID.User,
		cfg.RK7MID.Pass,
		cfg.XMLINTERFACE.Type,
		cfg.XMLINTERFACE.UserName,
		cfg.XMLINTERFACE.Password,
		cfg.XMLINTERFACE.Token,
		cfg.XMLINTERFACE.ProductID,
		cfg.XMLINTERFACE.Guid,
		cfg.XMLINTERFACE.URL,
		cfg.XMLINTERFACE.RestCode)
	if err != nil {
		loggerMain.Fatal(err)
	}

	_, err = crm.NewAPI(
		api.Advantshop(loggerMain,
			cfg.LOG.Debug,
			cfg.ADVANTSHOP.RPS,
			cfg.ADVANTSHOP.ApiKey,
			cfg.ADVANTSHOP.URL,
			cfg.ADVANTSHOP.OrderPrefix,
			cfg.ADVANTSHOP.OrderSource,
			cfg.ADVANTSHOP.Currency,
			cfg.ADVANTSHOP.CheckOrderItemExist,
			cfg.ADVANTSHOP.CheckOrderItemAvailable,
			cfg.ADVANTSHOP.Timeout,
			cfg.ADVANTSHOP.BonusInFio,
		))
	if err != nil {
		loggerMain.Fatal(err)
	}

	go telegram.BotStart(
		loggerTelegram,
		"telegram.db",
		cfg.TELEGRAM.BotToken,
		cfg.TELEGRAM.Debug)

	router := httprouter.New()
	router.GET("/GetCardInfoEx", handler.GetCardInfoEx)
	router.GET("/FindByEmail", handler.FindByEmail)
	router.POST("/TransactionsEx", handler.TransactionsEx)
	router.GET("/FindOwnerByNamePart", handler.FindOwnerByNamePart)
	//router.POST("/Update", handler.UpdateDiscount) TODO ручку для обновления скидок-грейдов

	loggerMain.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", cfg.SERVICE.PORT),
			RequestLogger{h: router, l: loggerMain},
		),
	)
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
	rl.l.Debug("ApiUrl: ", r.URL)
	rl.l.Debug("RequestURI: ", r.RequestURI)
	rl.l.Debug("path: ", r.URL.Path)
	rl.l.Debug("Form: ", r.Form)
	rl.l.Debug("MultipartForm: ", r.MultipartForm)
	rl.l.Debug("ContentLength: ", r.ContentLength)
	rl.l.Debug("Header: ", r.Header)
	rl.h.ServeHTTP(w, r)
	rl.l.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
}
