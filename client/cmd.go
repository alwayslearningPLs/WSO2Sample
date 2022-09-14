package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var req request

func newRootCmd() *cobra.Command {
  var rootCmd = &cobra.Command {
    Use: "client",
    Aliases: []string{"cl"},
    Short: "client used to fetch resources from the dogstore app using wso2 as API gateway to tests subscriptions stuff",
    Long: "client used to fetch resources from the dogstore app using wso2 as API gateway to tests subscriptions stuff",
    Run: func(cmd *cobra.Command, args []string) {
      ctx, cl := context.WithTimeout(context.TODO(), time.Minute)
      defer cl()

      r, err := req.newRequestWithCtx(ctx)
      if err != nil {
        panic(err)
      }

      var now = time.Now()
      defer func() {
        log.Println("time from the beginning till the end is: ", int(time.Since(now) / time.Second))
      }()

      success := exec(r, req.times)

      log.Printf("%d - %d = %d", req.times, success, req.times - int(success))
    },
  }

  rootCmd.PersistentFlags().StringVarP(&req.method, "method", "m", http.MethodGet, "method to execute the request")
  rootCmd.PersistentFlags().StringVarP(&req.url, "url", "u", "", "url to fetch")
  rootCmd.PersistentFlags().IntVarP(&req.times, "times", "t", 10, "amount of times is needed to execute the request")
  rootCmd.PersistentFlags().StringArrayVar(&req.headers, "header", []string{}, "headers to attach to the request")

  rootCmd.MarkPersistentFlagRequired("url") //nolint:errcheck

  //nolint:errcheck
  rootCmd.RegisterFlagCompletionFunc("method", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
    return []string{http.MethodGet, http.MethodDelete}, cobra.ShellCompDirectiveDefault
  })

  return rootCmd
}

