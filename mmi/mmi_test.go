package mmi

import "testing"

func TestBestEstimateOverlapUsesAllenWald12(t *testing.T) {
	m := 6.0
	r := 25.0
	want := AllenWald12(m, r)
	got := BestEstimate(m, r)
	if got != want {
		t.Errorf("BestEstimate(%v,%v) = %v, want %v", m, r, got, want)
	}
}
