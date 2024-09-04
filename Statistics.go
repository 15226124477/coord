package coord

const (
	GGA     = 0
	POS     = 1
	RINEX   = 2
	BoatPVT = 3
)

type FileData struct {
	GGAList     []DataGGA
	POSList     []DataPOS
	RinexList   []DataRinex
	BoatPVTList []DataBoatPVT
	FileType    int
}

// SolIntegrity 固定率统计
func (fd *FileData) SolIntegrity() {
	switch fd.FileType {
	case GGA:

	case POS:

	case RINEX:

	case BoatPVT:

	}

}
