# Hyperledger Fabric Go chaincode testing utilities

These utilities are helpful during chaincode testing since they include:

*  CheckInvoke - returns the result of an invoke if it succeeded, otherwise fails the test
*  CheckBadInvoke - expect an invoke to return an error, otherwise fails the test
*  CheckState - compares the passed value with the stored one and fails the test if they are not similar
*  CheckNoState - expects no value by the passed key and fails the test if one was found
*  CheckInit - calls Init() function and fails the test in case of an error

These functions receive a MockStub implemented in [this repository](https://bitbucket.org/yurii_uhlanov_intellecteu/fabric-go-mockstub-impl/).

### Installation
To include the module in your project just add it to your dependencies:
```Go
import "bitbucket.org/yurii_uhlanov_intellecteu/fabric-go-test-utils"
```

