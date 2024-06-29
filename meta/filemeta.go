package meta

// import "sort"
import (
	mydb "cloud-storage/db"
	"log"
)

// FileMeta -> struct for metadata of the file
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta: update/ modify metadata
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

// get metadata of the file
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

// get a batch of file meta in form of array
// func GetLastFileMetas(count int) []FileMeta {
// 	fMetaArray := make([]FileMeta, len(fileMetas))

// 	for _, v := range fileMetas {
// 		fMetaArray = append(fMetaArray, v)
// 	}

// 	sort.Sort(ByUploadTime(fMetaArray))
// 	return fMetaArray[0:count]
// }

func UpdateFileMetaDB(fmeta FileMeta) bool {
	return mydb.OnFileUploadFinished(fmeta.FileSha1,
		fmeta.FileName,
		fmeta.FileSize,
		fmeta.Location)
}

func RemoveFileMeta(filesha1 string) {
	delete(fileMetas, filesha1)
}

// get metadata from mysql
func GetFileMetaDB(fileSha1 string) (FileMeta, error) {
	tfile, err := mydb.GetFileMeta(fileSha1)
	if err != nil {
		log.Fatalln(err.Error())
		return FileMeta{}, err
	}
	fmeta := FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}
	return fmeta, nil

}
