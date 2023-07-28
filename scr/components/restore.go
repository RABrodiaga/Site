package components

import (
	"fmt"
	"io/ioutil"

	"github.com/brendonmatos/golive"
)

type Restore struct {
	golive.LiveComponentWrapper
	Login    string
	Email    string
	IsForgot bool
}

func NewRestore() *golive.LiveComponent {
	return golive.NewLiveComponent("Restore", &Restore{})
}

func (t *Restore) CanRestore() bool {
	return len(t.Login) > 7 && len(t.Email) > 7
}

func (t *Restore) RestorePassword() bool {
	return true
}

func (t *Restore) ForgotPassword() {
	t.IsForgot = !t.IsForgot
}

func (t *Restore) TemplateHandler(_ *golive.LiveComponent) string {
	fileBytes, err := ioutil.ReadFile("assets/html/restore.gohtml")
	if err != nil {
		panic(fmt.Errorf("read template file: %w", err))
	}
	return string(fileBytes)
}
