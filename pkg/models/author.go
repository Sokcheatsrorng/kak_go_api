package models

import "gorm.io/gorm"

type Author struct{
	gorm.Model
    
	// firstName of Author
	FirstName   string `json:"first_name"`
	// lastName of Author
    LastName    string `json:"last_name"`
	// fullName of Author
    FullName    string `json:"full_name"` 
	// DOB of Author
    DateOfBirth string `json:"date_of_birth"` 
	// DOD of Author
    DateOfDeath string `json:"date_of_death"` 
	// biolography of Author
    Biography   string `json:"biography"`
    // image of Author
    PhotoURL    string `json:"photo_url"`
	// books stand as foreign 
	Books []Book `json:"books" gorm:"foreignKey:AuthorID"`
}