package main

import (
	"net/http"

	"github.com/go-mego/cookie"
	"github.com/go-mego/mego"
)

func main() {
	e := mego.Default()
	e.GET("/", cookie.New(), func(c *mego.Context, j *cookie.Jar) {
		j.Set(&cookie.Cookie{
			Name:  "foo",
			Value: "bar",
		})
		v, err := j.Get("foo")
		if err == cookie.ErrNoCookie {
			c.String(http.StatusOK, "No cookie")
			return
		}
		c.String(http.StatusOK, "Cookie: %+v", v)
	})
	e.Run()
}
