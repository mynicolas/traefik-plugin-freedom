package ip2geo

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Address       string `json:"address"`
	Elasticsearch string `json:"elasticsearch"`
	ESUsername    string `json:"esusername"`
	ESPassword    string `json:"espassword"`
}

func CreateConfig() *Config {
	return &Config{
		Address:       "",
		Elasticsearch: "",
		ESUsername:    "",
		ESPassword:    "",
	}
}

type Ip2Geo struct {
	next http.Handler
	ctx  context.Context
	name string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.Address == "" {
		return nil, fmt.Errorf("address 参数不能为空")
	}
	if config.Elasticsearch == "" {
		return nil, fmt.Errorf("elasticsearch 参数不能为空")
	}

	return &Ip2Geo{
		ctx:  ctx,
		next: next,
		name: name,
	}, nil
}

func (a *Ip2Geo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	xfor := req.Header["X-Forward-For"]
	host := req.Host
	ips := strings.Join(xfor, " ")
	os.Stdout.WriteString(host)
	os.Stdout.WriteString(ips)
	a.next.ServeHTTP(rw, req)
}

