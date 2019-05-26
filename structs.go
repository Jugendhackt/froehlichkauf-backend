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
	Name                  string      `json:"name"`
	Supplier              string      `json:"supplier"`
	Country               string      `json:"country"`
	Code                  string      `json:"code"`
	Contents              []string    `json:"contents"`
	Nutritional           nutritional `json:"nutritional"`
	Reusable              bool        `json:"reusable"`
	Packaging             string      `json:"packaging"`
	Description           string      `json:"description"`
	Errors                string      `json:"errors"`
	ScoreUmwelt           int         `json:"scoreUmwelt"`
	ScoreUmweltHalve      bool        `json:"scoreUmweltHalve"`
	ScoreVerpackung       int         `json:"scoreVerpackung"`
	ScoreVerpackungHalve  bool        `json:"scoreVerpackungHalve"`
	ScoreHerkunft         int         `json:"scoreHerkunft"`
	ScoreHerkunftHalve    bool        `json:"scoreHerkunftHalve"`
	ScoreEthik            int         `json:"scoreEthik"`
	ScoreEthikHalve       bool        `json:"scoreEthikHalve"`
	ScoreHealth           int         `json:"scoreHealth"`
	ScoreHealthHalve      bool        `json:"scoreHealthHalve"`
	ScoreIngredients      int         `json:"scoreIngredients"`
	ScoreIngredientsHalve bool        `json:"scoreIngredientsHalve"`
	ScoreNutrition        int         `json:"scoreNutrition"`
	ScoreNutritionHalve   bool        `json:"scoreNutritionHalve"`
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
	Encryption     bool   `json:"encryption"`
}
type nutrition struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}

type companies struct {
	CocaCola              []string `json:"Coca-Cola"`
	Kellogs               []string `json:"Kellogs"`
	Mars                  []string `json:"Mars"`
	Nestle                []string `json:"Nestle"`
	PepsiCo               []string `json:"PepsiCo"`
	ProcterGambles        []string `json:"Procter & Gamble"`
	Unilever              []string `json:"Unilever"`
	MondelezInternational []string `json:"Mondelez International"`
}
