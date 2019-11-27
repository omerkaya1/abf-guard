package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"google.golang.org/grpc"

	"integration-test/api"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

type authorise struct {
	Login string `json:"login,omitempty"`
	Pwd   string `json:"password,omitempty"`
	IP    string `json:"ip,omitempty"`
}

type testABFG struct {
	c          api.ABFGuardClient
	authReq    *authorise
	loginLimit int
	pwdLimit   int
	ipLimit    int
	duration   time.Duration
	latestErr  error
	latestOk   bool
	respOkMap  map[int]bool
}

func fastFail(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func blacklistCheck(arg string) bool {
	if arg == "black" {
		return true
	}
	return false
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomIfEmpty(arg string) string {
	if arg != "" {
		return arg
	}
	b := make([]byte, 5)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func getNameByIndex(index int, auth authorise) string {
	switch index {
	case 0:
	case 3:
		return auth.Login
	case 1:
	case 4:
		return auth.Pwd
	case 2:
	case 5:
		return auth.IP
	}
	return ""
}

func newTestABFG() *testABFG {
	conn, err := grpc.Dial(os.Getenv("TEST_HOST")+":"+os.Getenv("TEST_PORT"), grpc.WithInsecure())
	fastFail(err)
	client := api.NewABFGuardClient(conn)
	loginLimit, err := strconv.Atoi(os.Getenv("LOGIN_LIMIT"))
	fastFail(err)
	pwdLimit, err := strconv.Atoi(os.Getenv("PASSWORD_LIMIT"))
	fastFail(err)
	ipLimit, err := strconv.Atoi(os.Getenv("IP_LIMIT"))
	fastFail(err)
	duration, err := time.ParseDuration(os.Getenv("EXPIRE_LIMIT"))
	fastFail(err)
	return &testABFG{
		c:          client,
		authReq:    nil,
		loginLimit: loginLimit,
		pwdLimit:   pwdLimit,
		ipLimit:    ipLimit,
		duration:   duration,
		respOkMap:  make(map[int]bool),
	}
}

func (abfg *testABFG) weSendARequestToAddSubnetToTheList(arg1, arg2 string) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	resp, err := abfg.c.AddIPToWhitelist(ctx, &api.SubnetRequest{Ip: arg1, List: blacklistCheck(arg2)})
	if err != nil {
		return err
	}
	if resp.GetError() != "" && resp.GetOk() {
		abfg.latestErr = fmt.Errorf("last returned error: %s and %v", resp.GetError(), resp.GetOk())
	}
	return nil
}

func (abfg *testABFG) theRequestIsCompletedWithoutErrors() error {
	if abfg.latestErr != nil {
		return abfg.latestErr
	}
	return nil
}

func (abfg *testABFG) theIpIsInTheList(arg1, arg2, arg3 string) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	resp, err := abfg.c.GetIPList(ctx, &api.ListRequest{ListType: blacklistCheck(arg3)})
	if err != nil {
		return err
	}
	if resp.GetError() != "" {
		log.Println(resp.GetError())
		return fmt.Errorf(resp.GetError())
	}

	// "not" is passed
	if arg2 != "" {
		for _, v := range resp.GetIps().GetList() {
			if v == arg1 {
				return fmt.Errorf("\033[34mthe %s address is in the %s list\033[0m\n", arg1, arg3)
			}
		}
		log.Printf("the %s address is %s in the %s list\n", arg1, arg2, arg3)
		return nil
	} else {
		for _, v := range resp.GetIps().GetList() {
			if v == arg1 {
				fmt.Printf("\033[34mthe %s address is in the %s list\033[0m\n", arg1, arg3)
				return nil
			}
		}
		return fmt.Errorf("the %s address is NOT IN in the %s list", arg1, arg3)
	}
}

func (abfg *testABFG) weSendAuthorisationRequestsForTimesOfTheAllowedLimitsWithParameters(
	arg1, arg2 string, arg3 *gherkin.DocString) error {
	// Function declaration was too long
	times, err := strconv.Atoi(arg2)
	if err != nil {
		return err
	}
	replacer := strings.NewReplacer("\n", "", "\t", "")
	cleanJson := replacer.Replace(arg3.Content)
	a := &authorise{}
	if err := json.Unmarshal([]byte(cleanJson), a); err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	switch arg1 {
	case "white":
		for i := 0; i < times*abfg.loginLimit; i++ {
			resp, err := abfg.c.Authorisation(ctx, &api.AuthorisationRequest{Login: a.Login, Password: a.Pwd, Ip: a.IP})
			if err != nil {
				fmt.Println(err)
				return err
			}
			abfg.respOkMap[i] = resp.GetOk()
		}
		break
	case "black":
		for i := 0; i < times*abfg.ipLimit; i++ {
			resp, err := abfg.c.Authorisation(ctx,
				&api.AuthorisationRequest{Login: randomIfEmpty(a.Login), Password: randomIfEmpty(a.Pwd), Ip: a.IP})
			if err != nil {
				fmt.Println(err)
				return err
			}
			abfg.respOkMap[i] = resp.GetOk()
		}
		break
	case "normal":
		for i := 0; i < times*abfg.loginLimit; i++ {
			resp, err := abfg.c.Authorisation(ctx, &api.AuthorisationRequest{Login: a.Login, Password: a.Pwd, Ip: a.IP})
			if err != nil {
				fmt.Println(err)
				return err
			}
			abfg.respOkMap[i] = resp.GetOk()
		}
		break
	default:
		return fmt.Errorf("unknown error: %s passed as an ergument", arg1)
	}
	return nil
}

func (abfg *testABFG) theyAllSucceed() error {
	for i, v := range abfg.respOkMap {
		if !v {
			return fmt.Errorf("%d request failed", i)
		}
	}
	for i := range abfg.respOkMap {
		delete(abfg.respOkMap, i)
	}
	fmt.Printf("\033[34m successfuly cleaned the respOkMap: %v\033[0m\n", abfg.respOkMap)
	return nil
}

func (abfg *testABFG) theyAllFail() error {
	for i, v := range abfg.respOkMap {
		if v {
			return fmt.Errorf("%d request DID NOT fail", i)
		}
	}
	for i := range abfg.respOkMap {
		delete(abfg.respOkMap, i)
	}
	fmt.Printf("\033[34m successfuly cleaned the respOkMap: %v\033[0m\n", abfg.respOkMap)
	return nil
}

func (abfg *testABFG) preciselyOfTheRequestsShouldHavePassedAndShouldNotHavePassed(arg1, arg2 string) error {
	passed, err := strconv.Atoi(arg1)
	if err != nil {
		return err
	}
	failed, err := strconv.Atoi(arg2)
	if err != nil {
		return err
	}
	actualPassed, actualFailed := 0, 0
	for _, v := range abfg.respOkMap {
		if v {
			actualPassed++
		} else {
			actualFailed++
		}
	}
	if actualPassed != passed {
		return fmt.Errorf("values of actualPassed and expected to pass requests do not match: %d and %d",
			actualPassed, passed)
	}
	if actualFailed != failed {
		return fmt.Errorf("values of actualFailed and expected to fail requests do not match: %d and %d",
			actualFailed, failed)
	}
	for i := range abfg.respOkMap {
		delete(abfg.respOkMap, i)
	}
	return nil
}

func (abfg *testABFG) weSendAnAuthorisationRequestWithParameters(arg1 *gherkin.DocString) error {
	replacer := strings.NewReplacer("\n", "", "\t", "")
	cleanJson := replacer.Replace(arg1.Content)
	a := &authorise{}
	if err := json.Unmarshal([]byte(cleanJson), a); err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	resp, err := abfg.c.Authorisation(ctx, &api.AuthorisationRequest{Login: a.Login, Password: a.Pwd, Ip: a.IP})
	if err != nil {
		fmt.Println(err)
		return err
	}
	if resp.GetOk() {
		return fmt.Errorf("the request failed %v: error: %s", resp.GetOk(), resp.GetError())
	}
	return nil
}

func (abfg *testABFG) sendingRequestTimesForTheFollowingBuckets(arg1, arg2 string, arg3 *gherkin.DocString) error {
	times, err := strconv.Atoi(arg2)
	if err != nil {
		return err
	}
	replacer := strings.NewReplacer("\n", "", "\t", "")
	cleanJson := replacer.Replace(arg3.Content)
	a := &authorise{}
	if err := json.Unmarshal([]byte(cleanJson), a); err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	for i := 0; i < times; i++ {
		switch arg1 {
		case "flush":
			resp, err := abfg.c.FlushBuckets(ctx, &api.FlushBucketRequest{Login: a.Login, Ip: a.IP})
			if err != nil {
				fmt.Println(err)
				return err
			}
			abfg.respOkMap[i] = resp.GetOk()
			break
		case "purge":
			resp, err := abfg.c.PurgeBucket(ctx, &api.PurgeBucketRequest{Name: getNameByIndex(i, *a)})
			if err != nil {
				fmt.Println(err)
				return err
			}
			abfg.respOkMap[i] = resp.GetOk()
			break
		default:
			return fmt.Errorf("unknown error: %s passed as an ergument", arg1)
		}
	}
	return nil
}

func (abfg *testABFG) theRequestsAreCompletedWithoutErrors() error {
	if abfg.latestErr != nil {
		return abfg.latestErr
	}
	return nil
}

func (abfg *testABFG) weSendARequestToGetAListOfIpsFromList(arg1 string) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	resp, err := abfg.c.GetIPList(ctx, &api.ListRequest{ListType: blacklistCheck(arg1)})
	if err != nil {
		log.Println(err)
		return err
	}
	if resp.GetError() != "" {
		log.Println(resp.GetError())
		return fmt.Errorf(resp.GetError())
	}
	if len(resp.GetIps().GetList()) != 1 {
		return fmt.Errorf("length of the %s list is not equal 1: actual: %d", arg1, len(resp.GetIps().GetList()))
	}
	return nil
}

func (abfg *testABFG) weSendARequestToDeleteIpFromList(arg1, arg2 string) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	resp := &api.Response{}
	var err error
	if arg2 == "black" {
		resp, err = abfg.c.DeleteIPFromBlacklist(ctx, &api.SubnetRequest{Ip: arg1, List: blacklistCheck(arg2)})
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		resp, err = abfg.c.DeleteIPFromWhitelist(ctx, &api.SubnetRequest{Ip: arg1, List: blacklistCheck(arg2)})
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if resp.GetError() != "" {
		log.Println(resp.GetError())
		return fmt.Errorf(resp.GetError())
	}

	if !resp.GetOk() {
		return fmt.Errorf("delete from %s request failed: %v: error: %s", arg1, resp.GetOk(), resp.GetError())
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	test := newTestABFG()
	s.Step(`^we send a request to add "([^"]*)" subnet to the "([^"]*)" list$`,
		test.weSendARequestToAddSubnetToTheList)
	s.Step(`^the request is completed without errors$`, test.theRequestIsCompletedWithoutErrors)
	s.Step(`^the "([^"]*)" ip is "([^"]*)" in the "([^"]*)" list$`, test.theIpIsInTheList)
	s.Step(`^we send "([^"]*)" authorisation requests for "([^"]*)" times of the allowed limits with parameters:$`,
		test.weSendAuthorisationRequestsForTimesOfTheAllowedLimitsWithParameters)
	s.Step(`^they all succeed$`, test.theyAllSucceed)
	s.Step(`^they all fail$`, test.theyAllFail)
	s.Step(`^precisely "([^"]*)" of the requests should have passed and "([^"]*)" should not have passed$`,
		test.preciselyOfTheRequestsShouldHavePassedAndShouldNotHavePassed)
	s.Step(`^we send an authorisation request with parameters:$`, test.weSendAnAuthorisationRequestWithParameters)
	s.Step(`^sending "([^"]*)" request "([^"]*)" times for the following buckets:$`,
		test.sendingRequestTimesForTheFollowingBuckets)
	s.Step(`^the requests are completed without errors$`, test.theRequestsAreCompletedWithoutErrors)
	s.Step(`^we send a request to get a list of ips from "([^"]*)" list$`,
		test.weSendARequestToGetAListOfIpsFromList)
	s.Step(`^we send a request to delete "([^"]*)" ip from "([^"]*)" list$`, test.weSendARequestToDeleteIpFromList)
}

func TestMain(m *testing.M) {
	fmt.Println("Waiting for all services to become available...")
	time.Sleep(time.Second * 10)

	status := godog.RunWithOptions("API", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:              "pretty",
		Paths:               []string{"./features"},
		Randomize:           0,
		ShowStepDefinitions: false,
	})

	os.Exit(status)
}
