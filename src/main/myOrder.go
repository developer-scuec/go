package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

func myOrderHandle(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	err:=r.ParseForm()
	if err!=nil{
		panic(err)
	}
	info:=r.Form
	id:=info["shopInfo[id]"]
	name:=info["shopInfo[name]"]
	price:=info["shopInfo[price]"]
	count:=info["shopInfo[count]"]
	p,_:=strconv.Atoi(info["shopInfo[price]"][0])
	c,_:=strconv.Atoi(info["shopInfo[count]"][0])
	allMoney:=strconv.Itoa(p*c)
	payToAddress:=info["payToAddress"]
	db:=GetDBLink()
	//将商品数据存储到数据库
	db.Put([]byte("id"),[]byte(id[0]),nil)
	db.Put([]byte("name"),[]byte(name[0]),nil)
	db.Put([]byte("price"),[]byte(price[0]),nil)
	db.Put([]byte("count"),[]byte(count[0]),nil)
	db.Put([]byte("allPrice"),[]byte(allMoney),nil)
	db.Put([]byte("payTo"),[]byte(payToAddress[0]),nil)
	db.Put([]byte("isSend"),[]byte("确认发送！"),nil)

}
func BroadcastHandle(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	db:=GetDBLink()
	id,_:=db.Get([]byte("id"),nil)
	name,_:=db.Get([]byte("name"),nil)
	price,_:=db.Get([]byte("price"),nil)
	count,_:=db.Get([]byte("count"),nil)
	payToAddress,_:=db.Get([]byte("payTo"),nil)
	allPrice,_:=db.Get([]byte("allPrice"),nil)
	m:=map[string]string{
		"id":string(id),
		"name":string(name),
		"price":string(price),
		"count":string(count),

	}
	result, _ := json.MarshalIndent(m, "", "")
	shopInfoMap:=string(result)
	StrPrk, _ := db.Get([]byte("privateKey"), nil)
	//StrPuk, _ := db.Get([]byte("publicKey"), nil)
	StrAddr, _ := db.Get([]byte("address"), nil)
	block := map[string]string{
		"fromAddress":string(StrAddr),
		"toAddress":string(payToAddress),
		"shopInfo":shopInfoMap,
		"money":string(allPrice),

	}
	blockString, _ := json.MarshalIndent(block, "", "")
	blockMap:=string(blockString)
	key:=UnStrKey(string(StrPrk))
	sign:=ToSign(blockMap,key)
	//将签名信息和商品信息广播
	block=map[string]string{
		"block":blockMap,
		"sign":sign,
	}
	blockString, _ = json.MarshalIndent(block, "", "")
	blockMap=string(blockString)
	broadcast(blockMap)
	w.Write([]byte("success"));
}
func broadcast(block string)  {
	resp,err:=http.PostForm("http://127.0.0.1:8080/Block/broadcast",url.Values{"blockInfo":{block}})
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()
}
func getOrderHandle(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	db:=GetDBLink()
	id,_:=db.Get([]byte("id"),nil)
	name,_:=db.Get([]byte("name"),nil)
	price,_:=db.Get([]byte("price"),nil)
	count,_:=db.Get([]byte("count"),nil)
	payToAddress,_:=db.Get([]byte("payTo"),nil)
	allPrice,_:=db.Get([]byte("allPrice"),nil)
	isSend,_:=db.Get([]byte("isSend"),nil)
	m:=map[string]string{
		"id":string(id),
		"name":string(name),
		"price":string(price),
		"count":string(count),
		"payToAddress":string(payToAddress),
		"allPrice":string(allPrice),
		"isSend":string(isSend),
	}
	db.Put([]byte("isSend"),[]byte("已确认"),nil)
	blockString, _ := json.MarshalIndent(m, "", "")
	w.Header().Set("Content-Type", "application/json")
	w.Write(blockString)
}