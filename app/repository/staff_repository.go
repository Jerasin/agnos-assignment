package repository

type StaffRepositoryInterface interface {
	GetBaseRepo() BaseRepositoryInterface
}

type StaffRepository struct {
	BaseRepository BaseRepositoryInterface
}

func StaffRepositoryInit(baseRepo *BaseRepository) *StaffRepository {
	return &StaffRepository{
		BaseRepository: baseRepo,
	}
}

func (p StaffRepository) GetBaseRepo() BaseRepositoryInterface {
	return p.BaseRepository
}
