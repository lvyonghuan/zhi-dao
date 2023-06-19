package model

type User struct {
	Id            int    `json:"id"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Avatar        string `json:"avatar"`        //头像，一串url链接。现阶段懒得做。
	Administrator bool   `json:"administrator"` //管理员权限，后台授予。现在还没打算做，放这里是为了将来打算做的时候做。大嘘。
}
