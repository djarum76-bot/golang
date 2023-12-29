package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Person struct {
	ID               int      `json:"id"`
	UserID           int      `json:"user_id"`
	Nickname         string   `json:"nickname"`
	FirstName        string   `json:"first_name"`
	LastName         string   `json:"last_name"`
	Email            []string `json:"email"`
	Job              string   `json:"job"`
	Phone            string   `json:"phone"`
	DateOfBirth      string   `json:"date_of_birth"`
	PlaceOfBirth     string   `json:"place_of_birth"`
	PlaceOfResidence []string `json:"place_of_residence"`
	Instagram        string   `json:"instagram"`
	Twitter          string   `json:"twitter"`
	Facebook         string   `json:"facebook"`
	LinkedIn         string   `json:"linkedin"`
	Major            []string `json:"major"`
	College          []string `json:"college"`
	Organization     []string `json:"organization"`
	Party            []string `json:"party"`
	LastActivity     []string `json:"last_activity"`
	FirstMeet        string   `json:"first_meet"`
	Category         string   `json:"category"`
	Tags             []string `json:"tags"`
	Image            string   `json:"image"`
	Images           []string `json:"images"`
	CreatedAt        string   `json:"created_at"`
}
