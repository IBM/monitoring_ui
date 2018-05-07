/*******************************************************************************
Copyright (c) 2016 IBM Corporation and other Contributors.


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.


Contributors:

Sumabala Nair - Initial Contribution
Kim Letkeman - Initial Contribution
Sumabala Nair - Updated for hyperledger May 2016
Sumabala Nair - Partial updates added May 2016
******************************************************************************/
//SN: March 2016

// IoT Blockchain Simple Smart Contract v 1.0

// This is a simple contract that creates a CRUD interface to
// create, read, update and delete an asset

package main

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "strings"
     "reflect"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)


// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

const CONTRACTSTATEKEY string = "ContractStateKey"
// store contract state - only version in this example
const MYVERSION string = "1.0"

// ************************************
// asset and contract state
// ************************************

type ContractState struct {
    Version      string                        `json:"version"`
}

type Geolocation struct {
    Latitude    *float64 `json:"latitude,omitempty"`
    Longitude   *float64 `json:"longitude,omitempty"`
}

type AssetState struct {
    AssetID        *string       `json:"assetID,omitempty"`        // all assets must have an ID, primary key of contract
    Location       *Geolocation  `json:"location,omitempty"`       // current asset location
    Temperature    *float64      `json:"temperature,omitempty"`    // asset temp
    Carrier        *string       `json:"carrier,omitempty"`        // the name of the carrier
}

var contractState = ContractState{MYVERSION}


// ************************************
// deploy callback mode
// ************************************
// func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) pb.Response {
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
    fmt.Println("Contract Is Starting Up")

    function, args := stub.GetFunctionAndParameters()
    fmt.Println(function)
    fmt.Println(args)

    var stateArg ContractState
    var err error
    var buffer bytes.Buffer

    if len(args) != 1 {
        // return nil, errors.New("init expects one argument, a JSON string with tagged version string")
        return shim.Error("init expects one argument, a JSON string with tagged version string")
    }
    err = json.Unmarshal([]byte(args[0]), &stateArg)
    if err != nil {
        // return nil, errors.New("Version argument unmarshal failed: " + fmt.Sprint(err))
        return shim.Error("Version argument unmarshal failed")
    }
    if stateArg.Version != MYVERSION {
        // return nil, errors.New("Contract version " + MYVERSION + " must match version argument: " + stateArg.Version)
        return shim.Error("Contract version ")
    }
    contractStateJSON, err := json.Marshal(stateArg)
    if err != nil {
        // return nil, errors.New("Marshal failed for contract state" + fmt.Sprint(err))
        return shim.Error("Marshal failed for contract state")
    }
    err = stub.PutState(CONTRACTSTATEKEY, contractStateJSON)
    if err != nil {
        // return nil, errors.New("Contract state failed PUT to ledger: " + fmt.Sprint(err))
        return shim.Error("Contract state failed PUT to ledger: ")
    }
    buffer.WriteString("Contract Initialized")
    return shim.Success(buffer.Bytes())
}

// ************************************
// deploy and invoke callback mode
// ************************************
// func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    function, args := stub.GetFunctionAndParameters()
    // Handle different functions
    if function == "createAsset" {
        // create assetID
        return t.createAsset(stub, args)
    } else if function == "updateAsset" {
        // create assetID
        return t.updateAsset(stub, args)
    } else if function == "deleteAsset" {
        // Deletes an asset by ID from the ledger
        return t.deleteAsset(stub, args)
    } else if function == "readAsset" {
      // gets the state for an assetID as a JSON struct
      return t.readAsset(stub, args)
    } else if function =="readAssetObjectModel" {
        return t.readAssetObjectModel(stub, args)
    } else if function == "readAssetSamples" {
		// returns selected sample objects
		return t.readAssetSamples(stub, args)
  	} else if function == "readAssetSchemas" {
  		// returns selected sample objects
  		return t.readAssetSchemas(stub, args)
  	}

    return shim.Error("Received unknown invocation: " + function )
}

// ************************************
// query callback mode
// ************************************
// func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface) pb.Response {
//
//     function, args := stub.GetFunctionAndParameters()
//     // Handle different functions
//     if function == "readAsset" {
//       // gets the state for an assetID as a JSON struct
//       return t.readAsset(stub, args)
//     } else if function =="readAssetObjectModel" {
//         return t.readAssetObjectModel(stub, args)
//     } else if function == "readAssetSamples" {
// 		// returns selected sample objects
// 		return t.readAssetSamples(stub, args)
//   	} else if function == "readAssetSchemas" {
//   		// returns selected sample objects
//   		return t.readAssetSchemas(stub, args)
//   	}
//     // return nil, errors.New("Received unknown invocation: " + function)
//     return shim.Error("Received unknown invocation: " + function )
// }

/**********main implementation *************/

func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple Chaincode: %s", err)
    }
}

/*****************ASSET CRUD INTERFACE starts here************/

/****************** 'deploy' methods *****************/

/******************** createAsset ********************/

// func (t *SimpleChaincode) createAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) pb.Response {
func (t *SimpleChaincode) createAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    erval:= t.createOrUpdateAsset(stub, args)
    return erval
}

//******************** updateAsset ********************/

// func (t *SimpleChaincode) updateAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) pb.Response {
func (t *SimpleChaincode) updateAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    erval:= t.createOrUpdateAsset(stub, args)
    return erval
}


//******************** deleteAsset ********************/

func (t *SimpleChaincode) deleteAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var assetID string // asset ID
    var err error
    var stateIn AssetState

    // validate input data for number of args, Unmarshaling to asset state and obtain asset id
    stateIn, err = t.validateInput(args)
    if err != nil {
      return shim.Error(err.Error())
    }
    assetID = *stateIn.AssetID
    // Delete the key / asset from the ledger
    err = stub.DelState(assetID)
    if err != nil {
        // err = errors.New("DELSTATE failed! : "+ fmt.Sprint(err))
       return shim.Error(err.Error())
    }
    return shim.Success(nil)
}

/******************* Query Methods ***************/

//********************readAsset********************/

func (t *SimpleChaincode) readAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var assetID string // asset ID
    var err error
    var state AssetState

     // validate input data for number of args, Unmarshaling to asset state and obtain asset id

    stateIn, err:= t.validateInput(args)
    if err != nil {
        // return nil, errors.New("Asset does not exist!")
        return shim.Error("Asset does not exist!")
    }
    assetID = *stateIn.AssetID
        // Get the state from the ledger
    assetBytes, err:= stub.GetState(assetID)
    if err != nil  || len(assetBytes) ==0{
        err = errors.New("Unable to get asset state from ledger")
        // return nil, err
        return shim.Error(err.Error())
    }
    err = json.Unmarshal(assetBytes, &state)
    if err != nil {
        err = errors.New("Unable to unmarshal state data obtained from ledger")
        return shim.Error(err.Error())
    }
    // return assetBytes, nil
    return shim.Success(assetBytes)
}

//*************readAssetObjectModel*****************/

func (t *SimpleChaincode) readAssetObjectModel(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var state AssetState = AssetState{}

    // Marshal and return
    stateJSON, err := json.Marshal(state)
    if err != nil {
        return shim.Error(err.Error())
    }
    // return stateJSON, nil
    return shim.Success(stateJSON)
}
//*************readAssetSamples*******************/

func (t *SimpleChaincode) readAssetSamples(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  return shim.Success([]byte(samples))
	// return []byte(samples), nil
}
//*************readAssetSchemas*******************/

func (t *SimpleChaincode) readAssetSchemas(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  return shim.Success([]byte(schemas))
	// return []byte(schemas), nil
}

// ************************************
// validate input data : common method called by the CRUD functions
// ************************************
func (t *SimpleChaincode) validateInput(args []string) (stateIn AssetState, err error)  {
    var assetID string // asset ID
    var state AssetState = AssetState{} // The calling function is expecting an object of type AssetState

    if len(args) !=1 {
        err = errors.New("Incorrect number of arguments. Expecting a JSON strings with mandatory assetID")
        return state, err
    }
    jsonData:=args[0]
    assetID = ""
    stateJSON := []byte(jsonData)
    err = json.Unmarshal(stateJSON, &stateIn)
    if err != nil {
        err = errors.New("Unable to unmarshal input JSON data")
        return state, err
        // shim.Error(err.Error())
        // state is an empty instance of asset state
    }
    // was assetID present?
    // The nil check is required because the asset id is a pointer.
    // If no value comes in from the json input string, the values are set to nil

    if stateIn.AssetID !=nil {
        assetID = strings.TrimSpace(*stateIn.AssetID)
        if assetID=="" {
            err = errors.New("AssetID not passed")
            return state, err
        }
    } else {
        err = errors.New("Asset id is mandatory in the input JSON data")
        return state, err
        // shim.Error(err.Error())
    }

    stateIn.AssetID = &assetID
    return stateIn, nil
}
//******************** createOrUpdateAsset ********************/

func (t *SimpleChaincode) createOrUpdateAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var assetID string                 // asset ID                    // used when looking in map
    var err error
    var stateIn AssetState
    var stateStub AssetState


    // validate input data for number of args, Unmarshaling to asset state and obtain asset id

    stateIn, err = t.validateInput(args)
    if err != nil {
        return shim.Error(err.Error())
    }
    assetID = *stateIn.AssetID
    // Partial updates introduced here
    // Check if asset record existed in stub
    assetBytes, err:= stub.GetState(assetID)
    if err != nil || len(assetBytes)==0{
        // This implies that this is a 'create' scenario
         stateStub = stateIn // The record that goes into the stub is the one that cme in
    } else {
        // This is an update scenario
        err = json.Unmarshal(assetBytes, &stateStub)
        if err != nil {
            err = errors.New("Unable to unmarshal JSON data from stub")
            // return nil, err
            return shim.Error(err.Error())
            // state is an empty instance of asset state
        }
          // Merge partial state updates
        stateStub, err =t.mergePartialState(stateStub,stateIn)
        if err != nil {
            err = errors.New("Unable to merge state")
            // return nil,err
            return shim.Error(err.Error())
        }
    }
    stateJSON, err := json.Marshal(stateStub)
    if err != nil {
        // return nil, errors.New("Marshal failed for contract state" + fmt.Sprint(err))
        return shim.Error(err.Error())
    }
    // Get existing state from the stub


    // Write the new state to the ledger
    err = stub.PutState(assetID, stateJSON)
    if err != nil {
        err = errors.New("PUT ledger state failed: "+ fmt.Sprint(err))
        return shim.Error(err.Error())
    }
    return shim.Success(nil)
}
/*********************************  internal: mergePartialState ****************************/
 func (t *SimpleChaincode) mergePartialState(oldState AssetState, newState AssetState) (AssetState,  error) {

    old := reflect.ValueOf(&oldState).Elem()
    new := reflect.ValueOf(&newState).Elem()
    for i := 0; i < old.NumField(); i++ {
        oldOne:=old.Field(i)
        newOne:=new.Field(i)
        if ! reflect.ValueOf(newOne.Interface()).IsNil() {
            oldOne.Set(reflect.Value(newOne))
        }
    }
    return oldState, nil
 }
