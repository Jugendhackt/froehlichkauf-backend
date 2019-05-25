package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
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
	mux.HandleFunc("/getProduct", productHandler)

	log.Fatal(srv.ListenAndServeTLS("configs/"+conf.Cert, "configs/"+conf.Key))

}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	setSecurityHeaders(w)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Everything is fine"))
}

func productHandler(w http.ResponseWriter, req *http.Request) {
	setSecurityHeaders(w)

	// decode incoming request (json)
	decoder := json.NewDecoder(req.Body)
	var code request
	err := decoder.Decode(&code)
	if err != nil {
		panic(err)
	}

	log.Printf("incoming request: %+v", code)

	var data datapoint

	parseJSONFile("data/"+code.Code+".json", &data)

	scoreUmwelt, scoreVerpackung, scoreHerkunft, errorUmwelt := umwelt(data.Packaging, code.Origin, data.Country)

	scoreEthik, errorEthik := ethik("")

	var nutritions []nutrition

	nutritions = append(nutritions, nutrition{"calories", data.Nutritional.Calories})
	nutritions = append(nutritions, nutrition{"glucides", data.Nutritional.Glucides})
	nutritions = append(nutritions, nutrition{"sugar", data.Nutritional.Sugar})
	nutritions = append(nutritions, nutrition{"lipides", data.Nutritional.Lipides})
	nutritions = append(nutritions, nutrition{"proteins", data.Nutritional.Proteins})
	nutritions = append(nutritions, nutrition{"salt", data.Nutritional.Salt})

	scoreHealth, scoreIngredients, scoreNutrition := gesundheit(nutritions, data.Contents)

	var resp response

	resp.Name = data.Name
	resp.Nutritional = data.Nutritional
	resp.Packaging = data.Packaging
	resp.Reusable = data.Reusable
	resp.Supplier = data.Supplier
	resp.Country = data.Country
	resp.Contents = data.Contents
	resp.Code = data.Code
	resp.Description = data.Description
	resp.ScoreEthik = scoreEthik
	resp.ScoreHealth = scoreHealth
	resp.ScoreHerkunft = scoreHerkunft
	resp.ScoreIngredients = scoreIngredients
	resp.ScoreNutrition = scoreNutrition
	resp.ScoreUmwelt = scoreUmwelt
	resp.ScoreVerpackung = scoreVerpackung
	resp.Errors = ""

	if errorUmwelt != nil {
		resp.Errors += ";Umwelt:" + errorUmwelt.Error()
	}
	if errorEthik != nil {
		resp.Errors += ";Ethik:" + errorEthik.Error()
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(resp)

	checkErr(err)
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

func fToString(f float32) string {
	ret := fmt.Sprintf("%g", f)
	return ret
}
