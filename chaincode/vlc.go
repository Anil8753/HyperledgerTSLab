/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const keyInsurance = "insurance_"
const keyRegistration = "registration_"
const keyService = "service_"

// SmartContract provides functions for managing the life cycle of the vehicles
type SmartContract struct {
	contractapi.Contract
}

type RegData struct {
	RegNumber      string `json:"regNumber"`
	ChassisNumber  string `json:"chassisNumber"`
	EngineNumber   string `json:"engineNumber"`
	MonthYearOfMfg string `json:"monthYearOfMfg"`
}

type ServiceData struct {
	RegNumber      string `json:"RegNumber"`
	ChassisNumber  string `json:"ChassisNumber"`
	EngineNumber   string `json:"EngineNumber"`
	MonthYearOfMfg string `json:"MonthYearOfMfg"`

	ServiceDetails string `json:"ServiceDetails"`
}

type InsuranceData struct {
	RegNumber string `json:"RegNumber"`

	UINNumber             string `json:"UINNumber"`
	PolicyNumber          string `json:"PolicyNumber"`
	InsuredNameAndAddress string `json:"InsuredNameAndAddress"`
	ContactNumber         string `json:"ContactNumber"`
	EmailId               string `json:"EmailId"`
	PeriodOfCover         string `json:"PeriodOfCover"`
	PremiumDetails        string `json:"PremiumDetails"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	return nil
}

///////////////////////////////////////////////////////////////
////////////////// Vehicle registration data  /////////////////
///////////////////////////////////////////////////////////////
func (s *SmartContract) SetRegData(
	ctx contractapi.TransactionContextInterface,
	regNumber string,
	chassisNumber string,
	engineNumber string,
	monthYearOfMfg string,
) error {

	data := RegData{
		RegNumber:      regNumber,
		ChassisNumber:  chassisNumber,
		EngineNumber:   engineNumber,
		MonthYearOfMfg: monthYearOfMfg,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("falied to parse : %s", err.Error())
		return err
	}

	key := keyRegistration + regNumber
	return ctx.GetStub().PutState(key, bytes)
}

func (s *SmartContract) GetRegData(
	ctx contractapi.TransactionContextInterface,
	regNumber string,
) (*RegData, error) {

	key := keyRegistration + regNumber
	bytes, err := ctx.GetStub().GetState(key)

	if err != nil {
		return nil, fmt.Errorf("failed to read from world state. %s", err.Error())
	}

	if bytes == nil {
		return nil, fmt.Errorf("%s does not exist", regNumber)
	}

	data := new(RegData)
	_ = json.Unmarshal(bytes, data)

	return data, nil
}

func (s *SmartContract) GetRegDataHistory(ctx contractapi.TransactionContextInterface, regNumber string) (string, error) {
	key := keyRegistration + regNumber
	return GetDataHistory(ctx, "GetRegDataHistory", key)
}

///////////////////////////////////////////////////////////////
////////////////// Vehicle insurance data  ////////////////////
///////////////////////////////////////////////////////////////
func (s *SmartContract) SetInsuranceData(
	ctx contractapi.TransactionContextInterface,
	regNumber string,
	uinNumber string,
	policyNumber string,
	insuredNameAndAddress string,
	contactNumber string,
	emailId string,
	periodOfCover string,
	premiumDetails string,
) error {

	data := InsuranceData{
		RegNumber:             regNumber,
		UINNumber:             uinNumber,
		PolicyNumber:          policyNumber,
		InsuredNameAndAddress: insuredNameAndAddress,
		ContactNumber:         contactNumber,
		EmailId:               emailId,
		PeriodOfCover:         periodOfCover,
		PremiumDetails:        premiumDetails,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("falied to parse : %s", err.Error())
		return err
	}

	key := keyInsurance + regNumber
	return ctx.GetStub().PutState(key, bytes)
}

func (s *SmartContract) GetInsuranceData(
	ctx contractapi.TransactionContextInterface,
	regNumber string,
) (*InsuranceData, error) {

	key := keyInsurance + regNumber
	bytes, err := ctx.GetStub().GetState(key)

	if err != nil {
		return nil, fmt.Errorf("failed to read from world state. %s", err.Error())
	}

	if bytes == nil {
		return nil, fmt.Errorf("%s does not exist", regNumber)
	}

	data := new(InsuranceData)
	_ = json.Unmarshal(bytes, data)

	return data, nil
}

func (s *SmartContract) GetInsuranceDataHistory(ctx contractapi.TransactionContextInterface, regNumber string) (string, error) {
	key := keyInsurance + regNumber
	return GetDataHistory(ctx, "GetInsuranceDataHistory", key)
}

///////////////////////////////////////////////////////////////
////////////////// Vehicle service data  //////////////////////
///////////////////////////////////////////////////////////////
func (s *SmartContract) SetServiceData(
	ctx contractapi.TransactionContextInterface,
	regNumber string,
	chassisNumber string,
	engineNumber string,
	monthYearOfMfg string,
	serviceDetails string,
) error {

	data := ServiceData{
		RegNumber:      regNumber,
		ChassisNumber:  chassisNumber,
		EngineNumber:   engineNumber,
		MonthYearOfMfg: monthYearOfMfg,
		ServiceDetails: serviceDetails,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("falied to parse : %s", err.Error())
		return err
	}

	key := keyService + regNumber
	return ctx.GetStub().PutState(key, bytes)
}

func (s *SmartContract) GetServiceData(
	ctx contractapi.TransactionContextInterface,
	regNumber string,
) (*ServiceData, error) {

	key := keyService + regNumber
	bytes, err := ctx.GetStub().GetState(key)

	if err != nil {
		return nil, fmt.Errorf("failed to read from world state. %s", err.Error())
	}

	if bytes == nil {
		return nil, fmt.Errorf("%s does not exist", regNumber)
	}

	data := new(ServiceData)
	_ = json.Unmarshal(bytes, data)

	return data, nil
}

// GetServiceDataHistory returns the complete servicing history
func (s *SmartContract) GetServiceDataHistory(ctx contractapi.TransactionContextInterface, regNumber string) (string, error) {
	key := keyService + regNumber
	return GetDataHistory(ctx, "GetServiceDataHistory", key)
}

//////////////////////////////////////////////////////////////
///////////////////////  Helpers //////////////////////////////
//////////////////////////////////////////////////////////////
func GetDataHistory(ctx contractapi.TransactionContextInterface, funcName string, identifier string) (string, error) {

	fmt.Printf("- start %s: %s\n", funcName, identifier)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(identifier)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return "", fmt.Errorf(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")

		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- %s returning:\n%s\n", funcName, buffer.String())

	return buffer.String(), nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create VLC chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting VLC chaincode: %s", err.Error())
	}
}
