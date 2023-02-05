package api

import (
	"anonichat-pgx-stdlib/repos"
	"anonichat-pgx-stdlib/utils"

	"github.com/gin-gonic/gin"
)

func GetAllGenders(ctx *gin.Context) {
	mes, err := repos.GetAllGenders()
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	utils.RetSucc(ctx, "berhasil GetAllGenders", mes)
}
