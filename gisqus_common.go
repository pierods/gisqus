package gisqus

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const disqusDateFormat = "2006-01-02T15:04:05"
const disqusDateFormatExact = "2006-01-02T15:04:05.000000"

var zeroDate time.Time

func fromDisqusTime(dT string) (time.Time, error) {
	if dT == "" {
		return zeroDate, nil
	}
	return time.Parse(disqusDateFormat, dT)
}

func fromDisqusTimeExact(dT string) (time.Time, error) {
	if dT == "" {
		return zeroDate, nil
	}
	return time.Parse(disqusDateFormatExact, dT)
}

func call(ctx context.Context, url string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP Response Error %d", resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (g *Gisqus) callAndInflate(ctx context.Context, url string, v interface{}) error {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	drl, err := decodeRateLimits(resp.Header)
	if err != nil {
		return err
	}
	g.limits = drl

	if resp.StatusCode != 200 {
		disqusErrorBytes, _ := ioutil.ReadAll(resp.Body)
		disqusError := string(disqusErrorBytes)
		return fmt.Errorf("http response error %s, code: %d, Message: %s``", resp.Status, resp.StatusCode, disqusError)
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func decodeRateLimits(header http.Header) (DisqusRateLimit, error) {

	drl := DisqusRateLimit{}

	var err error
	remaining := header["X-Ratelimit-Remaining"]
	if len(remaining) > 0 {
		drl.RatelimitRemaining, err = strconv.Atoi(remaining[0])
		if err != nil {
			return drl, err
		}
	}
	limit := header["X-Ratelimit-Limit"]
	if len(limit) > 0 {
		drl.RatelimitLimit, err = strconv.Atoi(limit[0])
		if err != nil {
			return drl, err
		}
	}
	reset := header["X-Ratelimit-Reset"]
	if len(reset) > 0 {
		unixDate, err := strconv.ParseInt(header["X-Ratelimit-Reset"][0], 10, 32)
		if err != nil {
			return drl, err
		}
		drl.RatelimitReset = time.Unix(unixDate, 0)
	}

	return drl, nil
}
