package cron

import (
	crawjud "github.com/anti-duhring/crawjud/pkg/crawler"
	"github.com/anti-duhring/goncurrency/pkg/logger"
)

type CRAWLER_ID string

var crawlers = map[CRAWLER_ID]func() (map[string]string, error){
	"TJPE_MOVIMENTACAO": crawjud.TJPE,
}
