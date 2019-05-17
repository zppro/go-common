package syntax2func

import (
	"github.com/zppro/go-common/constant/chars"
	"testing"
)

func TestIif(t *testing.T) {
	t.Log("测试IsZero函数")
	i0 := 0

	if Iif(IsZero(i0), 0, 1).(int) == 0 {
		t.Logf("	%s should return `0`." ,chars.CheckMark)

	} else {
		t.Errorf("	%s expected `0` but return `1`! " , chars.BallotX)
	}

	s0 := ""
	if Iif(IsZero(s0), "", "str").(string) == "" {
		t.Logf("	%s should return ``." ,chars.CheckMark)

	} else {
		t.Errorf("	%s expected `` but return `str`! " , chars.BallotX)
	}

	var struct0 struct{A int}
	if Iif(IsZero(struct0), 0, 1) == struct0.A {
		t.Logf("	%s struct0.A should return `0`." ,chars.CheckMark)

	} else {
		t.Errorf("	%s struct0.A expected `0` but return `1`! " , chars.BallotX)
	}

	var slice0 []int
	if Iif(IsZero(slice0), 0, 1) == len(slice0) {
		t.Logf("	%s len(slice0) should return `0`." ,chars.CheckMark)

	} else {
		t.Errorf("	%s len(slice0) expected `0` but return `1`! " , chars.BallotX)
	}

	map1 := map[string]int {"1": 2}

	if Iif(!IsZero(map1), 2, 0) == map1["1"] {
		t.Logf("	%s map1[\"1\"] should return `2`. " ,chars.CheckMark)

	} else {
		t.Errorf("	%s map1[\"1\"] expected `2` but return `0`! " , chars.BallotX)
	}

}