package mocks

type MockEnvManager struct{}

func (m *MockEnvManager) InitEnv()              {}
func (m *MockEnvManager) GetAddr() string       { return "ADDR" }
func (m *MockEnvManager) GetDBDSN() string      { return "DSN" }
func (m *MockEnvManager) GetJWTSignKey() string { return "my_secret_signing_key" }
