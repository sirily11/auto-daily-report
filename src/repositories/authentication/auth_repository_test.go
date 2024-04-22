package auth

import (
	"context"
	"io"
	"sme-demo/internal/config/constants/keys"
	"sme-demo/internal/repositories"
	dto "sme-demo/src/dto/authentication"
	"testing"

	"github.com/google/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tryvium-travels/memongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepositoryTestSuite struct {
	suite.Suite
	repo   *AuthRepository
	db     *mongo.Database
	server *memongo.Server
}

func (suite *AuthRepositoryTestSuite) SetupTest() {
	logger.Init("Logger", true, false, io.Discard)
	mongoServer, err := memongo.Start(keys.TestMongoDBVersion)
	if err != nil {
		suite.T().Fatal(err)
	}

	dbClient := repositories.NewDatabase()
	db := dbClient.ConnectWithUriAndDBName(mongoServer.URI(), keys.TestMongoDatabase)

	suite.repo = NewAuthRepository(db).(*AuthRepository) // Perform type assertion
	suite.db = db
	suite.server = mongoServer
}

func (suite *AuthRepositoryTestSuite) TearDownTest() {
	err := suite.db.Client().Disconnect(context.TODO())
	if err != nil {
		return
	}
	suite.server.Stop()
}

func (suite *AuthRepositoryTestSuite) TestCreateWithUsernamePassword() {
	user := dto.SignUpUserDto[any]{
		Source:      dto.SourceWebApp,
		Permissions: []string{},
		CredentialData: dto.SignUpUserDtoUserNamePasswordCredentialData{
			Username: "test",
			Password: "test",
		},
	}
	create, err := suite.repo.SignUp(user)
	if err != nil {
		suite.T().Fatal(err)
	}

	assert.NotNil(suite.T(), create)
	assert.Equal(suite.T(), user.Source, create.Source)
	assert.NotEmpty(suite.T(), create.Id)
	assert.NotEqual(suite.T(), create.CredentialData.(dto.SignUpUserDtoUserNamePasswordCredentialData).Password, "test")
	assert.Greater(suite.T(), len(create.CredentialData.(dto.SignUpUserDtoUserNamePasswordCredentialData).Password), 4)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AuthRepositoryTestSuite))
}
