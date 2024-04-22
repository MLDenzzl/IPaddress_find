package main

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	dbPath = "D:\\Go_project\\IPAddrtest\\ip2region.xdb"
	ipBuff []byte
)

func init() {
	var err error
	ipBuff, err = xdb.LoadContentFromFile(dbPath)
	if err != nil {
		fmt.Printf("加载数据库数据失败 `%s`: %s\n", dbPath, err)
		return
	}
}

func main() {

	resp, err := http.Get("http://ifconfig.me/ip")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("IP地址:", string(body))

	searcher, err := xdb.NewWithBuffer(ipBuff)
	if err != nil {
		fmt.Printf("创建searcher失败: %s\n", err.Error())
		return
	}

	defer searcher.Close()

	//addrs, err := net.InterfaceAddrs()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//for _, address := range addrs {
	//	if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	//		if ipnet.IP.To4() != nil {
	//			fmt.Printf("当前内网IP地址：%s\n", ipnet.IP.String())
	//			os.Exit(0)
	//		}
	//	}
	//}
	//
	//fmt.Println("未找到内网IP地址")
	//os.Exit(1)

	var ip = "45.76.225.49"
	var startTime = time.Now()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		fmt.Printf("查询ip失败(%s): %s\n", ip, err)
		return
	}

	fmt.Printf("addr: %s, took: %s\n", region, time.Since(startTime))

}
