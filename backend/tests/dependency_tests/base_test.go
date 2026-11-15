package dependency_tests

import (
	"lorekeeper/tests/helpers"
	"os"
	"testing"
	"time"
)

var (
	ts *helpers.TestServer
)

func TestMain(m *testing.M) {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	time.Local = utc

	ts = helpers.NewTestServer("../../../.env")

	// Close the database connection
	defer ts.S.DB.Close()

	// Run the test
	code := m.Run()

	os.Exit(code)
}
