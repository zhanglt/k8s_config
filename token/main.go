package main // import "github.com/k8s/token" 
import (
    "fmt"
    "os/exec"
    "net/http"
    "html/template"
)
func tokenHandle(w http.ResponseWriter, r *http.Request){

	 t := template.Must(template.ParseFiles("token.html"))

         token :=  getToken()
         t.Execute(w, token)
       // fmt.Fprintln(w,getToken())

}
     

func main(){
     
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
    http.HandleFunc("/token",tokenHandle)
    http.ListenAndServe(":9999",nil)  
}


func getToken()(token string){ 
    command := `./getToken.sh .`
    cmd := exec.Command("/bin/bash", "-c", command)
    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
        return
    }
    //fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
   return string(output)
}
