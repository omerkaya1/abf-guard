# Antibruteforce

## Features

1. Rate limit algorithm (leaky bucket) ```DONE```
2. Whitelist/blacklist (CLI to manage lists) ```DONE```
3. API ```DONE```
4. There should be no memory leaks ```DONE```
5. Code structure according to clean architecture style guides ```DONE```
6. Code passes go vet, golint and race detector checks ```DONE```
7. UTs ```DONE```
8. Try to reach 100% of the test coverage ```More or less DONE ~85%```
9. Integration tests should start by docker-compose ```DONE```
10. ```make run``` and ```make test``` presence ```DONE```

## TODO for the project

* ~~Update README.md file~~
* ~~Complete UTs~~
* ~~Complete integration tests~~

## TODO for the future

* Create a message queue that will send notifications for the service administrator about subnet
addresses that look suspicious;
* ~~Add context to every critical section of the programme so that OS signals could be propagated
and properly handled;~~ `DONE`
