package chaincode

import (
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) _assetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

func (s *SmartContract) _getstate(ctx contractapi.TransactionContextInterface, id string) ([]byte, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}
	return assetJSON. nill

}

func (s *SmartContract) _updateOwnerFungusCount(ctx contractapi.TransactionContextInterface, clientID string, increment int ) error {
	countBtye, err := s.GetState(ctx, clientId)
	if countBtye == nil {
		ctx.GetStub().Putstate(clientId, []byte(strconv.IToa(1)))
		return nil
	}
	if err !=nil {
		return err
	}
	ownerfunguscount, _ := strconv.Atoi(countBtye[:])
	ownerfunguscount += increment
	err = ctx.GetStub().Putstate(clientId, []byte(strconv.IToa(ownerfunguscount)))
	if err != nil {
		return fmt.Errorf("failed to put /ownerFungusCount state: %v", err)
	}
	return nil
}
