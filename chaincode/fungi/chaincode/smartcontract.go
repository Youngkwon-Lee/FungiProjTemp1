package chaincode

import (
	"strconv"
	"fmt"
	"time"
	"crypto/sha256"
	"encoding/binary"
	"math"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct{
	contractapi.Contract
}

// Fungus Asset describes basic details
type Fungus struct {
	FungusId 	uint	`json:"fungusid"`
	Name	 	string	`json:"name"`
	Owner    	string	`json:"owner"`
	Dna		 	uint	`json:"dna"`
	ReadyTime	uint32	`json:"readytime"`
}

// Define Key names for options
const fungusCountKey = "FungusCount"

// init the chaincode
func (s *SmartContract) Initialized (ctx contractapi.TransactionContextInterface) error {

	//예외처리를 한후!!

	//Check authorization
	clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get MSPID : %v ", err)
	}
	if clientMSPID != "Org1MSP" {
		return fmt.Errorf("client is not authorized to initalize fugusCount : %v", err)
	}

	//check contract is not already set
	fungusCount, err := ctx.GetStub().GetState(fungusCountKey)
	if err != nil {
		return fmt.Errorf("failed to get fungusCount : %v ", err)
	}
	if fungusCount != nil {
		return fmt.Errorf("fungusCount is already set : %v ", err)
	}
	// fungusId에 사용될 count 값을 0으로 초기화
	err = ctx.GetStub().PutState(fungusCountKey, []byte(strconv.Itoa(0)))
	if err != nil {
		return fmt.Errorf("failed to set fungusCount : %v ", err)
	}

	return err
}

func (s *SmartContract) CreateRandomFungus (ctx contractapi.TransactionContextInterface, name string) error {

		//Check aclientID
	clientId, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return fmt.Errorf("failed to get ClientID : %v ", err)
	}
	exists, err := s._assetExists(ctx. clientId)
	if err !=nil{
		return err
	}
	if exists {
		return fmt.Errorf("client has alreadyc created an initial fungus")
	}
	 // 랜덤 DNA 생성
	 // 버섯 생성 및 원장에 저장하는 함수
	 //clientID 를 key 값으로 하는 값을 +1 저장 (보유한 버섯의 수)
	return nil
}


func (s *SmartContract) _generateRandomDna (ctx contractapi.TransactionContextInterface, name string) uint {

	
	
	unixTime := time.Now().Unix()
	data := strconv.Itoa(int(unixTime)) + name
	hash := sha256.New()
	hash.Write([]byte(data))
	dnaHash := uint(binary.BigEndian.Uint64(hash.Sum(nil)))
	

	// make 14digits dna
	dna := dnaHash % uint(math.Pow(10, 10))
	dna = dna -(dna % 100)
	
	return dna

}



// 외부에서 호출 가능한 interface....
// 초기화에 관련된 함수
// 버섯생성 (CreateRandomFungus )
// 버섯조회 (GentFungiByOwner )
