package main

import (
	"net/http"
	"fmt"
	"time"
	"html/template"
	"github.com/machinebox/sdk-go/classificationbox"
	"context"
)

type PageVariables struct {
	Date         string
	Time         string
	Src 		 string
	Src2	string
}

func handleheloo(resp http.ResponseWriter, req *http.Request)  {
	fmt.Fprint(resp, "Pk was Here ....")
}

func getImageURL(str string)string  {
	cb := classificationbox.New("http://localhost:8080")
	cntx, _ := context.WithCancel(context.Background())
	arrFeatures := []classificationbox.Feature{}
	feature := classificationbox.Feature{"card",str,"image_url"}
	arrFeatures = append(arrFeatures, feature)
	predicationResponse := classificationbox.PredictRequest{Inputs: arrFeatures}
	resp,_ := cb.Predict(cntx,"CC", predicationResponse)
	if resp.Classes[0].Score > resp.Classes[1].Score{
		return "http://koditips.com/wp-content/uploads/uk-isp-blocked-kodi.jpg"
	}
	return str
}
func handleHome(resp http.ResponseWriter,req *http.Request) {

	now := time.Now()
	HomePageVars := PageVariables{
		
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
		//Src: getImageURL("https://www.enjoycompare.com/wp-content/uploads/2017/05/UOB-Yolo-Credit-Card-Large.gif"),
		//Src2:getImageURL("https://virunga.org/wp-content/uploads/2013/05/Kaboko.jpg"),
		Src: "https://www.enjoycompare.com/wp-content/uploads/2017/05/UOB-Yolo-Credit-Card-Large.gif",
		Src2: "https://virunga.org/wp-content/uploads/2013/05/Kaboko.jpg",
	}
	t, _ := template.ParseFiles("hompage.html")
	t.Execute(resp,HomePageVars)

}
func main(	) {

	http.HandleFunc("/heloo",handleheloo)
	http.HandleFunc("/",handleHome)
	http.ListenAndServe(":8090",nil)
}
