package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/odhiahmad/apiuser/dto"
	"github.com/odhiahmad/apiuser/helper"
	"github.com/odhiahmad/apiuser/service"
)

type RumahController interface {
	CreateRumah(ctx *gin.Context)
	UpdateRumah(ctx *gin.Context)
	Statistik(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindAllByKota(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Delete(ctx *gin.Context)
	DeleteByIds(ctx *gin.Context)
}

type rumahController struct {
	rumahService service.RumahService
	jwtService   service.JWTService
}

func NewRumahController(rumahService service.RumahService, jwtService service.JWTService) RumahController {
	return &rumahController{
		rumahService: rumahService,
		jwtService:   jwtService,
	}
}

func (c *rumahController) CreateRumah(ctx *gin.Context) {
	var rumahCreateDTO dto.RumahCreateDTO
	errDTO := ctx.ShouldBind(&rumahCreateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// if !c.rumahService.IsDuplicateNomorKK(rumahCreateDTO.NomorKK) {
	// 	response := helper.BuildErrorResponse("Failed to process request", "Duplicate Nomor KK", helper.EmptyObj{})
	// 	ctx.JSON(http.StatusConflict, response)
	// } else {

	// dec, err := base64.StdEncoding.DecodeString(rumahCreateDTO.File)
	// if err != nil {
	// 	panic(err)
	// }

	// f, err := os.Create("fileupload/" + createdRumah.ID.String() + rumahCreateDTO.Ekstensi)
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// if _, err := f.Write(dec); err != nil {
	// 	panic(err)
	// }
	// if err := f.Sync(); err != nil {
	// 	panic(err)
	// }
	rumahCreateDTO.NamaKepalaKeluarga = strings.ToLower(rumahCreateDTO.NamaKepalaKeluarga)
	rumahCreateDTO.Desa = strings.ToLower(rumahCreateDTO.Desa)
	rumahCreateDTO.Nagari = strings.ToLower(rumahCreateDTO.Nagari)
	rumahCreateDTO.Kecamatan = strings.ToLower(rumahCreateDTO.Kecamatan)
	rumahCreateDTO.Jorong = strings.ToLower(rumahCreateDTO.Jorong)
	rumahCreateDTO.Dusun = strings.ToLower(rumahCreateDTO.Dusun)
	createdRumah := c.rumahService.CreateRumah(rumahCreateDTO)
	response := helper.BuildResponse(true, "!OK", createdRumah)
	ctx.JSON(http.StatusCreated, response)

}

func (c *rumahController) UpdateRumah(ctx *gin.Context) {
	var rumahUpdateDTO dto.RumahUpdateDTO
	errDTO := ctx.ShouldBind(&rumahUpdateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// var getId = c.rumahService.IsDuplicateNomorKK(rumahUpdateDTO.NomorKK)

	// if getId
	// if !c.rumahService.IsDuplicateNomorKK(rumahUpdateDTO.NomorKK) {
	// 	response := helper.BuildErrorResponse("Failed to process request", "Duplicate Nomor KK", helper.EmptyObj{})
	// 	ctx.JSON(http.StatusConflict, response)
	// } else {

	rumahUpdateDTO.NamaKepalaKeluarga = strings.ToLower(rumahUpdateDTO.NamaKepalaKeluarga)

	rumahUpdateDTO.Desa = strings.ToLower(rumahUpdateDTO.Desa)
	rumahUpdateDTO.Kecamatan = strings.ToLower(rumahUpdateDTO.Kecamatan)
	rumahUpdateDTO.Nagari = strings.ToLower(rumahUpdateDTO.Nagari)
	rumahUpdateDTO.Jorong = strings.ToLower(rumahUpdateDTO.Jorong)
	rumahUpdateDTO.Dusun = strings.ToLower(rumahUpdateDTO.Dusun)

	updatedRumah := c.rumahService.UpdateRumah(rumahUpdateDTO)
	response := helper.BuildResponse(true, "!OK", updatedRumah)
	ctx.JSON(http.StatusCreated, response)

}

func (c *rumahController) FindById(ctx *gin.Context) {
	var idRumah dto.RumahFindIdDTO
	errDTO := ctx.ShouldBind(&idRumah)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	findRumah := c.rumahService.FindById(idRumah)
	response := helper.BuildResponse(true, "Data berhasil di select", findRumah)
	ctx.JSON(http.StatusOK, response)
}

func (c *rumahController) FindAll(ctx *gin.Context) {
	var pagination dto.Pagination

	errDTO := ctx.ShouldBind(&pagination)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	// paginationSet := helper.GeneratePaginationFromRequest(ctx)
	getRumah, err := c.rumahService.FindAll(pagination)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "!OK", getRumah)
	ctx.JSON(http.StatusOK, response)

}

func (c *rumahController) FindAllByKota(ctx *gin.Context) {
	var pagination dto.Pagination

	errDTO := ctx.ShouldBind(&pagination)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	// paginationSet := helper.GeneratePaginationFromRequest(ctx)
	getRumah, err := c.rumahService.FindAllByKota(pagination)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "!OK", getRumah)
	ctx.JSON(http.StatusOK, response)

}

func (c *rumahController) Delete(ctx *gin.Context) {
	var rumahDelete dto.RumahDeleteDTO
	errDTO := ctx.ShouldBind(&rumahDelete)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	deleteRumah := c.rumahService.Delete(rumahDelete.Id)
	response := helper.BuildResponse(true, "Data berhasil dihapus", deleteRumah)
	ctx.JSON(http.StatusOK, response)

}

func (c *rumahController) DeleteByIds(ctx *gin.Context) {
	var rumahDelete dto.RumahDeleteMultiID
	errDTO := ctx.ShouldBind(&rumahDelete)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	deleteRumah := c.rumahService.DeleteByIds(rumahDelete)

	// for i := 0; i < len(rumahDelete.NamaGambar); i++ {
	// 	os.Remove("fileupload/" + rumahDelete.NamaGambar[i])
	// }

	response := helper.BuildResponse(true, "Data berhasil dihapus", deleteRumah)
	ctx.JSON(http.StatusOK, response)

}

func (c *rumahController) Statistik(ctx *gin.Context) {
	var statistik dto.StatistikDTO

	errDTO := ctx.ShouldBind(&statistik)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	// paginationSet := helper.GeneratePaginationFromRequest(ctx)
	getRumah, err := c.rumahService.Statistik(statistik)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildResponse(true, "!OK", getRumah)
	ctx.JSON(http.StatusOK, response)

}
