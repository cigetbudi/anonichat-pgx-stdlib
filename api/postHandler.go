package api

import (
	"anonichat-pgx-stdlib/models"
	"anonichat-pgx-stdlib/repos"
	"anonichat-pgx-stdlib/utils"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPosts(ctx *gin.Context) {
	mes, err := repos.GetAllPost()
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	utils.RetSucc(ctx, "berhasil GetAllPosts", mes)
}

func CreatePost(ctx *gin.Context) {
	var err error
	p := models.Post{}
	if err := ctx.ShouldBindJSON(&p); err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	uid, err := utils.ExtractTokenID(ctx)
	if err != nil {
		utils.RetBadReq(ctx, errors.New("login tidak sah, harap login kembali"))
		return
	}
	p.UserID = int64(uid)
	err = repos.CreatePost(&p)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	utils.RetSucc(ctx, "berhasil membuat post", nil)
}

func DeletePost(ctx *gin.Context) {
	var err error
	user_id, err := utils.ExtractTokenID(ctx)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	err = repos.DeletePost(uint(id), user_id)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	utils.RetSucc(ctx, "berhasil menghapus post", nil)
}
