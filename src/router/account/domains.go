package account

type User struct {
	Name string `json:"name" form:"name" query:"name"`
	Age  int    `json:"age" form:"age" query:"age"`
}
