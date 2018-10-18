package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"video_server/api/defs"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare(`INSERT INTO sessions (session_id, TTL, login_name) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id=?")
	if err != nil {
		return nil, err
	}
	var ttl, uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}
	defer stmtOut.Close()

	var ttlint int64
	if ttlint, err = strconv.ParseInt(ttl, 10, 64); err != nil {
		return nil, err
	}
	ss.TTL = ttlint
	ss.Username = uname
	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, ttlstr, login_name string
		if err = rows.Scan(&id, &ttlstr, &login_name); err != nil {
			log.Printf("retrieve sessions error: %s", err)
			break
		}
		var ttl int64
		if ttl, err = strconv.ParseInt(ttlstr, 10, 64); err != nil {
			log.Printf("parse TTL error: %s", err)
			break
		}
		ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
		m.Store(id, ss)
		log.Printf("session id: %s, ttl: %d", id, ss.TTL)
	}
	return m, nil

}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id=?")
	if err != nil {
		return err
	}
	if _, err = stmtOut.Query(sid); err != nil {
		return err
	}
	defer stmtOut.Close()
	return nil
}