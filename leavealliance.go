package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/binding"
	"ssdb"
	"utils"
	"fmt"
	"os"
	"time"
)

type User struct {
	Name		string `form:"name"`
	Password	string `form:"password"`
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	ip := "127.0.0.1"
	port := 8888
	db, err := ssdb.Connect(ip, port)
	if err != nil {
		os.Exit(1)
	}

	defer db.Close()

	m.Get("/", func(r render.Render) {
		r.HTML(200, "login", nil)
	})

	m.Post("/login", binding.Form(User{}), func(user User, r render.Render) {
		//db.Set("usertest",`{"name":"测试用户","password":"password"}`)
		val, err1 := db.Get(user.Name)
		if err1 != nil {
			fmt.Printf("%s\n", "用户不存在！");
		}
		usrstr := val.(string)
		m := utils.JSONstringToMap(usrstr)

		password := m["password"]
		//os.Exit(1);
		if (user.Password == password) {
			id := time.Now()
			m["certificate"] = id.Format("20060102150405")
			db.Set("user1", utils.MapToJSONstring((map[string]interface{})(m)));
			r.HTML(200, "index", nil)
		} else {
			r.HTML(200, "login", nil)
		}
	})

	m.Run()
}
