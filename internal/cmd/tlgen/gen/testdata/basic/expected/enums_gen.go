// Code generated by generate-tl-files; DO NOT EDIT.

package telegram

type StorageFileType uint32

const (
	StorageFileGif     StorageFileType = 0xcae1aadf
	StorageFileJpeg    StorageFileType = 0x7efe0e
	StorageFileMov     StorageFileType = 0x4b09ebbc
	StorageFileMp3     StorageFileType = 0x528a0677
	StorageFileMp4     StorageFileType = 0xb3cea0e4
	StorageFilePartial StorageFileType = 0x40bc6f52
	StorageFilePdf     StorageFileType = 0xae1e508d
	StorageFilePng     StorageFileType = 0xa4f63c0
	StorageFileUnknown StorageFileType = 0xaa963b05
	StorageFileWebp    StorageFileType = 0x1081464c
)

func (e StorageFileType) String() string {
	switch e {
	case StorageFileType(0xcae1aadf):
		return "storage.fileGif"
	case StorageFileType(0x7efe0e):
		return "storage.fileJpeg"
	case StorageFileType(0x4b09ebbc):
		return "storage.fileMov"
	case StorageFileType(0x528a0677):
		return "storage.fileMp3"
	case StorageFileType(0xb3cea0e4):
		return "storage.fileMp4"
	case StorageFileType(0x40bc6f52):
		return "storage.filePartial"
	case StorageFileType(0xae1e508d):
		return "storage.filePdf"
	case StorageFileType(0xa4f63c0):
		return "storage.filePng"
	case StorageFileType(0xaa963b05):
		return "storage.fileUnknown"
	case StorageFileType(0x1081464c):
		return "storage.fileWebp"
	default:
		return "<UNKNOWN storage.FileType>"
	}
}

func (e StorageFileType) CRC() uint32 { return uint32(e) }
