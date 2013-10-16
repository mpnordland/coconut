package main

import (
	"github.com/hoisie/web"
	"math/rand"
	"strconv"
    "code.google.com/p/go.crypto/bcrypt"
)

type SessionManager struct {
	sessions []string
	users    map[string]string
}

func NewSessionManager(config *Config) *SessionManager {
	return &SessionManager{make([]string, 0), config.Users}
}

func (sm *SessionManager) SessionExists(id string) bool {
	for _, i := range sm.sessions {
		if i == id {
			return true
		}
	}
	return false
}

func (sm *SessionManager) LoggedIn(ctx *web.Context) bool {
	if id, ok := ctx.GetSecureCookie("TDB-user"); ok && sm.SessionExists(id) {
		return true
	}
	return false
}

func (sm *SessionManager) CreateSession(ctx *web.Context, user, pass string) bool {
	if pHash, ok := sm.users[user]; ok && bcrypt.CompareHashAndPassword([]byte(pHash), []byte(pass)) != nil {
		id := makeSessionId()
		sm.sessions = append(sm.sessions, id)
		ctx.SetSecureCookie("TDB-user", id, 60)
		return true
	}
	return false
}

func makeSessionId() string {
	return strconv.Itoa(rand.Int())
}
