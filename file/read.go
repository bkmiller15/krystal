package file

import (
	"errors"
	"os"
	"io/ioutil"
	"../ds"
	"encoding/json"
	"path/filepath"
	"sort"
	"log"
	"os/exec"
)

// Reads in JSON file and places into struct
func read_JSON_File(filePath string, i interface{}) error {
	var data []byte
	var err error

	if data, err = ioutil.ReadFile(filePath); err != nil {
		return errors.New("(Search) error doing ioutil.ReadFile " + err.Error())
	//		return errors.New("(tocLayout) error setting view: " + err.Error())
	}

	if err = json.Unmarshal(data, i); err != nil {
		return errors.New("(Search) error doing ioutil.ReadFile " + err.Error())
		//return nil, errors.New("(Search) error doing ioutil.ReadFile " + err.Error())
	}

	return nil
}

func Read_Log_Files() int {
	var fileCount int
	var logPath string

	logPath = ds.KConfig.LogPath

	filepath.Walk(logPath, func(path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			loadFile(path)
			fileCount++
		}
		return nil
	})

	//Stats.FileLoadCount = fileCount
	sort.Sort(ds.AllEntries)

	return fileCount
}

func Read_Config_File() {
	var newConfig []ds.Config
	read_JSON_File("./data/config.json", &newConfig)
	ds.KConfig = newConfig[0]
	ds.KConfig.ShowDisabled = false
	ds.KConfig.Mode = 0
	ds.KConfig.TodosOnly = false
}

func Read_Keyword_File(path string) {
	var newKeywordList []ds.Words
	read_JSON_File(path, &newKeywordList)
	ds.KeywordList = newKeywordList
}




func loadFile(path string) {
	var year, month, day string
	var newEntries []ds.Entry

	read_JSON_File(path, &newEntries)

	year = path[len(path)-15:len(path)-11]
	month = path[len(path)-10:len(path)-8]
	day = path[len(path)-7:len(path)-5]

	for i := range newEntries {

		newEntries[i].Year = year
		newEntries[i].Month = month
		newEntries[i].Day = day

		//Fix Numbers if missing
		if newEntries[i].Number == 0 {
			newEntries[i].Number = ds.Get_Next_Number(year, month, day)
		}

		ds.AllEntries = append(ds.AllEntries, newEntries[i])

	}
	
}








func OpenAttachment(fileName, year, month string) {




	attachmentsPath := (ds.KConfig.LogPath + "/" + year + "/" + month + "/attachments/"+fileName) 

   	cmd := exec.Command("xdg-open", attachmentsPath)
    	err := cmd.Start()
   	if err != nil {
    		log.Fatal(err)
    	}
    	//log.Printf("Waiting for command to finish...")
    	//err = cmd.Wait()

}

