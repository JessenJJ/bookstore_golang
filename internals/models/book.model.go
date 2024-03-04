package models

type BookModel struct {
	Id          int     `db:"Id"`
	Title       string  `db:"Title" json:"Title" form:"Title"`
	Description *string `db:"Description" json:"Description" form:"Description"`
	Author      string  `db:"Author" json:"Author" form:"Author"`
	Picture     *string `db:"Picture" json:"Picture" form:"Picture"`
	// diberikan * pada Description dan Picture agar ketika tIdak ada isinya maka akan bernilai nil, karena zero value pointer adalan nil
}
