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

func LoadAllCertificates() ([]businessTypes.Certificate, coreTypes.Result) {
	var certificates []businessTypes.Certificate

	dataHandler := dbTypes.DataHandler(func (dbRows [][]byte) coreTypes.Result {
		certificates = make([]businessTypes.Certificate, len(dbRows))
		
		for index, dataRow := range dbRows {
			certificates[index] = businessTypes.Certificate{}
			
			if unmarshalError := json.Unmarshal(dataRow, &certificates[index]); unmarshalError != nil {
				return coreTypes.NewResultFromErrorCode(errorCodes.ERROR_UNMARSHALING_CERTIFICATE_RECORD_FROM_DATABASE)
			}
		}
		
		return coreTypes.NewSuccessResult()
	})

	return certificates, database.Load(tableNames.CERTIFICATES, dataHandler)
}