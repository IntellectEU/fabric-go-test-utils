package testutils

import (
	"testing"

	cms "bitbucket.org/yurii_uhlanov_intellecteu/fabric-go-mockstub-impl"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var testLog = shim.NewLogger("test_logger")

func stringsAsBytes(strings []string) (bytes [][]byte) {
	bytes = make([][]byte, len(strings))
	for i, val := range strings {
		bytes[i] = []byte(val)
	}

	return
}

func CheckInvoke(t *testing.T, stub *cms.CustomMockStub, functionAndArgs []string, transactionID string) string {
	res := stub.MockInvoke(transactionID, stringsAsBytes(functionAndArgs))
	if res.Status != shim.OK {
		testLog.Error("Invoke", functionAndArgs, "failed", string(res.Message))
		t.FailNow()
	} else {
		testLog.Info("Invoke", functionAndArgs, "successful", string(res.Message))
	}

	return string(res.Payload)
}

func CheckBadInvoke(t *testing.T, stub *cms.CustomMockStub, functionAndArgs []string, transactionID string) {
	res := stub.MockInvoke(transactionID, stringsAsBytes(functionAndArgs))
	if res.Status == shim.OK {
		testLog.Error("Invoke", functionAndArgs, "unexpectedly succeeded")
		t.FailNow()
	} else {
		testLog.Info("Invoke", functionAndArgs, "failed as espected, with message", string(res.Message))
	}
}

func CheckState(t *testing.T, stub *cms.CustomMockStub, key string, expectedValue []byte) {
	realValue := stub.State[key]
	if realValue == nil {
		testLog.Error("State", key, "failed to get value")
		t.FailNow()
	}
	if string(realValue) != string(expectedValue) {
		testLog.Error("State value", key, "was", string(realValue), "and not", string(expectedValue), "as expected")
		t.FailNow()
	} else {
		testLog.Info("State value", key, "is", string(realValue), "as expected")
	}
}

func CheckNoState(t *testing.T, stub *cms.CustomMockStub, key string) {
	realValue := stub.State[key]
	if realValue == nil {
		testLog.Info("State value", key, "was not found as expected")
	}
}

func CheckInit(t *testing.T, stub *cms.CustomMockStub, args []string, transactionID string) {
	res := stub.MockInit(transactionID, stringsAsBytes(args))
	if res.Status != shim.OK {
		testLog.Error("Init", args, "failed", string(res.Message))
		t.FailNow()
	} else {
		testLog.Info("Init", args, "successful", string(res.Message))
	}
}

func PutState(t *testing.T, stub *cms.CustomMockStub, transactionID, key string, value []byte) {
	stub.MockTransactionStart(transactionID)

	// Put the value in the ledger
	if err := stub.PutState(key, value); err != nil {
		testLog.Error("Failed to wright the value", string(value), err.Error())
		t.FailNow()
	}

	stub.MockTransactionEnd(transactionID)
}
