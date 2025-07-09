package structs

// Struct ini digunakan untuk menampilkan data user sebagai response API
type UserResponse struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	RoleID    string  `json:"role_id"` // Tambahkan ini
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Token     *string `json:"token,omitempty"`
}

// Struct ini digunakan untuk menerima data saat proses create user
type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   string `json:"role_id" binding:"required"` // Tambahkan ini jika perlu
}

// Struct ini digunakan untuk menerima data saat proses update user
type UserUpdateRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password,omitempty"`
	RoleID   string `json:"role_id,omitempty"` // Tambahkan ini jika perlu
}

// Struct ini digunakan saat user melakukan proses login
type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

