package requests

type Header struct {
	PurchaseRequisition            int     `json:"PurchaseRequisition"`
	IsCancelled          	   	   *bool   `json:"IsCancelled"`
}
