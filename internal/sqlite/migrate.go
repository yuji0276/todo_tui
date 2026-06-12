package sqlite

import (
	"database/sql"
	"embed"
	"strconv"
	"strings"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func Migrate(db *sql.DB) error {
	//migrationsテーブルをなければ作成
	_, err := db.Exec("create table if not exists migrations(version integer)")
	if err != nil {
		return err
	}
	//migrationsテーブルからversionを取得してマイグレーション済みにする
	rows, err := db.Query("select version from migrations")
	if err != nil {
		return err
	}
	defer rows.Close()
	versions := make(map[int]bool)
	for rows.Next() {
		var version int
		if err := rows.Scan(&version); err != nil {
			return err
		}
		versions[version] = true
	}
	if err = rows.Err(); err != nil {
		return err
	}

	entries, err := migrationsFS.ReadDir("migrations")
	if err != nil {
		return err
	}
	for _, entry := range entries {
		fileName := entry.Name()
		fileName = "migrations/" + fileName
		data, err := migrationsFS.ReadFile(fileName)
		if err != nil {
			return err
		}
		content := string(data)
		parts := strings.Split(entry.Name(), "_")
		versionNum, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}

		if versions[versionNum] == true {
			continue
		} else {
			_, err := db.Exec(content)
			if err != nil {
				return err
			}
			_, err = db.Exec("insert into migrations (version) values(?)", versionNum)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
