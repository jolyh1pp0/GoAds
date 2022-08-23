package domain

const SessionDataKey = "session_data"

type SessionData struct {
	UserID    string `json:"user_id"`
	UserRoles []int  `json:"user_roles"`
}
