package integration_test

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/omerkaya1/abf-guard/test/integration-test/api"
	"github.com/omerkaya1/abf-guard/test/integration-test/bucket"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	status := 1
	fmt.Println("Waiting for all services to become available...")
	//time.Sleep(time.Second * 30)

	status = godog.RunWithOptions("API", func(s *godog.Suite) {
		api.FeatureContext(s)
	}, godog.Options{
		Format:              "pretty",
		Paths:               []string{"./features"},
		Randomize:           0,
		ShowStepDefinitions: false,
	})

	status = godog.RunWithOptions("Bucket", func(s *godog.Suite) {
		bucket.FeatureContext(s)
	}, godog.Options{
		Format:              "pretty",
		Paths:               []string{"./features"},
		Randomize:           0,
		ShowStepDefinitions: false,
	})
	os.Exit(status)
}
