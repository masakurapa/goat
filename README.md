# goat
`goat` is golang API testing library.

# Usage

- [using gin-gonic/gin](https://github.com/masakurapa/goat/blob/master/examples/gin/example_test.go)
- [using labstack/echo](https://github.com/masakurapa/goat/blob/master/examples/echo/example_test.go)
- [using net/http](https://github.com/masakurapa/goat/blob/master/examples/net/http/example_test.go)

# SetUp and TearDown

Use `SetUpBeforeTest` and `TearDownAfterTest` if you want to run the process before and after all test cases are executed.

Use `SetUp` and `TearDown` if you want the process to run before or after the test is executed.

See [here](https://github.com/masakurapa/goat/blob/master/examples/basic/example_test.go) for a example.

The result of running this test will look like this.

```sh
$ go test example_test.go -v
=== RUN   TestExample
run SetUpBeforeTest
=== RUN   TestExample/test1
run SetUp test1
=== RUN   TestExample/test1/CustomTestFunc1
run CustomTestFunc1
=== RUN   TestExample/test1/CustomTestFunc2
run CustomTestFunc2
run TearDown test1
=== RUN   TestExample/test2
run SetUp test2
=== RUN   TestExample/test2/CustomTestFunc#0
run CustomTestFunc3
=== RUN   TestExample/test2/CustomTestFunc#1
run CustomTestFunc4
run TearDown test2
run TearDownAfterTest
--- PASS: TestExample (0.00s)
    --- PASS: TestExample/test1 (0.00s)
        --- PASS: TestExample/test1/CustomTestFunc1 (0.00s)
        --- PASS: TestExample/test1/CustomTestFunc2 (0.00s)
    --- PASS: TestExample/test2 (0.00s)
        --- PASS: TestExample/test2/CustomTestFunc#0 (0.00s)
        --- PASS: TestExample/test2/CustomTestFunc#1 (0.00s)
PASS
ok      github.com/masakurapa/goat/examples/basic       0.250s
```
