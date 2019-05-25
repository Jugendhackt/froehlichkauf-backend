package main

type datapoint struct {
	Name        string      `json:"name"`
	Supplier    string      `json:"supplier"`
	Country     string      `json:"country"`
	Code        string      `json:"code"`
	Contents    []string    `json:"contents"`
	Nutritional nutritional `json:"nutritional"`
	Reusable    bool        `json:"reusable"`
	Packaging   string      `json:"packaging"`
	Description string      `json:"description"`
}

type response struct {
	Name             string      `json:"name"`
	Supplier         string      `json:"supplier"`
	Country          string      `json:"country"`
	Code             string      `json:"code"`
	Contents         []string    `json:"contents"`
	Nutritional      nutritional `json:"nutritional"`
	Reusable         bool        `json:"reusable"`
	Packaging        string      `json:"packaging"`
	Description      string      `json:"description"`
	Errors           string      `json:"errors"`
	ScoreUmwelt      float32     `json:"scoreUmwelt"`
	ScoreVerpackung  float32     `json:"scoreVerpackung"`
	ScoreHerkunft    float32     `json:"scoreHerkunft"`
	ScoreEthik       float32     `json:"scoreEthik"`
	ScoreHealth      float32     `json:"scoreHealth"`
	ScoreIngredients float32     `json:"scoreIngredients"`
	ScoreNutrition   float32     `json:"scoreNutrition"`
}

type nutritional struct {
	Calories float32 `json:"calories"`
	Glucides float32 `json:"glucides"`
	Sugar    float32 `json:"sugar"`
	Lipides  float32 `json:"lipides"`
	Proteins float32 `json:"proteins"`
	Salt     float32 `json:"salt"`
}

type request struct {
	Code   string `json:"code"`
	Origin string `json:"origin"`
}

type securitySetting struct {
	Header string `json:"header"`
	Option string `json:"option"`
}

type serverConfig struct {
	Port           string `json:"port"`
	Cert           string `json:"cert"`
	Key            string `json:"key"`
	SecurityConfig string `json:"securityConfig"`
}
type nutrition struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}
