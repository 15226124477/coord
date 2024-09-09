package coord

import (
	"github.com/15226124477/method"
	"math"
)

const (
	NEZ = 0
	BLH = 1
	XYZ = 2
)

// CoordinateNEZ NEZ坐标
type CoordinateNEZ struct {
	N float64
	E float64
	Z float64
}

// CoordinateBLH BLH坐标
type CoordinateBLH struct {
	B  float64
	L  float64
	H  float64
	L0 float64 // 中央经线
}

// CoordinateXYZ XYZ坐标
type CoordinateXYZ struct {
	X float64
	Y float64
	Z float64
}

type Coordinate struct {
	CoordinateNEZ
	CoordinateBLH
	CoordinateXYZ
	ConvertBefore int
	ConvertAfter  int
}

// BLH2XYZ 坐标转换BLH - > XYZ
func (pos *Coordinate) BLH2XYZ() {
	a := 6378137.0
	f := 1 / 298.257223563

	lon := pos.CoordinateBLH.L // 经度 L
	lat := pos.CoordinateBLH.B // 纬度 B
	H1 := pos.CoordinateBLH.H  // 高程

	b := a * (1.0 - f)
	e := math.Sqrt(a*a-b*b) / a
	N := a / math.Sqrt(1-e*e*math.Sin(lat*math.Pi/180)*math.Sin(lat*math.Pi/180))
	X := (N + H1) * math.Cos(lat*math.Pi/180) * math.Cos(lon*math.Pi/180)
	Y := (N + H1) * math.Cos(lat*math.Pi/180) * math.Sin(lon*math.Pi/180)
	Z := (N*(1-(e*e)) + H1) * math.Sin(lat*math.Pi/180)

	pos.CoordinateXYZ.X = method.Decimal(X, 4)
	pos.CoordinateXYZ.Y = method.Decimal(Y, 4)
	pos.CoordinateXYZ.Z = method.Decimal(Z, 4)
}

// BLH2NEZ 坐标转换BLH - > NEZ
func (pos *Coordinate) BLH2NEZ() {
	a := 6378137.0
	f := 1 / 298.257223563

	B := pos.CoordinateBLH.B
	L := pos.CoordinateBLH.L

	//计算中央子午线
	no := math.Floor((L + 1.5) / 3)
	L0 := no * 3
	// start.Info(L0)

	l := (L - L0) * 3600
	B = B * math.Pi / 180
	b := a * (1 - f)
	e2 := (a*a - b*b) / (a * a)
	e4 := (a*a - b*b) / (b * b)
	n2 := e4 * math.Pow(math.Cos(B), 2)
	t := math.Tan(B)
	c := a * a / b
	V := math.Sqrt(1 + e4*math.Pow(math.Cos(B), 2))
	N := c / V

	m0 := a * (1 - e2)
	m2 := 3 / 2.0 * e2 * m0
	m4 := 5 / 4.0 * e2 * m2
	m6 := 7 / 6.0 * e2 * m4
	m8 := 9 / 8.0 * e2 * m6

	a0 := m0 + 1/2.0*m2 + 3/8.0*m4 + 5/16.0*m6 + 35/128.0*m8
	a2 := 1/2.0*m2 + 1/2.0*m4 + 15/32.0*m6 + 7/16.0*m8
	a4 := 1/8.0*m4 + 3/16.0*m6 + 7/32.0*m8
	a6 := 1/32.0*m6 + 1/16.0*m8
	a8 := 1 / 128.0 * m8
	p := 206264.806247096355
	X := a0*B - a2/2*math.Sin(2*B) + a4/4*math.Sin(4*B) - a6/6*math.Sin(6*B) + a8/8*math.Sin(8*B)
	x := X + N/(2*math.Pow(p, 2))*math.Sin(B)*math.Cos(B)*l*l + N/(24*math.Pow(p, 4))*math.Sin(B)*math.Pow(math.Cos(B), 3)*(5-t*t+9*n2+4*n2*n2)*math.Pow(l, 4) + N/(720*math.Pow(p, 6))*math.Sin(B)*math.Pow(math.Cos(B), 5)*(61-58*t*t+math.Pow(t, 4))*math.Pow(l, 6)
	y := N/p*math.Cos(B)*l + N/(6*math.Pow(p, 3))*math.Pow(math.Cos(B), 3)*(1-t*t+n2)*math.Pow(l, 3) + N/(120*math.Pow(p, 5))*math.Pow(math.Cos(B), 5)*(5-18*t*t+math.Pow(t, 4)+14*n2-58*n2*t*t)*math.Pow(l, 5)
	y = y + 500000

	// NEZ := [...]float64{x, y, BLH_arr[2]}
	pos.CoordinateNEZ.N = method.Decimal(x, 4)
	pos.CoordinateNEZ.E = method.Decimal(y, 4)
}

// XYZ2BLH 坐标转换XYZ - > BLH
func (pos *Coordinate) XYZ2BLH() {
	a := 6378137.0
	f := 1 / 298.257223563

	N := 0.0
	X := pos.CoordinateXYZ.X
	Y := pos.CoordinateXYZ.Y
	Z := pos.CoordinateXYZ.Z
	b := a * (1 - f)
	e2 := (a*a - b*b) / (a * a)
	// e4 := (a * a - b * b) / (b * b)
	t0 := Z / math.Sqrt(X*X+Y*Y)
	t := t0
	for {
		lastT := t
		B := math.Atan(t)
		N = a / math.Sqrt(1-e2*math.Pow(math.Sin(B), 2))
		t = (Z + N*e2*math.Sin(B)) / math.Sqrt(X*X+Y*Y)
		if math.Abs(t-lastT) <= 5e-10 {
			break
		}
	}
	B := math.Atan(t)
	H := Z/math.Sin(B) - N*(1-e2)
	L := math.Atan(Y/X) * 180 / math.Pi
	if L < 0 {
		L = L + 180
	}
	B = B * 180 / math.Pi

	pos.CoordinateBLH.B = method.Decimal(B, 10)
	pos.CoordinateBLH.L = method.Decimal(L, 10)
	pos.CoordinateBLH.H = method.Decimal(H, 4)
}

func (pos *Coordinate) Convert() {
	switch pos.ConvertBefore {
	case NEZ:
		// Next
		pos.NEZ2BLH()
	case XYZ:
		pos.XYZ2BLH()
		pos.BLH2NEZ()
		pos.CoordinateNEZ.Z = pos.CoordinateBLH.H
	case BLH:
		pos.BLH2NEZ()
		pos.CoordinateNEZ.Z = pos.CoordinateBLH.H
	}
}

func (pos *Coordinate) NEZ2BLH() {
	x := pos.N
	y := pos.E
	a := 6378137.0
	f := 1 / 298.257223563
	//double width = 3;
	L0 := pos.L0
	y = y - 500000
	b := a * (1 - f)
	c := a * a / b
	e2 := (a*a - b*b) / (a * a)
	e4 := (a*a - b*b) / (b * b)

	m0 := a * (1 - e2)
	m2 := 3 / 2.0 * e2 * m0
	m4 := 5 / 4.0 * e2 * m2
	m6 := 7 / 6.0 * e2 * m4
	m8 := 9 / 8.0 * e2 * m6

	a0 := m0 + 1/2.0*m2 + 3/8.0*m4 + 5/16.0*m6 + 35/128.0*m8
	a2 := 1/2.0*m2 + 1/2.0*m4 + 15/32.0*m6 + 7/16.0*m8
	a4 := 1/8.0*m4 + 3/16.0*m6 + 7/32.0*m8
	a6 := 1/32.0*m6 + 1/16.0*m8
	a8 := 1 / 128.0 * m8
	//初值
	Bf0 := x / a0
	Bf := Bf0

	for {
		last_Bf := Bf
		Bf = 1 / a0 * (x + a2/2*math.Sin(2*Bf) - a4/4*math.Sin(4*Bf) + a6/6*math.Sin(6*Bf) - a8/8*math.Sin(8*Bf))
		if math.Abs(Bf-last_Bf) <= 5e-10 {
			break
		}
	}

	tf := math.Tan(Bf)
	nf2 := e4 * math.Pow(math.Cos(Bf), 2)
	Vf := math.Sqrt(1 + e4*math.Pow(math.Cos(Bf), 2))
	Nf := c / Vf
	Mf := c / math.Pow(Vf, 3)
	B := Bf - tf/(2*Mf*Nf)*y*y + tf/(24*Mf*math.Pow(Nf, 3))*(5+3*tf*tf+nf2-9*nf2*tf*tf)*math.Pow(y, 4) - tf/(720*Mf*math.Pow(Nf, 5))*(61+90*tf*tf+45*math.Pow(tf, 4))*math.Pow(y, 6)
	L := 1/(Nf*math.Cos(Bf))*y - 1/(6*math.Pow(Nf, 3)*math.Cos(Bf))*(1+2*tf*tf+nf2)*math.Pow(y, 3) + 1/(120*math.Pow(Nf, 5)*math.Cos(Bf))*(5+28*tf*tf+24*math.Pow(tf, 4)+6*nf2+8*nf2*tf*tf)*math.Pow(y, 5)
	B = B * 180 / math.Pi
	L = L*180/math.Pi + float64(L0)

	pos.B = B
	pos.L = L
	pos.H = pos.CoordinateNEZ.Z

}

func (pos *Coordinate) Value() {

}
