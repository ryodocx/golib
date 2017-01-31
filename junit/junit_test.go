package junit

import "testing"
import "runtime/debug"

func TestJUnit(t *testing.T) {

	junit := NewJUnit()
	suite := junit.NewTestSuite("sample_test_suite")

	testCase1 := suite.NewTestCase("sample_test_case1", "sample_class")
	testCase1.OK()

	testCase2 := suite.NewTestCase("sample_test_case2", "sample_class")
	testCase2.Fail("sample_error_type", "samplw_error_msg", string(debug.Stack()))

	testCase3 := suite.NewTestCase("sample_test_case3", "sample_class")
	testCase3.Error("sample_error_type", "samplw_error_msg", string(debug.Stack()))

	testCase4 := suite.NewTestCase("sample_test_case3", "sample_class")
	testCase4.Skip("sample_skip_msg")

	suite.Finish()

	if i := len(junit.Testsuites); i != 1 {
		t.Errorf("%d =! 1", i)
	}

	if suite.Tests != 4 {
		t.Errorf("%d != 4", suite.Tests)
	}

	if suite.Errors != 1 {
		t.Errorf("%d != 1", suite.Errors)
	}

	if suite.Failures != 1 {
		t.Errorf("%d != 1", suite.Failures)
	}

	if err := junit.Save("report.xml"); err != nil {
		t.Error(err)
	}

}
