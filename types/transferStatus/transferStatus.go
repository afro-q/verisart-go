package transferStatus

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types"
)

/*
 We put these in a subpackage so as to make it clearer to use them in the code; that is, the intended use is
  a := transferStatus.CREATED
*/

const (
	CREATED  coreTypes.TransferStatus = "CREATED"
	FAILED   coreTypes.TransferStatus = "FAILED"
	WAITING  coreTypes.TransferStatus = "WAITING"
	ACCEPTED coreTypes.TransferStatus = "ACCEPTED"
	ANNULED  coreTypes.TransferStatus = "ANNULED"
)

/*
 No business rule has been defined for the following status's

	CANCELED   coreTypes.TransferStatus = "CANCELED"	  // Originator decides to cancel the transfer
	REJECTED   coreTypes.TransferStatus = "REJECTED"    // Receiver decides to reject the transfer

 for a transfer that has been waiting indefinitely; as such no state is
 defined for it yet.
*/
