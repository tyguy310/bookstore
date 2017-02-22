package books

import "github.com/tyguy310/bookstore/config"

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func AllBooks() ([]*Book, error) {
	rows, err := config.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}
//
// func SingleBook() ([*Book], error) {
// 	row := config.DB.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)
//
// 	bk := new(Book)
// 	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
// 	if err == sql.ErrNoRows {
// 		http.NotFound(w, r)
// 		return
// 	} else if err != nil {
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}
// }
