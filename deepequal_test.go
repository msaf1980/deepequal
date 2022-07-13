// Derived from code Copyright 2009 The Go Authors. All rights reserved.

// Deep equality test via reflection

package deepequal

import (
	"math"
	"testing"
)

type testStruct struct {
	Name string
	S    []int
	M    map[int]string
}

func TestCompare(t *testing.T) {
	tests := []struct {
		name       string
		a1         interface{}
		a2         interface{}
		want       bool
		wantReason string
	}{
		{
			name:       "Equal array",
			a1:         []int{0, 1, 2},
			a2:         []int{0, 1, 2},
			want:       true,
			wantReason: "",
		},
		{
			name:       "Equal map",
			a1:         map[int]string{0: "0", 1: "1", 2: "2"},
			a2:         map[int]string{0: "0", 1: "1", 2: "2"},
			want:       true,
			wantReason: "",
		},
		{
			name: "Equal struct",
			a1: testStruct{
				Name: "s",
				S:    []int{0, 1, 2},
				M:    map[int]string{0: "0", 1: "1", 2: "2"},
			},
			a2: testStruct{
				Name: "s",
				S:    []int{0, 1, 2},
				M:    map[int]string{0: "0", 1: "1", 2: "2"},
			},
			want:       true,
			wantReason: "",
		},
		{
			name:       "int",
			a1:         2,
			a2:         2,
			want:       true,
			wantReason: "",
		},
		{
			name:       "int not equal",
			a1:         2,
			a2:         3,
			want:       false,
			wantReason: "scalar values differ",
		},
		{
			name:       "float64",
			a1:         2.0,
			a2:         2.0,
			want:       true,
			wantReason: "",
		},
		{
			name:       "float64 NaN",
			a1:         math.NaN(),
			a2:         math.NaN(),
			want:       true,
			wantReason: "",
		},
		{
			name:       "float64 NaN and number",
			a1:         math.NaN(),
			a2:         1.0,
			want:       false,
			wantReason: "scalar values differ",
		},
		{
			name: "Non Equal struct (slice elem)",
			a1: testStruct{
				Name: "s",
				S:    []int{0, 1, 2},
				M:    map[int]string{0: "0", 1: "1", 2: "2"},
			},
			a2: testStruct{
				Name: "s",
				S:    []int{0, 1, 4},
				M:    map[int]string{0: "0", 1: "1", 2: "2"},
			},
			want:       false,
			wantReason: "struct.S [2] scalar values differ",
		},
		{
			name: "Non Equal struct (slice elems len)",
			a1: testStruct{
				Name: "s",
				S:    []int{0, 1, 2},
				M:    map[int]string{0: "0", 1: "1", 2: "2"},
			},
			a2: testStruct{
				Name: "s",
				S:    []int{0, 1, 2, 5},
				M:    map[int]string{0: "0", 1: "1", 2: "2"},
			},
			want:       false,
			wantReason: "struct.S slices have different lengths",
		},
		{
			name: "Non Equal struct (map elems value mismatch)",
			a1: testStruct{
				Name: "s",
				S:    []int{0, 1, 2},
				M:    map[int]string{0: "0", 1: "1", 2: "2"},
			},
			a2: testStruct{
				Name: "s",
				S:    []int{0, 1, 2},
				M:    map[int]string{0: "0", 1: "1", 2: "1+1"},
			},
			want:       false,
			wantReason: "struct.M [2] scalar values differ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotReason := Compare(tt.a1, tt.a2)
			if got != tt.want {
				t.Errorf("Compare() got = %v, want %v", got, tt.want)
			}
			if gotReason != tt.wantReason {
				t.Errorf("Compare() got1 = '%v', want '%v'", gotReason, tt.wantReason)
			}
		})
	}
}
