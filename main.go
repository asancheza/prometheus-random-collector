package main

import (
	"os"
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	listenAddress = flag.String("web.listen", ":9118", "Address on which to expose metrics and web interface.")
	metricsPath   = flag.String("web.path", "/metrics", "Path under which to expose metrics.")
	namespace     = flag.String("namespace", "random", "Namespace for the Random metrics.")
)

func main() {
	flag.Parse()

	prometheus.MustRegister(NewExporter())
	prometheus.Unregister(prometheus.NewGoCollector())
	prometheus.Unregister(prometheus.NewProcessCollector(os.Getpid(), ""))

	log.Printf("Starting Server: %s", *listenAddress)
	handler := prometheus.UninstrumentedHandler()
	if *metricsPath == "" || *metricsPath == "/" {
		http.Handle(*metricsPath, handler)
	} else {
		http.Handle(*metricsPath, handler)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`<html>
			<head><title>Random Exporter</title></head>
			<body>
			<h1>Random Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
		})
	}

	_ = http.ListenAndServe(*listenAddress, nil)
	
}