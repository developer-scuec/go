package main

import (
	"encoding/json"
	"fmt"
	"github.com/prestonTao/upnp"
	"io/ioutil"
	"net/http"
	"net/url"
)
var mapping = new(upnp.Upnp)
func init() {
}

//获得公网ip地址
func ExternalIPAddr() (GatewayOutsideIP string){
	err := mapping.ExternalIPAddr()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("外网ip地址为：", mapping.GatewayOutsideIP)
		fmt.Println("本机ip地址为：",mapping.LocalHost)
	}
	AddPortMapping(9090,9090)
	return mapping.LocalHost+":9090"
}

/*
	添加一个端口映射
*/
func AddPortMapping(localPort, remotePort int) bool {
	//添加一个端口映射
	if err := mapping.AddPortMapping(localPort, remotePort, "TCP"); err == nil {
		fmt.Println("端口映射成功")
		return true
	} else {
		fmt.Println("端口映射失败",err)
		return false
	}
}
func ConnectPost(){
	db:=GetDBLink()
	//PublicKey,_:=db.Get([]byte("publicKey"),nil)
	WalletAddress,_:=db.Get([]byte("address"),nil)
	IpAddress:=ExternalIPAddr()
	//resp,err:=http.PostForm("http://127.0.0.1:8080/Block/connect",url.Values{"WalletAddress":{string(WalletAddress)},"IpAddress":{IpAddress},"PublicKey":{string(PublicKey)}})
	resp,err:=http.PostForm("http://127.0.0.1:8080/Block/connect",url.Values{"WalletAddress":{string(WalletAddress)},"IpAddress":{IpAddress}})

	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	str:=string(body)
	respMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(str), &respMap)
	if err != nil {
		fmt.Println(err)
	}
}

