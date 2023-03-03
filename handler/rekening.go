package handler

import (
	"net/http"
	"strconv"
	"tugas3rpl/common"
	"tugas3rpl/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RekeningHandler struct {
	DB *gorm.DB
}

func (h *RekeningHandler) HandleGetRekeningByNasabahID(ctx *gin.Context) {
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
			Message: "data rekening nasabah berhasil didapatkan",
			Data:    nasabah.Rekening,
		})
}

func (h *RekeningHandler) HandleGetRekening(ctx *gin.Context) {
	dbQuery := h.DB

	tipe := ctx.Query("type")
	moreThan := ctx.Query("more")
	lessThan := ctx.Query("less")
	if tipe != "" || moreThan != "" || lessThan != "" {
		if tipe != "" {
			dbQuery = dbQuery.Where("type = ?", tipe)
		}

		if moreThan != "" {
			more, err := strconv.ParseUint(moreThan, 10, 64)
			if err != nil {
				ctx.JSON(http.StatusBadRequest,
					common.Response{
						Status:  false,
						Message: err.Error(),
						Data:    nil,
					})
				return
			}

			dbQuery = dbQuery.Where("balance >= ?", more)
		}

		if lessThan != "" {
			less, err := strconv.ParseUint(lessThan, 10, 64)
			if err != nil {
				ctx.JSON(http.StatusBadRequest,
					common.Response{
						Status:  false,
						Message: err.Error(),
						Data:    nil,
					})
				return
			}

			dbQuery = dbQuery.Where("balance <= ?", less)
		}
	}

	var daftarRekening []entity.Rekening
	tx := dbQuery.Find(&daftarRekening)
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
			Message: "seluruh data Rekening berhasil didapatkan",
			Data:    daftarRekening,
		})
}

func (h *RekeningHandler) HandleGetRekeningByID(ctx *gin.Context) {
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

	var rekening entity.Rekening
	tx := h.DB.Where("id = ?", id).Take(&rekening)
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
			Message: "data rekening berhasil didapatkan",
			Data:    rekening,
		})
}

func (h *RekeningHandler) HandleInsertRekening(ctx *gin.Context) {
	var rekening entity.Rekening
	err := ctx.ShouldBind(&rekening)
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
	nasabahId := rekening.NasabahID
	tx := h.DB.Where("id = ?", nasabahId).Preload("NoTelp").Preload("Rekening").Take(&cari)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	tx = h.DB.Create(&rekening)
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
			Message: "rekening baru berhasil dimasukkan",
			Data:    rekening,
		})
}

func (h *RekeningHandler) HandleDeleteRekeningByID(ctx *gin.Context) {
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

	var cariRek entity.Rekening
	tx := h.DB.Where("id = ?", id).Take(&cariRek)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	var rekening entity.Rekening
	tx = h.DB.Where("id = ?", id).Delete(&rekening)
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
			Message: "data Rekening berhasil dihapus",
			Data:    nil,
		})
}

func (h *RekeningHandler) HandleEditRekeningByID(ctx *gin.Context) {
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

	var rekening entity.Rekening
	err = ctx.ShouldBind(&rekening)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	var cariRek entity.Rekening
	tx := h.DB.Where("id = ?", id).Take(&cariRek)
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
	nasabahId := rekening.NasabahID
	tx = h.DB.Where("id = ?", nasabahId).Preload("NoTelp").Preload("Rekening").Take(&cariNas)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	tx = h.DB.Model(&cariRek).Updates(&rekening)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	tx = h.DB.Where("id = ?", id).Take(&rekening)
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
			Message: "data rekening berhasil diubah",
			Data:    rekening,
		})
}
