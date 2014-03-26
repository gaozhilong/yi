package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"model"
	"utils"
	"fmt"
	"time"
	"ssdb"
)



func main() {
	m := martini.Classic()
	m.Use(utils.Static("static"))
	m.Use(render.Renderer())
	//m.Use(utils.DB())

	ip := "127.0.0.1"
	port := 8888
	db, err := ssdb.Connect(ip, port)
	if err != nil {
		panic(err)
	}

	m.Get("/", func(r render.Render) {
		r.HTML(200, "login", "")
	})

	m.Get("/register", func(r render.Render) {
		r.HTML(200, "signin", "")
	})

	m.Post("/login", binding.Form(model.User{}), func (user model.User, r render.Render) {
		val, err1 := db.Get(user.Name)
		if err1 != nil {
			fmt.Printf("%s\n", "用户不存在！");
		}
		usrstr := val.(string)
		m := utils.JSONstringToMap(usrstr)

		password := m["Password"]
		//os.Exit(1);
		if (user.Password == password) {
			id := time.Now()
			m["Certificate"] = id.Format("20060102150405")
			db.Set("user1", utils.MapToJSONstring((map[string]interface{})(m)));
			r.HTML(200, "index", nil)
		} else {
			r.HTML(200, "login", nil)
		}
	})

	m.Post("/register", binding.Form(model.User{}), func (user model.User, r render.Render) {
		//db.Set(user.Name,user)
		urlValues := utils.StructToMap(&user)
		db.Set(urlValues["Name"].(string), utils.MapToJSONstring((map[string]interface{})(urlValues)));
		r.Redirect("/user/"+urlValues["Name"].(string))
	})

	m.Get("/user/:name", func(r render.Render, params martini.Params) {
		val, err1 := db.Get(params["name"])
		if err1 != nil {
			fmt.Printf("%s\n", "用户不存在！");
		}
		usrstr := val.(string)
		mp := utils.JSONstringToMap(usrstr)
		r.HTML(200, "showuser", mp)
	})

	m.Run()
}
