package reference

import (
	"github.com/fullstack-lang/gongfly/go/models"
)

var Sat1 = (&models.Satellite{
	Line1: "1 00005U 58002B   00179.78495062  .00000023  00000-0  28098-4 0  4753",
	Line2: "2 00005  34.2682 348.7242 1859667 331.7664  19.3264 10.82419157413667",
	Name:  "Sat 1",
}).InitFromTLE().Stage(&models.Stage).Register()

var Sat2 = (&models.Satellite{
	Line1: "1 04632U 70093B   04031.91070959 -.00000084  00000-0  10000-3 0  9955",
	Line2: "2 04632  11.4628 273.1101 1450506 207.6000 143.9350  1.20231981 44145",
	Name:  "Sat 2",
}).InitFromTLE().Stage(&models.Stage).Register()

var Sat3 = (&models.Satellite{
	Line1: "1 06251U 62025E   06176.82412014  .00008885  00000-0  12808-3 0  3985",
	Line2: "2 06251  58.0579  54.0425 0030035 139.1568 221.1854 15.56387291  6774",
	Name:  "Sat 3",
}).InitFromTLE().Stage(&models.Stage).Register()

var Sat4 = (&models.Satellite{
	Line1: "1 88888U          80275.98708465  .00073094  13844-3  66816-4 0    8",
	Line2: "2 88888  72.8435 115.9689 0086731  52.6988 110.5714 16.05824518  105",
	Name:  "Sat 4",
}).InitFromTLE().Stage(&models.Stage).Register()

var ISS = (&models.Satellite{
	Line1: "1 25544U 98067A   08264.51782528 -.00002182  00000-0 -11606-4 0  2927",
	Line2: "2 25544  51.6416 247.4627 0006703 130.5360 325.0288 15.72125391563537",
	Name:  "ISS",
}).InitFromTLE().Stage(&models.Stage).Register()

var Sat5 = (&models.Satellite{
	Line1: "1 04632U 70093B   04031.91070959 -.00000084  00000-0  10000-3 0  9955",
	Line2: "2 04632  11.4628 273.1101 1450506 207.6000 143.9350  1.20231981 44145",
	Name:  "Sat 5",
}).InitFromTLE().Stage(&models.Stage).Register()

var Sat6 = (&models.Satellite{
	Line1: "1 24208U 96044A   06177.04061740 -.00000094  00000-0  10000-3 0  1600",
	Line2: "2 24208   3.8536  80.0121 0026640 311.0977  48.3000  1.00778054 36119",
	Name:  "Sat 6",
}).InitFromTLE().Stage(&models.Stage).Register()

var Sat7 = (&models.Satellite{
	Line1: "1 06251U 62025E   06176.82412014  .00008885  00000-0  12808-3 0  3985",
	Line2: "2 06251  58.0579  54.0425 0030035 139.1568 221.1854 15.56387291  6774",
	Name:  "Sat 7",
}).InitFromTLE().Stage(&models.Stage).Register()

var Sat8 = (&models.Satellite{
	Line1: "1 23599U 95029B   06171.76535463  .00085586  12891-6  12956-2 0  2905",
	Line2: "2 23599   6.9327   0.2849 5782022 274.4436  25.2425  4.47796565123555",
	Name:  "Sat 8",
}).InitFromTLE().Stage(&models.Stage).Register()
