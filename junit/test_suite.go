package junit

import "time"

type TestSuite struct {
	Name      string     `xml:"name,attr"`
	Tests     int64      `xml:"tests,attr"`
	Errors    int64      `xml:"errors,attr"`
	Failures  int64      `xml:"failures,attr"`
	Time      float64    `xml:"time,attr"`
	TestCases []TestCase `xml:"testcase"`
	SystemOut cdata      `xml:"system-out"`
	ErrorOut  cdata      `xml:"error-out"`
	beginTime time.Time
}

type cdata struct {
	Value string `xml:",cdata"`
}

func (p *JUnit) NewTestSuite(suiteName string) *TestSuite {
	suite := TestSuite{
		Name:      suiteName,
		Tests:     0,
		Errors:    0,
		Failures:  0,
		Time:      0,
		TestCases: []TestCase{},
		SystemOut: cdata{Value: ""},
		ErrorOut:  cdata{Value: ""},
		beginTime: time.Now(),
	}
	p.Testsuites = append(p.Testsuites, suite)

	return &p.Testsuites[len(p.Testsuites)-1]
}

func (p *TestSuite) AddSystemOut(msg string) {
	p.SystemOut.Value += "\n" + msg
}

func (p *TestSuite) AddErrorOut(msg string) {
	p.ErrorOut.Value += "\n" + msg
}

func (p *TestSuite) Finish() {
	p.Time = time.Now().Sub(p.beginTime).Seconds()
	p.build()
}

func (p *TestSuite) build() {
	p.Tests = 0
	p.Errors = 0
	p.Failures = 0
	for _, testCase := range p.TestCases {
		p.Tests++
		if testCase.Failure != nil {
			p.Failures++
		}
		if testCase.Err != nil {
			p.Errors++
		}
	}
}
