package session

import (
	"fmt"
	"sync"
	"time"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/utils"
)

var sessionMap *sync.Map //支持线程安全的 map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionsFromDB() *sync.Map{
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return nil
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
	return sessionMap
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := time.Now().UnixNano() / 1e6 // ms
	ttl := ct + 30 * 60 * 1000 //serverside session valid time: 30 min
	ss := &defs.SimpleSession{Username: un, TTL:ttl}
	sessionMap.Store(id, ss)
	err := dbops.InsertSession(id, ttl, un)
	if err != nil {
		return fmt.Sprintf("Error of GenerateNewSessionId: %s", err)
	}
	return id
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := time.Now().UnixNano() / 1e6
		if ss.(*defs.SimpleSession).TTL < ct {
			// delete expired session
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true
}


