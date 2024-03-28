package gqlcontext

import (
	"context"
	"net/http"
    "crypto/rand"
    "encoding/base64"

	"github.com/go-pkgz/auth/token"
	"study-gator-backend/graph/StudyGroup"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user_sg"}

type contextKey struct {
	name string
}

func GenerateID() (string, error) {
    byteSize := 16
    randomBytes := make([]byte, byteSize)
    _, err := rand.Read(randomBytes)
    if err != nil {
        return "", err
    }
    ID := base64.URLEncoding.EncodeToString(randomBytes)
    return ID, nil
}

var userDB = make(map[string]token.User)
var groupDB = make(map[string]StudyGroup.Group)

func InsertUser(ID string, user token.User) bool {
	_, exists := userDB[ID]
	if (exists) {
		return false
	}
	userDB[ID] = user
	return true
}

func MakeGroup(n string) StudyGroup.Group {
	var newID string
	for true {
		newID, _ = GenerateID()
		newID = "group_"+newID
		_, exists := groupDB[newID]
		if (!exists) {
			break
		}
	}
	newGroup := StudyGroup.Group{
		ID: newID,
		Name: n,
		Users: make([]token.User, 0)}
	groupDB[newID] = newGroup
	return newGroup
}

func GetUser(ID string) (token.User, bool) {
	a, b := userDB[ID]
	return a, b
}

func GetUserDB() map[string]token.User {
	return userDB
}

func GetGroup(ID string) (StudyGroup.Group, bool) {
	a, b := groupDB[ID]
	return a, b
}

func GetGroupDB() map[string]StudyGroup.Group {
	return groupDB
}

// Auth Middleware decodes the share session cookie and packs the session into context
func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := token.GetUserInfo(r)

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)
			InsertUser(user.ID, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func UserFromContext(ctx context.Context) token.User {
	raw, _ := ctx.Value(userCtxKey).(token.User)
	return raw
}
