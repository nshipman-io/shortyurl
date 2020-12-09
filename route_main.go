package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nshipman-io/shortyurl/data"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/index.html","./templates/urlform.html","./templates/navbar.html", "./templates/footer.html")
	if err != nil {
		log.Println(err)
	}
	t.ExecuteTemplate(w, "index", "")
}

func New(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	url := data.ShortURL{}
	url.OriginalUrl = r.Form.Get("original_url")
	err = data.CreateUrl(&url)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintf(w,"Access the new url at /api/shorturl/%d", url.ID)
}

func Process(w http.ResponseWriter, r *http.Request){
	urlId,err := strconv.Atoi(mux.Vars(r)["url_id"])
	if err != nil {
		log.Println(err)
		return
	}
	shortUrl,err := data.GetUrl(urlId)
	if err != nil {
		log.Println(err)
		return
	}

	http.Redirect(w, r, shortUrl.OriginalUrl, 301)

}