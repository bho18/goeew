// Package mmi provides several empirically–derived Modified Mercalli
// Intensity (MMI) prediction equations (Intensity Prediction Equations – IPEs)
// and a simple helper for selecting the most appropriate one at runtime.
//
// References (coefficients reproduced verbatim):
//   - Bakun & Wentworth (1997) – Bull. Seismol. Soc. Am. 87, 1502‑1521
//   - Atkinson & Wald (2007)   – Seismol. Res. Lett. 78, 362‑372
//   - Allen & Wald (2012)      – J. Seismol. 16, 409‑433 (global active crust)
//   - "Best Practices" (Cua et al., 2010) – GEM Tech. Rep. 2010‑4
//
// Each function returns the predicted MMI (decimal, not rounded to the nearest
// integer) given moment magnitude M (Mw) and source‑to‑site distance R (km).
//
// All equations are valid for shallow crustal earthquakes; see individual
// comments for regional applicability and distance limits.
package mmi

import "math"

// BakunWentworth97 implements equation (4) of Bakun & Wentworth (1997):
//
//	MMI = 3.67 + 1.17·M – 3.19·log10(R)
//
// Best performance is for Californian events at R > 20 km.
func BakunWentworth97(M, R float64) float64 {
	if R <= 0 {
		R = 0.1 // avoid log10(0) – outside published range but safer than NaN
	}
	return 3.67 + 1.17*M - 3.19*math.Log10(R)
}

// AtkinsonWald07 implements equation (1) with California coefficients from
// Atkinson & Wald (2007). It is recommended for R ≤ 30 km.
//
//	let r = √(R² + h²)          with h = 14 km
//	    ΔM = M – 6
//	    p = log10(r/30) if r > 30 km, else 0
//	MMI = c1 + c2·ΔM + c3·ΔM² + c4·log10(r) + c5·r + c6·p + c7·M·log10(r)
//
// where (California coefficients):
//
//	c1 = 12.27,  c2 =  2.270,  c3 = 0.1304,  c4 = ‑1.30,
//	c5 = –0.0007070, c6 = 1.95, c7 = ‑0.577
func AtkinsonWald07(M, R float64) float64 {
	if R <= 0 {
		R = 0.1
	}
	const h = 14.0
	r := math.Hypot(R, h) // √(R² + h²)

	var p float64
	if r > 30 {
		p = math.Log10(r / 30.0)
	}
	dM := M - 6.0

	return 12.27 +
		2.270*dM +
		0.1304*dM*dM +
		-1.30*math.Log10(r) +
		-0.0007070*r +
		1.95*p +
		-0.577*M*math.Log10(r)
}

// AllenWald12 implements the global active‑crust IPE of Allen & Wald (2012)
// (formerly "Allen & Wald 2010, in prep.").
//
//	d* = √( R² + [1 + 0.72·e^{M‑5}]² )
//	MMI = 3.15 + 1.03·M – 1.11·ln(d*)                     (natural log)
func AllenWald12(M, R float64) float64 {
	if R < 0 {
		R = 0
	}
	dStar := math.Hypot(R, 1.0+0.72*math.Exp(M-5.0))
	return 3.15 + 1.03*M - 1.11*math.Log(dStar)
}

// BestEstimate selects an IPE based on simple distance heuristics suggested by
// Cua et al. (2010):
//
//	– At R < 30 km, use Atkinson–Wald (2007)
//	– At R > 20 km, use Bakun–Wentworth (1997)
//	– Where ranges overlap, fall back to the global Allen–Wald model.
//
// Callers needing stricter regional control should invoke the specific models
// directly.
func BestEstimate(M, R float64) float64 {
	switch {
	case R <= 20:
		return AtkinsonWald07(M, R)
	case R >= 30:
		return BakunWentworth97(M, R)
	default:
		return AllenWald12(M, R)
	}
}
