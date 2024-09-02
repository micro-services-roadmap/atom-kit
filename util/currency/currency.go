package currency

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/micro-services-roadmap/kit-common/kg"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
)

const keyUrl = "https://v6.exchangerate-api.com/v6/%s/latest/%s"

// Cache holds the exchange rates with a TTL
type Cache struct {
	Rates     map[string]float64
	Timestamp time.Time
}

var (
	CacheTTL     = time.Minute * 30    // Cache time-to-live
	DefaultRates = map[string]float64{ // 202409002 rates
		"usd": 1,
		"aed": 3.6725,
		"afn": 70.5967,
		"all": 90.0267,
		"amd": 387.6498,
		"ang": 1.7900,
		"aoa": 919.1545,
		"ars": 952.8300,
		"aud": 1.4769,
		"awg": 1.7900,
		"azn": 1.6994,
		"bam": 1.7698,
		"bbd": 2.0000,
		"bdt": 119.4636,
		"bgn": 1.7698,
		"bhd": 0.3760,
		"bif": 2878.9290,
		"bmd": 1.0000,
		"bnd": 1.3057,
		"bob": 6.9190,
		"brl": 5.6129,
		"bsd": 1.0000,
		"btn": 83.9074,
		"bwp": 13.2731,
		"byn": 3.2564,
		"bzd": 2.0000,
		"cad": 1.3486,
		"cdf": 2846.4764,
		"chf": 0.8504,
		"clp": 913.9279,
		"cny": 7.1011,
		"cop": 4145.0138,
		"crc": 519.0559,
		"cup": 24.0000,
		"cve": 99.7792,
		"czk": 22.6377,
		"djf": 177.7210,
		"dkk": 6.7515,
		"dop": 59.6098,
		"dzd": 133.9012,
		"egp": 48.5959,
		"ern": 15.0000,
		"etb": 111.1551,
		"eur": 0.9050,
		"fjd": 2.1999,
		"fkp": 0.7619,
		"fok": 6.7512,
		"gbp": 0.7619,
		"gel": 2.6899,
		"ggp": 0.7619,
		"ghs": 15.6729,
		"gip": 0.7619,
		"gmd": 70.1383,
		"gnf": 8699.0228,
		"gtq": 7.7292,
		"gyd": 209.2370,
		"hkd": 7.7982,
		"hnl": 24.7883,
		"hrk": 6.8180,
		"htg": 131.7185,
		"huf": 354.9326,
		"idr": 15503.9359,
		"ils": 3.6469,
		"imp": 0.7619,
		"inr": 83.9045,
		"iqd": 1309.3857,
		"irr": 42012.9876,
		"isk": 138.1787,
		"jep": 0.7619,
		"jmd": 157.0576,
		"jod": 0.7090,
		"jpy": 146.2394,
		"kes": 128.7571,
		"kgs": 85.1136,
		"khr": 4067.4058,
		"kid": 1.4765,
		"kmf": 445.1834,
		"krw": 1335.9617,
		"kwd": 0.3053,
		"kyd": 0.8333,
		"kzt": 481.9196,
		"lak": 21913.9659,
		"lbp": 89500.0000,
		"lkr": 298.9345,
		"lrd": 195.0560,
		"lsl": 17.8092,
		"lyd": 4.7601,
		"mad": 9.7440,
		"mdl": 17.4009,
		"mga": 4541.9700,
		"mkd": 55.3955,
		"mmk": 2101.1640,
		"mnt": 3352.2975,
		"mop": 8.0321,
		"mru": 39.7709,
		"mur": 46.3954,
		"mvr": 15.4389,
		"mwk": 1735.9596,
		"mxn": 19.7225,
		"myr": 4.3202,
		"mzn": 63.6764,
		"nad": 17.8092,
		"ngn": 1594.4098,
		"nio": 36.8990,
		"nok": 10.6019,
		"npr": 134.2518,
		"nzd": 1.6006,
		"omr": 0.3845,
		"pab": 1.0000,
		"pen": 3.7477,
		"pgk": 3.8976,
		"php": 56.2018,
		"pkr": 278.7689,
		"pln": 3.8718,
		"pyg": 7600.7611,
		"qar": 3.6400,
		"ron": 4.4919,
		"rsd": 105.7406,
		"rub": 90.7650,
		"rwf": 1336.6490,
		"sar": 3.7500,
		"sbd": 8.5130,
		"scr": 13.8141,
		"sdg": 458.8843,
		"sek": 10.2653,
		"sgd": 1.3041,
		"shp": 0.7619,
		"sle": 22.6041,
		"sll": 22604.0599,
		"sos": 571.3137,
		"srd": 28.9949,
		"ssp": 3130.9772,
		"stn": 22.1701,
		"syp": 12852.7576,
		"szl": 17.8092,
		"thb": 34.0695,
		"tjs": 10.6061,
		"tmt": 3.4995,
		"tnd": 3.0512,
		"top": 2.2984,
		"try": 34.0982,
		"ttd": 6.7706,
		"tvd": 1.4765,
		"twd": 31.9248,
		"tzs": 2719.8335,
		"uah": 41.0456,
		"ugx": 3720.1853,
		"uyu": 40.3458,
		"uzs": 12697.0095,
		"ves": 36.6454,
		"vnd": 24877.7719,
		"vuv": 117.8039,
		"wst": 2.7036,
		"xaf": 593.5778,
		"xcd": 2.7000,
		"xdr": 0.7425,
		"xof": 593.5778,
		"xpf": 107.9839,
		"yer": 250.2550,
		"zar": 17.8116,
		"zmw": 26.2067,
		"zwl": 13.8546,
	}
)

type ExchangeRateResp struct {
	Result          string           `json:"result"`
	BaseCode        string           `json:"base_code"`
	ConversionRates *ConversionRates `json:"conversion_rates"`
}

type ConversionRates struct {
	USD float64 `json:"USD"`
	AED float64 `json:"AED"`
	AFN float64 `json:"AFN"`
	ALL float64 `json:"ALL"`
	AMD float64 `json:"AMD"`
	ANG float64 `json:"ANG"`
	AOA float64 `json:"AOA"`
	ARS float64 `json:"ARS"`
	AUD float64 `json:"AUD"`
	AWG float64 `json:"AWG"`
	AZN float64 `json:"AZN"`
	BAM float64 `json:"BAM"`
	BBD float64 `json:"BBD"`
	BDT float64 `json:"BDT"`
	BGN float64 `json:"BGN"`
	BHD float64 `json:"BHD"`
	BIF float64 `json:"BIF"`
	BMD float64 `json:"BMD"`
	BND float64 `json:"BND"`
	BOB float64 `json:"BOB"`
	BRL float64 `json:"BRL"`
	BSD float64 `json:"BSD"`
	BTN float64 `json:"BTN"`
	BWP float64 `json:"BWP"`
	BYN float64 `json:"BYN"`
	BZD float64 `json:"BZD"`
	CAD float64 `json:"CAD"`
	CDF float64 `json:"CDF"`
	CHF float64 `json:"CHF"`
	CLP float64 `json:"CLP"`
	CNY float64 `json:"CNY"`
	COP float64 `json:"COP"`
	CRC float64 `json:"CRC"`
	CUP float64 `json:"CUP"`
	CVE float64 `json:"CVE"`
	CZK float64 `json:"CZK"`
	DJF float64 `json:"DJF"`
	DKK float64 `json:"DKK"`
	DOP float64 `json:"DOP"`
	DZD float64 `json:"DZD"`
	EGP float64 `json:"EGP"`
	ERN float64 `json:"ERN"`
	ETB float64 `json:"ETB"`
	EUR float64 `json:"EUR"`
	FJD float64 `json:"FJD"`
	FKP float64 `json:"FKP"`
	FOK float64 `json:"FOK"`
	GBP float64 `json:"GBP"`
	GEL float64 `json:"GEL"`
	GGP float64 `json:"GGP"`
	GHS float64 `json:"GHS"`
	GIP float64 `json:"GIP"`
	GMD float64 `json:"GMD"`
	GNF float64 `json:"GNF"`
	GTQ float64 `json:"GTQ"`
	GYD float64 `json:"GYD"`
	HKD float64 `json:"HKD"`
	HNL float64 `json:"HNL"`
	HRK float64 `json:"HRK"`
	HTG float64 `json:"HTG"`
	HUF float64 `json:"HUF"`
	IDR float64 `json:"IDR"`
	ILS float64 `json:"ILS"`
	IMP float64 `json:"IMP"`
	INR float64 `json:"INR"`
	IQD float64 `json:"IQD"`
	IRR float64 `json:"IRR"`
	ISK float64 `json:"ISK"`
	JEP float64 `json:"JEP"`
	JMD float64 `json:"JMD"`
	JOD float64 `json:"JOD"`
	JPY float64 `json:"JPY"`
	KES float64 `json:"KES"`
	KGS float64 `json:"KGS"`
	KHR float64 `json:"KHR"`
	KID float64 `json:"KID"`
	KMF float64 `json:"KMF"`
	KRW float64 `json:"KRW"`
	KWD float64 `json:"KWD"`
	KYD float64 `json:"KYD"`
	KZT float64 `json:"KZT"`
	LAK float64 `json:"LAK"`
	LBP float64 `json:"LBP"`
	LKR float64 `json:"LKR"`
	LRD float64 `json:"LRD"`
	LSL float64 `json:"LSL"`
	LYD float64 `json:"LYD"`
	MAD float64 `json:"MAD"`
	MDL float64 `json:"MDL"`
	MGA float64 `json:"MGA"`
	MKD float64 `json:"MKD"`
	MMK float64 `json:"MMK"`
	MNT float64 `json:"MNT"`
	MOP float64 `json:"MOP"`
	MRU float64 `json:"MRU"`
	MUR float64 `json:"MUR"`
	MVR float64 `json:"MVR"`
	MWK float64 `json:"MWK"`
	MXN float64 `json:"MXN"`
	MYR float64 `json:"MYR"`
	MZN float64 `json:"MZN"`
	NAD float64 `json:"NAD"`
	NGN float64 `json:"NGN"`
	NIO float64 `json:"NIO"`
	NOK float64 `json:"NOK"`
	NPR float64 `json:"NPR"`
	NZD float64 `json:"NZD"`
	OMR float64 `json:"OMR"`
	PAB float64 `json:"PAB"`
	PEN float64 `json:"PEN"`
	PGK float64 `json:"PGK"`
	PHP float64 `json:"PHP"`
	PKR float64 `json:"PKR"`
	PLN float64 `json:"PLN"`
	PYG float64 `json:"PYG"`
	QAR float64 `json:"QAR"`
	RON float64 `json:"RON"`
	RSD float64 `json:"RSD"`
	RUB float64 `json:"RUB"`
	RWF float64 `json:"RWF"`
	SAR float64 `json:"SAR"`
	SBD float64 `json:"SBD"`
	SCR float64 `json:"SCR"`
	SDG float64 `json:"SDG"`
	SEK float64 `json:"SEK"`
	SGD float64 `json:"SGD"`
	SHP float64 `json:"SHP"`
	SLE float64 `json:"SLE"`
	SLL float64 `json:"SLL"`
	SOS float64 `json:"SOS"`
	SRD float64 `json:"SRD"`
	SSP float64 `json:"SSP"`
	STN float64 `json:"STN"`
	SYP float64 `json:"SYP"`
	SZL float64 `json:"SZL"`
	THB float64 `json:"THB"`
	TJS float64 `json:"TJS"`
	TMT float64 `json:"TMT"`
	TND float64 `json:"TND"`
	TOP float64 `json:"TOP"`
	TRY float64 `json:"TRY"`
	TTD float64 `json:"TTD"`
	TVD float64 `json:"TVD"`
	TWD float64 `json:"TWD"`
	TZS float64 `json:"TZS"`
	UAH float64 `json:"UAH"`
	UGX float64 `json:"UGX"`
	UYU float64 `json:"UYU"`
	UZS float64 `json:"UZS"`
	VES float64 `json:"VES"`
	VND float64 `json:"VND"`
	VUV float64 `json:"VUV"`
	WST float64 `json:"WST"`
	XAF float64 `json:"XAF"`
	XCD float64 `json:"XCD"`
	XDR float64 `json:"XDR"`
	XOF float64 `json:"XOF"`
	XPF float64 `json:"XPF"`
	YER float64 `json:"YER"`
	ZAR float64 `json:"ZAR"`
	ZMW float64 `json:"ZMW"`
	ZWL float64 `json:"ZWL"`
}

func DoExchangeRate4Usd(key ...string) (*ExchangeRateResp, error) {
	return DoExchangeRate("USD", key...)
}

func DoExchangeRate(base string, key ...string) (*ExchangeRateResp, error) {
	var apiKey string
	if len(key) == 0 {
		apiKey = kg.C.System.CurrencyKey
	} else {
		apiKey = key[0]
	}

	url := fmt.Sprintf(keyUrl, apiKey, base)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to make request: %v\n", err)
		kg.L.Error("Failed to make request: %v", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		kg.L.Error("Failed to read response body: %v", zap.Error(err))
		return nil, err
	}

	rateResp := &ExchangeRateResp{}
	err = json.Unmarshal(body, rateResp)
	if err != nil {
		return nil, err
	}

	if rateResp.Result != "success" {
		return nil, errors.New("failed to get exchange rate")
	}

	return rateResp, nil
}

func Convert2Map(rates *ConversionRates) (map[string]float64, error) {

	if rates == nil {
		return nil, errors.New("rates is nil")
	}

	return map[string]float64{
		"usd": rates.USD,
		"aed": rates.AED,
		"afn": rates.AFN,
		"all": rates.ALL,
		"amd": rates.AMD,
		"ang": rates.ANG,
		"aoa": rates.AOA,
		"ars": rates.ARS,
		"aud": rates.AUD,
		"awg": rates.AWG,
		"azn": rates.AZN,
		"bam": rates.BAM,
		"bbd": rates.BBD,
		"bdt": rates.BDT,
		"bgn": rates.BGN,
		"bhd": rates.BHD,
		"bif": rates.BIF,
		"bmd": rates.BMD,
		"bnd": rates.BND,
		"bob": rates.BOB,
		"brl": rates.BRL,
		"bsd": rates.BSD,
		"btn": rates.BTN,
		"bwp": rates.BWP,
		"byn": rates.BYN,
		"bzd": rates.BZD,
		"cad": rates.CAD,
		"cdf": rates.CDF,
		"chf": rates.CHF,
		"clp": rates.CLP,
		"cny": rates.CNY,
		"cop": rates.COP,
		"crc": rates.CRC,
		"cup": rates.CUP,
		"cve": rates.CVE,
		"czk": rates.CZK,
		"djf": rates.DJF,
		"dkk": rates.DKK,
		"dop": rates.DOP,
		"dzd": rates.DZD,
		"egp": rates.EGP,
		"ern": rates.ERN,
		"etb": rates.ETB,
		"eur": rates.EUR,
		"fjd": rates.FJD,
		"fkp": rates.FKP,
		"fok": rates.FOK,
		"gbp": rates.GBP,
		"gel": rates.GEL,
		"ggp": rates.GGP,
		"ghs": rates.GHS,
		"gip": rates.GIP,
		"gmd": rates.GMD,
		"gnf": rates.GNF,
		"gtq": rates.GTQ,
		"gyd": rates.GYD,
		"hkd": rates.HKD,
		"hnl": rates.HNL,
		"hrk": rates.HRK,
		"htg": rates.HTG,
		"huf": rates.HUF,
		"idr": rates.IDR,
		"ils": rates.ILS,
		"imp": rates.IMP,
		"inr": rates.INR,
		"iqd": rates.IQD,
		"irr": rates.IRR,
		"isk": rates.ISK,
		"jep": rates.JEP,
		"jmd": rates.JMD,
		"jod": rates.JOD,
		"jpy": rates.JPY,
		"kes": rates.KES,
		"kgs": rates.KGS,
		"khr": rates.KHR,
		"kid": rates.KID,
		"kmf": rates.KMF,
		"krw": rates.KRW,
		"kwd": rates.KWD,
		"kyd": rates.KYD,
		"kzt": rates.KZT,
		"lak": rates.LAK,
		"lbp": rates.LBP,
		"lkr": rates.LKR,
		"lrd": rates.LRD,
		"lsl": rates.LSL,
		"lyd": rates.LYD,
		"mad": rates.MAD,
		"mdl": rates.MDL,
		"mga": rates.MGA,
		"mkd": rates.MKD,
		"mmk": rates.MMK,
		"mnt": rates.MNT,
		"mop": rates.MOP,
		"mru": rates.MRU,
		"mur": rates.MUR,
		"mvr": rates.MVR,
		"mwk": rates.MWK,
		"mxn": rates.MXN,
		"myr": rates.MYR,
		"mzn": rates.MZN,
		"nad": rates.NAD,
		"ngn": rates.NGN,
		"nio": rates.NIO,
		"nok": rates.NOK,
		"npr": rates.NPR,
		"nzd": rates.NZD,
		"omr": rates.OMR,
		"pab": rates.PAB,
		"pen": rates.PEN,
		"pgk": rates.PGK,
		"php": rates.PHP,
		"pkr": rates.PKR,
		"pln": rates.PLN,
		"pyg": rates.PYG,
		"qar": rates.QAR,
		"ron": rates.RON,
		"rsd": rates.RSD,
		"rub": rates.RUB,
		"rwf": rates.RWF,
		"sar": rates.SAR,
		"sbd": rates.SBD,
		"scr": rates.SCR,
		"sdg": rates.SDG,
		"sek": rates.SEK,
		"sgd": rates.SGD,
		"shp": rates.SHP,
		"sle": rates.SLE,
		"sll": rates.SLL,
		"sos": rates.SOS,
		"srd": rates.SRD,
		"ssp": rates.SSP,
		"stn": rates.STN,
		"syp": rates.SYP,
		"szl": rates.SZL,
		"thb": rates.THB,
		"tjs": rates.TJS,
		"tmt": rates.TMT,
		"tnd": rates.TND,
		"top": rates.TOP,
		"try": rates.TRY,
		"ttd": rates.TTD,
		"tvd": rates.TVD,
		"twd": rates.TWD,
		"tzs": rates.TZS,
		"uah": rates.UAH,
		"ugx": rates.UGX,
		"uyu": rates.UYU,
		"uzs": rates.UZS,
		"ves": rates.VES,
		"vnd": rates.VND,
		"vuv": rates.VUV,
		"wst": rates.WST,
		"xaf": rates.XAF,
		"xcd": rates.XCD,
		"xdr": rates.XDR,
		"xof": rates.XOF,
		"xpf": rates.XPF,
		"yer": rates.YER,
		"zar": rates.ZAR,
		"zmw": rates.ZMW,
		"zwl": rates.ZWL,
	}, nil
}
