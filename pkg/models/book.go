package models

import "gorm.io/gorm"

type Book struct{

	gorm.Model
    
	// declare Title for Book as string type
	Title          string `json:"title"`
	// declare AuthorID for Book as unsigned integer type
    AuthorID       uint   `json:"author_id"`
	// declare PublishedDate for Book as string type
    PublishedDate  string `json:"published_date"`
	// Unique identifier for the book.
    ISBN           string `json:"isbn"`
    // Number of pages.
    PageCount      int    `json:"page_count"`
    // Language 
    Language       string `json:"language"`
	// Average rating, if you want to include this feature.
    Rating         float64 `json:"rating"`
}