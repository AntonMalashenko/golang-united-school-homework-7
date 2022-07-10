package coverage

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func matchSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for idx := range a {
		if a[idx] != b[idx] {
			return false
		}
	}
	return true
}

func matchMatrix(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for idx := range a {
		if !matchSlice(a[idx], b[idx]) {
			return false
		}
	}
	return true
}

func TestPeopleLength(t *testing.T) {
	var testPeople People
	const expectedLen = 3

	for idx := range [expectedLen]int{} {
		testPeople = append(
			testPeople,
			Person{
				firstName: fmt.Sprintf("first%d", idx),
				lastName:  fmt.Sprintf("last%d", idx),
				birthDay:  time.Now(),
			},
		)
	}

	peopleLen := testPeople.Len()
	if peopleLen != expectedLen {
		t.Errorf("Expected len == %d; but got %d", expectedLen, peopleLen)
	}
}

func TestPeopleSwap(t *testing.T) {
	leftPerson := Person{
		firstName: "firstLeft",
		lastName:  "lastLeft",
		birthDay:  time.Now(),
	}
	rightPerson := Person{
		firstName: "firstRight",
		lastName:  "lastRight",
		birthDay:  time.Now(),
	}
	testPeople := People{leftPerson, rightPerson}
	testPeople.Swap(0, 1)

	if testPeople[0] != rightPerson || testPeople[1] != leftPerson {
		t.Errorf(
			"Expected testPeople[0].firstName == \"firstRight\" and testPeople[1].firstName == \"firstLeft\" \n" +
				"but got \"firstLeft\", \"firstRight\"",
		)
	}
}

func TestCompareBirthDay(t *testing.T) {
	baseTimeForCompare := time.Now()
	leftPerson := Person{
		firstName: "firstLeft",
		lastName:  "lastLeft",
		birthDay:  baseTimeForCompare.Add(-10 * time.Hour),
	}
	rightPerson := Person{
		firstName: "firstRight",
		lastName:  "lastRight",
		birthDay:  baseTimeForCompare.Add(-20 * time.Hour),
	}
	testPeople := People{leftPerson, rightPerson}
	isRightLess := testPeople.Less(0, 1)

	if !isRightLess {
		t.Errorf("Expected right less then left == true, but got False")
	}
}

func TestCompareFirstName(t *testing.T) {
	baseTimeForCompare := time.Now()
	leftPerson := Person{
		firstName: "aaaaaa",
		lastName:  "lastLeft",
		birthDay:  baseTimeForCompare,
	}
	rightPerson := Person{
		firstName: "bbbbbb",
		lastName:  "lastRight",
		birthDay:  baseTimeForCompare,
	}
	testPeople := People{leftPerson, rightPerson}
	isRightLess := testPeople.Less(0, 1)

	if !isRightLess {
		t.Errorf("Expected \"aaaaaa\" less then \"bbbbbb\" == true, but got False")
	}
}

func TestCompareLastName(t *testing.T) {
	baseTimeForCompare := time.Now()
	leftPerson := Person{
		firstName: "aaaaaa",
		lastName:  "aaaaaa",
		birthDay:  baseTimeForCompare,
	}
	rightPerson := Person{
		firstName: "aaaaaa",
		lastName:  "bbbbbb",
		birthDay:  baseTimeForCompare,
	}
	testPeople := People{leftPerson, rightPerson}
	isRightLess := testPeople.Less(0, 1)

	if !isRightLess {
		t.Errorf("Expected right less then left == true, but got False")
	}
}

////////////////////////////////////////////////////////////////////////////////////

func TestCreateMatrix(t *testing.T) {
	inputString := "1 2 3\n4 5 6\n7 8 9"
	expectedData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expectedRows, expectedCols := 3, 3

	matrix, err := New(inputString)
	if err != nil {
		t.Error(err.Error())
	}

	if matrix.rows != expectedRows {
		t.Errorf("Expected %d rows, but got %d", expectedRows, matrix.rows)
	}
	if matrix.cols != expectedCols {
		t.Errorf("Expected %d rows, but got %d", expectedCols, matrix.cols)
	}
	if !matchSlice(expectedData, matrix.data) {
		t.Errorf("Expected %v rows, but got %v", expectedData, matrix.data)
	}
}

func TestCreateMatrixColLenError(t *testing.T) {
	inputString := "1 2 3\n4 5 6\n7 8 9 9"
	expectedErrorMessage := "Rows need to be the same length"

	_, err := New(inputString)
	if err == nil || err.Error() != expectedErrorMessage {
		t.Errorf("Expected Error \"%s\", but got \"%v\"", expectedErrorMessage, err)
	}
}

func TestCreateMatrixValueError(t *testing.T) {
	inputString := "1 2 3\n4 5 6\n7 8 q"
	expectedErrorMessage := "strconv.Atoi: parsing \"q\": invalid syntax"
	_, err := New(inputString)
	if err == nil || err.Error() != expectedErrorMessage {
		t.Errorf("Expected \"%s\", but got, \"%s\"", expectedErrorMessage, err)
	}
}

func TestRows(t *testing.T) {
	inputString := "1 2 3\n4 5 6\n7 8 9"
	expectedResult := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrix, _ := New(inputString)
	result := matrix.Rows()
	if !matchMatrix(result, expectedResult) {
		t.Errorf("Expected %v, but got %v", expectedResult, result)
	}
}

func TestCols(t *testing.T) {
	inputString := "1 2 3\n4 5 6\n7 8 9"
	expectedResult := [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}

	matrix, _ := New(inputString)
	result := matrix.Cols()
	if !matchMatrix(result, expectedResult) {
		t.Errorf("Expected %v, but got %v", expectedResult, result)
	}
}

func TestSet(t *testing.T) {
	inputString := "1 2 3\n4 5 6\n7 8 9"
	expectedResult := []int{1, 2, 3, 4, 99, 6, 7, 8, 9}

	matrix, _ := New(inputString)
	_ = matrix.Set(1, 1, 99)
	if !matchSlice(matrix.data, expectedResult) {
		t.Errorf("Expected %v, but got %v", expectedResult, matrix.data)
	}
}

func TestSetError(t *testing.T) {
	inputString := "1 2 3\n4 5 6\n7 8 9"
	matrix, _ := New(inputString)
	result := matrix.Set(3, 3, 99)
	if result {
		t.Error("Expected result == false, but got true")
	}
}
