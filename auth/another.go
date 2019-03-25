package auth

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/filebrowser/filebrowser/v2/settings"
	"github.com/filebrowser/filebrowser/v2/users"
)

const MethodAnotherAuth settings.AuthMethod = "another"

type anotherCred struct {
	Password  string `json:"password"`
	Username  string `json:"username"`
	ReCaptcha string `json:"recaptcha"`
}

type AnotherAuth struct {
}

func (a AnotherAuth) Auth(r *http.Request, sto *users.Storage, root string) (*users.User, error) {
	var cred anotherCred

	if r.Body == nil {
		return nil, os.ErrPermission
	}

	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		return nil, os.ErrPermission
	}

	u := &users.User{}
	u.Username = cred.Username
	u.Password = cred.Password

	err = u.Clean(root)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// LoginPage tells that json auth doesn't require a login page.
func (a AnotherAuth) LoginPage() bool {
	return true
}
