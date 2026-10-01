package hotel

import "time"

type Hotel struct {
	ID            int64     `db:"id"`
	CategoryID    int64     `db:"category_id"`
	Name          string    `db:"name"`
	Slug          string    `db:"slug"`
	City          string    `db:"city"`
	Address       string    `db:"address"`
	Description   string    `db:"description"`
	ImageURL      string    `db:"image_url"`
	StartingPrice *float64  `db:"starting_price"`
	Currency      string    `db:"currency"`
	Rating        *float64  `db:"rating"`
	ReviewCount   int64     `db:"review_count"`
	Badge         *string   `db:"badge"`
	IsPublished   bool      `db:"is_published"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
