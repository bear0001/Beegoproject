package models



//构建用户结构体，用于不表示数据
//type User struct {
//	User     string `json:"name"`
//	Birthday string`json:"birthday"`
//	Address  string`json:"address"`
//	Nick     string`json:"nick"`
//}

type User struct {
	User string `from:"name"`
	Birthday string`form:"birthday"`
	Address string`from:"address"`
	Nick string`from:"nick"`

}