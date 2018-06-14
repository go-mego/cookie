package cookie

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-mego/mego"
)

// ErrNoCookie is returned by Request's Cookie method when a cookie is not found.
var ErrNoCookie = errors.New("cookie: named cookie not present")

func New() mego.HandlerFunc {
	return func(c *mego.Context) {
		c.Map(&Jar{
			context: c,
		})
		c.Next()
	}
}

type Cookie struct {
	Name  string
	Value string

	Path    string    // optional
	Domain  string    // optional
	Expires time.Time // optional

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HTTPOnly bool
}

type Jar struct {
	context *mego.Context
}

func (j *Jar) Set(cookie *Cookie) {
	http.SetCookie(j.context.Writer, &http.Cookie{
		Name:     cookie.Name,
		Value:    cookie.Value,
		Path:     cookie.Path,
		Domain:   cookie.Domain,
		Expires:  cookie.Expires,
		MaxAge:   cookie.MaxAge,
		Secure:   cookie.Secure,
		HttpOnly: cookie.HTTPOnly,
	})
}

func (j *Jar) Has(name string) bool {
	_, err := j.Get(name)
	if err == http.ErrNoCookie {
		return false
	}
	return true
}

func (j *Jar) Get(name string) (string, error) {
	cookie, err := j.context.Request.Cookie(name)
	if err != nil {
		if err == http.ErrNoCookie {
			return "", ErrNoCookie
		}
		return "", err
	}
	return cookie.Value, nil
}

func (j *Jar) Delete(name string) error {
	cookie, err := j.context.Request.Cookie(name)
	if err != nil {
		if err == http.ErrNoCookie {
			return ErrNoCookie
		}
		return err
	}
	http.SetCookie(j.context.Writer, &http.Cookie{
		Name:     cookie.Name,
		Value:    cookie.Value,
		Path:     cookie.Path,
		Domain:   cookie.Domain,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		Secure:   cookie.Secure,
		HttpOnly: cookie.HttpOnly,
	})
	return nil
}
