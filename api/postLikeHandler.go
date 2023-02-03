package api

import (
	"anonichat-pgx-stdlib/models"
	"anonichat-pgx-stdlib/repos"
	"anonichat-pgx-stdlib/utils"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetLikesByPostID(ctx *gin.Context) {
	pidStr := ctx.Param("pid")
	pid, err := uuid.Parse(pidStr)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	pls, err := repos.GetLikesByPostID(pid)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	utils.RetSucc(ctx, "berhasil GetPostLikesByID", pls)
}

func AddLikeToPostID(ctx *gin.Context) {
	pl := models.PostLike{}

	pidStr := ctx.Param("pid")
	pid, err := uuid.Parse(pidStr)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	pl.PostId = pid
	uid, err := utils.ExtractTokenID(ctx)
	if err != nil {
		utils.RetBadReq(ctx, errors.New("terjadi kendala sesi login, harap login kembali"))
		return
	}
	if uid == uuid.Nil {
		utils.RetBadReq(ctx, errors.New("login tidak sah, harap login kembali"))
		return
	}
	pl.UserId = uid
	err = repos.AddLikeToPostID(pl.PostId, pl.UserId)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	utils.RetSucc(ctx, "berhasil menyukai post", nil)
}

func UnLikeFromPostID(ctx *gin.Context) {
	pl := models.PostLike{}

	pidStr := ctx.Param("pid")
	pid, err := uuid.Parse(pidStr)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	pl.PostId = pid
	uid, err := utils.ExtractTokenID(ctx)
	if err != nil {
		utils.RetBadReq(ctx, errors.New("terjadi kendala sesi login, harap login kembali"))
		return
	}
	if uid == uuid.Nil {
		utils.RetBadReq(ctx, errors.New("login tidak sah, harap login kembali"))
		return
	}
	pl.UserId = uid
	err = repos.UnLikeFromPostID(pl.PostId, pl.UserId)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	utils.RetSucc(ctx, "berhasil unlike post", nil)
}

func CountLikesPostID(ctx *gin.Context) {
	pidStr := ctx.Param("pid")
	cl := models.CountLike{}
	pid, err := uuid.Parse(pidStr)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	count, err := repos.CountLikePostID(pid)
	if err != nil {
		utils.RetBadReq(ctx, err)
		return
	}
	cl.Likes = int32(count)
	utils.RetSucc(ctx, "berhasil count likes", cl)
}
