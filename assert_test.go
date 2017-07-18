package assert

import "testing"

func TestEqual(t *testing.T) {
	Equal(t, 100, 100)
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, 100, 101)
}

func TestLesserThan(t *testing.T) {
	LesserThan(t, 10, 11)
}

func TestLesserEqual(t *testing.T) {
	LesserEqual(t, 10, 10)
}

func TestGreaterThan(t *testing.T) {
	GreaterThan(t, 11, 10)
}

func TestGreaterEqual(t *testing.T) {
	GreaterEqual(t, 10, 10)
}

func TestNotNil(t *testing.T) {
	NotNil(t, 1)
}

func TestNil(t *testing.T) {
	Nil(t, nil)
}

func TestLen(t *testing.T) {
	l := [2]string{}
	Len(t, l, 2)
}

func TestContains(t *testing.T) {
	l := [2]string{
		"str",
		"int",
	}
	Contains(t, l, "str")
}

func TestNotContains(t *testing.T) {
	l := [2]string{
		"str",
		"int",
	}
	NotContains(t, l, "float")
}

func TestTrue(t *testing.T) {
	True(t, true)
}

func TestFalse(t *testing.T) {
	False(t, false)
}
