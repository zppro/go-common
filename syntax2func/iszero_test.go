package syntax2func

import (
	"github.com/zppro/go-common/constant/chars"
	"testing"
)

func TestIsZero(t *testing.T) {
	t.Log("测试IsZero函数")
	i0, i1 := 0, 1

	if IsZero(i0) {
		t.Logf("	%s int %d should return `true`." ,chars.CheckMark, i0)

	} else {
		t.Errorf("	%s int %d expected `false` but return `true`! " , chars.BallotX, i0)
	}

	if !IsZero(i1) {
		t.Logf("	%s int %d should return `false`." ,chars.CheckMark, i1)

	} else {
		t.Errorf("	%s int %d expected `false` but return `true`! " , chars.BallotX, i1)
	}


	s0, s1 := "", "str"
	if IsZero(s0) {
		t.Logf("	%s string '%s' should return `true`." ,chars.CheckMark, s0)

	} else {
		t.Errorf("	%s string '%s' expected `false` but return `true`! " , chars.BallotX, s0)
	}

	if !IsZero(s1) {
		t.Logf("	%s string '%s' should return `false`. " ,chars.CheckMark, s1)

	} else {
		t.Errorf("	%s string '%s' expected `false` but return `true`! " , chars.BallotX, s1)
	}

	var struct0 struct{A int}
	struct1 := struct{A int}{1}
	if IsZero(struct0) {
		t.Logf("	%s struct %v should return `true`." ,chars.CheckMark, struct0)

	} else {
		t.Errorf("	%s struct %v expected `false` but return `true`! " , chars.BallotX, struct0)
	}

	if !IsZero(struct1) {
		t.Logf("	%s struct %v should return `false`. " ,chars.CheckMark, struct1)

	} else {
		t.Errorf("	%s struct %v expected `false` but return `true`! " , chars.BallotX, struct1)
	}

	var slice0 []int
	slice1 := []int {1, 2}
	if IsZero(slice0) {
		t.Logf("	%s slice %v should return `true`." ,chars.CheckMark, slice0)

	} else {
		t.Errorf("	%s slice %v expected `false` but return `true`! " , chars.BallotX, slice0)
	}

	if !IsZero(slice1) {
		t.Logf("	%s slice %v should return `false`. " ,chars.CheckMark, slice1)

	} else {
		t.Errorf("	%s slice %v expected `false` but return `true`! " , chars.BallotX, slice1)
	}

	var map0 map[string]int
	map1 := map[string]int {"1": 2}
	if IsZero(map0) {
		t.Logf("	%s map %v should return `true`." ,chars.CheckMark, map0)

	} else {
		t.Errorf("	%s map %v expected `false` but return `true`! " , chars.BallotX, map0)
	}

	if !IsZero(map1) {
		t.Logf("	%s map %v should return `false`. " ,chars.CheckMark, map1)

	} else {
		t.Errorf("	%s map %v expected `false` but return `true`! " , chars.BallotX, map1)
	}

}