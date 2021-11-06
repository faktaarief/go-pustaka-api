package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(id int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(id int, book BookRequest) (Book, error)
	Delete(id int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindById(id int) (Book, error) {
	book, err := s.repository.FindById(id)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(id int, bookRequest BookRequest) (Book, error) {
	book, _ := s.repository.FindById(id)

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Rating = int(rating)

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(id int) (Book, error) {
	book, _ := s.repository.FindById(id)
	delBook, err := s.repository.Delete(book)

	return delBook, err
}
