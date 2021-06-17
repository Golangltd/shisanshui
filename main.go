package main

import (
	. "G13s/shi_san_shui"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type pokerInfo struct {
	Pokers    string
	Desc      string
	IsTest    bool
	TestPoker []*Poker
}

func GenTestPokers(Pokers string) []*Poker {
	testPokers := make([]*Poker, 0)
	pokerStrList := strings.Split(Pokers, ",")
	for _, str := range pokerStrList {
		huaDesc := str[0:3]
		hua := DescHua(huaDesc).ToPokerHua()
		pointDesc := str[3:]
		point := DescPoint(pointDesc).ToPokerPoint()
		poker := NewPoker(point, hua)
		testPokers = append(testPokers, poker)
	}
	return testPokers
}

func main() {

	//读取json文件中的测试表
	jsonFile, err := os.Open("TestCase.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println(string(byteValue))

	var infos []pokerInfo

	json.Unmarshal(byteValue, &infos)
	for _, info := range infos {

		info.TestPoker = GenTestPokers(info.Pokers)
		if len(info.TestPoker) != 13 {
			log.Printf("json配牌出错，类型：%s\n", info.Desc)
			return
		}
		if info.IsTest {
			startTime := time.Now().Nanosecond()
			fmt.Printf("测试牌型=%+s,开始时间=%d\n", info.Desc, startTime)
			CalResult(info.TestPoker)
			endTime := time.Now().Nanosecond()
			costTime := endTime - startTime
			fmt.Printf("结束时间=%d,AI算法耗时【%d】微秒\n\n", endTime, costTime/1000)
		}
	}
}