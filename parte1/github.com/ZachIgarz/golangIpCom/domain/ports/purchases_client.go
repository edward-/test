package ports

import "github.com/ZachIgarz/golangIpCom/infrastructure/entities"

//PurchasesClient ..
type PurchasesClient interface {
	Get(purchaseResumeRequest entities.PurchaseResumeRequest) ([]byte, error)
}
