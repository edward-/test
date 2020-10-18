package application

import (
	"bytes"
	"encoding/json"
	domainEntities "github.com/ZachIgarz/golangIpCom/domain/entities"
	"github.com/ZachIgarz/golangIpCom/domain/ports"
	"github.com/ZachIgarz/golangIpCom/infrastructure/entities"
)

//PurchasesUseCase ...
type PurchasesUseCase interface {
	Handler(purchaseResumeRequest entities.PurchaseResumeRequest) (domainEntities.Statistics, error)
}

//PurchasesApplication ...
type PurchasesApplication struct {
	purchaseClient  ports.PurchasesClient
	purchasesEntity []domainEntities.Purchases
}

//NewPurchasesApplication ....
func NewPurchasesApplication(purchaseClient ports.PurchasesClient,
	purchasesEntity []domainEntities.Purchases) *PurchasesApplication {

	return &PurchasesApplication{
		purchaseClient:  purchaseClient,
		purchasesEntity: purchasesEntity,
	}

}

/*Handler ...*/
func (purchasesApplication *PurchasesApplication) Handler(purchaseResumeRequest entities.PurchaseResumeRequest) (statistics domainEntities.Statistics, err error) {
	responsePurchaseClient, error := purchasesApplication.purchaseClient.Get(purchaseResumeRequest)

	if error != nil {
		err := error

		return statistics, err
	}

	dec := json.NewDecoder(bytes.NewReader(responsePurchaseClient))
	dec.Decode(&purchasesApplication.purchasesEntity)

	obtainStatistics(&statistics, purchasesApplication.purchasesEntity)

	return statistics, error
}

func obtainStatistics(statistics *domainEntities.Statistics, purchasesEntity []domainEntities.Purchases) {
	statistics.GetTotalPurchases(purchasesEntity)
	statistics.WithoutPurchases(purchasesEntity)
	statistics.HighestPurchases(purchasesEntity)
	statistics.PurchasesByCreditCards(purchasesEntity)
}
