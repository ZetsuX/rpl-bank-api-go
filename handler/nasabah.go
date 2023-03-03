package handler

import (
	"net/http"
	"strconv"
	"tugas3rpl/common"
	"tugas3rpl/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NasabahHandler struct {
	DB *gorm.DB
}

func (h *NasabahHandler) HandleGetNasabah(ctx *gin.Context) {
	q := ctx.Query("q")

	dbQuery := h.DB
	if q != "" {
		q = "%" + q + "%"
		dbQuery = dbQuery.Where("name ILIKE $1 OR email ILIKE $1 OR address ILIKE $1", q)
	}

	var daftarNasabah []entity.Nasabah
	tx := dbQuery.Preload("NoTelp").Preload("Rekening").Find(&daftarNasabah)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}
	ctx.JSON(http.StatusOK,
		common.Response{
			Status:  true,
			Message: "seluruh data nasabah berhasil didapatkan",
			Data:    daftarNasabah,
		})
}

func (h *NasabahHandler) HandleGetNasabahByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	var nasabah entity.Nasabah
	tx := h.DB.Where("id = ?", id).Preload("NoTelp").Preload("Rekening").Take(&nasabah)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}
	ctx.JSON(http.StatusOK,
		common.Response{
			Status:  true,
			Message: "data nasabah berhasil didapatkan",
			Data:    nasabah,
		})
}

func (h *NasabahHandler) HandleInsertNasabah(ctx *gin.Context) {
	var nasabah entity.Nasabah

	err := ctx.ShouldBind(&nasabah)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}
	tx := h.DB.Create(&nasabah)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	ctx.JSON(http.StatusOK,
		common.Response{
			Status:  true,
			Message: "nasabah berhasil dibuat",
			Data:    nasabah,
		})
}

func (h *NasabahHandler) HandleDeleteNasabahByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	var cariNasabah entity.Nasabah
	tx := h.DB.Where("id = ?", id).Preload("NoTelp").Preload("Rekening").Take(&cariNasabah)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	var nasabah entity.Nasabah
	tx = h.DB.Where("id = ?", id).Preload("NoTelp").Preload("Rekening").Delete(&nasabah)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}
	ctx.JSON(http.StatusOK,
		common.Response{
			Status:  true,
			Message: "data nasabah berhasil dihapus",
			Data:    nil,
		})
}

func (h *NasabahHandler) HandleEditNasabahByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	var nasabah entity.Nasabah
	err = ctx.ShouldBind(&nasabah)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	var cari entity.Nasabah
	tx := h.DB.Where("id = ?", id).Preload("NoTelp").Preload("Rekening").Take(&cari)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	tx = h.DB.Model(&cari).Updates(&nasabah)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	tx = h.DB.Where("id = ?", id).Preload("NoTelp").Preload("Rekening").Take(&nasabah)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}
	ctx.JSON(http.StatusOK,
		common.Response{
			Status:  true,
			Message: "data nasabah berhasil diubah",
			Data:    nasabah,
		})
}
