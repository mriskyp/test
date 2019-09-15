This repo is test for Gojek 

<!-- include -->
include main.go and main_test.go

<!-- define -->
main_test.go is unit test for function that will return actual and expected condition.

<!-- building -->
Set initial build in bin/setup which generate docker from dep ensure -v and go build 
In the first when clone the repo, you need to `dep ensure -v` first just in case there will be update from the repo.

<!-- run -->
There are some option to run the code:
-You have to manually `go build` and then run `go run main.go`
-run the file by doing `go run main.go` or run by script bin/parking_lot and there will be STDOUT from println

<!-- run test case -->
To run test case, you have to run `go test -v` which will run all the test case in the repo.
If your test case failed, just edit code and run the unit test 

for example:
    TestLeaveWithoutError return error
    to test only the test case, just fix the code and click the func. Then there will be single run unit test


Project Requirements:

-Docker
-Golang
-Git