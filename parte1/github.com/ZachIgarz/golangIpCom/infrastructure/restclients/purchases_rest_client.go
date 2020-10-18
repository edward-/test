package restclients

import (
	"io/ioutil"
	"net/http"

	"github.com/ZachIgarz/golangIpCom/infrastructure/entities"
)

const purchaseURL = "https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/"

//PurchaseRestClient...
type PurchaseRestClient struct {
}

/*Get consume api y retorna respuesta*/
func (purchaseRestClient PurchaseRestClient) Get(purchaseResumeRequest entities.PurchaseResumeRequest) ([]byte, error) {

	purchaseRestClientURL := (purchaseURL) + (purchaseResumeRequest.RealDate())

	response, err := http.Get(purchaseRestClientURL)
	var respuesta []byte
	if err != nil {
		return respuesta, err
	}
	//DEFER
	defer response.Body.Close()

	respuesta, _ = ioutil.ReadAll(response.Body)

	return respuesta, err
}
