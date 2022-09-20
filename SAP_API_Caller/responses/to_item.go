package responses

type ToItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			FiscalYear                     string `json:"FiscalYear"`
			PhysicalInventoryDocument      string `json:"PhysicalInventoryDocument"`
			PhysicalInventoryDocumentItem  string `json:"PhysicalInventoryDocumentItem"`
			Plant                          string `json:"Plant"`
			StorageLocation                string `json:"StorageLocation"`
			Material                       string `json:"Material"`
			Batch                          string `json:"Batch"`
			InventorySpecialStockType      string `json:"InventorySpecialStockType"`
			PhysicalInventoryStockType     string `json:"PhysicalInventoryStockType"`
			SalesOrder                     string `json:"SalesOrder"`
			SalesOrderItem                 string `json:"SalesOrderItem"`
			Supplier                       string `json:"Supplier"`
			Customer                       string `json:"Customer"`
			WBSElement                     string `json:"WBSElement"`
			LastChangeDate                 string `json:"LastChangeDate"`
			CountedByUser                  string `json:"CountedByUser"`
			PhysicalInventoryLastCountDate string `json:"PhysicalInventoryLastCountDate"`
			AdjustmentPostingMadeByUser    string `json:"AdjustmentPostingMadeByUser"`
			PostingDate                    string `json:"PostingDate"`
			PhysicalInventoryItemIsCounted bool   `json:"PhysicalInventoryItemIsCounted"`
			PhysInvtryDifferenceIsPosted   bool   `json:"PhysInvtryDifferenceIsPosted"`
			PhysInvtryItemIsRecounted      bool   `json:"PhysInvtryItemIsRecounted"`
			PhysInvtryItemIsDeleted        bool   `json:"PhysInvtryItemIsDeleted"`
			IsHandledInAltvUnitOfMsr       bool   `json:"IsHandledInAltvUnitOfMsr"`
			CycleCountType                 string `json:"CycleCountType"`
			IsValueOnlyMaterial            bool   `json:"IsValueOnlyMaterial"`
			PhysInventoryReferenceNumber   string `json:"PhysInventoryReferenceNumber"`
			MaterialDocument               string `json:"MaterialDocument"`
			MaterialDocumentYear           string `json:"MaterialDocumentYear"`
			MaterialDocumentItem           string `json:"MaterialDocumentItem"`
			PhysInvtryRecountDocument      string `json:"PhysInvtryRecountDocument"`
			PhysicalInventoryItemIsZero    bool   `json:"PhysicalInventoryItemIsZero"`
			ReasonForPhysInvtryDifference  string `json:"ReasonForPhysInvtryDifference"`
			MaterialBaseUnit               string `json:"MaterialBaseUnit"`
			BookQtyBfrCountInMatlBaseUnit  string `json:"BookQtyBfrCountInMatlBaseUnit"`
			Quantity                       string `json:"Quantity"`
			UnitOfEntry                    string `json:"UnitOfEntry"`
			QuantityInUnitOfEntry          string `json:"QuantityInUnitOfEntry"`
			Currency                       string `json:"Currency"`
			LastChangeDateTime             string `json:"LastChangeDateTime"`
		} `json:"results"`
	} `json:"d"`
}
