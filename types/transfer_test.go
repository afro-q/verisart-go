package types

import (
	"testing"
)

func Test_NewTransferGeneratesATimeStamp(t *testing.T) {
	// Not going to validate that this generates the right timestamp
	transfer := NewTransfer()

	if len(transfer.CreatedAt) == 0 {
		t.Error("On initialization, the transfer should always have a value for `CreatedAt`")
	}
}

func Test_IsValidSucceedsWhenItShould(t *testing.T) {
}

func Test_IsValidFailsWhenCreatedAtIsEmpty(t *testing.T) {
}

func Test_IsValidFailsWhenTheUserEmailAddressIsInvalid(t *testing.T) {
}

func Test_IsValidFailsWhenTheUserEmailIsNotFound(t *testing.T) {
}
