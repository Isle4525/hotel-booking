package category

import "time"

type CategoryReq struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CategoryRes struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
