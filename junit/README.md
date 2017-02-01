# summary
* JUnit xml builder for golang
* `not` converter go test result to JUnit

## usage

* add to GOPATH
```bash
go get github.com/ryodocx/golib/junit
```

* sample code
```go
import "github.com/ryodocx/golib/junit"

func sammpleJUnit() {

	testReport := junit.NewJUnit()
	suite := testReport.NewTestSuite("sample_test_suite")

	testCase1 := suite.NewTestCase("sample_test_case1", "sample_class")
	testCase1.OK()

	testCase2 := suite.NewTestCase("sample_test_case2", "sample_class")
	testCase2.Fail("sample_error_type", "sample_error_msg", string(debug.Stack()))
  
	testCase3 := suite.NewTestCase("sample_test_case3", "sample_class")
	testCase3.Skip("sample_skip_msg")

	suite.Finish()

	if err := testReport.Save("report.xml"); err != nil {
		panic(err)
	}

}

```
