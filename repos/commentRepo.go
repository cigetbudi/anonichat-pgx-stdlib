package repos

import (
	"anonichat-pgx-stdlib/db"
	"anonichat-pgx-stdlib/models"
	"anonichat-pgx-stdlib/utils"
	"errors"
	"time"
)

func AddComment(c *models.Comment) error {
	defer utils.Timer(time.Now(), "addComment")
	var err error
	res, err := db.DB.Exec("INSERT INTO comments (post_id, user_id, comment, created_at, location) VALUES ($1, $2, $3, $4, $5)", c.PostId, c.UserId, c.Comment, time.Now(), c.Location)
	if err != nil {
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

func GetCommentsFromPostID(post_id uint) (*[]models.Comment, error) {
	defer utils.Timer(time.Now(), "getCommentsFromPostID")
	var (
		err error
		cs  []models.Comment
		c   models.Comment
	)

	rows, err := db.DB.Query("SELECT id,  user_id, post_id, comment, location FROM comments WHERE post_id = $1 AND deleted_at IS NULL", post_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&c.Id, &c.UserId, &c.PostId, &c.Comment, &c.Location)
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	return &cs, nil
}

func DeleteCommentFromID(id, user_id uint) error {
	defer utils.Timer(time.Now(), "deleteCommentFromID")
	var err error
	res, err := db.DB.Exec("UPDATE comments SET deleted_at = $1 WHERE id = $2 AND user_id = $3", time.Now(), id, user_id)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row == 0 {
		return errors.New("data tidak ditemukan/gagal menghapus")
	} else if row != 1 {
		return errors.New("data terhapus ada lebih dari 1")
	}
	return nil

}
