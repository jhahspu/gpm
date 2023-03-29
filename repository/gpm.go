package repository

import "gpm/models"

func (s *sqliteRepository) InsertGPM(line models.GPM) error {
	insertRow := `INSERT INTO gpm(website, username, email, password, notes) VALUES (?, ?, ?, ?, ?)`
	statement, err := s.db.Prepare(insertRow)
	if err != nil {
		return err
	}
	_, err = statement.Exec(line.Website, line.Username, line.Email, line.Password, line.Notes)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqliteRepository) DisplayGPMs() (models.GPMs, error) {
	var gpms models.GPMs

	rows, err := s.db.Query("SELECT * FROM gpm ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var gpm models.GPM
		err := rows.Scan(&gpm.ID, &gpm.Website, &gpm.Username, &gpm.Email, &gpm.Password, &gpm.Notes)
		if err != nil {
			return nil, err
		}
		gpms = append(gpms, gpm)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return gpms, nil
}
