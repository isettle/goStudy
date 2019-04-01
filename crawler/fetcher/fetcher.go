package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var requestInterval = time.Tick(time.Millisecond * 10)

func Fetch(url string) ([]byte, error) {
	<-requestInterval
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0.3 Safari/605.1.15")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error status code:", resp.StatusCode)
		return nil, fmt.Errorf("wrong http statusCode %d", resp.StatusCode)
	}

	// 获取body
	return ioutil.ReadAll(resp.Body)
}
