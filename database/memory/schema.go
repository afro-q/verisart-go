package memory

import (
	goMemDb "github.com/hashicorp/go-memdb"
)

var dbSchema goMemDb.DBSchema = goMemDb.DBSchema{
	Tables: map[string]*goMemDb.TableSchema{
		"users": &goMemDb.TableSchema{
			Name:    "users",
			Indexes: getIndexesForUsersTable(),
		},

		"certificates": &goMemDb.TableSchema{
			Name:    "certificates",
			Indexes: getIndexesForCertificatesTable(),
		},

		"transfers": &goMemDb.TableSchema{
			Name: "transfers",
			Indexes: getIndexesForTransfersTable(), 
		},
	},
}

func getIndexesForUsersTable() map[string]*goMemDb.IndexSchema {
	return map[string]*goMemDb.IndexSchema{
		"id": &goMemDb.IndexSchema{
			Name:    "id",
			Unique:  true,
			Indexer: &goMemDb.StringFieldIndex{Field: "id"},
		},
	}
}

func getIndexesForCertificatesTable() map[string]*goMemDb.IndexSchema {
	return map[string]*goMemDb.IndexSchema{
		"id": &goMemDb.IndexSchema{
			Name:    "id",
			Unique:  true,
			Indexer: &goMemDb.StringFieldIndex{Field: "id"},
		},

		"userId": &goMemDb.IndexSchema{
			Name:    "userId",
			Unique:  false,
			Indexer: &goMemDb.StringFieldIndex{Field: "ownerId"},
		},
	}
}

func getIndexesForTransfersTable() map[string]*goMemDb.IndexSchema {
	return map[string]*goMemDb.IndexSchema{
		"id": &goMemDb.IndexSchema{
			Name:    "id",
			Unique:  true,
			Indexer: &goMemDb.StringFieldIndex{Field: "id"},
		},

		"userId": &goMemDb.IndexSchema{
			Name:    "userId",
			Unique:  false,
			Indexer: &goMemDb.StringFieldIndex{Field: "createdBy"},
		},
	}
}
