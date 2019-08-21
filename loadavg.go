package proc

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type LoadAvg struct {
	Last1Min  float64 `json:"last1min"`
	Last5Min  float64 `json:"last5min"`
	Last15Min float64 `json:"last15min"`
}

func (self *LoadAvg) Last1() float64 {
	return self.Last1Min
}

func (self *LoadAvg) Last5() float64 {
	return self.Last5Min
}

func (self *LoadAvg) Last15() float64 {
	return self.Last15Min
}

func ReadLoadAvg(path string) (*LoadAvg, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	content := string(b)
	fields := strings.Fields(content)
	loadavg := LoadAvg{}
	loadavg.Last1Min, _ = strconv.ParseFloat(fields[0], 64)
	loadavg.Last5Min, _ = strconv.ParseFloat(fields[1], 64)
	loadavg.Last15Min, _ = strconv.ParseFloat(fields[2], 64)
	return &loadavg, nil
}
