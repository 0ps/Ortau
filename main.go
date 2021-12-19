package main

import (
	"Ortau/conf"
	. "Ortau/reverseproxy"
	"Ortau/static"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type ReverseProxy struct {
	RedirectUrl string
	Ua          string
}

type Cfx struct {
	Host        string
	Port        string
	RedirectUrl string
}

func (p *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//log.Println("[Info] UserAgent: ", r.UserAgent())
	remote, err := url.Parse(p.RedirectUrl)
	if err != nil {
		log.Println("[Error] Error is : ", err)
	}

	c2remote, err := url.Parse(conf.GetCfgSectionKey("default", "c2Url"))
	log.Println(c2remote)

	//log.Println(conf.GetCfgSectionKey("filter", "uaKey"))
	//根据uaKey判断转发
	if strings.Contains(r.UserAgent(), conf.GetCfgSectionKey("filter", "uaKey")) == true {
		log.Println("[Info] Try to redirect...")
		proxy := NewProxy(c2remote)
		proxy.ServeHTTP(w, r)
	} else {
		proxy := NewProxy(remote)
		proxy.ServeHTTP(w, r)
	}

}

func Runing() {
	fmt.Printf("\033[1;31;40m%s\033[0m", static.Banner)
	fmt.Println("\n\n")

	for {
		conf.MakeCfg()
		if conf.MakeCfg() == true {
			break
		}
	}
	ua := conf.GetCfgSectionKey("filter", "uaKey")
	if ua == "" {
		log.Println("[Error] UaKey is null,Please check config.ini ...")
		os.Exit(1)
	}

	cfx := &Cfx{
		Host:        conf.GetCfgSectionKey("default", "host"),
		Port:        conf.GetCfgSectionKey("default", "port"),
		RedirectUrl: conf.GetCfgSectionKey("default", "redirectUrl"),
	}

	localIpAddress := cfx.Host + ":" + cfx.Port

	proxyHandle := &ReverseProxy{RedirectUrl: cfx.RedirectUrl}
	log.Printf("[Info] Proxy Addr: %v, RedirectUrl: %v\n", localIpAddress, proxyHandle)
	err := http.ListenAndServe(localIpAddress, proxyHandle)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}

}

func main() {
	Runing()
}
