package entity

type User struct {
	Id             int    `db:"id"`
	PassportNumber string `db:"passport_number"`
	Name           string `db:"name"`
	Surname        string `db:"surname"`
	Patronymic     string `db:"patronymic"`
	Address        string `db:"address"`
}
