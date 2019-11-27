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

	# NOTE: Checked. Ready to be used.
	Scenario: Add an IP to the whitelist
		When we send a request to add "10.0.0.1" subnet to the "white" list
		Then the request is completed without errors
		And the "10.0.0.1" ip is "" in the "white" list

	# NOTE: Checked. Ready to be used.
	Scenario: Add an IP to the blacklist
		When we send a request to add "10.0.0.2" subnet to the "black" list
		Then the request is completed without errors
		And the "10.0.0.2" ip is "" in the "black" list

    # NOTE: Checked. Ready to be used.
	Scenario: Check happy path for a subnet
		When we send "white" authorisation requests for "2" times of the allowed limits with parameters:
			"""
			{ "login": "Morty", "password": "1234", "ip": "10.0.0.1" }
			"""
		Then they all succeed

	# NOTE: Checked. Ready to be used.
	Scenario: Check blocking for a subnet
		When we send "black" authorisation requests for "2" times of the allowed limits with parameters:
			"""
			{ "login": "Morty", "password": "1234", "ip": "10.0.0.2" }
			"""
		Then they all fail

	# NOTE: Checked. Ready to be used.
	Scenario: Normal authorisation
		When we send "normal" authorisation requests for "2" times of the allowed limits with parameters:
			"""
			{ "login": "Morty", "password": "1234", "ip": "10.0.0.3" }
			"""
		Then half of the requests pass and half do not

	# NOTE: Checked. Ready to be used.
	Scenario: Send a flush buckets request
		Given we send an authorisation request with parameters:
			"""
			{ "login": "Morty", "password": "1234", "ip": "10.0.0.3" }
			"""
		When sending "flush" request "2" times for the following buckets:
			"""
			{ "login": "Morty", "password": "", "ip": "10.0.0.3" }
			"""
		Then half of the requests pass and half do not

	# NOTE: Checked. Ready to be used.
	Scenario: Send a purge bucket request
		Given we send an authorisation request with parameters:
			"""
			{ "login": "Morty", "password": "1234", "ip": "10.0.0.3" }
			"""
		When sending "purge" request "6" times for the following buckets:
		"""
			{ "login": "Morty", "password": "1234", "ip": "10.0.0.3" }
			"""
		Then half of the requests pass and half do not

	# NOTE: Checked. Ready to be used.
	Scenario: Get a list of IPs that belong to the whitelist
		When we send a request to get a list of ips from "white" list
		Then the "10.0.0.1" ip is "" in the "white" list

	# NOTE: Checked. Ready to be used.
	Scenario: Get a list of IPs that belong to the blacklist
		When we send a request to get a list of ips from "black" list
		Then the "10.0.0.2" ip is "" in the "black" list

	# NOTE: Checked. Ready to be used.
	Scenario: Delete an IP from the whitelist
		When we send a request to delete "10.0.0.1" ip from "white" list
		Then the request is completed without errors

	# NOTE: Checked. Ready to be used.
	Scenario: Delete an IP from the blacklist
		When we send a request to delete "10.0.0.2" ip from "black" list
		Then the request is completed without errors

