package main

func main(){
    go heartbeat.ListenHearteat()
    http.HandleFunc("/objects/", objects.Handler)
    http.HandleFunc("/locate/", locate.Handler)
    err := http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"). nil)
    if err != nil {
        log.Fatal(err)
    }
}
