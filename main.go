package main

import (
	"fmt"
	"go_cron/pkg/jobs"
	"go_cron/pkg/logging"
	"net/http"
	//l4g "github.com/alecthomas/log4go"
	"go_cron/pkg/setting"
	"go_cron/routers"
)

func main() {
	logging.Info("run...")
	jobs.InitJobs()
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
