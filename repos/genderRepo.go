package repos

import (
	"anonichat-pgx-stdlib/db"
	"anonichat-pgx-stdlib/models"
	"anonichat-pgx-stdlib/utils"
	"time"
)

func GetAllGenders() (*[]models.Gender, error) {
	defer utils.Timer(time.Now(), "getAllGender")
	var gs []models.Gender

	rows, err := db.DBN.Query("SELECT gender_code, gender_name from genders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var g models.Gender
		err = rows.Scan(&g.GenderCode, &g.GenderName)
		if err != nil {
			return nil, err
		}
		gs = append(gs, g)
	}
	return &gs, nil
}
