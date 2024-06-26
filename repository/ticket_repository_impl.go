package repository

import (
	"gorm.io/gorm"
	"mini_project/constant"
	"mini_project/model/domain"
)

type TicketRepositoryImpl struct {
	DB *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &TicketRepositoryImpl{
		DB: db,
	}
}

func (repository *TicketRepositoryImpl) FindByID(id string) (ticket *domain.Ticket, err error) {
	if err = repository.DB.Preload("User").Preload("TouristAttraction").Preload("TouristAttraction.TouristAttractionType").Where("id LIKE ?", id).First(&ticket).Error; err != nil {
		return ticket, constant.ErrDataNotFound
	}
	return ticket, nil
}

func (repository *TicketRepositoryImpl) FindByUserID(userID int) (tickets *[]domain.Ticket, err error) {
	if err = repository.DB.Preload("TouristAttraction").Preload("TouristAttraction.TouristAttractionType").Preload("User").Where("user_id = ?", userID).Find(&tickets).Error; err != nil {
		return tickets, constant.ErrDataNotFound
	}
	return tickets, nil
}

func (repository *TicketRepositoryImpl) FindByTouristAttractionID(touristAttractionID int) (tickets *[]domain.Ticket, err error) {
	if err = repository.DB.Preload("User").Preload("TouristAttraction").Preload("TouristAttraction.TouristAttractionType").Where("tourist_attraction_id = ?", touristAttractionID).Find(&tickets).Error; err != nil {
		return tickets, constant.ErrDataNotFound
	}
	return tickets, nil
}

func (repository *TicketRepositoryImpl) Insert(ticket *domain.Ticket) error {
	if err := repository.DB.Create(&ticket).Error; err != nil {
		return constant.ErrInsertData
	}
	return nil
}
