package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"encoding/json"
	//"github.com/gin-gonic/gin"
)

const DEFALUT_PORT = 10805

func Version(w http.ResponseWriter, r *http.Request) {
	desPath := r.URL.Path[1:len(r.URL.Path)]
	fileData, err := ioutil.ReadFile(desPath)
	if err != nil {
		log.Println("Read File ERR:", err.Error())
	} else {
		log.Println("Send File", desPath)
		w.Write(fileData)
	}
}

func handleGetFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("handleGetFile:", r.RemoteAddr)
	desPath := r.URL.Path[1:len(r.URL.Path)]
	fmt.Println(desPath)
	desStat, err := os.Stat(desPath)
	if err != nil {
		log.Println("FILE Not Exit", desPath)
		http.NotFoundHandler().ServeHTTP(w, r)
	} else if desStat.IsDir() {
		log.Println("FIle is Dir", desPath)
		http.NotFoundHandler().ServeHTTP(w, r)
	} else {
		fileData, err := ioutil.ReadFile(desPath)
		if err != nil {
			log.Println("Read File ERR:", err.Error())
		} else {
			log.Println("Send File", desPath)
			w.Write(fileData)
		}
	}
}

func handleGetExecute(w http.ResponseWriter, r *http.Request) {
	log.Println("handleGetExecute:", r.RemoteAddr)
	desPath := r.URL.Path[1:len(r.URL.Path)]
	fmt.Println(desPath)
	desStat, err := os.Stat(desPath)
	if err != nil {
		log.Println("FILE Not Exit", desPath)
		http.NotFoundHandler().ServeHTTP(w, r)
	} else if desStat.IsDir() {
		log.Println("FIle is Dir", desPath)
		http.NotFoundHandler().ServeHTTP(w, r)
	} else {
		fileData, err := ioutil.ReadFile(desPath)
		if err != nil {
			log.Println("Read File ERR:", err.Error())
		} else {
			log.Println("Send File", desPath)
			w.Write(fileData)
		}
	}
}

func GetData(w http.ResponseWriter, r *http.Request) {
	desPath := r.URL.Path[1:len(r.URL.Path)]
	if strings.Contains(desPath, "api/") {
		fmt.Println("api") // 下发版本
		
		w.Write()
	}
	if strings.Contains(desPath, "dist") {
		fmt.Println("dist") // 下发文件
		w.Write()
	}
}

func main() {
	port := flag.Int("p", DEFALUT_PORT, "Set The Http Port")
	flag.Parse()

	http.HandleFunc("/", GetData)
	//http.HandleFunc("/dist", handleGetExecute)
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if nil != err {
		log.Fatal("Get Dir err", err.Error())
	}
}
