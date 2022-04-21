package main
import(
	//"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World from GO \n"))
	})
    http.ListenAndServe(":80",nil)
}
