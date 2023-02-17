package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var c = &http.Client{}
var sid string
var list = []string{}
var pric = []float64{}

var goods struct {
	Data struct {
		Items []struct {
			BuyMaxPrice    string `json:"buy_max_price"`
			BuyNum         int    `json:"buy_num"`
			ID             int    `json:"id"`
			MarketHashName string `json:"market_hash_name"`
			SellMinPrice   string `json:"sell_min_price"`
			SellNum        int    `json:"sell_num"`
		} `json:"items"`
	} `json:"data"`
}

var price struct {
	Data struct {
		GoodsInfos struct {
			Num struct {
				MarketHashName string `json:"market_hash_name"`
			}
		} `json:"goods_infos"`
		Items []struct {
			Price string `json:"price"`
		} `json:"items"`
	} `json:"data"`
}

func auth() {
	sid = f_reader("reading /sessionid.txt/ ...")
	url := "https://buff.163.com/"
	r, _ := http.NewRequest("GET", url, nil)

	r.Header.Add("Cookie", "session="+sid)

	res, _ := c.Do(r)
	defer res.Body.Close()

	cookies := res.Header.Values("Set-Cookie")
	if strings.Contains(fmt.Sprint(fmt.Sprint(cookies)), "session=;") == false {
		fmt.Println("authenticated")
	} else {
		fmt.Println("not authenticated")
	}
}

func modules() {

	fmt.Println("popeye - 0.1")

	fmt.Println("[1] flipping")
	fmt.Println("[2] sniping")

	i, _ := strconv.Atoi(i_reader("enter_ "))

	switch i {
	case 1:
		search(1)
	case 2:
		search(2)
	default:
		os.Exit(1)
	}

}

func search(in int) {

	url := "https://buff.163.com/api/market/goods"
	r, _ := http.NewRequest("GET", url, nil)

	r.Header.Add("Cookie", "session="+sid)

	q := r.URL.Query()
	q.Add("game", "csgo")
	q.Add("page_num", "1")
	r.URL.RawQuery = q.Encode()

	for {
		res, _ := c.Do(r)
		defer res.Body.Close()

		switch in {
		case 1:
			flipping(res.Body)
		case 2:
			sniping(res.Body)
		default:
			os.Exit(1)
		}

		time.Sleep(1 * time.Second)
	}
}

func flipping(ri io.ReadCloser) {
	var sellorders = 100
	var minprice float64 = 100

	var _ = json.NewDecoder(ri).Decode(&goods)

	var sell float64
	var buy float64

	for i, _ := range goods.Data.Items {
		sell, _ = strconv.ParseFloat(goods.Data.Items[i].SellMinPrice, 64)
		buy, _ = strconv.ParseFloat(goods.Data.Items[i].BuyMaxPrice, 64)
		if goods.Data.Items[i].SellNum > sellorders && (sell*0.975-buy)/buy*100 > 5 && sell > minprice {

			tmp := fmt.Sprint(list)
			if strings.Contains(tmp, goods.Data.Items[i].MarketHashName) == false {
				list = append(list, goods.Data.Items[i].MarketHashName)
				fmt.Println(goods.Data.Items[i].MarketHashName)
			}
		}
	}
}

func sniping(rin io.ReadCloser) {
	var _ = json.NewDecoder(rin).Decode(&goods)
	for i, _ := range goods.Data.Items {
		fmt.Println(strconv.Itoa(goods.Data.Items[i].ID))

		item(strconv.Itoa(goods.Data.Items[i].ID))
	}
}

func item(id string) {
	url := "https://buff.163.com/api/market/goods/sell_order"
	r, _ := http.NewRequest("GET", url, nil)

	r.Header.Add("Cookie", "session="+sid)

	q := r.URL.Query()
	q.Add("game", "csgo")
	q.Add("goods_id", id)
	q.Add("page_num", "1")
	r.URL.RawQuery = q.Encode()

	res, _ := c.Do(r)
	defer res.Body.Close()

	var _ = json.NewDecoder(res.Body).Decode(&price)

}
