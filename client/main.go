package main

func main() {
  if err := newRootCmd().Execute(); err != nil {
    panic(err)
  }
}
