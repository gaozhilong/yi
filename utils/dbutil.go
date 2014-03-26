package utils

import (
	"github.com/codegangsta/martini"
	"ssdb"
	"reflect"
	"strconv"
)

// DB Returns a martini.Handler
func DB() martini.Handler {
	ip := "127.0.0.1"
	port := 8888
	db, err := ssdb.Connect(ip, port)
	if err != nil {
		panic(err)
	}

	return func(c martini.Context) {
		c.Map(db)
		defer db.Close()
		c.Next()
	}
}

func StructToMap(i interface{}) map[string]interface{} {
	values := make(map[string]interface{})
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		var v string
		switch f.Interface().(type) {
			case int, int8, int16, int32, int64:
				v = strconv.FormatInt(f.Int(), 10)
			case uint, uint8, uint16, uint32, uint64:
				v = strconv.FormatUint(f.Uint(), 10)
			case []byte:
				v = string(f.Bytes())
			case string:
				v = f.String()
		}
		values[typ.Field(i).Name] = interface {}(v)
	}
	return values
}
