package main

import (
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"log"
	"maxtv_middleware/pkg/api"
	. "maxtv_middleware/pkg/db_interface"
	"maxtv_middleware/pkg/jobs"
	. "maxtv_middleware/pkg/log_processing"
	"os"
	"time"
)

func init() {
	jobs.Sch = gocron.NewScheduler(time.Local)
	jobs.Sch.StartAsync()

	jobs.JobInit()
}

func main() {

	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("ERROR opening file: %v", err)
	}
	defer f.Close()

	err = godotenv.Load("p.env")
	if err != nil {
		log.Fatal("ERROR loading .env file")
	}

	InitLog(f)
	InitDatabase()
	api.Processing()

}
