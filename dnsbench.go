package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func lookup(resolver string, domain string) string {
	var retstr string
	var stringValue string
	start := time.Now()
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, resolver)
		},
	}
	r.LookupHost(context.Background(), domain)
	elapsed := time.Since(start).Microseconds()
	stringValue = strconv.FormatInt(elapsed, 10)
	retstr = start.String()[:23] + "," + resolver + "," + domain + "," + stringValue
	println(retstr)
	return retstr
}

func main() {
	var list_of_domain [10]string
	var filename string
	now := time.Now()

	filename = "dnsbench-output-" + strings.ReplaceAll(strings.ReplaceAll(now.String()[:19], " ", "_"), ":", "_") + ".txt"

	fmt.Println(filename)

	list_of_domain[0] = "vnexpress.net"
	list_of_domain[1] = "google.com"
	list_of_domain[2] = "youtube.com"
	list_of_domain[3] = "facebook.com"
	list_of_domain[4] = "shopee.vn"
	list_of_domain[5] = "zalo.me"
	list_of_domain[6] = "24h.com.vn"
	list_of_domain[7] = "dantri.com.vn"
	list_of_domain[8] = "baomoi.com"
	list_of_domain[9] = "tiktok.com"

	var list_of_resolver [4]string
	list_of_resolver[0] = "116.97.90.124:53"
	list_of_resolver[1] = "116.97.90.125:53"
	list_of_resolver[2] = "203.113.131.2:53"
	list_of_resolver[3] = "203.113.188.6:53"

	file, fileErr := os.Create(filename)
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	for {
		for i := 0; i < len(list_of_domain); i++ {
			for j := 0; j < len(list_of_resolver); j++ {
				fmt.Fprintf(file, "%v\n", lookup(list_of_resolver[j], list_of_domain[i]))
			}
		}
	}
}
