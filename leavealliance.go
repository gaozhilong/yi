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
	"net/http"
	"os"
	"io"
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

	m.Get("/upload", binding.Form(model.User{}), func (r render.Render) {
			r.HTML(200, "uploadfile", "")
	})

	m.Post("/upload", binding.Form(model.User{}), func (w http.ResponseWriter, req *http.Request, r render.Render) {
		//db.Set(user.Name,user)
			file, header, err := req.FormFile("file")
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			defer file.Close()

			out, err := os.Create("D:/GO/temp/"+header.Filename)
			if err != nil {
				fmt.Fprintf(w, "Failed to open the file for writing")
				return
			}
			defer out.Close()
			_, err = io.Copy(out, file)
			if err != nil {
				fmt.Fprintln(w, err)
			}

			// the header contains useful info, like the original file name
			fmt.Fprintf(w, "File %s uploaded successfully.", header.Filename)
			//r.Redirect("/user")
	})

	m.Run()
}
