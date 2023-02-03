package api

import (
	"anonichat-pgx-stdlib/models"
	"anonichat-pgx-stdlib/repos"
	"anonichat-pgx-stdlib/utils"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllPosts(ctx *gin.Context) {
	mes, err := repos.GetAllPost()
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	utils.RetSucc(ctx, "berhasil GetAllPosts", mes)
}

func GetAllPostsByUserID(ctx *gin.Context) {
	uidStr := ctx.Param("user_id")
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	mes, err := repos.GetAllPostByUserID(uid)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	utils.RetSucc(ctx, "berhasil GetAllPosts by UserID", mes)
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
	p.UserID = uid
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
	idStr := ctx.Param("pid")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.RetBadReq(ctx, errors.New("post_id tidak valid"))
		return
	}
	err = repos.DeletePost(id, user_id)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	utils.RetSucc(ctx, "berhasil menghapus post", nil)
}
