package following

import (
	"database/sql"
)

type Repository interface {
	Store(following Following) error
}

// ------------------- connect to database -------------------
type MysqlRepository struct {
	sqlDB *sql.DB
}

func NewMysqlRepository(sqlDB *sql.DB) Repository {
	return MysqlRepository{sqlDB: sqlDB}
}

func (sql MysqlRepository) Store(following Following) error {
	stmt, err := sql.sqlDB.Prepare("INSERT INTO following (username, full_name) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(following.Username, following.FullName); err != nil {
		return err
	}
	return nil
}

// ------------------- service  -------------------
type Service struct {
	followingRepository Repository
}

func NewService(followingRepository Repository) Service {
	return Service{followingRepository}
}

func (service Service) Insert(username, fullName string) error {
	err := service.followingRepository.Store(NewFollowing(username, fullName))
	if err != nil {
		return err
	}
	return nil
}
