package ptr_test

import (
	"fmt"
	"testing"

	"github.com/ryodocx/golib/v2/ptr"
)

func TestToPtr(t *testing.T) {

	check := func(v interface{}) error {
		p := ptr.ToPtr(v)
		if *p != v {
			return fmt.Errorf("not match")
		}
		return nil
	}

	if e := check(1); e != nil {
		t.Error(e)
	}
	if e := check("hello"); e != nil {
		t.Error(e)
	}
	if e := check(nil); e != nil {
		t.Error(e)
	}
	if e := check(0.01); e != nil {
		t.Error(e)
	}
}

func TestFromPtr(t *testing.T) {

	check := func(p interface{}) error {
		v := ptr.FromPtr(&p, nil)
		if p != v {
			return fmt.Errorf("not match")
		}
		v2 := ptr.FromPtr(&p, p)
		if p != v2 {
			return fmt.Errorf("not match")
		}
		v3 := ptr.FromPtr(nil, p)
		if p != v3 {
			return fmt.Errorf("not match")
		}
		return nil
	}

	if e := check(1); e != nil {
		t.Error(e)
	}
	if e := check("hello"); e != nil {
		t.Error(e)
	}
	if e := check(nil); e != nil {
		t.Error(e)
	}
	if e := check(0.01); e != nil {
		t.Error(e)
	}
}
