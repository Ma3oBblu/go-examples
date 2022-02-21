package internal

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	type args struct {
		l []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "пустое значение",
			args: args{},
			want: make([]int64, 0),
		},
		{
			name: "пустое массив",
			args: args{},
			want: make([]int64, 0),
		},
		{
			name: "неуникальные значения",
			args: args{
				l: []int64{1, 2, 3, 7, 2, 8, 1, 7, 7},
			},
			want: []int64{1, 2, 3, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := Unique(tt.args.l); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Unique() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestUpdateSlice(t *testing.T) {
	type fields struct {
		SliceNums []int64
	}
	type args struct {
		updater Updater
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantHasChange bool
		wantSliceNums []int64
	}{
		{
			name: "нет изменений",
			fields: fields{
				SliceNums: []int64{1, 2, 3},
			},
			args: args{
				updater: func(int64s []int64) []int64 {
					return int64s
				},
			},
			wantHasChange: false,
			wantSliceNums: []int64{1, 2, 3},
		},
		{
			name: "есть изменения",
			fields: fields{
				SliceNums: []int64{1, 2, 3},
			},
			args: args{
				updater: func(int64s []int64) []int64 {
					return []int64{2}
				},
			},
			wantHasChange: true,
			wantSliceNums: []int64{2},
		},
		{
			name: "есть изменения",
			fields: fields{
				SliceNums: []int64{1, 2, 3, 4, 228},
			},
			args: args{
				updater: func(int64s []int64) []int64 {
					return append(int64s, []int64{228, 1090296}...)
				},
			},
			wantHasChange: true,
			wantSliceNums: []int64{1, 2, 3, 4, 228, 1090296},
		},
		{
			name: "есть изменения",
			fields: fields{
				SliceNums: []int64{1, 2, 3, 4, 225, 228},
			},
			args: args{
				updater: func(int64s []int64) []int64 {
					return append(int64s, []int64{225, 228, 1090296}...)
				},
			},
			wantHasChange: true,
			wantSliceNums: []int64{1, 2, 3, 4, 225, 228, 1090296},
		},
		{
			name: "есть изменения",
			fields: fields{
				SliceNums: []int64{1, 2, 3},
			},
			args: args{
				updater: func(int64s []int64) []int64 {
					if isArraysIntersects(int64s, []int64{7, 8}) {
						return append(int64s, []int64{11, 12}...)
					}
					return int64s
				},
			},
			wantHasChange: false,
			wantSliceNums: []int64{1, 2, 3},
		},
		{
			name: "есть изменения",
			fields: fields{
				SliceNums: []int64{1, 2, 7, 3},
			},
			args: args{
				updater: func(int64s []int64) []int64 {
					if isArraysIntersects(int64s, []int64{7, 8}) {
						return append(int64s, []int64{11, 12}...)
					}
					return int64s
				},
			},
			wantHasChange: true,
			wantSliceNums: []int64{1, 2, 7, 3, 11, 12},
		},
		{
			name: "есть изменения",
			fields: fields{
				SliceNums: []int64{1, 2, 3},
			},
			args: args{
				updater: func(int64s []int64) []int64 {
					if isArraysIntersects(int64s, []int64{1, 4}) {
						return append(int64s, []int64{5}...)
					}
					return int64s
				},
			},
			wantHasChange: true,
			wantSliceNums: []int64{1, 2, 3, 5},
		},
		{
			name: "есть изменения",
			fields: fields{
				SliceNums: []int64{1, 2, 3},
			},
			args: args{
				updater: func(int64s []int64) []int64 {
					if isArraysIntersects(int64s, []int64{1}) {
						return append(int64s, []int64{1, 2}...)
					}
					return int64s
				},
			},
			wantHasChange: false,
			wantSliceNums: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				p := &Slice{
					SliceNums: tt.fields.SliceNums,
				}
				if gotHasChange := p.UpdateSlice(tt.args.updater); gotHasChange != tt.wantHasChange {
					t.Errorf("UpdateMicroCategoryIDs() = %v, want %v", gotHasChange, tt.wantHasChange)
				}
				if !reflect.DeepEqual(p.SliceNums, tt.wantSliceNums) {
					t.Errorf("SliceNums = %v, want %v", p.SliceNums, tt.wantSliceNums)
				}
			},
		)
	}
}
