package repository

import (
	"gorm.io/gorm"
	"mini_project/model/domain"
	"mini_project/repository/interface"
)

type TicketRepositoryImpl struct {
	DB *gorm.DB
}

func NewTicketRepository(db *gorm.DB) _interface.TicketRepository {
	return &TicketRepositoryImpl{
		DB: db,
	}
}

func (repository *TicketRepositoryImpl) FindByID(id string) (ticket *domain.Ticket, err error) {
	if err = repository.DB.Where("id LIKE ?", id).First(&ticket).Error; err != nil {
		return nil, err
	}
	return ticket, nil
}

func (repository *TicketRepositoryImpl) FindByUserID(userID int) (tickets *[]domain.Ticket, err error) {
	if err = repository.DB.Preload("User").Where("user_id = ?", userID).Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (repository *TicketRepositoryImpl) FindByTouristAttractionID(touristAttractionID int) (tickets *[]domain.Ticket, err error) {
	if err = repository.DB.Preload("TouristAttraction").Where("tourist_attraction_id = ?", touristAttractionID).Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (repository *TicketRepositoryImpl) Insert(ticket *domain.Ticket) error {
	if err := repository.DB.Create(&ticket).Error; err != nil {
		return err
	}
	return nil
}
