package model

import (

)

type User struct {
	Name		 string `form:"name"`
	Password	 string `form:"password"`
	//RePassword	 string `form:"repassword"`
}
