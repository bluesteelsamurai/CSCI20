// Lab 15 -- structs/methods
//
// Programmer name: Nigel Adams
// Date completed:  12/3/19

package main
import f"fmt"

// Movie represents a movie for a database of movies.
type Movie struct {
	Title    string
	Genre    string
	Likes    int
	Dislikes int
	Views    int
}

// NewMovie returns a Movie with the given title and genre.
// Views, Likes, and Dislikes all initialized to 0.
func NewMovie(title, genre string) *Movie {
	return &Movie{title,genre,0,0,0}
}

// AddLikeView increments Views and Likes.
func (m *Movie) AddLikeView() {
	m.Views++
	m.Likes++
}

// AddDislikeView increments Views and Dislikes.
func (m *Movie) AddDislikeView() {
	m.Views++
	m.Dislikes++
}

// GetPercentageLikes returns the percentage of likes against
// the Views for the Movie. The value is between 0.0 and 100.0.
func (m *Movie) GetPercentageLikes() float64 {
	if m.Views<=0{
		return 0.0
	}else{
	return (float64(m.Likes)/float64(m.Views)*100)
	}
}

// GetPercentageDislikes returns the percentage of dislikes against
// the Views for the Movie. The value is between 0.0 and 100.0.
func (m *Movie) GetPercentageDislikes() float64 {
	if m.Views<=0{
		return 0.0
	}else{
	return (float64(m.Dislikes)/float64(m.Views)*100)
	}
}

// ToString returns a string representation of the Movie.
//
// TEMPLATE: title, genre, __ views, __.__% likes
// EXAMPLE:  Godzilla, action, 100 views, 80.00% likes
func (m *Movie) ToString() string {
	return f.Sprintf("%s, %s, %d views, %.2f%% likes",m.Title,m.Genre,m.Views,m.GetPercentageLikes())
}

func main() {

}
