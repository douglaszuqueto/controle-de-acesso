package models

import "time"

// Log log model
type Log struct {
	ID        string
	Device    string
	User      string
	Tag       string
	State     int
	CreatedAt time.Time
}

// Logs logs
type Logs []Log

// AllLogs allLogs
func AllLogs() Logs {
	rows, err := Query("SELECT * FROM sistema.history")
	defer rows.Close()

	checkErr(err)

	logs := Logs{}
	log := Log{}

	for rows.Next() {
		err = rows.Scan(&log.ID, &log.Device, &log.User, &log.Tag, &log.State, &log.CreatedAt)
		checkErr(err)

		logs = append(logs, log)
	}

	return logs
}

// CreateLog createLog
func CreateLog(p *Ping) error {

	stmt, err := db.Prepare(`INSERT INTO sistema.history (device, "user", tag, state) VALUES ($1,$2,$3,$4)`)
	checkErr(err)

	_, err = stmt.Exec(p.Device, p.User, p.Tag, p.State)
	defer stmt.Close()

	if err != nil {
		panic(err)
	}

	return nil
}
