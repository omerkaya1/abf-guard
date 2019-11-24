# Main test suite for the bucket feature
# The following features must be tested:
########################################
#  1) Add buckets until the login bucket is full
#  2) Add buckets until the password bucket is full
#  3) Add buckets until the ip bucket is full
#  4) Flush buckets
#  5) Remove a bucket
#  6) The monitoring routine check
########################################
Feature: Leaky bucket
	In order to prove that the leaky bucket algorithm is implemented
	As a server that manages buckets through a bucket service
	The bucket service should be able to do the following

	Scenario: Accept authorisation requests until the login bucket is full
        Given the bucket service received the following settings:
			"""
			{ "login": 5, "password": 10, "ip": 100, "expire": "20s" }
			"""
        When "10" requests contain the same login:
			"""
			{ "login": "Morty" }
			"""
        Then "5" requests pass and "5" requests fail

	Scenario: Accept authorisation requests until the password bucket is full
        Given the bucket service received the following settings:
			"""
			{ "login": 5, "password": 10, "ip": 100, "expire": "20s" }
			"""
        When "15" requests contain the same password:
			"""
			{ "password": "12345" }
			"""
        Then "10" requests pass and "5" requests fail

	Scenario: Accept authorisation requests until the ip bucket is full
        Given the bucket service received the following settings:
			"""
			{ "login": 5, "password": 10, "ip": 100, "expire": "30s" }
			"""
        When "105" requests contain the same ip:
			"""
			{ "login": "10.0.0.1" }
			"""
        Then "100" requests passes and "5" requests fail

	Scenario: Flush existing buckets
        Given the bucket service received the following settings:
			"""
			{ "login": 5, "password": 10, "ip": 100, "expire": "30s" }
			"""
		And "5" buckets are created with the following parameters:
			"""
			{ "login": "morty", "password": "1234", "ip": "10.0.0.1" }
			"""
        When the request to flush "morty" and "10.0.0.1" buckets is received
        Then the request is completed without errors

	Scenario: Purge existing bucket
        Given the bucket service received the following settings:
			"""
			{ "login": 5, "password": 10, "ip": 100, "expire": "30s" }
			"""
		And "5" buckets are created with the following parameters:
			"""
			{ "login": "morty", "password": "1234", "ip": "10.0.0.1" }
			"""
        When the request to remove "morty" bucket is received
        Then the request is completed without errors
