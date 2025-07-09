package controllers

import (
	"net/http"
	"candra/backend-api/helpers"
	"candra/backend-api/models"
	"candra/backend-api/repositories"
	"candra/backend-api/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	var users []models.User
	if err := repositories.GetAllUsers(&users, limit, offset); err != nil {
		// Log error detail, tampilkan pesan umum ke client
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Gagal mengambil data user",
			Errors:  map[string]string{"database": "Terjadi kesalahan pada server, silakan coba lagi nanti"},
		})
		return
	}

	total, err := repositories.CountAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Gagal menghitung total user",
			Errors:  map[string]string{"database": "Terjadi kesalahan pada server, silakan coba lagi nanti"},
		})
		return
	}

	userResponses := make([]structs.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = structs.UserResponse{
			Id:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			RoleID:    user.Role.Name,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	baseURL := "http://" + c.Request.Host + c.Request.URL.Path
	pagination := helpers.BuildPaginationResponse(baseURL, page, limit, total, userResponses)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Berhasil mengambil data user",
		Data:    pagination,
	})
}

// GetUserByID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "ID user tidak boleh kosong",
			Errors:  map[string]string{"id": "ID user tidak boleh kosong"},
		})
		return
	}

	var user models.User
	if err := repositories.FindUserByID(id, &user); err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "User tidak ditemukan",
			Errors:  map[string]string{"database": err.Error()},
		})
		return
	}

	userResponse := structs.UserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		RoleID:    user.Role.Name,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Berhasil mengambil data user",
		Data:    userResponse,
	})
}
//UpdateUser
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "ID user tidak boleh kosong",
			Errors:  map[string]string{"id": "ID user tidak boleh kosong"},
		})
		return
	}

	var user models.User
	if err := repositories.FindUserByID(id, &user); err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "User tidak ditemukan",
			Errors:  map[string]string{"database": err.Error()},
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "Data yang dikirim tidak valid",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := repositories.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Gagal memperbarui data user",
			Errors:  map[string]string{"database": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Berhasil memperbarui data user",
		Data:    user,
	})
}
// DeleteUser
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "ID user tidak boleh kosong",
			Errors:  map[string]string{"id": "ID user tidak boleh kosong"},
		})
		return
	}

	var user models.User
	if err := repositories.FindUserByID(id, &user); err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "User tidak ditemukan",
			Errors:  map[string]string{"database": err.Error()},
		})
		return
	}

	if err := repositories.DeleteUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Gagal menghapus data user",
			Errors:  map[string]string{"database": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Berhasil menghapus data user",
	})
}