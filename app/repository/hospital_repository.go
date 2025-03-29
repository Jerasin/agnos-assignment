package repository

type HospitalRepositoryInterface interface {
	GetBaseRepo() BaseRepositoryInterface
}

type HospitalRepository struct {
	BaseRepository BaseRepositoryInterface
}

func HospitalRepositoryInit(baseRepo *BaseRepository) *HospitalRepository {
	return &HospitalRepository{
		BaseRepository: baseRepo,
	}
}

func (p HospitalRepository) GetBaseRepo() BaseRepositoryInterface {
	return p.BaseRepository
}
