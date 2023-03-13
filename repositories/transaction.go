package repositories

import (
	"holyways/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	FindTransactionDonation(donationId int) ([]models.Transaction, error)
	FindTransactionUser(userId int) ([]models.Transaction, error)
	GetTransactionId(transactionId int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, orderId int) (models.Transaction, error)
	GetTransactionByDonation(ID int) ([]models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Donation.User").Preload("User").Find(&transactions).Error

	return transactions, err
}

func (r *repository) FindTransactionDonation(donationId int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Where("donation_id=?", donationId).Where("status=?", "success").Preload("Donation").Preload("User").Find(&transactions).Error

	return transactions, err
}
func (r *repository) FindTransactionUser(userId int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Where("user_id", userId).Where("status=?", "success").Preload("Donation").Preload("User").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Donation").Preload("User").First(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransactionId(transactionId int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Donation").Preload("User").First(&transaction, transactionId).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {

	err := r.db.Preload("Donation").Preload("User").Create(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(status string, orderId int) (models.Transaction, error) {
	var transaction models.Transaction

	r.db.Preload("Donation").Preload("User").First(&transaction, orderId)
	if status != transaction.Status && status == "success" {
		var donation models.Donation
		r.db.First(&donation, transaction.Donation.ID)
		donation.CurrentGoal = donation.CurrentGoal + transaction.Total
		donation.TotalDonation = donation.TotalDonation + 1
		r.db.Save(&donation)
	}

	transaction.Status = status
	err := r.db.Save(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransactionByDonation(ID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Where("donation_id", ID).Find(&transactions).Error

	return transactions, err
}
