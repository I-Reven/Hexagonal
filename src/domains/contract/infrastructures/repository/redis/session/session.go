package session

import "github.com/gin-contrib/sessions"

type Session interface {
	Store() sessions.Store
}
