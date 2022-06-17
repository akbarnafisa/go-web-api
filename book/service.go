package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(ID int, book BookRequest) (Book, error)
	Delete(ID int) (Book, error)
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

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	b, err := s.repository.FindById(ID)

	price, _ := bookRequest.Price.Int64()

	b.Title = bookRequest.Title
	b.Price = int(price)
	b.Description = bookRequest.Description
	b.Rating = bookRequest.Rating

	newBook, err := s.repository.Update(b)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	b, _ := s.repository.FindById(ID)
	deletedBook, err := s.repository.Delete(b)
	return deletedBook, err
}
