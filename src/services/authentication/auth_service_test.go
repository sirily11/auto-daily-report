package services

import (
	authDto "sme-demo/src/dto/authentication"
	auth "sme-demo/src/repositories/authentication"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type AuthTestSuite struct {
	suite.Suite
	ctl *gomock.Controller
}

func (suite *AuthTestSuite) SetupTest() {
	ctl := gomock.NewController(suite.T())
	suite.ctl = ctl
}

func (suite *AuthTestSuite) TestSignUp() {
	repo := auth.NewMockAuthRepositoryInterface(suite.ctl)
	service := NewAuthService(nil, repo)

	repo.EXPECT().SignUp(gomock.Any()).Return(nil, nil)
	_, err := service.SignUp(
		authDto.SignUpUserDto[any]{
			Source:      authDto.SourceWebApp,
			Permissions: []string{},
			CredentialData: authDto.SignUpUserDtoUserNamePasswordCredentialData{
				Username: "test",
				Password: "test",
			},
		})
	suite.Nil(err)
}

func UserTestNewTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
