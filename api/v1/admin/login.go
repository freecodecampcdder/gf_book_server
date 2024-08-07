package admin

// for gtoken
type LoginRes struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
	IsAdmin  int    `json:"is_admin"` //是否超管
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
}
