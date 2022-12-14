package users

import (
	"context"
	"log"

	"github.com/vireocloud/property-pros-service/common"
	"github.com/vireocloud/property-pros-service/interfaces"
	"github.com/vireocloud/property-pros-service/interop"
)

type UserModel struct {
	*interop.User
	*common.BaseModel[interop.User]
	gateway interfaces.IUsersGateway
}

func (model *UserModel) GetContext() context.Context {
	return model.Context
}

func (model *UserModel) Save() (interfaces.IUserModel, error) {
	return model.gateway.SaveUser(model)
}

func (model *UserModel) HasAuthenticIdentity() (bool, error) {
	userModel, err := model.gateway.GetUserByUsername(model)

	if err != nil {
		return false, nil
	}

	return model.MatchCredentials(userModel)
}

func (model *UserModel) MatchCredentials(identity interfaces.IUserModel) (bool, error) {
	return model.GetPassword() == identity.GetPassword(), nil
}

func (model *UserModel) HasAuthorization() (bool, error) {
	log.Default().Println("model: ", model)
	return model.GetId() != "", nil
}

func NewUserModel(gateway interfaces.IUsersGateway) *UserModel {
	return &UserModel{gateway: gateway}
}
