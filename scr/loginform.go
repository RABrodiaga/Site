package scr

import (
	"context"
	"fmt"
	"io/ioutil"
	"site/scr/components"
	"text/template"
	"time"

	"github.com/brendonmatos/golive"
)

type LoginForm struct {
	golive.LiveComponentWrapper
	Login    *golive.LiveComponent
	Restore  *golive.LiveComponent
	IsForgot bool
}

func NewLoginForm() func(ctx context.Context) *golive.LiveComponent {
	return func(ctx context.Context) *golive.LiveComponent {
		var c LoginForm
		c.Login = components.NewLogin()
		c.Restore = components.NewRestore()
		return golive.NewLiveComponent("LoginForm", &c)
	}
}

func (t *LoginForm) Mounted(l *golive.LiveComponent) {
	go func() {
		for {
			if l.Exited {
				return
			}
			if !t.IsForgot {
				t.IsForgot = t.Login.GetFieldFromPath("IsForgot").Bool()
			} else {
				t.IsForgot = !t.Restore.GetFieldFromPath("IsForgot").Bool()
			}
			if t.Login.GetFieldFromPath("IsForgot").Bool() == !t.Restore.GetFieldFromPath("IsForgot").Bool() {
				t.Login.SetValueInPath("false", "IsForgot")
				t.Restore.SetValueInPath("false", "IsForgot")
			}

			time.Sleep(time.Microsecond * 200)
			t.Commit()
		}
	}()
}

func (t *LoginForm) TemplateHandler(_ *golive.LiveComponent) string {
	fileBytes, err := ioutil.ReadFile("assets/html/loginform.gohtml")
	if err != nil {
		panic(fmt.Errorf("read template file: %w", err))
	}
	css, _ := template.ParseFiles("assets/css/loginform.css")
	page := string(fileBytes) + `<style>` + css.Root.String() + `</style>`
	return page
}