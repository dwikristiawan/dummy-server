package auth_dto

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UpdateUsersRequest struct {
	Id       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}
type RemoveUsersRequest struct {
	Id       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}
