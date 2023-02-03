package repos

import (
	"anonichat-pgx-stdlib/db"
	"anonichat-pgx-stdlib/models"
	"anonichat-pgx-stdlib/utils"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GetLikesByPostID(postId uuid.UUID) (*[]models.PostLike, error) {
	defer utils.Timer(time.Now(), "getLikesByPostID")
	var pls []models.PostLike
	rows, err := db.DBN.Query("SELECT id, post_id, user_id, created_at FROM post_likes WHERE post_id = $1", postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var pl models.PostLike
		err = rows.Scan(&pl.Id, &pl.PostId, &pl.UserId, &pl.CreatedAt)
		if err != nil {
			return nil, err
		}
		pls = append(pls, pl)
	}
	return &pls, nil
}

func AddLikeToPostID(post_id, user_id uuid.UUID) error {
	defer utils.Timer(time.Now(), "addLikeToID")
	var err error
	res, err := db.DBN.Exec("INSERT INTO post_likes (post_id, user_id, created_at) VALUES ($1, $2, $3)", post_id, user_id, time.Now())
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return errors.New("anda sudah menyukai post ini sebelumnya")
		}
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row != 1 {
		return errors.New("data terinsert ada lebih dari 1")
	}
	return nil
}

func UnLikeFromPostID(post_id, user_id uuid.UUID) error {
	defer utils.Timer(time.Now(), "UnLikeFromID")
	var err error
	res, err := db.DBN.Exec("DELETE FROM post_likes where post_id = $1 AND user_id = $2", post_id, user_id)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("data tidak ditemukan")
	} else if row != 1 {
		return errors.New("data terhapus ada lebih dari 1")
	}
	return nil
}

func CountLikePostID(post_id uuid.UUID) (int, error) {
	defer utils.Timer(time.Now(), "CountLikePostID")
	var count int
	err := db.DBN.QueryRow("SELECT COUNT (*) FROM post_likes WHERE post_id = $1", post_id).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
