package components

import (
	"fmt"
	"io/ioutil"

	"github.com/brendonmatos/golive"
)

type Login struct {
	golive.LiveComponentWrapper
	Login    string
	Password string
	IsForgot bool
}

func NewLogin() *golive.LiveComponent {
	return golive.NewLiveComponent("Login", &Login{})
}

func (t *Login) CanLogin() bool {
	return len(t.Login) > 7 && len(t.Password) > 7
}

func (t *Login) ForgotPassword() {
	t.IsForgot = !t.IsForgot
	t.Commit()
}
func (t *Login) TemplateHandler(_ *golive.LiveComponent) string {
	fileBytes, err := ioutil.ReadFile("assets/html/login.gohtml")
	if err != nil {
		panic(fmt.Errorf("read template file: %w", err))
	}
	return string(fileBytes)
}
