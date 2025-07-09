package controllers

import (
	"net/http"
	"candra/backend-api/helpers"
	"candra/backend-api/models"
	"candra/backend-api/repositories"
	"candra/backend-api/structs"

	"github.com/gin-gonic/gin"
)

// Register menangani proses registrasi user baru
func Register(c *gin.Context) {
	var req structs.UserCreateRequest

	// Validasi input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validasi gagal",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Hash password
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Gagal mengenkripsi password",
			Errors:  map[string]string{"password": "Gagal hash password"},
		})
		return
	}

	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		RoleID:   req.RoleID, // Pastikan role_id valid
		Password: hashedPassword,
	}

	// Simpan user ke database lewat repository
	if err := repositories.CreateUser(&user); err != nil {
		if helpers.IsDuplicateEntryError(err) {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "Username atau email sudah terdaftar",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		} else {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Gagal membuat user",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		}
		return
	}

	// Ambil nama role berdasarkan RoleID
	var role models.Role
	if err := repositories.FindRoleByID(user.RoleID, &role); err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Gagal mengambil data role",
			Errors:  map[string]string{"role": "Role tidak ditemukan"},
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User berhasil dibuat",
		Data: structs.UserResponse{
			Id:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			RoleID:    role.Name, // Tampilkan nama role, bukan ID
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

// Login menangani proses autentikasi user
func Login(c *gin.Context) {
	var req structs.UserLoginRequest

	// Validasi input
	if err := c.ShouldBindJSON(&req); err != nil {
		// Ambil error validasi per field
		errors := helpers.TranslateErrorMessage(err)
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validasi gagal",
			Errors:  errors, // Akan tampil: {"username": "username wajib diisi", "password": "password wajib diisi"}
		})
		return
	}

	// Cari user berdasarkan username
	var user models.User
	if err := repositories.FindUserByUsername(req.Username, &user); err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Username atau password salah",
			Errors:  map[string]string{"login": "Username atau password salah"},
		})
		return
	}

	// Cek password
	if err := helpers.CheckPassword(req.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Username atau password salah",
			Errors:  map[string]string{"login": "Username atau password salah"},
		})
		return
	}

	// Ambil nama role
	var role models.Role
	if err := repositories.FindRoleByID(user.RoleID, &role); err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Role tidak ditemukan",
			Errors:  map[string]string{"role": "Role tidak ditemukan"},
		})
		return
	}

	// Generate token JWT
	token, err := helpers.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Gagal membuat token",
			Errors:  map[string]string{"token": "Gagal membuat token"},
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Login berhasil",
		Data: structs.UserResponse{
			Id:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			RoleID:    role.Name, // Tampilkan nama role
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
			Token:     &token,
		},
	})
}
