package pkg

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"os"
	"time"
)

var sqliteDB = os.Getenv("SQLITE_DB")
var selectQuery = os.Getenv("SELECT_QUERY")

type Repository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewRepository(logger *zap.Logger) (*Repository, error) {
	db, err := sql.Open("sqlite3", sqliteDB)
	if err != nil {
		logger.Error("error::NewRepository:18", zap.Error(err))
		return nil, err
	}

	return &Repository{DB: db, Logger: logger}, nil
}

func (r *Repository) FetchData() ([]string, error) {
	data := make([]string, 0)

	rows, err := r.DB.Query(selectQuery)
	if err != nil {
		r.Logger.Error("error::FetchData:29", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			r.Logger.Error("error::FetchData:37", zap.Error(err))
			return nil, err
		}
		data = append(data, value)
	}

	if err := rows.Err(); err != nil {
		r.Logger.Error("error::FetchData:44", zap.Error(err))
		return nil, err
	}

	r.Logger.Info("Data fetched successfully", zap.Int("Count", len(data)))
	return data, nil
}

func (r *Repository) FetchIDs() (map[string]bool, error) {
	ids := make(map[string]bool)

	rows, err := r.DB.Query("SELECT id FROM vehicle_list")
	if err != nil {
		r.Logger.Error("error::FetchIDs:57", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			r.Logger.Error("error::FetchIDs:65", zap.Error(err))
			return nil, err
		}
		ids[id] = true
	}

	if err := rows.Err(); err != nil {
		r.Logger.Error("error::FetchIDs:72", zap.Error(err))
		return nil, err
	}

	r.Logger.Info("IDs fetched successfully", zap.Int("Count", len(ids)))
	return ids, nil
}

func (r *Repository) Insert(id, url string) error {
	_, err := r.DB.Exec("INSERT INTO vehicle_list (id, url, posting_time) VALUES (?, ?, ?)", id, url, time.Now())
	if err != nil {
		r.Logger.Error("error::Insert:82", zap.String("ID", id), zap.String("URL", url), zap.Error(err))
		return err
	}

	r.Logger.Info("Record inserted successfully", zap.String("URL", url))
	return nil
}
