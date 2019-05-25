package main

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var securityHeaders []securitySetting

func main() {

	var conf serverConfig
	parseJSONFile("configs/main.conf", &conf)

	parseJSONFile("configs/"+conf.SecurityConfig, &securityHeaders)

	mux := http.NewServeMux()

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		MaxVersion:               tls.VersionTLS13,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites:             getCiphers(),
	}

	srv := &http.Server{
		Addr:         "0.0.0.0:" + conf.Port,
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	mux.HandleFunc("/", rootHandler)

	log.Fatal(srv.ListenAndServeTLS("configs/"+conf.Cert, "configs/"+conf.Key))

}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	setSecurityHeaders(w)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Everything is fine"))
}

func getCiphers() []uint16 {
	return []uint16{
		tls.TLS_AES_128_GCM_SHA256,
		tls.TLS_AES_256_GCM_SHA384,
		tls.TLS_CHACHA20_POLY1305_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
	}
}

func setSecurityHeaders(w http.ResponseWriter) {
	for i := range securityHeaders {
		w.Header().Set(securityHeaders[i].Header, securityHeaders[i].Option)
	}
}

func parseJSONFile(file string, i interface{}) {
	// Import Configuration
	file = filepath.Clean(file)
	files, err := os.Open(file) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 10000)
	count, err := files.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(data[0:count]))

	err = json.Unmarshal(data[0:count], i)

	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
