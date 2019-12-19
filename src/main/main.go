package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v7"
	"io/ioutil"
	"log"
	"models"
	"net/http"
	"os"
	"strings"
	"time"
)

var host = "http://172.17.12.47:9200"
var client *elastic.Client

func init() {
	client, err := elastic.NewClient(
		elastic.SetURL(host),
		//elastic.SetBasicAuth("user", "secret"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))

	if err != nil {
		panic(err)
	}

	fmt.Println("conn es succ", client)
}

func TermQuery(index, type_, fieldName, fieldValue string) *elastic.SearchResult {
	query := elastic.NewTermQuery(fieldName, fieldValue)
	//_ = elastic.NewQueryStringQuery(fieldValue) //关键字查询

	searchResult, err := client.Search().
		Index(index).Type(type_).
		Query(query).
		From(0).Size(10).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		panic(err)
	}
	fmt.Printf("query cost %d millisecond.\n", searchResult.TookInMillis)

	return searchResult
}

type NetEaseNews struct {
	Summary     string   `json:"summary"`
	PublishTime string   `json:"publishTime"`
	InfoId      string   `json:"infoId"`
	Link        string   `json:"link"`
	ShowType    string   `json:"showType"`
	Source      string   `json:"source"`
	Title       string   `json:"title"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	Covers      []string `json:"covers"`
	Quality     float64  `json:"quality"`
}

type NetEase struct {
	RequestId string        `json:"requestId"`
	Code      int           `json:"code"`
	Message   string        `json:"message"`
	Data      []NetEaseNews `json:"data"`
}

func main() {
	//getNetEase()
}

func getNetEase() {
	thirdSource := "wangyi"
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	client := &http.Client{}
	netEaseUrl := "https://yiyouliao.com/rss/common/xingqiuunion/list/json?num=200"
	request, err := http.NewRequest("GET", netEaseUrl, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Appkey", "3607f554e48446d19ca3281c175a81cc")

	//处理返回结果
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	var dataArr NetEase
	err = json.Unmarshal([]byte(string(body)), &dataArr)
	if err != nil {
		panic(err)
	}
	if dataArr.Code == 0 && dataArr.Message == "success" {
		news := make([]models.TmpNewsCrawalContent, len(dataArr.Data))
		for i, datum := range dataArr.Data {
			news[i].NewsId = GetNewsId(datum.InfoId, thirdSource)
			news[i].ThirdSource = thirdSource
			news[i].InfoId = datum.InfoId
			news[i].IsPure = 1
			news[i].ContentType = "news"
			news[i].PublishDate = time.Now()
			news[i].Title = datum.Title
			news[i].Summary = datum.Summary
			publishTime, _ := time.ParseInLocation(timeLayout, datum.PublishTime, loc)
			news[i].PublishDate = publishTime
			news[i].Source = datum.Source
			news[i].PureUrl = datum.Link
			news[i].NewsTag = strings.Join(datum.Tags, ",")
			news[i].Type = datum.Category
			news[i].CreateTime = time.Now()
			news[i].UpdateTime = time.Now()

			if len(datum.Covers) == 1 {
				news[i].BigPic = 1
				covers := map[string]string{"src": datum.Covers[0]}
				largePic, _ := json.Marshal(covers)
				news[i].LargePic = string(largePic)
				news[i].MiniPic = "[]"
				news[i].MiniImgSize = 0
			} else if len(datum.Covers) > 1 {
				news[i].BigPic = 0
				covers := make([]map[string]string, len(datum.Covers))
				for j, pic := range datum.Covers {
					covers[j] = map[string]string{"src": pic}
				}
				miniPic, _ := json.Marshal(covers)
				news[i].MiniPic = string(miniPic)
				news[i].LargePic = "[]"
				news[i].MiniImgSize = len(datum.Covers)
			}
		}
		rows, _ := InsertNews(news)
		fmt.Println(rows)
	}
}
