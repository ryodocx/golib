package junit

import (
	"encoding/xml"
	"os"
	"strings"
)

type JUnit struct {
	XMLName    xml.Name    `xml:"testsuites"`
	Testsuites []TestSuite `xml:"testsuite"`
}

func NewJUnit() *JUnit {
	return &JUnit{
		Testsuites: []TestSuite{},
	}
}

func (p *JUnit) Build() *JUnit {
	for _, testSuite := range p.Testsuites {
		testSuite.build()
	}
	return p
}

func (p *JUnit) Save(filePath string) error {

	p.Build()

	if !strings.HasSuffix(filePath, ".xml") {
		filePath += ".xml"
	}

	if _, err := os.Stat(filePath); err == nil {
		if err := os.Remove(filePath); err != nil {
			return err
		}
	}

	var encoder *xml.Encoder
	if f, err := os.Create(filePath); err != nil {
		return err
	} else {
		encoder = xml.NewEncoder(f)
		encoder.Indent("", "\t")
	}

	if err := encoder.Encode(*p); err != nil {
		return err
	}
	if err := encoder.Flush(); err != nil {
		return err
	}

	return nil
}
