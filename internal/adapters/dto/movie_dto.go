package dto

type MovieDto struct {
	ID          string   `json:"id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Overview    string   `json:"overview,omitempty"`
	ReleaseDate string   `json:"release_date,omitempty"`
	Image       string   `json:"image,omitempty"`
	Rating      float32  `json:"rating,omitempty"`
	Duration    int16    `json:"duration,omitempty"`
	Budget      float64  `json:"budget,omitempty"`
	Genres      []string `json:"genres,omitempty"`
}
