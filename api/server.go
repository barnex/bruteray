package api

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"runtime"
	"sync"

	imagef "github.com/barnex/bruteray/image"
	"github.com/barnex/bruteray/sampler"
	//. "github.com/barnex/bruteray/tracer/types"
)

func Serve(addr string, spec Spec) error {
	s := newServer(addr, spec)
	return s.listenAndServe()
}

type server struct {
	spec       Spec
	mux        *http.ServeMux
	httpServer http.Server

	mu        sync.Mutex
	cancel    chan struct{}
	smplr     *sampler.Adaptive
	smplrView View // View currently being rendred by smplr
	smplrNum  int
}

func newServer(addr string, spec Spec) *server {
	mux := http.NewServeMux()
	s := &server{
		spec: spec,
		mux:  mux,
		httpServer: http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
	s.handle("/gob", s.handleGOB)
	s.handle("/cancel", s.handleCancel)
	return s

}

func (s *server) handleGOB(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	// (1) Read view settings from request body (JSON)
	var v View
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return err
	}

	// (2) Render (may return "canceled" error)
	img, err := s.renderView(v)
	if err != nil {
		return err
	}

	// (3) Send encoded image as response
	return gob.NewEncoder(w).Encode(img)
}

func (s *server) renderView(v View) (imagef.Image, error) {

	cancel, err := s.prepareBakery(v)
	if err != nil {
		return nil, err
	}
	defer func() { // wrong
		s.mu.Lock()
		defer s.mu.Unlock()
		s.cancel = nil
		s.smplrNum++
	}()

	nCPU := runtime.NumCPU()

	s.smplr.SampleNumCPUWithCancel(nCPU, s.smplrNum, cancel)
	img := s.smplr.StoredImage()

	if img == nil {
		return nil, errors.New("baking cancelled")
	}
	return img, nil
}

func (s *server) prepareBakery(v View) (chan struct{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cancel != nil {
		return nil, errors.New("already baking")
	}
	s.cancel = make(chan struct{})

	if s.smplr == nil || s.smplrView != v {
		spec := v.ApplyTo(s.spec)
		s.smplr = sampler.NewAdaptive(spec.ImageFunc(), v.Width, v.Height, v.AntiAlias)
		s.smplrView = v
		s.smplrNum = 1
	}

	return s.cancel, nil
}

func (s *server) handleCancel(w http.ResponseWriter, r *http.Request) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.cancel == nil {
		return errors.New("already canceled")
	}
	close(s.cancel)
	s.cancel = nil
	return nil
}

func (s *server) handle(prefix string, h handler) {
	s.mux.Handle(prefix, h)
}

func (s *server) listenAndServe() error {
	return s.httpServer.ListenAndServe()
}

type handler func(w http.ResponseWriter, r *http.Request) error

func (f handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	err := f(w, r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
