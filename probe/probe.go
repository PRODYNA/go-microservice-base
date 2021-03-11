package probe

import "net/http"

type ProbeService struct {
	Start []func()(string,bool)
	Live  []func()(string,bool)
	Ready []func()(string,bool)
}

func NewProbeService() *ProbeService {
	p := &ProbeService{
		Start: make([]func()(string,bool), 0),
		Ready: make([]func()(string,bool), 0),
		Live:  make([]func()(string,bool), 0)}

	return p
}

func (p *ProbeService) AddStart(probe func()(string,bool)) {
	p.Start = append(p.Start, probe)
}

func (p *ProbeService) AddLive(probe func()(string,bool)) {
	p.Live = append(p.Live, probe)
}

func (p *ProbeService) AddReady(probe func()(string,bool)) {
	p.Ready = append(p.Ready, probe)
}


func (p ProbeService) verifyStart() bool {
	var result bool = true

	for _,p := range p.Start {
		_,ok := p()
		result = result && ok
	}

	return result
}

func (p ProbeService) verifyReady() bool {
	var result bool = true

	for _,p := range p.Ready {
		_,ok := p()
		result = result && ok
	}

	return result
}

func (p ProbeService) verifyLive() bool {
	var result bool = true

	for _,p := range p.Live {
		_,ok := p()
		result = result && ok
	}

	return result
}

func (p *ProbeService) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var result bool = false

	if r.URL.Path == "/start" && r.Method == "GET" {
		result = p.verifyStart()
	} else if r.URL.Path == "/ready" && r.Method == "GET" {
		result = p.verifyReady()
	} else if r.URL.Path == "/live" && r.Method == "GET" {
		result = p.verifyLive()
	} else {
		w.WriteHeader(404)
		w.Write([]byte("{ error: \"Not Found\" }"))
	}

	if result {
		w.WriteHeader(200)
		w.Write([]byte("{ ok: true }"))
	} else  {
		w.WriteHeader(500)
		w.Write([]byte("{ ok: false }"))
	}
}
