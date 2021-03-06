package common

import "testing"

func TestIsSetCover(t *testing.T) {
	arr1 := []string{"a", "B", "c", "d"}
	arr2 := []string{"a", "b", "c"}
	fPln(CanSetCover(arr1, arr2))
	arr1 = []string{"c", "b", "a"}
	arr2 = []string{"a", "b", "c"}
	fPln(CanSetCover(arr1, arr2))
	// arr3 := 6
	// arr4 := 7
	// fPln(CanSetCover(arr3, arr4))
}

func TestToSet(t *testing.T) {
	fPln(ToSet(nil))
	fPln(ToSet([]int{1, 3, 2, 1, 3, 5}))
	fPln(ToSet([]string{"1", "2", "3", "4", "1", "3", "2"}))
}

func TestSetIntersect(t *testing.T) {
	arr1 := []string{"A", "B", "c", "d"}
	arr2 := []string{"C", "d", "a"}
	fPln(SetIntersect(arr1, arr2).([]string))
}

func TestSetUnion(t *testing.T) {
	arr1 := []string{"A", "B", "c", "d"}
	arr2 := []string{"C", "d", "A"}
	fPln(SetUnion(arr1, arr2).([]string))
}

func TestToGeneralSlc(t *testing.T) {
	fn := func(arr []interface{}) {
		for _, e := range arr {
			fPln(e)
		}
	}
	slc1 := []string{"a", "b", "C"}
	fn(ToGeneralSlc(slc1))
	// fn(slc1)
}
