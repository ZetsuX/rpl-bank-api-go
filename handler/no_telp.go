package handler

import (
	"net/http"
	"strconv"
	"tugas3rpl/common"
	"tugas3rpl/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NoTelpHandler struct {
	DB *gorm.DB
}

func (h *NoTelpHandler) HandleGetNoTelpByNasabahID(ctx *gin.Context) {
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
			Message: "data nomor telepon nasabah berhasil didapatkan",
			Data:    nasabah.NoTelp,
		})
}

func (h *NoTelpHandler) HandleGetNoTelp(ctx *gin.Context) {
	dbQuery := h.DB

	country_code := ctx.Query("country")
	if country_code != "" {
		country_code = "+" + country_code
		dbQuery = dbQuery.Where("country_code = ?", country_code)
	}

	var daftarNoTelp []entity.NoTelp
	tx := dbQuery.Find(&daftarNoTelp)
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
			Message: "seluruh data nomor telepon berhasil didapatkan",
			Data:    daftarNoTelp,
		})
}

func (h *NoTelpHandler) HandleGetNoTelpByID(ctx *gin.Context) {
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

	var noTelp entity.NoTelp
	tx := h.DB.Where("id = ?", id).Take(&noTelp)
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
			Message: "data nomor telepon berhasil didapatkan",
			Data:    noTelp,
		})
}

func (h *NoTelpHandler) HandleInsertNoTelp(ctx *gin.Context) {
	var noTelp entity.NoTelp
	err := ctx.ShouldBind(&noTelp)
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
	nasabahId := noTelp.NasabahID
	tx := h.DB.Where("id = ?", nasabahId).Preload("NoTelp").Preload("NoTelp").Take(&cari)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	tx = h.DB.Create(&noTelp)
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
			Message: "nomor telepon berhasil dimasukkan",
			Data:    noTelp,
		})
}

func (h *NoTelpHandler) HandleDeleteNoTelpByID(ctx *gin.Context) {
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

	var cariTelp entity.NoTelp
	tx := h.DB.Where("id = ?", id).Take(&cariTelp)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	var noTelp entity.NoTelp
	tx = h.DB.Where("id = ?", id).Delete(&noTelp)
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
			Message: "data nomor telepon berhasil dihapus",
			Data:    nil,
		})
}

func (h *NoTelpHandler) HandleEditNoTelpByID(ctx *gin.Context) {
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

	var noTelp entity.NoTelp
	err = ctx.ShouldBind(&noTelp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	var cariTelp entity.NoTelp
	tx := h.DB.Where("id = ?", id).Take(&cariTelp)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	var cariNas entity.Nasabah
	nasabahId := noTelp.NasabahID
	tx = h.DB.Where("id = ?", nasabahId).Preload("NoTelp").Preload("NoTelp").Take(&cariNas)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	tx = h.DB.Model(&cariTelp).Updates(&noTelp)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	tx = h.DB.Where("id = ?", id).Take(&noTelp)
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
			Message: "data nomor telepon berhasil diubah",
			Data:    noTelp,
		})
}
