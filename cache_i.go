package coingecko

import (
	"strconv"
	"os"
	"log"
	"io/fs"
	"time"
	"encoding/json"
)

func GetFromCache(id string, exp int64, out any) (found bool) {
	found, err := getFromCache(id, exp, out)
	if err != nil { log.Print(err) }
	return
}

func SaveToCache(id string, in any) {
	err := saveToCache(id, in)
	if err != nil { log.Print(err) }
}

func getFromCache(id string, exp int64, out any) (found bool, err error) {
	var tmpFile   string
	var file     *os.File
	var info      fs.FileInfo
	var now, then int64
	
	tmpFile = os.TempDir() + "/" + strconv.Itoa(os.Getuid()) + "-" + id + ".json"
	info, err = os.Stat(tmpFile)
	if err != nil && os.IsNotExist(err) {
		err = nil
		return
	}
	if err != nil { return }
	
	now = time.Now().Unix()
	then = info.ModTime().Unix()
	if ((now - then) > exp) {
		return
	}
	
	file, err = os.Open(tmpFile)
	if err != nil { return }
	defer file.Close()
	
	err = json.NewDecoder(file).Decode(out)
	if err != nil { return }
	
	found = true
	return
}

func saveToCache(id string, in any) (err error) {
	var tmpFile   string
	var file     *os.File
	
	tmpFile = os.TempDir() + "/" + strconv.Itoa(os.Getuid()) + "-" + id + ".json"
	
	file, err = os.Create(tmpFile)
	if err != nil { return err }
	defer file.Close()
	
	err = json.NewEncoder(file).Encode(in)
	if err != nil { return err }
	
	return nil
}

