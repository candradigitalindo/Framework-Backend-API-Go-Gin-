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

// Tampilkan semua role dengan pagination
func GetAllRoles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	var roles []models.Role
	if err := repositories.FindAllRoles(&roles, limit, offset); err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Gagal mengambil data role",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	roleResponses := make([]structs.RoleResponse, len(roles))
	for i, role := range roles {
		roleResponses[i] = structs.RoleResponse{
			ID:        role.ID,
			Name:      role.Name,
			CreatedAt: role.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: role.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Berhasil mengambil data role",
		Data:    roleResponses,
	})
}

// Fungsi untuk membuat role baru
func CreateRole(c *gin.Context) {
	var req structs.RoleCreateRequest

	// Validasi input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validasi gagal",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	role := models.Role{Name: req.Name}

	if err := repositories.CreateRole(&role); err != nil {
		if helpers.IsDuplicateEntryError(err) {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "Role sudah ada",
				Errors:  map[string]string{"name": "Role dengan nama ini sudah ada"},
			})
		} else {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Gagal membuat role",
				Errors:  helpers.TranslateErrorMessage(err),
			})
		}
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Role berhasil dibuat",
		Data: structs.RoleResponse{
			ID:        role.ID,
			Name:      role.Name,
			CreatedAt: role.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: role.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

// Fungsi untuk mengambil detail role berdasarkan ID
func GetRoleByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "ID role tidak boleh kosong",
			Errors:  map[string]string{"id": "ID role tidak boleh kosong"},
		})
		return
	}

	var role models.Role
	if err := repositories.FindRoleByID(id, &role); err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Role tidak ditemukan",
			Errors:  map[string]string{"role": "Role dengan ID tersebut tidak ditemukan"},
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Berhasil mengambil data role",
		Data: structs.RoleResponse{
			ID:        role.ID,
			Name:      role.Name,
			CreatedAt: role.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: role.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

// Fungsi untuk mengupdate role
func UpdateRole(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "ID role tidak boleh kosong",
			Errors:  map[string]string{"id": "ID role tidak boleh kosong"},
		})
		return
	}

	var req structs.RoleCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validasi gagal",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	updatedData := map[string]interface{}{"name": req.Name}

	if err := repositories.UpdateRole(id, updatedData); err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Gagal mengupdate role",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Ambil data role terbaru setelah update
	var updatedRole models.Role
	if err := repositories.FindRoleByID(id, &updatedRole); err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Gagal mengambil data role setelah update",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Role berhasil diupdate",
		Data: structs.RoleResponse{
			ID:        updatedRole.ID,
			Name:      updatedRole.Name,
			CreatedAt: updatedRole.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: updatedRole.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

// Fungsi untuk menghapus role
func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "ID role tidak boleh kosong",
			Errors:  map[string]string{"id": "ID role tidak boleh kosong"},
		})
		return
	}

	err := repositories.DeleteRole(id)
	if err != nil {
		status := http.StatusBadRequest
		msg := "Gagal menghapus role"
		if err.Error() == "Role sedang dipakai oleh user, tidak bisa dihapus" {
			status = http.StatusConflict
			msg = "Role sedang dipakai oleh user"
		}
		c.JSON(status, structs.ErrorResponse{
			Success: false,
			Message: msg,
			Errors:  map[string]string{"role": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Role berhasil dihapus",
		Data:    nil,
	})
}
