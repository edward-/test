package entities

//PurchaseResumeRequest ...
type PurchaseResumeRequest struct {
	realDate string
	days     int
}

//NewPurchaseResumeRequest ...
func NewPurchaseResumeRequest(realDate string, days int) *PurchaseResumeRequest {
	return &PurchaseResumeRequest{
		realDate: realDate,
		days:     days,
	}
}

//RealDate ...
func (purchase *PurchaseResumeRequest) RealDate() string {
	return purchase.realDate
}
