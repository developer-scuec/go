package main

import (
	"encoding/json"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"net/http"
)

type Account struct {
	StrPuk string `json:"publicKey"`
	StrPrk string `json:"privateKey"`
	StrAdr string `json:"address"`
}
var DbLink *leveldb.DB
func LinkDB(){
	//创建用于存储数据的数据库
	DbLink,_= leveldb.OpenFile("BlockChainDb", nil)
}
func GetDBLink()  *leveldb.DB{
	return DbLink
}
func NewAccount(w http.ResponseWriter, r *http.Request) {
	StrPrk, _ := DbLink.Get([]byte("privateKey"), nil)
	StrPuk, _ := DbLink.Get([]byte("publicKey"), nil)
	StrAddr, _ := DbLink.Get([]byte("address"), nil)
	if StrPrk != nil && StrPuk != nil && StrAddr != nil {
		newCount := Account{
			StrPuk: string(StrPuk),
			StrPrk: string(StrPrk),
			StrAdr: string(StrAddr),
		}
		jsonBytes, err := json.Marshal(newCount)
		if err != nil {
			panic(err)
		} else {

			w.Write(jsonBytes)
		}
		//fmt.Println("db:" + string(jsonBytes))

	} else {
		Randstring := RandStringBytes(40)
		privateKey, err := MakeNewKey(Randstring)
		if err != nil {
			fmt.Println(err)
		}
		StrPrk, StrPuk := StrKey(privateKey)
		StrAddr := GetAddress(privateKey.PublicKey)
		newCount := Account{
			StrPuk: StrPuk,
			StrPrk: StrPrk,
			StrAdr: StrAddr,
		}

		DbLink.Put([]byte("privateKey"), []byte(StrPrk), nil)
		DbLink.Put([]byte("publicKey"), []byte(StrPuk), nil)
		DbLink.Put([]byte("address"), []byte(StrAddr), nil)
		jsonBytes, err := json.Marshal(newCount)
		//fmt.Println("new:" + string(jsonBytes))
		if err != nil {
			w.Write(jsonBytes)
		}

	}

}
