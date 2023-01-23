package api

import (
	"anonichat-pgx-stdlib/models"
	"anonichat-pgx-stdlib/repos"
	"anonichat-pgx-stdlib/utils"
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	u := models.User{}
	if err := ctx.ShouldBindJSON(&u); err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	err := repos.AddUser(&u)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			utils.RetBadReq(ctx, errors.New("username atau email sudah terdaftar"))
			return
		}
		utils.RetBadReq(ctx, err)
		return
	}
	utils.RetSucc(ctx, "berhasil mendaftar akun", nil)
}

func Login(ctx *gin.Context) {
	u := models.User{}
	if err := ctx.ShouldBindJSON(&u); err != nil {
		utils.RetBadReq(ctx, err)
		return
	}

	logCount, err := repos.CheckLoginAttemp(u.Username)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	if logCount > 3 {
		utils.RetBadReq(ctx, errors.New("akun anda terkunci, mohon menghubungi admin"))
		return
	}

	token, err := repos.LoginCheck(u.Username, u.Password)
	if err != nil {
		err = repos.AddLoginAttemp(u.Username)
		if err != nil {
			log.Fatal()
		}
		utils.RetBadReq(ctx, errors.New("username atau password tidak sesuai"))

		return
	}

	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}

	t := models.TokenResp{
		Token: token,
	}

	utils.RetSucc(ctx, "berhasil login", t)

}
