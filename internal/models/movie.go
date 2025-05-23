package models

// Movie struct represents a movie entity
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director struct represents a director entity
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
