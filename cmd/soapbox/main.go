package main
import ("fmt";"log";"net/http";"os";"github.com/stockyard-dev/stockyard-soapbox/internal/server";"github.com/stockyard-dev/stockyard-soapbox/internal/store")
func main(){port:=os.Getenv("PORT");if port==""{port="9700"};dataDir:=os.Getenv("DATA_DIR");if dataDir==""{dataDir="./soapbox-data"}
db,err:=store.Open(dataDir);if err!=nil{log.Fatalf("soapbox: %v",err)};defer db.Close();srv:=server.New(db,server.DefaultLimits())
fmt.Printf("\n  Soapbox — Self-hosted internal Q&amp;A and Stack Overflow\n  Dashboard:  http://localhost:%s/ui\n  API:        http://localhost:%s/api\n\n",port,port)
log.Printf("soapbox: listening on :%s",port);log.Fatal(http.ListenAndServe(":"+port,srv))}
