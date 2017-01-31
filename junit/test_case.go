package junit

import (
	"fmt"
	"runtime/debug"
	"time"
)

type TestCase struct {
	Classname string   `xml:"classname,attr"`
	Name      string   `xml:"name,attr"`
	Time      float64  `xml:"time,attr"`
	Failure   *testErr `xml:"failure,omitempty"`
	Err       *testErr `xml:"error,omitempty"`
	Skipped   *skipped `xml:"skipped,omitempty"`
	beginTime time.Time
}

type testErr struct {
	Type       string `xml:"type,attr"`
	Message    string `xml:"message,attr"`
	StackTrace string `xml:",chardata"`
}

type skipped struct {
	Message string `xml:"message,attr"`
}

func (p *TestSuite) NewTestCase(testName string, className ...string) *TestCase {
	name := fmt.Sprintf("%s.test%d", p.Name, p.Tests)
	if className != nil {
		name = className[0]
	}
	p.TestCases = append(p.TestCases, TestCase{
		Classname: name,
		Name:      testName,
		Time:      0,
		Failure:   nil,
		Err:       nil,
		Skipped:   nil,
		beginTime: time.Now(),
	})
	return &p.TestCases[len(p.TestCases)-1]
}

func (p *TestCase) OK() {
	p.Time = time.Now().Sub(p.beginTime).Seconds()
}

func (p *TestCase) Error(errType string, errMsg string, stackTrace ...string) {
	p.Time = time.Now().Sub(p.beginTime).Seconds()

	trace := string(debug.Stack())
	if stackTrace != nil {
		trace = stackTrace[0]
	}

	p.Err = &testErr{
		Type:       errType,
		Message:    errMsg,
		StackTrace: trace,
	}
}

func (p *TestCase) Fail(errType string, errMsg string, stackTrace ...string) {
	p.Time = time.Now().Sub(p.beginTime).Seconds()

	trace := string(debug.Stack())
	if stackTrace != nil {
		trace = stackTrace[0]
	}

	p.Failure = &testErr{
		Type:       errType,
		Message:    errMsg,
		StackTrace: trace,
	}
}

func (p *TestCase) Skip(msg string) {
	p.Time = time.Now().Sub(p.beginTime).Seconds()
	p.Skipped = &skipped{
		Message: msg,
	}
}
