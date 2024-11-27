package generics

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	go startServer()

	os.Exit(m.Run())
}

func TestGetUsers(t *testing.T) {
	users, err := getApiData[user]("http://127.0.0.1:8000/users")
	require.NoError(t, err)

	expectedUsers := []user{
		{ID: 1, Name: "John Doe", Email: "john_doe@example.org"},
		{ID: 2, Name: "Jane Doe", Email: "jane_doe@example.org"},
	}
	require.Equal(t, expectedUsers, users)
}

func TestGetCompanies(t *testing.T) {
	companies, err := getApiData[company]("http://127.0.0.1:8000/companies")
	require.NoError(t, err)

	expectedCompanies := []company{
		{ID: 1, Name: "Green Tech LLC", Address: "123 Main Street"},
		{ID: 2, Name: "Blue Tech LLC", Address: "456 Main Street"},
	}
	require.Equal(t, expectedCompanies, companies)
}
