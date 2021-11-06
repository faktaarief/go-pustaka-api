package helper

import "faktaarief/go-pustaka-api/book"

func BookReponse(b *book.Book) book.BookResponse {
	response := book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
	}

	return response
}

func BookReponses(books []book.Book) []book.BookResponse {
	var bookResponses []book.BookResponse

	for _, b := range books {
		response := book.BookResponse{
			ID:          b.ID,
			Title:       b.Title,
			Description: b.Description,
			Price:       b.Price,
			Rating:      b.Rating,
		}

		bookResponses = append(bookResponses, response)
	}

	return bookResponses
}
