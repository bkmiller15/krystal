package file

import (
	"os"
	"io"
	"fmt"
	"regexp"
	"encoding/json"
	"io/ioutil"
	"../ds"
	"bytes"
)
func Write_Log(year, month, day string) {
	var subEntries []ds.Entry

	path := (ds.KConfig.LogPath + "/" + year)

	// Make sure that the log dir exists
	if ! exists(path) {
		create_log_dir(path)
	}

	path = path+"/"+month

	// Make sure that the log dir exists
	if ! exists(path) {
		create_log_dir(path)
	}

	path = path+"/"+ day + ".json"

	// Make sure that the log file exists
	if ! exists(path) {
		create_log_file(path)
	}

	subEntries = ds.Get_AllEntries_On_Date(year, month, day)

	if subEntries == nil {
		removeFile(path)
	} else {
		write_JSON_File(path, subEntries)
	}


}

func write_JSON_File(path string, i interface{}) {
	text, err := json.Marshal(i)

	var out bytes.Buffer
	json.Indent(&out, text, "", "\t")

	if err != nil {
	  fmt.Println("error:", err)
	  os.Exit(1)
	}

	err2 := ioutil.WriteFile(path, out.Bytes(), 0644)
	if err2 != nil {
		panic(err2)
	}
}


func Write_Keyword_File(path string) {
		write_JSON_File(path, &ds.KeywordList)
}


func exists(path string) (bool) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}



func Copy_Attachments(path, year, month string) string {

	var fileName string

	regxName := regexp.MustCompile(`\w+\.\w+`)
	fileName = regxName.FindString(path)

	attachmentsPath := (ds.KConfig.LogPath + "/" + year + "/" + month + "/attachments") 

	// Make sure attachments directory exists
	if ! exists(attachmentsPath) {
		create_log_dir(attachmentsPath)
	}


	//Copy file?
	copy_File_Contents(path, (attachmentsPath+"/"+fileName))
	
	return fileName
	
} 





func copy_File_Contents(src, dst string) (err error) {
    in, err := os.Open(src)
    if err != nil {
        return
    }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil {
        return
    }
    defer func() {
        cerr := out.Close()
        if err == nil {
            err = cerr
        }
    }()
    if _, err = io.Copy(out, in); err != nil {
        return
    }
    err = out.Sync()
    return
}




func create_log_dir(path string)  {

	err := os.Mkdir(path, 0777)
	if err != nil {
		panic(err)
	}
}

func create_log_file(path string)  {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	file.Close()
}

func removeFile(path string) {
	os.Remove(path)
}

