package database

import "github.com/suite911/query911/query"

func Register(email string) error {
	q := query.Query{ DB: DB() }
	q.SQL = `INSERT OR IGNORE INTO "RegisteredUsers"("email") VALUES(?);`
	q.Exec(email)
	if !q.OK() {
		return q.LastError()
	}
	return nil
}
