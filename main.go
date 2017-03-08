package main

import (
	"fmt"
	"net/http"
	"os"

	pushbullet "github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
)

func pushHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	token := os.Args[1]
	pb := pushbullet.New(token)
	n := requests.NewNote()
	n.Title = "Ping"
	n.Body = r.Form.Get("message")

	if _, err := pb.PostPushesNote(n); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		http.Redirect(w, r, "/failure", 301)
		return
	}

	http.Redirect(w, r, "/success", 301)
}
func main() {
	http.HandleFunc("/push", pushHandler)
	http.Handle("/", http.FileServer(assetFS()))
	http.ListenAndServe(":8080", nil)
}
