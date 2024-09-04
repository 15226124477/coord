package coord

import (
	"github.com/15226124477/method"
)

// Sat 卫星数
type Sat struct {
	SatNum int
}

// Diff 差分龄期
type Diff struct {
	DiffValue float64
}

// Sol 解状态
type Sol struct {
	SolValue int
	SolMode  int
}

// Heading 定向 度,弧度
type Heading struct {
	HeadingRadian  float64
	HeadingDegrees float64
	InputType      int
}

func (h *Heading) Value() float64 {
	var result float64
	switch h.InputType {
	case method.DEGREES:
		// 返回度
		result = h.HeadingDegrees
	case method.RADIAN:
		// 弧度转换
		h.HeadingDegrees = method.Radians2Degrees(h.HeadingRadian)
		result = h.HeadingDegrees
	}
	return result
}

// Speed 速度
type Speed struct {
	SpeedValue float64
}

// DataRinex 20240202
// 01 时间
// 02 卫星数
type DataRinex struct {
	*MTime
	*Sat
}

// DataGGA NMEA语句 20240202
// 01 时间
// 02 解状态
// 03 卫星数
// 04 差分龄期
type DataGGA struct {
	*MTime
	*Sol
	*Sat
	*Diff
}

// DataPOS POS文件 20240202
// 01 时间
// 02 解状态
// 03 卫星数
// 04 差分龄期
type DataPOS struct {
	*MTime
	*Sol
	*Sat
	*Diff
}

// DataNormalCSV 标准CSV 20240202
// 01 时间
// 02 解状态
// 03 卫星数
// 04 差分龄期
type DataNormalCSV struct {
	*MTime
	*Sol
	*Sat
	*Diff
}

// DataAutoTestCSV 自动化测试固件CSV 20240202
// 01 时间
// 02 解状态
// 03 卫星数
// 04 差分龄期
type DataAutoTestCSV struct {
	*MTime
	*Sol
	*Sat
	*Diff
}

// DataBoatPVT 无人船PVT数据 20240202
// 01 时间
// 02 解状态
// 03 卫星数
// 04 差分龄期
// 05 定向
// 06 速度
type DataBoatPVT struct {
	*MTime
	*Sol
	*Sat
	*Diff
	*Heading
	*Speed
}
