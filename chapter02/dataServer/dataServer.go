package main
import(
    "log"
    "net/http"
    "os"
    "cuimeng/chapter02/dataServer/locate"
    "cuimeng/chapter02/dataServer/heartbeat"
    "cuimeng/chapter02/dataServer/objects"
)
func main(){
    go heartbeat.ListenHeartbeat()
    go locate.StartLocate()
    http.HandleFunc("/objects/", objects.Handler)
    log.Fatal(http.ListenAndServe(os.Getevn("LISTEN_ADDRESS"), nil))
}

