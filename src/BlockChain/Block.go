package BlockChain

import (
	"Block"
	"encoding/json"
	"fmt"
	"strconv"
)

//通过链表的形式维护区块链的业务

type Node struct {
	//指针域
	NextNode *Node
	//数据域
	Data *Block.Block
}

var HeadNode *Node


//创建头节点，保存创世区块
func CreatHeadNode(data *Block.Block) *Node {

	var headNode *Node = new(Node)
	headNode.NextNode = nil

	headNode.Data = data

	HeadNode = headNode
	return headNode
}
//添加节点，当挖矿成功以后添加区块
func AddNode(data *Block.Block,prefNode *Node) *Node {

	var newNode *Node = new(Node)
	newNode.Data = data
	newNode.NextNode = nil
	prefNode.NextNode = newNode
	return newNode
}
//添加节点
func AddNewNode(data string){
	if HeadNode == nil{
		block:=Block.GenerateFirstBlock(data)
		HeadNode=CreatHeadNode(&block)
	}else {
		node := HeadNode
		//nodeData:=HeadNode.Data
		for {
			if node.NextNode!= nil {
				node = node.NextNode
				//nodeData=node.Data
			}else if node.NextNode== nil{
				preBlockPointer:=node.Data
				preBlock:=Block.Block{
					Index:preBlockPointer.Index,
					TimeStamp:preBlockPointer.TimeStamp,
					Diff:preBlockPointer.Diff,
					PreHash:preBlockPointer.PreHash,
					HashCode:preBlockPointer.HashCode,
					Nonce:preBlockPointer.Nonce,
					Data:preBlockPointer.Data,
				}
				newblock:=Block.GenerateNextBlock(data,preBlock)
				var newNode *Node=new(Node)
				newNode.Data=&newblock
				newNode.NextNode=nil
				node.NextNode=newNode
				break
			}
		}
	}
}
//遍历节点
func ShowBlockChain()  {
	var node *Node = HeadNode
	for {
		if node.NextNode != nil {
			fmt.Println(node.Data)
			node = node.NextNode
		}else if node.NextNode == nil{
			fmt.Println(node.Data)
			break
		}
	}
    GetBlockChain()
}
func GetBlockChain() []byte {
	BlockMap := make(map[string]string)
	node:=HeadNode
	for {
		if node.NextNode != nil {
			block:=Block.Block{
				//区块链高度
				Index:node.Data.Index,
				//时间戳
				TimeStamp:node.Data.TimeStamp,
				//难度
				Diff:node.Data.Diff,
				//上一个区块哈希
				PreHash:node.Data.PreHash,
				//当前区块哈希
				HashCode:node.Data.HashCode,
				//随机数
				Nonce:node.Data.Nonce,
				//交易信息
				Data:node.Data.Data,
			}
			jsons, _:= json.Marshal(block)
			BlockMap[strconv.Itoa(node.Data.Index)]=string(jsons)
			node = node.NextNode
		}else if node.NextNode == nil{
			block:=Block.Block{
				//区块链高度
				Index:node.Data.Index,
				//时间戳
				TimeStamp:node.Data.TimeStamp,
				//难度
				Diff:node.Data.Diff,
				//上一个区块哈希
				PreHash:node.Data.PreHash,
				//当前区块哈希
				HashCode:node.Data.HashCode,
				//随机数
				Nonce:node.Data.Nonce,
				//交易信息
				Data:node.Data.Data,
			}
			jsons, _:= json.Marshal(block)
			BlockMap[strconv.Itoa(node.Data.Index)]=string(jsons)
			break
		}
	}
	blockJson,_:=json.Marshal(BlockMap)
	return blockJson
}

