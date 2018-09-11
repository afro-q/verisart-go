package certificate

import (
	"encoding/json"
	
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"

	businessTypes "github.com/quinlanmorake/verisart-go/types"

	dbTypes "github.com/quinlanmorake/verisart-go/database/types"	
	database "github.com/quinlanmorake/verisart-go/database"
	
	tableNames "github.com/quinlanmorake/verisart-go/database/types/tableNames"
)

func LoadAllTransfers() ([]businessTypes.Transfer, coreTypes.Result) {
	var transfers []businessTypes.Transfer

	dataHandler := dbTypes.DataHandler(func (dbRows [][]byte) coreTypes.Result {
		transfers = make([]businessTypes.Transfer, len(dbRows))
		
		for index, dataRow := range dbRows {
			transfers[index] = businessTypes.Transfer{}
			
			if unmarshalError := json.Unmarshal(dataRow, &transfers[index]); unmarshalError != nil {
				return coreTypes.NewResultFromErrorCode(errorCodes.ERROR_UNMARSHALING_TRANSFER_RECORD_FROM_DATABASE)
			}
		}
		
		return coreTypes.NewSuccessResult()
	})

	return transfers, database.Load(tableNames.TRANSFERS, dataHandler)
}
