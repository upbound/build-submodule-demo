package generics

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestContains(t *testing.T) {
	type args struct {
		something []string
		check     string
	}
	type want struct {
		retVal bool
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"Valid": {
			reason: "If slice contains the value return true.",
			args: args{
				something: []string{"a", "b", "c"},
				check:     "a",
			},
			want: want{
				retVal: true,
			},
		},
		"Invalid": {
			reason: "If slice does not contains the value return false.",
			args: args{
				something: []string{"a", "b", "c"},
				check:     "X",
			},
			want: want{
				retVal: false,
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			s := Contains(tc.args.something, tc.args.check)
			if diff := cmp.Diff(tc.want.retVal, s); diff != "" {
				t.Errorf("\n%s\nContains(...): -want Val, +got Val:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type someStruct struct {
		a string
		b bool
		c int
	}
	type args struct {
		something []someStruct
		selector  func(s someStruct) string
	}
	type want struct {
		retVal []string
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"Valid": {
			reason: "If valid we should return a slice of items with properties selected by the selector func.",
			args: args{
				something: []someStruct{
					{
						a: "a",
						b: true,
						c: 1,
					},
					{
						a: "b",
						b: false,
						c: 2,
					},
					{
						a: "c",
						b: true,
						c: 3,
					},
				},
				selector: func(s someStruct) string { return s.a },
			},
			want: want{
				retVal: []string{"a", "b", "c"},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			s := Reduce(tc.args.something, tc.args.selector)
			if diff := cmp.Diff(tc.want.retVal, s, cmpopts.SortSlices(func(a, b string) bool { return a < b })); diff != "" {
				t.Errorf("\n%s\nReduce(...): -want Val, +got Val:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		something []string
		selector  func(s string) bool
	}
	type want struct {
		retVal []string
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"Valid": {
			reason: "If valid we should return a slice of items matching selector func.",
			args: args{
				something: []string{"a", "b", "c"},
				selector:  func(s string) bool { return s == "b" || s == "c" },
			},
			want: want{
				retVal: []string{"b", "c"},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			s := Filter(tc.args.something, tc.args.selector)
			if diff := cmp.Diff(tc.want.retVal, s, cmpopts.SortSlices(func(a, b string) bool { return a < b })); diff != "" {
				t.Errorf("\n%s\nFilter(...): -want Val, +got Val:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type someStruct struct {
		X string
		Y bool
		Z int
	}
	type args struct {
		something []someStruct
		selector  func(s someStruct) string
	}
	type want struct {
		retVal map[string]someStruct
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"Valid": {
			reason: "If valid we return a map of structs where the key is the value returned by the selector func.",
			args: args{
				something: []someStruct{
					{
						X: "a",
						Y: true,
						Z: 1,
					},
					{
						X: "b",
						Y: false,
						Z: 2,
					},
					{
						X: "c",
						Y: true,
						Z: 3,
					},
				},
				selector: func(s someStruct) string { return s.X },
			},
			want: want{
				retVal: map[string]someStruct{
					"a": {
						X: "a",
						Y: true,
						Z: 1,
					},
					"b": {
						X: "b",
						Y: false,
						Z: 2,
					},
					"c": {
						X: "c",
						Y: true,
						Z: 3,
					},
				},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			s := Map(tc.args.something, tc.args.selector)
			if diff := cmp.Diff(tc.want.retVal, s, cmpopts.SortSlices(func(a, b string) bool { return a < b })); diff != "" {
				t.Errorf("\n%s\nMap(...): -want Val, +got Val:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type args struct {
		something []string
	}
	type want struct {
		retVal []string
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"Valid": {
			reason: "If valid we should return a slice of items where the items are all unique.",
			args: args{
				something: []string{"a", "a", "b", "b", "b", "b", "b", "b", "c"},
			},
			want: want{
				retVal: []string{"a", "b", "c"},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			s := Unique(tc.args.something)
			if diff := cmp.Diff(tc.want.retVal, s, cmpopts.SortSlices(func(a, b string) bool { return a < b })); diff != "" {
				t.Errorf("\n%s\nUnique(...): -want Val, +got Val:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestMapReduce(t *testing.T) {
	type testStruct struct {
		A string
		B bool
		C uint
	}
	type args struct {
		something map[uint]testStruct
		ks        func(k uint, v testStruct) string
		vs        func(k uint, v testStruct) uint
	}
	type want struct {
		retVal map[string]uint
	}
	cases := map[string]struct {
		reason string
		args   args
		want   want
	}{
		"Valid": {
			reason: "If valid we should return a map of items where the are keyed by the selector and values are the return of value selector.",
			args: args{
				something: map[uint]testStruct{
					0: {
						A: "a",
						B: false,
						C: 10,
					},
					1: {
						A: "b",
						B: false,
						C: 11,
					},
					2: {
						A: "c",
						B: false,
						C: 12,
					},
				},
				ks: func(k uint, v testStruct) string { return v.A },
				vs: func(k uint, v testStruct) uint { return v.C },
			},
			want: want{
				retVal: map[string]uint{
					"a": 10,
					"b": 11,
					"c": 12,
				},
			},
		},
		"ValidUseKey": {
			reason: "If valid we should return a map of items where the are keyed by the selector and values are the return of value selector.",
			args: args{
				something: map[uint]testStruct{
					0: {
						A: "a",
						B: false,
						C: 10,
					},
					1: {
						A: "b",
						B: false,
						C: 11,
					},
					2: {
						A: "c",
						B: false,
						C: 12,
					},
				},
				ks: func(k uint, v testStruct) string { return v.A },
				vs: func(k uint, v testStruct) uint { return k },
			},
			want: want{
				retVal: map[string]uint{
					"a": 0,
					"b": 1,
					"c": 2,
				},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			s := FilterMap(tc.args.something, tc.args.ks, tc.args.vs)
			if diff := cmp.Diff(tc.want.retVal, s, cmpopts.SortMaps(func(a, b string) bool { return a < b })); diff != "" {
				t.Errorf("\n%s\nUnique(...): -want Val, +got Val:\n%s", tc.reason, diff)
			}
		})
	}
}
