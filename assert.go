package assert

import (
	"fmt"
	"reflect"
)

func NotNil(t TestingT, object interface{}) {
	if isNil(object) == false {
		return
	}

	failureMessage := "Nil"

	fail(t, object, "Not Nil", failureMessage)
	

}

func Nil(t TestingT, object interface{}) {
	if isNil(object) {
		return
	}

	failureMessage := "Not Nil"

	fail(t, object, "Nil", failureMessage)
}

func True(t TestingT, b bool) {
	if b {
		return
	}

	failureMessage := "Not True"

	fail(t, b, "true", failureMessage)
}

func False(t TestingT, b bool) {
	if b == false {
		return
	}

	failureMessage := "Not False"

	fail(t, b, "false", failureMessage)
}

// assert.Len(t, rtnList, 2)
func Len(t TestingT, object interface{}, length int) {
	ok, l := getLen(object)
	if ok == false {
		failureMessage := "Can't Use Len"

		actual := fmt.Sprintf("%T(%#v)", object, object)
		fail(t, actual, length, failureMessage)
		return
	}

	if l != length {
		failureMessage := "Incorrect Len"
		fail(t, l, length, failureMessage)
		return
	}
}

// assert.Contains(t, rtnList, "abc")
func Contains(t TestingT, s, contains interface{}) {
	ok, found := includeElement(s, contains)
	if ok == false {
		failureMessage := "Can't Use Contains"

		actual := fmt.Sprintf("%T(%#v)", s, s)
		fail(t, actual, contains, failureMessage)
		return
	}
	if found == false {
		failureMessage := "Not Contains"
		fail(t, s, contains, failureMessage)
		return
	}
}

// assert.NotContains(t, rtnList, "abc")
func NotContains(t TestingT, s, contains interface{}) {
	ok, found := includeElement(s, contains)
	if ok == false {
		failureMessage := "Can't Use NotContains"

		actual := fmt.Sprintf("%T(%#v)", s, s)
		fail(t, actual, contains, failureMessage)
		return
	}
	if found {
		failureMessage := "Contains"
		fail(t, s, contains, failureMessage)
		return
	}
}

// actual == expected
func Equal(t TestingT, actual, expected interface{}) {
	if objectsAreEqual(actual, expected) {
		return
	}

	actual, expected = formatUnequalValues(actual, expected)

	failureMessage := "Not Equal"

	fail(t, actual, expected, failureMessage)
}

// actual != expected
func NotEqual(t TestingT, actual, expected interface{}) {
	if objectsAreEqual(actual, expected) == false {
		return
	}

	actual, expected = formatUnequalValues(actual, expected)

	failureMessage := "Equal"

	fail(t, actual, expected, failureMessage)
}

// actual < expected
func LesserThan(t TestingT, actual, expected interface{}) {
	if isNumbericType(expected) == false ||
		isNumbericType(actual) == false {
		failureMessage := "Can't Use LesserThan"
		fail(t, actual, expected, failureMessage)
		return
	}

	var f float64
	if reflect.ValueOf(actual).Convert(reflect.TypeOf(f)).Float() <
		reflect.ValueOf(expected).Convert(reflect.TypeOf(f)).Float() {
		return
	}

	actual, expected = formatUnequalValues(actual, expected)

	failureMessage := "Greater Equal"

	fail(t, actual, expected, failureMessage)
}

// actual <= expected
func LesserEqual(t TestingT, actual, expected interface{}) {
	if isNumbericType(expected) == false ||
		isNumbericType(actual) == false {
		failureMessage := "Can't Use LesserThan"
		fail(t, actual, expected, failureMessage)
		return
	}

	var f float64
	if reflect.ValueOf(actual).Convert(reflect.TypeOf(f)).Float() <=
		reflect.ValueOf(expected).Convert(reflect.TypeOf(f)).Float() {
		return
	}

	actual, expected = formatUnequalValues(actual, expected)

	failureMessage := "Greater Than"

	fail(t, actual, expected, failureMessage)
}

// actual > expected
func GreaterThan(t TestingT, actual, expected interface{}) {
	if isNumbericType(expected) == false ||
		isNumbericType(actual) == false {
		failureMessage := "Can't Use LesserThan"
		fail(t, actual, expected, failureMessage)
		return
	}

	var f float64
	if reflect.ValueOf(actual).Convert(reflect.TypeOf(f)).Float() >
		reflect.ValueOf(expected).Convert(reflect.TypeOf(f)).Float() {
		return
	}

	actual, expected = formatUnequalValues(actual, expected)

	failureMessage := "Lesser Equal"

	fail(t, actual, expected, failureMessage)
}

// actual >= expected
func GreaterEqual(t TestingT, actual, expected interface{}) {
	if isNumbericType(expected) == false ||
		isNumbericType(actual) == false {
		failureMessage := "Can't Use LesserThan"
		fail(t, actual, expected, failureMessage)
		return
	}

	var f float64
	if reflect.ValueOf(actual).Convert(reflect.TypeOf(f)).Float() >=
		reflect.ValueOf(expected).Convert(reflect.TypeOf(f)).Float() {
		return
	}

	actual, expected = formatUnequalValues(actual, expected)

	failureMessage := "Lesser Than"

	fail(t, actual, expected, failureMessage)
}
