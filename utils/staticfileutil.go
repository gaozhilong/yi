package utils

import (
	"net/http"
	"strings"
	"github.com/codegangsta/martini"
)

/*加载静态文件的工具类，返回的是 martini.Handler中间对象。 martini原有的加载静态文件的工具类使用有问题，在我的环境下无法正常的找到静态文件
Load tools static files, returns martini.Handler object.
 Tools martini original use of static files loaded in my normal environment can not find static files
*/
func Static(directory string) martini.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		file := r.URL.Path[1:]
		if (strings.HasPrefix(file, directory)) {
			http.ServeFile(w, r, r.URL.Path[1:])
		}
	}
}
