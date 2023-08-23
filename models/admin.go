package models

type Administrator struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`  // Por seguridad, no devolver la contraseña en las respuestas al frontend
	RoleID   int    `json:"role_id"`
}