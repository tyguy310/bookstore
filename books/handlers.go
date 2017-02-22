package books

import (
	"fmt"
	"net/http"
)

func BooksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}

// func BooksShow(w http.ResponseWriter, r *http.Request) {
//   if r.Method != "GET" {
//     http.Error(w, http.StatusText(405), 405)
//     return
//   }
//
//   isbn := r.FormValue("isbn")
//   bk, err := SingleBook()
//   if isbn == "" {
//     http.Error(w, http.StatusText(400), 400)
//     return
//   }
//
//   fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
//
// }
