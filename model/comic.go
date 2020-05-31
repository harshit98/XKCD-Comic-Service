package model

import (
	"encoding/json"
	"fmt"
)

// Response from the API
type ComicResponse struct {
	Month		string		`json:"month"`
	Num			int			`json:"num"`
	Link		string		`json:"link"`
	Year		string		`json:"year"`
	News		string		`json:"news"`
	SafeTitle	string		`json:"safe_title"`
	Transcript	string		`json:"transcript"`
	Alt			string		`json:"alt"`
	Img			string		`json:"img"`
	Title		string		`json:"title"`
	Day			string		`json:"day"`
}

// Output data format from our CLI application
type Comic struct {
	Title		string		`json:"title"`
	Number		int			`json:"num"`
	Date		string		`json:"date"`
	Description	string		`json:"description"`
	Image		string		`json:"image"`
}

// Create a formatted date from API response
func (cr ComicResponse) FormatDate() string {
	return fmt.Sprintf("%s-%s-%s", cr.Day, cr.Month, cr.Year)
}

// Comic converts ComicResponse that we receive from API to our application
func (cr ComicResponse) Comic() Comic {
	return Comic{
		Title: cr.Title,
		Number: cr.Num,
		Date: cr.FormatDate(),
		Description: cr.Alt,
		Image: cr.Img,
	}
}

// Prettify output from our CLI application
func (c Comic) PrettyString() string {
	prettyString := fmt.Sprintf("Title: %s\nComic No: %d\nDate: %s\nDescription: %s\nImage: %s\n",
								c.Title, c.Number, c.Date, c.Description, c.Image)
	return prettyString
}

func (c Comic) JSON() string {
	comicJSON, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(comicJSON)
}
