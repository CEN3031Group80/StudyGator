package StudyGroup

import(
	"github.com/go-pkgz/auth/token"
)

type Group struct {
	ID string
	Name string
	Users []token.User
}