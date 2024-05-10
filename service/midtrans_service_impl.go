package service

import (
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"mini_project/config"
	"mini_project/model/domain"
	"mini_project/service/interface"
)

type MidtransServiceImpl struct {
	midtransConfig config.Midtrans
	envi           midtrans.EnvironmentType
}

func NewMidtransService(cnf *config.Config) _interface.MidtransService {
	envi := midtrans.Sandbox
	if cnf.Midtrans.IsProd {
		envi = midtrans.Production
	}
	return &MidtransServiceImpl{
		midtransConfig: cnf.Midtrans,
		envi:           envi,
	}
}

func (service *MidtransServiceImpl) GenerateSnapURL(payment *domain.Payment) error {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  payment.ID.String(),
			GrossAmt: int64(payment.Amount),
		},
	}

	var client snap.Client
	client.New(service.midtransConfig.Key, service.envi)
	// 3. Request create Snap transaction to Midtrans
	snapResp, err := client.CreateTransaction(req)
	if err != nil {
		return err
	}
	payment.SnapURL = snapResp.RedirectURL
	return nil
}

func (service *MidtransServiceImpl) VerifyPayment(orderID string) (bool, error) {
	var client coreapi.Client
	client.New(service.midtransConfig.Key, service.envi)

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := client.CheckTransaction(orderID)
	if e != nil {
		return false, e
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					//return true, nil
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				return true, nil
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
				fmt.Println("payment pending")
			}
		}
	}
	return false, e
}
