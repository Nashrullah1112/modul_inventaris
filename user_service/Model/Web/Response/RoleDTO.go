package Response

// RoleCreateRequest represents the request payload for creating a new Role
type RoleCreateRequest struct {
	Nama string `json:"nama" validate:"required"`
}

// RoleUpdateRequest represents the request payload for updating an existing Role
type RoleUpdateRequest struct {
	ID   int64  `json:"id" validate:"required"`
	Nama string `json:"nama" validate:"required"`
}

// RoleResponse represents the response payload for returning Role data
type RoleResponse struct {
	ID   int64  `json:"id"`
	Nama string `json:"nama"`
}
