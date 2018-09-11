package transferStatus

/*
 We put these in a subpackage so as to make it clearer to use them in the code; that is, the intended use is
  a := transferStatus.CREATED
*/

type TransferStatus string

const (
	CREATED  TransferStatus = "CREATED"
	ACCEPTED TransferStatus = "ACCEPTED"
	ANNULED  TransferStatus = "ANNULED"
)

/*
 No business rule has been defined for the following status's

	CANCELED   coreTypes.TransferStatus = "CANCELED"	  // Originator decides to cancel the transfer
	REJECTED   coreTypes.TransferStatus = "REJECTED"    // Receiver decides to reject the transfer

 for a transfer that has been waiting indefinitely; as such no state is
 defined for it yet.
*/
