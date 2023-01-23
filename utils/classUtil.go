package utils

import (
	"anonichat-pgx-stdlib/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("gagal membaca file .env ", err)
	}
	return os.Getenv(key)
}

func RetSucc(ctx *gin.Context, desc string, data interface{}) {
	r := models.Response{}
	r.StatusCode = "01"
	r.Description = desc
	r.Data = data
	ctx.JSON(http.StatusOK, r)
}

func RetBadReq(ctx *gin.Context, err error) {
	r := models.Response{}
	r.StatusCode = "01"
	r.Description = err.Error()
	ctx.JSON(http.StatusBadRequest, r)
}

func Timer(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
