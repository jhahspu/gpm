package repository

import "gpm/models"

type SqliteRepository interface {
	InsertGPM(line models.GPM) error
	DisplayGPMs() (models.GPMs, error)
}
