package models

type BaseModel struct {
	ID        int `json:"id" db:"id"`
	CreatedAt int `json:"created_at" db:"created_at"`
	UpdatedAt int `json:"updated_at" db:"updated_at"`
}

type User struct {
	UserID       int    `json:"user_id" db:"user_id"`
	Username     string `json:"username" db:"username"`
	FirstName    string `json:"first_name" db:"first_name"`
	LastName     string `json:"last_name" db:"last_name"`
	PhoneNumber  string `json:"phone_number" db:"phone_number"`
	LanguageCode string `json:"language_code" db:"language_code"`
	IsBanned     bool   `json:"is_banned" db:"is_banned"`
	BaseModel
}

type City struct {
	NameUz    string     `json:"name_uz" db:"name_uz"`
	NameRu    string     `json:"name_ru" db:"name_ru"`
	NameEn    string     `json:"name_en" db:"name_en"`
	Districts []District `json:"districts"`
	BaseModel
}

type District struct {
	NameUz string `json:"name_uz" db:"name_uz"`
	NameRu string `json:"name_ru" db:"name_ru"`
	NameEn string `json:"name_en" db:"name_en"`
	// City   *City  `json:"city"`
	BaseModel
}

// type Message struct {
// 	UserID   int    `json:"user_id" db:"user_id"`
// 	City     int    `json:"city" db:"region"`
// 	District int    `json:"district" db:"district"`
// 	Text     string `json:"text" db:"text"`
// 	BaseModel
// }

// type Chat struct {
// 	IsAdminSender bool   `json:"is_admin_sender" db:"is_admin_sender"`
// 	Text          string `json:"text" db:"text"`
// 	MessageID     int    `json:"message_id" db:"message_id"`
// 	BaseModel
// }
