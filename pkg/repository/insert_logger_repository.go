package repository

import (
	"test-stockbit/pkg/model"
)

func (repo *repository) InsertLoggerRepository(logger *model.Logger) error {
	query := `INSERT INTO test.logger	
	(transport, logging, ` + "`level`" + `, created_at, updated_at)
	VALUES("", ?, ?, NOW(), NOW());`

	_, err := repo.db.Exec(query, logger.Logging, logger.Level)
	if err != nil {
		repo.log.Error(err)
		return err
	}
	return nil
}
