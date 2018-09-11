package memory

import (
	goMemDb "github.com/hashicorp/go-memdb"

	tableNames "github.com/quinlanmorake/verisart-go/database/types/tableNames"
)

var dbSchema goMemDb.DBSchema = goMemDb.DBSchema{
	Tables: map[string]*goMemDb.TableSchema{
		tableNames.USERS: &goMemDb.TableSchema{
			Name:    tableNames.USERS,
			Indexes: getIndexesForUsersTable(),
		},

		tableNames.CERTIFICATES: &goMemDb.TableSchema{
			Name:    tableNames.CERTIFICATES,
			Indexes: getIndexesForCertificatesTable(),
		},

		tableNames.TRANSFERS: &goMemDb.TableSchema{
			Name:    tableNames.TRANSFERS,
			Indexes: getIndexesForTransfersTable(),
		},
	},
}

// This library uses reflection to determine how the rows map to columns,
// so the value for the "Field" in the indexer needs to be the name of the property on the object
func getIndexesForUsersTable() map[string]*goMemDb.IndexSchema {
	return map[string]*goMemDb.IndexSchema{
		"id": &goMemDb.IndexSchema{
			Name:    "id",
			Unique:  true,
			Indexer: &goMemDb.StringFieldIndex{Field: "Id"},
		},
	}
}

func getIndexesForCertificatesTable() map[string]*goMemDb.IndexSchema {
	return map[string]*goMemDb.IndexSchema{
		"id": &goMemDb.IndexSchema{
			Name:    "id",
			Unique:  true,
			Indexer: &goMemDb.StringFieldIndex{Field: "Id"},
		},

		"userId": &goMemDb.IndexSchema{
			Name:    "userId",
			Unique:  false,
			Indexer: &goMemDb.StringFieldIndex{Field: "OwnerId"},
		},
	}
}

func getIndexesForTransfersTable() map[string]*goMemDb.IndexSchema {
	return map[string]*goMemDb.IndexSchema{
		"id": &goMemDb.IndexSchema{
			Name:    "id",
			Unique:  true,
			Indexer: &goMemDb.StringFieldIndex{Field: "Id"},
		},

		"userId": &goMemDb.IndexSchema{
			Name:    "userId",
			Unique:  false,
			Indexer: &goMemDb.StringFieldIndex{Field: "CreatedBy"},
		},
	}
}
