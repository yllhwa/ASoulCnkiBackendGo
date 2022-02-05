package db

import (
	"GoBackend/pkg/utils"
	"bufio"
	"fmt"
	"io"
	"os"
)

var indexMap = make(map[int64]*[2]int64)

func LoadDbDes() {
	keyFile, _ := os.Open("database/test.des")
	keyFileReader := bufio.NewReaderSize(keyFile, 6*1024)
	defer keyFile.Close()
	fmt.Println("Loading db...")
	var lastKey int64 = -1
	keyBuffer := make([]byte, 6)
	for {
		_, err := keyFileReader.Read(keyBuffer)
		if err == io.EOF {
			var pos int64 = 0 //TODO
			indexMap[lastKey][1] = pos
			break
		}
		key := utils.BytesToInt(keyBuffer)

		keyFileReader.Read(keyBuffer)
		offset := utils.BytesToInt(keyBuffer)
		if lastKey != -1 {
			indexMap[lastKey][1] = offset - indexMap[lastKey][0]
		}
		indexMap[key] = &[2]int64{offset, -1}
		lastKey = key
		//fmt.Println(indexMap)
	}
	fmt.Println("Finished loading db.")
}

func GetHitsByHashs(hashs []int64) []int64 {
	var passages = make(map[int64]int)
	dbFile, _ := os.Open("database/test.db")
	//dbFileReader := bufio.NewReader(dbFile)
	//dbFile := os.OpenFile("database/test.db", syscall, 0666)
	defer dbFile.Close()
	for _, hash := range hashs {
		data, ok := indexMap[hash]
		if !ok {
			continue
		}
		offset := data[0]
		length := data[1]
		dbFile.Seek(offset, 0)
		cache := make([]byte, length)
		dbFile.Read(cache)
		var i int64 = 0
		for ; i < length; i += 6 {
			key := utils.BytesToInt(cache[i : i+6])
			passages[key]++
		}
	}
	passageList := utils.NewValSorter(passages)
	passageList.Sort()
	var ret []int64
	for _, key := range passageList.Keys {
		ret = append(ret, key)
	}
	if len(ret) >= 5 {
		ret = ret[:10]
	} else {
		return ret
	}
	return ret
}
