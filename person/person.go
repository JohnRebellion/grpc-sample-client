package person

type PersonJSON struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}
