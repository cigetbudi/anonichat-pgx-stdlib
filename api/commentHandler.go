package api

import (
	"anonichat-pgx-stdlib/models"
	"anonichat-pgx-stdlib/repos"
	"anonichat-pgx-stdlib/utils"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddComment(ctx *gin.Context) {
	var err error
	c := models.Comment{}
	if err := ctx.ShouldBindJSON(&c); err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	uid, err := utils.ExtractTokenID(ctx)
	if err != nil {
		utils.RetBadReq(ctx, errors.New("login tidak sah, harap login kembali"))
		return
	}

	idStr := ctx.Param("pid")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	c.UserId = int64(uid)
	c.PostId = int64(id)
	err = repos.AddComment(&c)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	utils.RetSucc(ctx, "berhasil menambah komentar", nil)
}

func GetCommentsFromPostID(ctx *gin.Context) {
	var err error
	pidStr := ctx.Param("pid")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	cs, err := repos.GetCommentsFromPostID(uint(pid))
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	utils.RetSucc(ctx, "berhasil getCommentsByPostID", cs)
}

func DeleteCommentFromID(ctx *gin.Context) {
	var err error
	cidStr := ctx.Param("cid")
	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	uid, err := utils.ExtractTokenID(ctx)
	if err != nil {
		utils.RetBadReq(ctx, errors.New("login tidak sah, harap login kembali"))
		return
	}
	err = repos.DeleteCommentFromID(uint(cid), uid)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	utils.RetSucc(ctx, "berhasil menghapus comment", nil)

}
