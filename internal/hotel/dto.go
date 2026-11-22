package hotel

import "time"


type HotelReq{
	Name string `json: name`
	Slug string `json: slug`
	City string `json: city`
	Address string `json: adress`
	Description string `json: description`
	ImageURL string `json: image_url` 
	StartingPrice *float64  `json:"starting_price"`
	Currency      string    `json:"currency"`
	Rating        *float64  `json:"rating"`
	ReviewCount   int64     `json:"review_count"`
	Badge         *string   `json:"badge"`
	IsPublished   bool      `json:"is_published"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}



type HotelRes{

	Name string `json: name`
	Slug string `json: slug`
	City string `json: city`
	Address string `json: adress`
	Description string `json: description`
	ImageURL string `json: image_url` 
	StartingPrice *float64  `json:"starting_price"`
	Currency      string    `json:"currency"`
	Rating        *float64  `json:"rating"`
	ReviewCount   int64     `json:"review_count"`
	Badge         *string   `json:"badge"`
	IsPublished   bool      `json:"is_published"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}