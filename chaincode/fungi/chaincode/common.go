package chaincode

import (
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) _AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}


// 외부에서 호출 가능한 interface....
// 초기화에 관련된 함수
// 버섯생성 (CreateRandomFungus )
// 버섯조회 (GentFungiByOwner )