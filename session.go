package main

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/hoisie/web"
	"math/rand"
	"strconv"
	"time"
)

type session struct {
	id        string
	expiresOn time.Time
}

type SessionManager struct {
	sessions []session
	users    map[string]string
}

func NewSessionManager(config *Config) *SessionManager {
	return &SessionManager{make([]session, 0), config.Users}
}

func (sm *SessionManager) SessionExists(id string) bool {
	for _, s := range sm.sessions {
		if s.id == id && time.Now().Before(s.expiresOn) {
			return true
		}
	}
	return false
}

func (sm *SessionManager) removeExpired() {
	w := 0
	for _, s := range sm.sessions {
		if time.Now().After(s.expiresOn) {
			continue
		}
		sm.sessions[w] = s
		w++
	}
	sm.sessions = sm.sessions[:w]
}

func (sm *SessionManager) LoggedIn(ctx *web.Context) bool {
	if id, ok := ctx.GetSecureCookie("TDB-user"); ok && sm.SessionExists(id) {
		return true
	}
	return false
}

func (sm *SessionManager) Login(ctx *web.Context, user, pass string) bool {
	sm.removeExpired()
	if sm.LoggedIn(ctx) {
		return true
	}
	if pHash, ok := sm.users[user]; ok && bcrypt.CompareHashAndPassword([]byte(pHash), []byte(pass)) == nil {
		s := session{makeSessionId(), time.Now().Add(2 * time.Minute)}
		sm.sessions = append(sm.sessions, s)
		ctx.SetSecureCookie("TDB-user", s.id, 120)
		return true
	}
	return false
}

func makeSessionId() string {
	return strconv.Itoa(rand.Int())
}
