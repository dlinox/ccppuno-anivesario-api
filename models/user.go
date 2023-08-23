package models

type User struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"` // Nota: En un escenario real, asegura esta contraseña con un hash.
}