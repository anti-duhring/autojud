package processes

import "github.com/google/uuid"

type Court string

const (
	COURT_UNKNOWN Court = "UNKNOWN"
	COURT_TJPE    Court = "TJPE"
)

type ProcessFollow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	ProcessID uuid.UUID `json:"process_id"`
	CreatedAt string    `json:"created_at"`
	DeletedAt *string   `json:"deleted_at"`
}

type PendingProcess struct {
	ID         uuid.UUID `json:"id"`
	ProcessID  uuid.UUID `json:"process_id"`
	CreatedAt  string    `json:"created_at"`
	InsertedAt *string   `json:"inserted_at"`
	DeletedAt  *string   `json:"deleted_at"`
}

type Process struct {
	ID            uuid.UUID `json:"id"`
	ProcessNumber string    `json:"process_number"`
	Court         Court     `json:"court"`
	Origin        *string   `json:"origin"`
	Judge         *string   `json:"judge"`
	ActivePart    *string   `json:"active_part"`
	PassivePart   *string   `json:"passive_part"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
	DeletedAt     *string   `json:"deleted_at"`
}
