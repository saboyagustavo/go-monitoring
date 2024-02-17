package database

import (
	"database/sql"

	"github.com/saboyagustavo/go-monitoring/internal/entity"
)

type ResourceDB struct {
	db *sql.DB
}

func NewResourceDB(db *sql.DB) *ResourceDB {
	return &ResourceDB{db: db}
}

func (rDB *ResourceDB) GetVersion() (string, error) {
	var version string
	err := rDB.db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		return "", err
	}

	return version, nil

}
func (rDB *ResourceDB) GetResources() ([]*entity.Resource, error) {
	rows, err := rDB.db.Query(`
		SELECT 
			id, 
			url,
			status_code
		FROM resources
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var resources []*entity.Resource

	for rows.Next() {
		var resource entity.Resource

		err := rows.Scan(&resource.ID, &resource.URL, &resource.StatusCode)
		if err != nil {
			return nil, err
		}

		resources = append(resources, &resource)
	}

	return resources, nil
}

func (rDB *ResourceDB) GetResource(id string) (*entity.Resource, error) {
	var resource entity.Resource
	err := rDB.db.QueryRow(`
		SELECT 
			id, 
			url,
			status_code
		FROM resources
		WHERE id = ?
	`, id,
	).Scan(
		&resource.ID,
		&resource.URL,
		&resource.StatusCode,
	)

	if err != nil {
		return nil, err
	}

	return &resource, nil
}

func (rDB *ResourceDB) SaveResource(resource *entity.Resource) (string, error) {
	_, err := rDB.db.Exec(`
		INSERT INTO resources (
			id, url, status_code
		) VALUES (
			?, ?, ?
		)
	`,
		resource.ID,
		resource.URL,
		resource.StatusCode,
	)

	if err != nil {
		return "", err
	}

	return resource.ID, nil
}

func (rDB *ResourceDB) DeleteResource(id string) (string, error) {
	_, err := rDB.db.Query(`
		DELETE 
		FROM resources
		WHERE id = ?
	`, id,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}