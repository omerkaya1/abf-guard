# Main test suite for the api feature
# The following features must be tested:
########################################
#  1) Add an IP to the whitelist;
#  2) Add an IP to the blacklist;
#  3) Delete an IP from the whitelist
#  4) Delete an IP from the blacklist
#  5) Get the list of IPs from the whitelist
#  6) Get the list of IPs from the whitelist
#  7) Send an authorise request
#  8) Send a flush buckets request
#  9) Send a purge bucket request
#  10) Error handling
########################################
Feature: Functional ABFG API
	In order to test the abf-guard application
	As a user that operates with the service through API
	The service should allow the following

	Scenario: Add an IP to the whitelist
		When we send a request to add "10.0.0.1" ip to "white" list
		Then the request is completed without errors

	Scenario: Add an IP to the blacklist
		When we send a request to add "10.0.0.1" ip to "black" list
		Then the request is completed without errors

	Scenario: Delete an IP from the whitelist
		When we send a request to delete "10.0.0.1" ip from "white" list
		Then the request is completed without errors

	Scenario: Delete an IP from the blacklist
		When we send a request to delete "10.0.0.1" ip from "black" list
		Then the request is completed without errors

	Scenario: Get a list of IPs that belong to the whitelist
		When we send a request to get a list of ips from "white" list
		Then the request is completed without errors

	Scenario: Get a list of IPs that belong to the blacklist
		When we send a request to get a list of ips from "black" list
		Then the request is completed without errors

	Scenario: Send an authorisation request
		When we send an authorisation request with parameters:
			"""
			{ "login": "Morty", "password": "1234", "ip": "10.0.0.1" }
			"""
		Then the request is completed without errors

	Scenario: Send a flush buckets request
		When we send an authorisation request with parameters:
			"""
			{ "login": "Morty", "password": "1234", "ip": "10.0.0.1" }
			"""
		And send a flush request for the login "Morty" and ip "10.0.0.1" buckets
		Then the request is completed without errors

	Scenario: Send a purge bucket request
		When we send an authorisation request with parameters:
			"""
			{ "login": "Morty", "password": "1234", "ip": "10.0.0.1" }
			"""
		And send "1" purge request for the login "Morty" bucket
		Then the request is completed without errors

	Scenario: Send a purge bucket request
		When we send an authorisation request with parameters:
			"""
			{ "login": "Morty", "password": "1234", "ip": "10.0.0.1" }
			"""
		And send "2" purge request for the login "Morty" bucket
		Then the request fails
