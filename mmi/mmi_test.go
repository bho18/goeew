package mmi

import (
	"testing"
)

func almostEqual(a, b, tol float64) bool {
	if a > b {
		return a-b <= tol
	}
	return b-a <= tol
}

func TestBakunWentworth97(t *testing.T) {
	tests := []struct {
		M, R float64
		want float64
	}{
		{6.0, 10.0, 7.5},
		{7.0, 25.0, 7.400571372336199},
		{5.5, 100.0, 3.7250000000000005},
	}
	const tol = 1e-6
	for _, tt := range tests {
		got := BakunWentworth97(tt.M, tt.R)
		if !almostEqual(got, tt.want, tol) {
			t.Errorf("BakunWentworth97(%v,%v)=%.12g want %.12g", tt.M, tt.R, got, tt.want)
		}
	}
}

func TestAtkinsonWald07(t *testing.T) {
	tests := []struct {
		M, R float64
		want float64
	}{
		{6.0, 10.0, 6.373690748041067},
		{5.0, 5.0, 5.214248856709551},
		{6.0, 50.0, 4.529311815861192},
	}
	const tol = 1e-6
	for _, tt := range tests {
		got := AtkinsonWald07(tt.M, tt.R)
		if !almostEqual(got, tt.want, tol) {
			t.Errorf("AtkinsonWald07(%v,%v)=%.12g want %.12g", tt.M, tt.R, got, tt.want)
		}
	}
}

func TestAllenWald12(t *testing.T) {
	tests := []struct {
		M, R float64
		want float64
	}{
		{6.0, 10.0, 6.72760279786184},
		{5.5, 5.0, 6.931358655429676},
		{6.0, 40.0, 5.232318708429628},
	}
	const tol = 1e-6
	for _, tt := range tests {
		got := AllenWald12(tt.M, tt.R)
		if !almostEqual(got, tt.want, tol) {
			t.Errorf("AllenWald12(%v,%v)=%.12g want %.12g", tt.M, tt.R, got, tt.want)
		}
	}
}

func TestBestEstimateSelection(t *testing.T) {
	const tol = 1e-6
	type test struct {
		M, R float64
		want float64
	}
	tests := []test{
		{6.0, 15.0, AtkinsonWald07(6.0, 15.0)},
		{6.0, 25.0, AllenWald12(6.0, 25.0)},
		{6.0, 35.0, BakunWentworth97(6.0, 35.0)},
	}
	for _, tt := range tests {
		got := BestEstimate(tt.M, tt.R)
		if !almostEqual(got, tt.want, tol) {
			t.Errorf("BestEstimate(%v,%v)=%.12g want %.12g", tt.M, tt.R, got, tt.want)
		}
	}
}
