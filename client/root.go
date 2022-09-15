package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
)

var errStatusCodeInvalid = errors.New("status code invalid")

type request struct {
  method string
  url string
  // body maybe later
  times int
  headers []string
}

func (r request) newRequestWithCtx(ctx context.Context) (*http.Request, error) {
  req, err := http.NewRequestWithContext(ctx, r.method, r.url, nil)
  if err != nil {
    return nil, err
  }

  for _, h := range r.headers {
    if hspl := strings.Split(h, ":"); len(hspl) == 2 {
      req.Header.Add(hspl[0], hspl[1])
    }
  }

  return req, nil
}

func exec(req *http.Request, times int) int32 {
  var (
    wg sync.WaitGroup

    success int32
  )

  for i := 0; i < times; i++ {
    wg.Add(1)
    go func(i int) {
      defer wg.Done()
      if err := fetch(req); err == nil {
        atomic.AddInt32(&success, 1)
      }
    }(i)
  }

  wg.Wait()
  return success
}

func fetch(req *http.Request) error {
  var (
    res *http.Response
    err error
  )

  defer func() {
    if res != nil {
      log.Println(req.Method, " ", req.URL.String(), " ", res.Status)
    }
  }()

  if res, err = (&http.Client{}).Do(req); err != nil {
    return err
  } else if res.StatusCode >= 400 {
    return errStatusCodeInvalid
  }

  return nil
}
