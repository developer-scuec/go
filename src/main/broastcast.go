package main

import (
	"BlockChain"
	"fmt"
	"net/http"
)

func BroastCastHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	block := r.Form["block"][0]
	sign := r.Form["sign"][0]
	fmt.Println(block)
	fmt.Println(sign)
	////创建用于存储数据的数据库
	db := GetDBLink()
	StrPrk, err := db.Get([]byte("privateKey"), nil)
	if err != nil {
		fmt.Println("err:", err)
	}
	Prk := UnStrKey(string(StrPrk))
	isTrue, error := Verify(block, sign, Prk.PublicKey)
	if error != nil {
		panic(error)
	}
	fmt.Println("验证签名：", isTrue)
	if isTrue{
		BlockChain.AddNewNode(block)
	}
	//fmt.Println("block:",block)
	//fmt.Println("sign:",sign)
}
