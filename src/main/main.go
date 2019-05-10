package main

import (
	"BlockChain"
	"net/http"
)

func main() {
	LinkDB()
	ConnectPost()
	server:=http.Server{
		Addr:"127.0.0.1:9090",
	}
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer((http.Dir("static")))))
	http.HandleFunc("/account",NewAccount)
	http.HandleFunc("/myorder",myOrderHandle)
	http.HandleFunc("/broastcast",BroastCastHandler)
	http.HandleFunc("/getmyoder",getOrderHandle)
	http.HandleFunc("/Broadcast",BroadcastHandle)
	http.HandleFunc("/GetBlockChain",BlockChain.GetBlockChainHandle)
	server.ListenAndServe()

}
