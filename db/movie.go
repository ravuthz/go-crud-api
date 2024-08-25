package db

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func CreateMovie(movie *Movie) (*Movie, error) {
	movie.ID = uuid.New().String()
	res := db.Create(&movie)
	if res.Error != nil {
		return nil, res.Error
	}
	return movie, nil
 }
 
 func GetMovie(id string) (*Movie, error) {
	var movie Movie
	res := db.First(&movie, "id = ?", id)
	if res.RowsAffected == 0 {
	  return nil, errors.New(fmt.Sprintf("movie of id %s not found", id))
	}
   return &movie, nil
  }

  func GetMovies() ([]*Movie, error) {
	var movies []*Movie
	res := db.Find(&movies)
	if res.Error != nil {
		return nil, errors.New("no movies found")
	}
	return movies, nil
 }

 func UpdateMovie(movie *Movie) (*Movie, error) {
	var movieToUpdate Movie
	result := db.Model(&movieToUpdate).Where("id = ?", movie.ID).Updates(movie)
	if result.RowsAffected == 0 {
		return &movieToUpdate, errors.New("movie not updated")
	}
	return movie, nil
 }

 func DeleteMovie(id string) error {
	var deletedMovie Movie
	result := db.Where("id = ?", id).Delete(&deletedMovie)
	if result.RowsAffected == 0 {
		return errors.New("movie not deleted")
	}
	return nil
 }