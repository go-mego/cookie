# Cookie

Cookie 套件可以讓開發者針對客戶端的瀏覽器 Cookie 進行操作。

# 索引

* [安裝方式](#安裝方式)
* [使用方式](#使用方式)

# 安裝方式

打開終端機並且透過 `go get` 安裝此套件即可。

```bash
$ go get github.com/go-mego/cookie
```

# 使用方式

先將 `cookie.Cooker` 傳入 Mego 的 `Use` 來初始化一個餅乾廚師後方能於路由中使用 Cookie 的所有功能。

```go
package main

import (
	"github.com/go-mego/cookie"
	"github.com/go-mego/mego"
)

func main() {
	m := mego.New()
	m.Use(cookie.Cooker())
	m.Get("/set", func(j *cookie.Jar) {
        // 設置一個名為 `myCookie` 的 Cookie 至客戶端瀏覽器。
		j.Set(cookie.Cookie{
			Key:   "myCookie",
			Value: "我要開動了！",
		})
	})
	m.Get("/get", func(j *cookie.Jar) string {
        // 取得客戶端瀏覽器中的 `myCookie` 內容。
		return j.Get("myCookie")
	})
	m.Run()
}
```