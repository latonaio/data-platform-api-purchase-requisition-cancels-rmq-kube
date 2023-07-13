package requests

type Item struct {
	PurchaseRequisition    		int     `json:"PurchaseRequisition"`
	PurchaseRequisitionItem		int     `json:"PurchaseRequisitionItem"`
	IsCancelled        			*bool   `json:"IsCancelled"`
}
