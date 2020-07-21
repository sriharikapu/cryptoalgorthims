package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type nodeInfo struct {
	id string
	path string
	writer http.ResponseWriter
}

var nodeTable = make(map[string]string)

func (node *nodeInfo) request(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm() 
	if len(request.Form["warTime"]) > 0 {
		node.writer = writer
		fmt.Println("", request.Form["warTime"][0])
		node.broadcast(request.Form["warTime"][0], "/prePrepare")
	}
}

func (node *nodeInfo) broadcast(msg string, path string) {
	fmt.Println("", path)
	for nodeId, url := range nodeTable {
		if nodeId == node.id {
			continue
		}
		http.Get("http://" + url + path + "?warTime=" + msg + "&nodeId=" + node.id)
	}
}

func (node *nodeInfo) prePrepare(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println("", request.Form["warTime"][0])
	if len(request.Form["warTime"]) > 0 {
		node.broadcast(request.Form["warTime"][0], "/prepare")
	}
}

func (node *nodeInfo) prepare(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println("", request.Form["warTime"][0])
	if len(request.Form["warTime"]) > 2/3*len(nodeTable) {
		node.authentication(request)
	}
}

var authenticationNodeMap = make(map[string]string)
var authenticationSuceess = false

func (node *nodeInfo) authentication(request *http.Request) {
	if !authenticationSuceess {
		if len(request.Form["nodeId"]) > 0 {
			authenticationNodeMap[request.Form["nodeId"][0]] = "OK"
			if len(authenticationNodeMap) > len(nodeTable)/3 {
				authenticationSuceess = true
				node.broadcast(request.Form["warTime"][0], "/commit")
			}
		}
	}
}

func (node *nodeInfo) commit(writer http.ResponseWriter, request *http.Request) {
	if writer != nil {
		fmt.Println("")
		io.WriteString(node.writer, "ok")
	}
}

func main() {
	userId := os.Args[1]
	fmt.Println(userId)
	nodeTable = map[string]string{
		"Apple":  "localhost:1111",
		"MS":     "localhost:1112",
		"Google": "localhost:1113",
		"IBM":    "localhost:1114",
	}

	node := nodeInfo{id: userId, path: nodeTable[userId]}

	http.HandleFunc("/req", node.request)
	http.HandleFunc("/prePrepare", node.prePrepare)
	http.HandleFunc("/prepare", node.prepare)
	http.HandleFunc("/commit", node.commit)
	if err := http.ListenAndServe(node.path, nil); err != nil {
		fmt.Println(err)
	}
}
