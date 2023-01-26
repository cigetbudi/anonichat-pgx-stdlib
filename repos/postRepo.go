package repos

import (
	"anonichat-pgx-stdlib/db"
	"anonichat-pgx-stdlib/models"
	"anonichat-pgx-stdlib/utils"
	"errors"
	"time"
)

func GetAllPost() (*[]models.Post, error) {
	defer utils.Timer(time.Now(), "getAllPost")
	var ps []models.Post
	rows, err := db.DB.Query("SELECT id, content, location, user_id FROM posts WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT 50")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p models.Post
		err = rows.Scan(&p.Id, &p.Content, &p.Location, &p.UserID)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	return &ps, nil
}

func CreatePost(p *models.Post) error {
	defer utils.Timer(time.Now(), "createPost")
	var err error
	res, err := db.DB.Exec("INSERT INTO posts (content, location, created_at, user_id) values ($1, $2, $3, $4)", p.Content, p.Location, time.Now(), p.UserID)
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

func DeletePost(pid, userId uint) error {
	defer utils.Timer(time.Now(), "deletePost")
	var err error
	res, err := db.DB.Exec("UPDATE posts SET deleted_at = $1 WHERE id = $2 AND user_id = $3", time.Now(), pid, userId)
	if err != nil {
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if row != 1 {
		return errors.New("data terupdate ada lebih dari 1")
	}
	return nil
}
