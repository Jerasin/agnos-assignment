package repository

type PatientRepositoryInterface interface {
	GetBaseRepo() BaseRepositoryInterface
}

type PatientRepository struct {
	BaseRepository BaseRepositoryInterface
}

func PatientRepositoryInit(baseRepo *BaseRepository) *PatientRepository {
	return &PatientRepository{
		BaseRepository: baseRepo,
	}
}

func (p PatientRepository) GetBaseRepo() BaseRepositoryInterface {
	return p.BaseRepository
}
