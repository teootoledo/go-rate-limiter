package config_test

import (
	"rateLimiter/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ConfigSuite struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}

func (suite *ConfigSuite) TestCustomValuesAreReturnedProperly() {
	conf := config.GetConfig()

	require.Equal(suite.T(), conf.ServerPort, config.DefaultServerPort)
}
