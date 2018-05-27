package mmi

import ("math"
		
		)

/*
Bakun and Wentworth (1997): California, USA - best used on sources > 20km (Cua et al, 2010)
Accepts Moment Magnitude and distance in KM and returns a Modified Mercalli Intensity 
https://earthquake.usgs.gov/learn/topics/mercalli.php 
*/
func GetIntensity(magnitude, distance float64) float64 {
	return (3.67 + 1.17 * magnitude) - (3.19*(math.Log10(distance)))
}


/*
Atkinson and Wald (2007): California, USA - best for sources < 30km (Cua et al, 2010)
Accepts Moment Magnitude and distance in KM and returns a Modified Mercalli Intensity 
https://earthquake.usgs.gov/learn/topics/mercalli.php 
*/
func GetIntensity2(magnitude, distance float64) float64 {
	var r float64 = math.Sqrt(math.Pow(distance, 2) + 196)
	var b float64
	p := &b

	if  r <= 30 {
		*p = 0
	} else {
 		*p = math.Log10(r/30)
	}
	return 12.27 + 2.270 * (magnitude - 6.0) + 0.1304 * math.Pow((magnitude - 6.0), 2) - 1.30 * math.Log10(r) - 0.0007070 * r + 1.95 * *p - 0.577 * magnitude * math.Log10(r)
}


/*
Allen and Wald (2010): Global Active Crust - best general performance (Cua et al, 2010)
Accepts Moment Magnitude and distance in KM and returns a Modified Mercalli Intensity 
https://earthquake.usgs.gov/learn/topics/mercalli.php 
*/
func GetIntensity3 (magnitude, distance float64) float64 {
  return 3.15 + 1.03 * magnitude - 1.11 * (math.Log(math.Sqrt(math.Pow(distance, 2) + (math.Pow(1 + 0.72 * (math.Pow(math.E, (magnitude - 5.0))), 2)))))
}

/*
References:
G. Cua, D. J. Wald, T. I. Allen, D. Garcia, C. B. Worden, M. Gerstenberger, K. Lin, K. Marano (2010) 
“Best Practices” for Using Macroseismic Intensity and Ground Motion Intensity Conversion Equations for Hazard and Loss Models in GEM1,
GEM Technical Report 2010-4
http://www.globalquakemodel.org/media/publication/GEM-TechnicalReport_2010-4.pdf
  
Allen T. I., and Wald D. J. [2010] “Prediction of macroseismic intensities for global active crustal earthquakes“, J. Seismol. in
prep.

Bakun W. H. and Wentworth C. M. [1997] “Estimating earthquake location and magnitude from seismic intensity data”, Bull.
Seism. Soc. Am. 87, 1502-1521.
*/