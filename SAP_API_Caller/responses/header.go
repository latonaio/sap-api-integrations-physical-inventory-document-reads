package responses

type Header struct {
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
			InventoryTransactionType       string `json:"InventoryTransactionType"`
			Plant                          string `json:"Plant"`
			StorageLocation                string `json:"StorageLocation"`
			InventorySpecialStockType      string `json:"InventorySpecialStockType"`
			DocumentDate                   string `json:"DocumentDate"`
			PhysInventoryPlannedCountDate  string `json:"PhysInventoryPlannedCountDate"`
			PhysicalInventoryLastCountDate string `json:"PhysicalInventoryLastCountDate"`
			PostingDate                    string `json:"PostingDate"`
			FiscalPeriod                   string `json:"FiscalPeriod"`
			PostingIsBlockedForPhysInvtry  bool   `json:"PostingIsBlockedForPhysInvtry"`
			PhysicalInventoryCountStatus   string `json:"PhysicalInventoryCountStatus"`
			PhysInvtryAdjustmentPostingSts string `json:"PhysInvtryAdjustmentPostingSts"`
			PhysInvtryDeletionStatus       string `json:"PhysInvtryDeletionStatus"`
			PhysInvtryDocHasQtySnapshot    bool   `json:"PhysInvtryDocHasQtySnapshot"`
			PhysicalInventoryGroupType     string `json:"PhysicalInventoryGroupType"`
			PhysicalInventoryGroup         string `json:"PhysicalInventoryGroup"`
			PhysicalInventoryNumber        string `json:"PhysicalInventoryNumber"`
			PhysInventoryReferenceNumber   string `json:"PhysInventoryReferenceNumber"`
			PhysicalInventoryDocumentDesc  string `json:"PhysicalInventoryDocumentDesc"`
			PhysicalInventoryType          string `json:"PhysicalInventoryType"`
			LastChangeDateTime             string `json:"LastChangeDateTime"`
			ToItem                         struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PhysicalInventoryDocumentItem"`
		} `json:"results"`
	} `json:"d"`
}
