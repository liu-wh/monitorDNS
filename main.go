package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	url            string
	dns            []net.IP
	interval       int
	intervalTime   time.Duration
	changeInterval int
	temp           net.IP
	consolePrint   bool
	dnsSlice       []string
	count          int
	Max            int
	Min            int
	Avg            int
	changeSlice    []int
)

func removeDuplicateElement(slice []string) []string {
	result := make([]string, 0, len(slice))
	temp := map[string]struct{}{}
	for _, item := range slice {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func main() {
	flag.IntVar(&interval, "i", 10, "dns解析地址, 时间间隔, 单位秒")
	flag.StringVar(&url, "d", "", "要解析的url")
	flag.BoolVar(&consolePrint, "p", false, "是否打印在标准输出")
	flag.Parse()
	if url == "" {
		flag.Usage()
		os.Exit(1)
	}
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("没有权限创建日志文件")
		os.Exit(1)
	}
	startTime := time.Now()
	intervalTime = time.Duration(interval) * time.Second
	var lastDns []string = nil
	var ipSlice []string = nil
	log.SetOutput(file)
	for {
		dns, err = net.LookupIP(url)
		if err != nil {
			log.Printf("解析错误: %v", err)
			continue
		}
		now := time.Now().Format("2006-01-02 15:04:05")

		//把解析的ipv4 地址构成数组
		dnsSlice = []string{}
		for _, i := range dns {
			if temp = i.To4(); temp == nil {
				continue
			}
			dnsSlice = append(dnsSlice, i.String())
			ipSlice = append(ipSlice, dnsSlice...)
		}

		if lastDns == nil {
			lastDns = dnsSlice
			if consolePrint {
				fmt.Printf("%s %s的解析是: %v\n", now, url, dnsSlice)
			}
			changeInterval = 0
			log.Printf("%s的解析是: %v\n", url, dnsSlice)
		} else {
			var flag bool = false
			lastDnsLen := len(lastDns)
			dnsSliceLen := len(dnsSlice)
			if lastDnsLen != dnsSliceLen {
				flag = true
			} else {
				tNum := 0
				for _, i := range lastDns {
					for _, j := range dnsSlice {
						if i == j {
							tNum += 1
							break
						}
					}
				}
				if tNum != lastDnsLen {
					flag = true
				}

			}
			if flag {
				ipSlice = append(ipSlice, dnsSlice...)
				ipSlice = removeDuplicateElement(ipSlice)
				count += 1
				if Min == 0 {
					Min = changeInterval
				}
				if changeInterval < Min {
					Min = changeInterval
				}
				if changeInterval > Max {
					Max = changeInterval
				}
				changeSlice = append(changeSlice, changeInterval)
				x := 0
				for _, i := range changeSlice {
					x += i
				}
				Avg = x / len(changeSlice)

				if consolePrint {
					fmt.Printf("%s %s解析发生了变化: %v -> %v, 间隔了%d秒\n", now, url, lastDns, dnsSlice, changeInterval)
					fmt.Printf("%s 当前已经变化过的ip地址有%d个, 列表: %v\n", now, len(ipSlice), ipSlice)
					fmt.Printf("%s 已经监控了%d秒, 解析变化了%d次, 解析变化间隔列表:%v, 解析变化最长间隔%d秒, 最短%d秒, 平均%d秒\n", now, int(time.Since(startTime).Seconds()), count, changeSlice, Max, Min, Avg)
				}

				log.Printf("%s解析发生了变化: %v -> %v, 间隔了%d秒\n", url, lastDns, dnsSlice, changeInterval)
				log.Printf("当前已经变化过的ip地址有%d个, 列表: %v", len(ipSlice), ipSlice)
				log.Printf("已经监控了%d秒, 解析变化了%d次, 解析变化间隔列表:%v, 解析变化最长间隔%d秒, 最短%d秒, 平均%d秒\n", int(time.Since(startTime).Seconds()), count, changeSlice, Max, Min, Avg)
				lastDns = dnsSlice
				changeInterval = 0
			} else {
				if consolePrint {
					fmt.Printf("%s %s的解析是: %v\n", now, url, dnsSlice)
				}
				log.Printf("%s的解析是: %v\n", url, dnsSlice)
			}
		}

		time.Sleep(intervalTime)
		changeInterval += interval
	}
}
