package manticore

type Result struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Hits     Hits `json:"hits"`
}

type Hits struct {
	Total         int      `json:"total"`
	TotalRelation string   `json:"total_relation"`
	Details       []Detail `json:"hits"`
}

type Detail struct {
	ID     string                 `json:"_id"`
	Score  int                    `json:"_score"`
	Source map[string]interface{} `json:"_source"`
}

type Source struct {
	Title       string  `json:"title"`
	CategoryID  int     `json:"category_id"`
	ReleaseYear int64   `json:"release_year"`
	Price       float64 `json:"price"`
	CreatedAt   int     `json:"created_at"`
	UpdatedAt   int     `json:"updated_at"`
	DeletedAt   int     `json:"deleted_at"`
	IsFlag      bool    `json:"is_flag"`
	Description string  `json:"description"`
}