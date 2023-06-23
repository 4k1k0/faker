package phone

var (
	phoneFormats = []string{
		// International format
		"+1-{{areaCode}}-{{exchangeCode}}-####",
		"+1 ({{areaCode}}) {{exchangeCode}}-####",
		"+1-{{areaCode}}-{{exchangeCode}}-####",
		"+1.{{areaCode}}.{{exchangeCode}}.####",
		"+1{{areaCode}}{{exchangeCode}}####",
		// Standard formats
		"{{areaCode}}-{{exchangeCode}}-####",
		"({{areaCode}}) {{exchangeCode}}-####",
		"1-{{areaCode}}-{{exchangeCode}}-####",
		"{{areaCode}}.{{exchangeCode}}.####",
		"{{areaCode}}-{{exchangeCode}}-####",
		"({{areaCode}}) {{exchangeCode}}-####",
		"1-{{areaCode}}-{{exchangeCode}}-####",
		"{{areaCode}}.{{exchangeCode}}.####",
		// Extensions
		"{{areaCode}}-{{exchangeCode}}-#### x###",
		"({{areaCode}}) {{exchangeCode}}-#### x###",
		"1-{{areaCode}}-{{exchangeCode}}-#### x###",
		"{{areaCode}}.{{exchangeCode}}.#### x###",
		"{{areaCode}}-{{exchangeCode}}-#### x####",
		"({{areaCode}}) {{exchangeCode}}-#### x####",
		"1-{{areaCode}}-{{exchangeCode}}-#### x####",
		"{{areaCode}}.{{exchangeCode}}.#### x####",
		"{{areaCode}}-{{exchangeCode}}-#### x#####",
		"({{areaCode}}) {{exchangeCode}}-#### x#####",
		"1-{{areaCode}}-{{exchangeCode}}-#### x#####",
		"{{areaCode}}.{{exchangeCode}}.#### x#####"}

	tollFreeAreaCodes = []string{"800", "844", "855", "866", "877", "888"}

	tollFreeFormats = []string{ // Standard formats
		"{{tollFreeAreaCode}}-{{exchangeCode}}-####",
		"({{tollFreeAreaCode}}) {{exchangeCode}}-####",
		"1-{{tollFreeAreaCode}}-{{exchangeCode}}-####",
		"{{tollFreeAreaCode}}.{{exchangeCode}}.####"}
)
