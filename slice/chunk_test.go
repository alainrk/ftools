package slice_test

import (
	"reflect"
	"testing"

	"github.com/alainrk/ftools/slice"
)

func TestChunkEmpty(t *testing.T) {
	l := []int{}
	res := slice.Chunk(l, 0)
	exp := []int{}
	if len(res) != 0 {
		t.Errorf("Expected %v, got %v", exp, res)
	}

	res = slice.Chunk(l, 3)
	if len(res) != 0 {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}

func TestChunk(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := slice.Chunk(l, 2)
	exp := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}

	l = []int{1, 2, 3}
	res = slice.Chunk(l, 1)
	exp = [][]int{{1}, {2}, {3}}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}

	l = []int{1, 2, 3}
	res = slice.Chunk(l, 5)
	exp = [][]int{{1, 2, 3}}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}

	l = []int{1, 2, 3, 4}
	res = slice.Chunk(l, 3)
	exp = [][]int{{1, 2, 3}, {4}}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}

	l = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res = slice.Chunk(l, 4)
	exp = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10}}

	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
