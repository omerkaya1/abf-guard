package bucket

import (
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

func theBucketServiceReceivedTheFollowingSettings(arg1 *gherkin.DocString) error {
	return godog.ErrPending
}

func requestsContainTheSameLogin(arg1 string, arg2 *gherkin.DocString) error {
	return godog.ErrPending
}

func requestsPassAndRequestsFail(arg1, arg2 string) error {
	return godog.ErrPending
}

func requestsContainTheSamePassword(arg1 string, arg2 *gherkin.DocString) error {
	return godog.ErrPending
}

func requestsContainTheSameIp(arg1 string, arg2 *gherkin.DocString) error {
	return godog.ErrPending
}

func requestsPassesAndRequestsFail(arg1, arg2 string) error {
	return godog.ErrPending
}

func bucketsAreCreatedWithTheFollowingParameters(arg1 string, arg2 *gherkin.DocString) error {
	return godog.ErrPending
}

func theRequestToFlushAndBucketsIsReceived(arg1, arg2 string) error {
	return godog.ErrPending
}

func theRequestIsCompletedWithoutErrors() error {
	return godog.ErrPending
}

func theRequestToRemoveBucketIsReceived(arg1 string) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^the bucket service received the following settings:$`, theBucketServiceReceivedTheFollowingSettings)
	s.Step(`^"([^"]*)" requests contain the same login:$`, requestsContainTheSameLogin)
	s.Step(`^"([^"]*)" requests pass and "([^"]*)" requests fail$`, requestsPassAndRequestsFail)
	s.Step(`^"([^"]*)" requests contain the same password:$`, requestsContainTheSamePassword)
	s.Step(`^"([^"]*)" requests contain the same ip:$`, requestsContainTheSameIp)
	s.Step(`^"([^"]*)" requests passes and "([^"]*)" requests fail$`, requestsPassesAndRequestsFail)
	s.Step(`^"([^"]*)" buckets are created with the following parameters:$`, bucketsAreCreatedWithTheFollowingParameters)
	s.Step(`^the request to flush "([^"]*)" and "([^"]*)" buckets is received$`, theRequestToFlushAndBucketsIsReceived)
	s.Step(`^the request is completed without errors$`, theRequestIsCompletedWithoutErrors)
	s.Step(`^the request to remove "([^"]*)" bucket is received$`, theRequestToRemoveBucketIsReceived)
}
