package integration_test

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	status := 1
	fmt.Println("Waiting for all services to become available...")
	time.Sleep(time.Second * 15)

	status = godog.RunWithOptions("API", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:              "pretty",
		Paths:               []string{"./features"},
		Randomize:           0,
		ShowStepDefinitions: false,
	})

	os.Exit(status)
}

func weSendARequestToAddIpToList(arg1, arg2 string) error {
	return godog.ErrPending
}

func theRequestIsCompletedWithoutErrors() error {
	return godog.ErrPending
}

func weSendARequestToDeleteIpFromList(arg1, arg2 string) error {
	return godog.ErrPending
}

func weSendARequestToGetAListOfIpsFromList(arg1 string) error {
	return godog.ErrPending
}

func weSendAnAuthorisationRequestWithParameters(arg1 *gherkin.DocString) error {
	return godog.ErrPending
}

func sendAFlushRequestForTheLoginAndIpBuckets(arg1, arg2 string) error {
	return godog.ErrPending
}

func sendPurgeRequestForTheLoginBucket(arg1, arg2 string) error {
	return godog.ErrPending
}

func theRequestFails() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^we send a request to add "([^"]*)" ip to "([^"]*)" list$`, weSendARequestToAddIpToList)
	s.Step(`^the request is completed without errors$`, theRequestIsCompletedWithoutErrors)
	s.Step(`^we send a request to delete "([^"]*)" ip from "([^"]*)" list$`, weSendARequestToDeleteIpFromList)
	s.Step(`^we send a request to get a list of ips from "([^"]*)" list$`, weSendARequestToGetAListOfIpsFromList)
	s.Step(`^we send an authorisation request with parameters:$`, weSendAnAuthorisationRequestWithParameters)
	s.Step(`^send a flush request for the login "([^"]*)" and ip "([^"]*)" buckets$`, sendAFlushRequestForTheLoginAndIpBuckets)
	s.Step(`^send "([^"]*)" purge request for the login "([^"]*)" bucket$`, sendPurgeRequestForTheLoginBucket)
	s.Step(`^the request fails$`, theRequestFails)
}
