package skill

import "database/sql"

type skillRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) skillRepository {
	return skillRepository{db: db}
}

func (repository *skillRepository) getSkillByKey(key string) *sql.Row {
	sql := "SELECT key, name, description, logo, tags FROM skill where key=$1"
	return repository.db.QueryRow(sql, key)
}

func (repository *skillRepository) getAllSkill() (*sql.Rows, error) {
	rows, err := repository.db.Query("SELECT key, name, description, logo, tags FROM skill ORDER BY key")
	return rows, err
}
