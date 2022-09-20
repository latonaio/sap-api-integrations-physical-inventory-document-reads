package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-physical-inventory-document-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			FiscalYear:                     data.FiscalYear,
			PhysicalInventoryDocument:      data.PhysicalInventoryDocument,
			InventoryTransactionType:       data.InventoryTransactionType,
			Plant:                          data.Plant,
			StorageLocation:                data.StorageLocation,
			InventorySpecialStockType:      data.InventorySpecialStockType,
			DocumentDate:                   data.DocumentDate,
			PhysInventoryPlannedCountDate:  data.PhysInventoryPlannedCountDate,
			PhysicalInventoryLastCountDate: data.PhysicalInventoryLastCountDate,
			PostingDate:                    data.PostingDate,
			FiscalPeriod:                   data.FiscalPeriod,
			PostingIsBlockedForPhysInvtry:  data.PostingIsBlockedForPhysInvtry,
			PhysicalInventoryCountStatus:   data.PhysicalInventoryCountStatus,
			PhysInvtryAdjustmentPostingSts: data.PhysInvtryAdjustmentPostingSts,
			PhysInvtryDeletionStatus:       data.PhysInvtryDeletionStatus,
			PhysInvtryDocHasQtySnapshot:    data.PhysInvtryDocHasQtySnapshot,
			PhysicalInventoryGroupType:     data.PhysicalInventoryGroupType,
			PhysicalInventoryGroup:         data.PhysicalInventoryGroup,
			PhysicalInventoryNumber:        data.PhysicalInventoryNumber,
			PhysInventoryReferenceNumber:   data.PhysInventoryReferenceNumber,
			PhysicalInventoryDocumentDesc:  data.PhysicalInventoryDocumentDesc,
			PhysicalInventoryType:          data.PhysicalInventoryType,
			LastChangeDateTime:             data.LastChangeDateTime,
			ToItem:                         data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
			FiscalYear:                     data.FiscalYear,
			PhysicalInventoryDocument:      data.PhysicalInventoryDocument,
			PhysicalInventoryDocumentItem:  data.PhysicalInventoryDocumentItem,
			Plant:                          data.Plant,
			StorageLocation:                data.StorageLocation,
			Material:                       data.Material,
			Batch:                          data.Batch,
			InventorySpecialStockType:      data.InventorySpecialStockType,
			PhysicalInventoryStockType:     data.PhysicalInventoryStockType,
			SalesOrder:                     data.SalesOrder,
			SalesOrderItem:                 data.SalesOrderItem,
			Supplier:                       data.Supplier,
			Customer:                       data.Customer,
			WBSElement:                     data.WBSElement,
			LastChangeDate:                 data.LastChangeDate,
			CountedByUser:                  data.CountedByUser,
			PhysicalInventoryLastCountDate: data.PhysicalInventoryLastCountDate,
			AdjustmentPostingMadeByUser:    data.AdjustmentPostingMadeByUser,
			PostingDate:                    data.PostingDate,
			PhysicalInventoryItemIsCounted: data.PhysicalInventoryItemIsCounted,
			PhysInvtryDifferenceIsPosted:   data.PhysInvtryDifferenceIsPosted,
			PhysInvtryItemIsRecounted:      data.PhysInvtryItemIsRecounted,
			PhysInvtryItemIsDeleted:        data.PhysInvtryItemIsDeleted,
			IsHandledInAltvUnitOfMsr:       data.IsHandledInAltvUnitOfMsr,
			CycleCountType:                 data.CycleCountType,
			IsValueOnlyMaterial:            data.IsValueOnlyMaterial,
			PhysInventoryReferenceNumber:   data.PhysInventoryReferenceNumber,
			MaterialDocument:               data.MaterialDocument,
			MaterialDocumentYear:           data.MaterialDocumentYear,
			MaterialDocumentItem:           data.MaterialDocumentItem,
			PhysInvtryRecountDocument:      data.PhysInvtryRecountDocument,
			PhysicalInventoryItemIsZero:    data.PhysicalInventoryItemIsZero,
			ReasonForPhysInvtryDifference:  data.ReasonForPhysInvtryDifference,
			MaterialBaseUnit:               data.MaterialBaseUnit,
			BookQtyBfrCountInMatlBaseUnit:  data.BookQtyBfrCountInMatlBaseUnit,
			Quantity:                       data.Quantity,
			UnitOfEntry:                    data.UnitOfEntry,
			QuantityInUnitOfEntry:          data.QuantityInUnitOfEntry,
			Currency:                       data.Currency,
			LastChangeDateTime:             data.LastChangeDateTime,
		})
	}

	return item, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
			FiscalYear:                     data.FiscalYear,
			PhysicalInventoryDocument:      data.PhysicalInventoryDocument,
			PhysicalInventoryDocumentItem:  data.PhysicalInventoryDocumentItem,
			Plant:                          data.Plant,
			StorageLocation:                data.StorageLocation,
			Material:                       data.Material,
			Batch:                          data.Batch,
			InventorySpecialStockType:      data.InventorySpecialStockType,
			PhysicalInventoryStockType:     data.PhysicalInventoryStockType,
			SalesOrder:                     data.SalesOrder,
			SalesOrderItem:                 data.SalesOrderItem,
			Supplier:                       data.Supplier,
			Customer:                       data.Customer,
			WBSElement:                     data.WBSElement,
			LastChangeDate:                 data.LastChangeDate,
			CountedByUser:                  data.CountedByUser,
			PhysicalInventoryLastCountDate: data.PhysicalInventoryLastCountDate,
			AdjustmentPostingMadeByUser:    data.AdjustmentPostingMadeByUser,
			PostingDate:                    data.PostingDate,
			PhysicalInventoryItemIsCounted: data.PhysicalInventoryItemIsCounted,
			PhysInvtryDifferenceIsPosted:   data.PhysInvtryDifferenceIsPosted,
			PhysInvtryItemIsRecounted:      data.PhysInvtryItemIsRecounted,
			PhysInvtryItemIsDeleted:        data.PhysInvtryItemIsDeleted,
			IsHandledInAltvUnitOfMsr:       data.IsHandledInAltvUnitOfMsr,
			CycleCountType:                 data.CycleCountType,
			IsValueOnlyMaterial:            data.IsValueOnlyMaterial,
			PhysInventoryReferenceNumber:   data.PhysInventoryReferenceNumber,
			MaterialDocument:               data.MaterialDocument,
			MaterialDocumentYear:           data.MaterialDocumentYear,
			MaterialDocumentItem:           data.MaterialDocumentItem,
			PhysInvtryRecountDocument:      data.PhysInvtryRecountDocument,
			PhysicalInventoryItemIsZero:    data.PhysicalInventoryItemIsZero,
			ReasonForPhysInvtryDifference:  data.ReasonForPhysInvtryDifference,
			MaterialBaseUnit:               data.MaterialBaseUnit,
			BookQtyBfrCountInMatlBaseUnit:  data.BookQtyBfrCountInMatlBaseUnit,
			Quantity:                       data.Quantity,
			UnitOfEntry:                    data.UnitOfEntry,
			QuantityInUnitOfEntry:          data.QuantityInUnitOfEntry,
			Currency:                       data.Currency,
			LastChangeDateTime:             data.LastChangeDateTime,
		})
	}

	return toItem, nil
}
