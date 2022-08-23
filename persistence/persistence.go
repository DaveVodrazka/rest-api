package persistence

import "github.com/DaveVodrazka/rest-api/models"

var DB = models.BooksDB{Hp1, Hp2}

var Jkr = models.Author{
	ID:          "a1",
	FirstName:   "JK",
	LastName:    "Rowling",
	Nationality: "GB",
	Born:        1965,
}

var Hp1 = models.Book{
	ID:     "1",
	Isbn:   "9078",
	Title:  "Harry Potter and Sorcerer's stone",
	Author: &Jkr,
}

var Hp2 = models.Book{
	ID:     "2",
	Isbn:   "9079",
	Title:  "Harry Potter and Chamber of Secrets",
	Author: &Jkr,
}
