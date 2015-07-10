package main

import (
	"fmt"
	"github.com/hashicorp/mdns"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var templates = template.Must(template.ParseFiles("views/index.html"))

type EnvVar struct {
	Name  string
	Value string
}

type Context struct {
	Hostname string
	Env      []EnvVar
	Peers    []string
}

var c Context
var port int

func renderIndex(wc web.C, w http.ResponseWriter, r *http.Request) {
	_ = templates.ExecuteTemplate(w, "index.html", &c)
}

func main() {
	c.Hostname, _ = os.Hostname()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		c.Env = append(c.Env, EnvVar{pair[0], pair[1]})
	}
	fmt.Println(c)
	static := web.New()
	static.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/static/", static)

	goji.Get("/", renderIndex)

	// Setup our service export
	host, _ := os.Hostname()
	info := []string{"simple peer discovery test"}
	service, _ := mdns.NewMDNSService(host, "_http._tcp", "", "", 8000, nil, info)

	// Create the mDNS server, defer shutdown
	fmt.Println("Creating mDNS server")
	server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	defer server.Shutdown()

	// Make a channel for results and start listening
	peers := make(chan *mdns.ServiceEntry, 4)
	go func() {
		fmt.Println("Waiting for peers")
		for entry := range peers {
            known := false
            for _, e := range c.Peers {
                if entry.Name == e {
                    known = true
                }
            }
            if known == false {
			    fmt.Println("Got new peer", entry.Name)
			    c.Peers = append(c.Peers, entry.Name)
            }
		}
	}()

	fmt.Println("Creating mDNS listener")
	//mdns.Lookup("_googlecast._tcp", peers)
    mdns.Lookup("_http._tcp", peers)
	fmt.Println("Creating HTTP server")
	goji.Serve() // port 8000 by default
	close(peers)
}
