package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"testamus/logger"
	"testamus/thief"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func main() {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", false),
	)
	nelctx, nelcancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer nelcancel()

	ctx, cancel := chromedp.NewContext(nelctx)
	defer cancel()

	if err := chromedp.Run(ctx,
		network.Enable(),

		//Настроить фильтрацию по url, методу, типу ресурса, стадии запроса (смотри WithPatterns)
		fetch.Enable().WithPatterns([]*fetch.RequestPattern{
			{
				URLPattern: "*/api/v2*",
				//RequestStage: fetch.RequestStageRequest,
				//ResourceType: network.ResourceTypeXHR,
			},
			// {
			// 	URLPattern:   "*moon.examus.net/api/v2/*",
			// 	RequestStage: fetch.RequestStageResponse,
			// 	//ResourceType: network.ResourceTypeXHR,
			// },
		}),
	); err != nil {
		log.Fatal("Error enabling fetch with patterns:", err)
	}

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *fetch.EventRequestPaused:
			go func() {
				var mutex sync.Mutex
				jsonRequest, err := json.Marshal(ev)
				var postData []byte
				for _, entry := range ev.Request.PostDataEntries {
					decoded, err := base64.StdEncoding.DecodeString(entry.Bytes)
					if err != nil {
						log.Fatal("Ошибка при декодировании данных:", err)
					}
					postData = append(postData, decoded...)
				}
				if err != nil {
					fmt.Println("Ошибка при маршалинге запроса:", err)
				}
				err = logger.Logging(jsonRequest, &mutex)
				if err != nil {
					fmt.Println("Ошибка при записи запроса:", err)
				}

				c := chromedp.FromContext(ctx)
				exec := cdp.WithExecutor(ctx, c.Target)

				if ev.Request.Method == "POST" {
					flag, err := thief.CheckJsonFilter(postData)
					if err != nil {
						fmt.Println("Ошибка при проверке запроса:", err)
					}

					if flag {
						fmt.Println("Есть детект списывания")
						fetch.FailRequest(ev.RequestID, network.ErrorReasonTimedOut).Do(exec)
					}
				}
				err = fetch.ContinueRequest(ev.RequestID).Do(exec)
				if err != nil {
					fmt.Println("Ошибка при продолжении запроса:", err)
				}
			}()
		}
	})

	if err := chromedp.Run(ctx,
		chromedp.Tasks{
			chromedp.Navigate("https://lms.demo.examus.net/demo_examus/"),
			chromedp.Sleep(2 * time.Second),
			chromedp.Click(".start-demo", chromedp.NodeVisible),
			chromedp.Sleep(3 * time.Second),
			chromedp.Evaluate(`console.clear = () => {}`, nil),
		},
	); err != nil {
		log.Fatal("Появление ошибки: ", err)
	}

	//Ожидание завершения программы
	select {}
}
