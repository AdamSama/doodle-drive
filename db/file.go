package db

import (
	mydb "cloud-storage/db/mysql"
	"database/sql"
	"fmt"
	"log"
)

// upload file finished, store meta data into mysql database
func OnFileUploadFinished(filehash string, filename string, filesize int64, fileaddr string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_file(`file_sha1`, `file_name`, `file_size`, " +
			"`file_addr`, `status`) values(?,?,?,?,1)",
	)
	if err != nil {
		fmt.Println("Failed to prepare statement, err: ")
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Printf("File with hash: %s has been uploaded before\n", filehash)
		}
		return true
	}
	return false
}

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

// get file metadata from mysql
func GetFileMeta(filehash string) (*TableFile, error) {
	stmt, err := mydb.DBConn().Prepare(
		`SELECT file_sha1, file_addr, file_name, file_size 
		 FROM tbl_file 
		 WHERE file_sha1=? AND status=1 
		 LIMIT 1`,
	)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	defer stmt.Close()

	tfile := TableFile{}
	err = stmt.QueryRow(filehash).Scan(&tfile.FileHash, &tfile.FileAddr, &tfile.FileName, &tfile.FileSize)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	return &tfile, nil
}
