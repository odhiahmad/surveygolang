package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/odhiahmad/apiuser/dto"
	"github.com/odhiahmad/apiuser/entity"
	"github.com/odhiahmad/apiuser/repository"
)

type RumahService interface {
	CreateRumah(rumah dto.RumahCreateDTO) entity.Rumah
	UpdateRumah(rumah dto.RumahUpdateDTO) entity.Rumah
	Statistik(statistik dto.StatistikDTO) (*dto.StatistikDTO, error)
	FindAll(pagination dto.Pagination) (*dto.Pagination, error)
	FindById(id dto.RumahFindIdDTO) entity.Rumah
	FindAllByKota(pagination dto.Pagination) (*dto.Pagination, error)
	Delete(id uint64) entity.Rumah
	DeleteByIds(id dto.RumahDeleteMultiID) entity.Rumah
	IsDuplicateNomorKK(nomorKK string) bool
}

type rumahService struct {
	rumahRepository repository.RumahRepository
}

func NewRumahService(rumahRepo repository.RumahRepository) RumahService {
	return &rumahService{
		rumahRepository: rumahRepo,
	}
}

func (service *rumahService) CreateRumah(rumah dto.RumahCreateDTO) entity.Rumah {
	rumahToCreate := entity.Rumah{}
	err := smapping.FillStruct(&rumahToCreate, smapping.MapFields(&rumah))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	rumahToCreate.Prepare()
	res := service.rumahRepository.InsertRumah((rumahToCreate))
	return res
}

func (service *rumahService) UpdateRumah(rumah dto.RumahUpdateDTO) entity.Rumah {
	rumahToUpdate := entity.Rumah{}
	err := smapping.FillStruct(&rumahToUpdate, smapping.MapFields(&rumah))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}

	res := service.rumahRepository.UpdateRumah((rumahToUpdate))
	return res
}

func (service *rumahService) Delete(id uint64) entity.Rumah {

	res := service.rumahRepository.Delete(id)
	return res
}

func (service *rumahService) FindById(id dto.RumahFindIdDTO) entity.Rumah {

	res := service.rumahRepository.FindById(id.Id)
	return res
}

func (service *rumahService) DeleteByIds(multiId dto.RumahDeleteMultiID) entity.Rumah {

	// var panjang = len(multiId.Ids)
	// for i := 0; i < panjang; i++ {
	// 	resGetId = service.rumahRepository.FindById(multiId.Ids[i]);
	// 	e := os.Remove("/Users/anki/Documents/new_folder/GeeksforGeeks.txt")
	// 	if e != nil {
	// 		log.Fatal(e)

	// 	}
	// }

	res := service.rumahRepository.DeleteByIds(&multiId.Ids)
	return res
}

func (service *rumahService) FindAll(pagination dto.Pagination) (*dto.Pagination, error) {
	return service.rumahRepository.FindAll(pagination)
}

func (service *rumahService) FindAllByKota(pagination dto.Pagination) (*dto.Pagination, error) {
	return service.rumahRepository.FindAllByKota(pagination)
}

func (service *rumahService) Statistik(statistik dto.StatistikDTO) (*dto.StatistikDTO, error) {
	return service.rumahRepository.Statistik(statistik)
}

func (service *rumahService) IsDuplicateNomorKK(nomorKK string) bool {
	res := service.rumahRepository.IsDuplicateNomorKK(nomorKK)
	return !(res.Error == nil)
}
