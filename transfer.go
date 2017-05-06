package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type Bid struct {
	companyName				string			`json:"companyname"`
	logMessage              string           `json:"logmessage"`
	DateTime				string			`json:"datetime"`
	TxID                    string          `json:"txid"`
}
var bidLogIndexStr = "bidLogs"
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "logging" {
		return t.logging(stub, args)
	}   
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" {
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}



func (t *SimpleChaincode) logging(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("running logging()")
	var txid string
	val := new(Bid)
    
	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3 (companyName, logMessage, DateTime).")
	}
	
	
	txid = stub.GetTxID()
		val.companyName = args[0]
		val.logMessage = args[1]
		val.DateTime = args[2]
		val.TxID = txid
	bidAsBytes, _ := json.Marshal(val)
	
	stub.PutState(bidLogIndexStr, bidAsBytes)
		
		
	return nil, nil	
}


func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	bidIndexBytes, err := stub.GetState(bidLogIndexStr)
	if err != nil { return nil, errors.New("Failed to read bids index")}

	var bidIndex []string
	err = json.Unmarshal(bidIndexBytes, &bidIndex)
	if err != nil { return nil, errors.New("Could not marshal bid indexes") }
	bidsJson, err := json.Marshal(bidIndex)
	if err != nil { return nil, errors.New("Failed to marshal bids to JSON")}

	return bidsJson, nil

}
