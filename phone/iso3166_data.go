package phone

import (
	c "github.com/dicefm/extraterrestrial/countries"
)

var (
	ISO3166_PhoneData []*PhoneData = []*PhoneData{
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameUnitedStates],
			MobileBeginsWith:   []string{"201", "202", "203", "205", "206", "207", "208", "209", "210", "212", "213", "214", "215", "216", "217", "218", "219", "224", "225", "227", "228", "229", "231", "234", "239", "240", "248", "251", "252", "253", "254", "256", "260", "262", "267", "269", "270", "272", "274", "276", "278", "281", "283", "301", "302", "303", "304", "305", "307", "308", "309", "310", "312", "313", "314", "315", "316", "317", "318", "319", "320", "321", "323", "325", "327", "330", "331", "334", "336", "337", "339", "341", "346", "347", "351", "352", "360", "361", "364", "369", "380", "385", "386", "401", "402", "404", "405", "406", "407", "408", "409", "410", "412", "413", "414", "415", "417", "419", "423", "424", "425", "430", "432", "434", "435", "440", "442", "443", "445", "447", "458", "464", "469", "470", "475", "478", "479", "480", "484", "501", "502", "503", "504", "505", "507", "508", "509", "510", "512", "513", "515", "516", "517", "518", "520", "530", "531", "534", "539", "540", "541", "551", "557", "559", "561", "562", "563", "564", "567", "570", "571", "573", "574", "575", "580", "582", "585", "586", "601", "602", "603", "605", "606", "607", "608", "609", "610", "612", "614", "615", "616", "617", "618", "619", "620", "623", "626", "627", "628", "630", "631", "636", "641", "646", "650", "651", "657", "659", "660", "661", "662", "667", "669", "678", "679", "681", "682", "689", "701", "702", "703", "704", "706", "707", "708", "712", "713", "714", "715", "716", "717", "718", "719", "720", "724", "725", "727", "730", "731", "732", "734", "737", "740", "747", "752", "754", "757", "760", "762", "763", "764", "765", "769", "770", "772", "773", "774", "775", "779", "781", "785", "786", "801", "802", "803", "804", "805", "806", "808", "810", "812", "813", "814", "815", "816", "817", "818", "828", "830", "831", "832", "835", "843", "845", "847", "848", "850", "856", "857", "858", "859", "860", "862", "863", "864", "865", "870", "872", "878", "901", "903", "904", "906", "907", "908", "909", "910", "912", "913", "914", "915", "916", "917", "918", "919", "920", "925", "927", "928", "929", "931", "935", "936", "937", "938", "940", "941", "947", "949", "951", "952", "954", "956", "957", "959", "970", "971", "972", "973", "975", "978", "979", "980", "984", "985", "989"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameAruba],
			MobileBeginsWith:   []string{"5", "6", "7", "9"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameAfghanistan],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameAngola],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameAnguilla],
			MobileBeginsWith:   []string{"5", "7"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameAlandIslands],
			MobileBeginsWith:   []string{"18"},
			PhoneNumberLengths: []int{6, 7, 8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameAlbania],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameAndorra],
			MobileBeginsWith:   []string{"3", "4", "6"},
			PhoneNumberLengths: []int{6, 9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameUnitedArabEmirates],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameArgentina],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameArmenia],
			MobileBeginsWith:   []string{"5", "7", "9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameAmericanSamoa],
			MobileBeginsWith:   []string{"733", "258"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameAntiguaandBarbuda],
			MobileBeginsWith:   []string{"4", "7"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameAustralia],
			MobileBeginsWith:   []string{"4"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameAustria],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameAzerbaijan],
			MobileBeginsWith:   []string{"4", "5", "6", "7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBurundi],
			MobileBeginsWith:   []string{"7", "29"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBelgium],
			MobileBeginsWith:   []string{"4"},
			PhoneNumberLengths: []int{8, 9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBenin],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBurkinaFaso],
			MobileBeginsWith:   []string{"6", "7"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBangladesh],
			MobileBeginsWith:   []string{"1"},
			PhoneNumberLengths: []int{8, 9, 10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBulgaria],
			MobileBeginsWith:   []string{"87", "88", "89", "98", "99", "43"},
			PhoneNumberLengths: []int{7, 8, 9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBahrain],
			MobileBeginsWith:   []string{"3"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBahamas],
			MobileBeginsWith:   []string{"3", "4", "5", "6", "7"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBosniaandHerzegovina],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBelarus],
			MobileBeginsWith:   []string{"25", "29", "33", "44"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBelize],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBermuda],
			MobileBeginsWith:   []string{"3", "5", "7"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBolivia],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBrazil],
			MobileBeginsWith:   []string{"119", "129", "139", "149", "159", "169", "179", "189", "199", "219", "229", "249", "279", "289", "31", "32", "34", "38", "41", "43", "44", "45", "47", "48", "51", "53", "54", "55", "61", "62", "65", "67", "68", "69", "71", "73", "75", "77", "79", "81", "82", "83", "84", "85", "86", "91", "92", "95", "96", "98"},
			PhoneNumberLengths: []int{10, 11},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBarbados],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBruneiDarussalam],
			MobileBeginsWith:   []string{"7", "8"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBhutan],
			MobileBeginsWith:   []string{"17"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameBotswana],
			MobileBeginsWith:   []string{"71", "72", "73", "74", "75", "76"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameCentralAfricanRepublic],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameCanada],
			MobileBeginsWith:   []string{"204", "226", "236", "249", "250", "289", "306", "343", "365", "403", "416", "418", "431", "437", "438", "450", "506", "514", "519", "579", "581", "587", "600", "604", "613", "639", "647", "705", "709", "778", "780", "807", "819", "867", "873", "902", "905"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSwitzerland],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameChile],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameChina],
			MobileBeginsWith:   []string{"13", "14", "15", "18"},
			PhoneNumberLengths: []int{11},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameCoteDIvoire],
			MobileBeginsWith:   []string{"0", "4", "5", "6"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameCameroon],
			MobileBeginsWith:   []string{"7", "9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameCongo],
			MobileBeginsWith:   []string{"8", "9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameCongo],
			MobileBeginsWith:   []string{"0"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameCookIslands],
			MobileBeginsWith:   []string{"5", "7"},
			PhoneNumberLengths: []int{5},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameColombia],
			MobileBeginsWith:   []string{"3"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameComoros],
			MobileBeginsWith:   []string{"3", "76"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameCapeVerde],
			MobileBeginsWith:   []string{"5", "9"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameCostaRica],
			MobileBeginsWith:   []string{"5", "6", "7", "8"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameCuba],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameCaymanIslands],
			MobileBeginsWith:   []string{"345"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameCyprus],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameCzechRepublic],
			MobileBeginsWith:   []string{"6", "7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameGermany],
			MobileBeginsWith:   []string{"15", "16", "17"},
			PhoneNumberLengths: []int{3, 4, 5, 6, 7, 8, 9, 10, 11},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameDjibouti],
			MobileBeginsWith:   []string{"77"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameDominica],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameDenmark],
			MobileBeginsWith:   []string{"2", "30", "31", "40", "41", "42", "50", "51", "52", "53", "60", "61", "71", "81", "91", "92", "93"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameDominicanRepublic],
			MobileBeginsWith:   []string{"809", "829", "849"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameAlgeria],
			MobileBeginsWith:   []string{"5", "6", "7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameEcuador],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameEgypt],
			MobileBeginsWith:   []string{"1"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameEritrea],
			MobileBeginsWith:   []string{"1", "7", "8"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSpain],
			MobileBeginsWith:   []string{"6", "7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameEstonia],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{7, 8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameEthiopia],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameFinland],
			MobileBeginsWith:   []string{"4", "5"},
			PhoneNumberLengths: []int{4, 5, 6, 7, 8, 9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameFiji],
			MobileBeginsWith:   []string{"7", "9"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameFalklandIslands],
			MobileBeginsWith:   []string{"5", "6"},
			PhoneNumberLengths: []int{5},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameFrance],
			MobileBeginsWith:   []string{"6", "7"},
			PhoneNumberLengths: []int{8, 9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameFaroeIslands],
			MobileBeginsWith:   []string{"21", "22", "23", "24", "25", "26", "27", "28", "29", "5", "71", "72", "73", "74", "75", "76", "77", "78", "79", "91", "92", "93", "94", "95", "96", "97", "98", "99"},
			PhoneNumberLengths: []int{6},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMicronesiaFederatedStatesOf],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameGabon],
			MobileBeginsWith:   []string{"05", "06", "07"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameUnitedKingdom],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{7, 9, 10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameGeorgia],
			MobileBeginsWith:   []string{"5", "7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameGhana],
			MobileBeginsWith:   []string{"2", "5"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameGibraltar],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameGuinea],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameGuadeloupe],
			MobileBeginsWith:   []string{"690"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameGambia],
			MobileBeginsWith:   []string{"7", "9"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameGuineaBissau],
			MobileBeginsWith:   []string{"5", "6", "7"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameEquatorialGuinea],
			MobileBeginsWith:   []string{"222", "551"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameGreece],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameGrenada],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameGreenland],
			MobileBeginsWith:   []string{"4", "5"},
			PhoneNumberLengths: []int{6},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameGuatemala],
			MobileBeginsWith:   []string{"3", "4", "5"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameFrenchGuiana],
			MobileBeginsWith:   []string{"694"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameGuam],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameGuyana],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameHongKong],
			MobileBeginsWith:   []string{"5", "6", "9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameHonduras],
			MobileBeginsWith:   []string{"3", "7", "8", "9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameCroatia],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8, 9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameHaiti],
			MobileBeginsWith:   []string{"3", "4"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameHungary],
			MobileBeginsWith:   []string{"20", "30", "31", "70"},
			PhoneNumberLengths: []int{8, 9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameIndonesia],
			MobileBeginsWith:   []string{"8"},
			PhoneNumberLengths: []int{9, 10, 11},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameIndia],
			MobileBeginsWith:   []string{"7", "8", "9"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameIreland],
			MobileBeginsWith:   []string{"82", "83", "84", "85", "86", "87", "88", "89"},
			PhoneNumberLengths: []int{7, 8, 9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameIran],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameIraq],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameIceland],
			MobileBeginsWith:   []string{"6", "7", "8"},
			PhoneNumberLengths: []int{7, 9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameIsrael],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameItaly],
			MobileBeginsWith:   []string{"3"},
			PhoneNumberLengths: []int{8, 9, 10, 11, 12},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameJamaica],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameJordan],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameJapan],
			MobileBeginsWith:   []string{"70", "80", "90"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameKazakhstan],
			MobileBeginsWith:   []string{"70", "77"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameKenya],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameKyrgyzstan],
			MobileBeginsWith:   []string{"5", "7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameCambodia],
			MobileBeginsWith:   []string{"1", "6", "7", "8", "9"},
			PhoneNumberLengths: []int{8, 9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameKiribati],
			MobileBeginsWith:   []string{"9", "30"},
			PhoneNumberLengths: []int{5},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSaintKittsAndNevis],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameKorea],
			MobileBeginsWith:   []string{"1"},
			PhoneNumberLengths: []int{9, 10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameKuwait],
			MobileBeginsWith:   []string{"5", "6", "9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameLao],
			MobileBeginsWith:   []string{"20"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameLebanon],
			MobileBeginsWith:   []string{"3", "7"},
			PhoneNumberLengths: []int{7, 8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameLiberia],
			MobileBeginsWith:   []string{"4", "5", "6", "7"},
			PhoneNumberLengths: []int{7, 8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameLibyanArabJamahiriya],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSaintLucia],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameLiechtenstein],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{7, 8, 9, 10, 11},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSriLanka],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameLesotho],
			MobileBeginsWith:   []string{"5", "6"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameLithuania],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameLuxembourg],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{6, 7, 8, 9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameLatvia],
			MobileBeginsWith:   []string{"2"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMacao],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMorocco],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMonaco],
			MobileBeginsWith:   []string{"4", "6"},
			PhoneNumberLengths: []int{8, 9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMoldova],
			MobileBeginsWith:   []string{"6", "7"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMadagascar],
			MobileBeginsWith:   []string{"3"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMaldives],
			MobileBeginsWith:   []string{"7", "9"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMexico],
			MobileBeginsWith:   []string{""},
			PhoneNumberLengths: []int{10, 11},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMarshallIslands],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMacedonia],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMali],
			MobileBeginsWith:   []string{"6", "7"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMalta],
			MobileBeginsWith:   []string{"79", "99"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMyanmar],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMontenegro],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMongolia],
			MobileBeginsWith:   []string{"5", "8", "9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameNorthernMarianaIslands],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMozambique],
			MobileBeginsWith:   []string{"8"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMauritania],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMontserrat],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMartinique],
			MobileBeginsWith:   []string{"696"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMauritius],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMalawi],
			MobileBeginsWith:   []string{"77", "88", "99"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMalaysia],
			MobileBeginsWith:   []string{"1"},
			PhoneNumberLengths: []int{9, 10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameMayotte],
			MobileBeginsWith:   []string{"639"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameNamibia],
			MobileBeginsWith:   []string{"60", "81", "82", "85"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameNewCaledonia],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{6},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameNiger],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameNorfolkIsland],
			MobileBeginsWith:   []string{"5", "8"},
			PhoneNumberLengths: []int{5},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameNigeria],
			MobileBeginsWith:   []string{"70", "80", "81"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameNicaragua],
			MobileBeginsWith:   []string{"8"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameNiue],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{4},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameNetherlands],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{9, 10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameNorway],
			MobileBeginsWith:   []string{"4", "9"},
			PhoneNumberLengths: []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameNepal],
			MobileBeginsWith:   []string{"97", "98"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameNauru],
			MobileBeginsWith:   []string{"555"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameNewZealand],
			MobileBeginsWith:   []string{"2"},
			PhoneNumberLengths: []int{8, 9, 10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameOman],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNamePakistan],
			MobileBeginsWith:   []string{"3"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNamePanama],
			MobileBeginsWith:   []string{"5", "6"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNamePeru],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNamePhilippines],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNamePalau],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNamePapuaNewGuinea],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNamePoland],
			MobileBeginsWith:   []string{"5", "6", "7", "8"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNamePuertoRico],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNamePortugal],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameParaguay],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNamePalestinia],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameFrenchPolynesia],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{6},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameQatar],
			MobileBeginsWith:   []string{"33", "55", "66", "77"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameReunion],
			MobileBeginsWith:   []string{"692", "693"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameRomania],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameRussianFederation],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameRwanda],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSaudiArabia],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSudan],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSenegal],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSingapore],
			MobileBeginsWith:   []string{"8", "9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSaintHelena],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{4},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSvalbardAndJanMayen],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSolomonIslands],
			MobileBeginsWith:   []string{"7", "8"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSierraLeone],
			MobileBeginsWith:   []string{"21", "25", "30", "33", "34", "40", "44", "50", "55", "76", "77", "78", "79", "88"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameElSalvador],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSanMarino],
			MobileBeginsWith:   []string{"3", "6"},
			PhoneNumberLengths: []int{4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSomalia],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSaintPierreAndMiquelon],
			MobileBeginsWith:   []string{"55"},
			PhoneNumberLengths: []int{6},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSerbia],
			MobileBeginsWith:   []string{"6"},
			PhoneNumberLengths: []int{8, 9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSaoTomeandPrincipe],
			MobileBeginsWith:   []string{"98", "99"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSuriname],
			MobileBeginsWith:   []string{"6", "7", "8"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSlovakia],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSlovenia],
			MobileBeginsWith:   []string{"3", "4", "5", "6", "7"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSweden],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{6, 7, 8, 9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSeychelles],
			MobileBeginsWith:   []string{"2"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSyrianArabRepublic],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameTurksAndCaicosIslands],
			MobileBeginsWith:   []string{"2", "3", "4"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameChad],
			MobileBeginsWith:   []string{"6", "7", "9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameTogo],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameThailand],
			MobileBeginsWith:   []string{"8", "9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameTajikistan],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameTokelau],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{4},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameTurkmenistan],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameTimorLeste],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameTonga],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{5},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameTrinidadAndTobago],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameTunisia],
			MobileBeginsWith:   []string{"2", "9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameTurkey],
			MobileBeginsWith:   []string{"5"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameTuvalu],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{5},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameTaiwan],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameTanzania],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameUganda],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameUkraine],
			MobileBeginsWith:   []string{"39", "50", "63", "66", "67", "68", "9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameUruguay],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{8},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameUzbekistan],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSaintVincentAndTheGrenedines],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameVenezuela],
			MobileBeginsWith:   []string{"4"},
			PhoneNumberLengths: []int{10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameVirginIslandsBritish],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameVirginIslandsUS],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameVietNam],
			MobileBeginsWith:   []string{"9", "1"},
			PhoneNumberLengths: []int{9, 10},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameVanuatu],
			MobileBeginsWith:   []string{"5", "7"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameWallisAndFutuna],
			MobileBeginsWith:   []string{},
			PhoneNumberLengths: []int{6},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSamoa],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{7},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameYemen],
			MobileBeginsWith:   []string{"7"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameSouthAfrica],
			MobileBeginsWith:   []string{"6", "7", "8"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameZambia],
			MobileBeginsWith:   []string{"9"},
			PhoneNumberLengths: []int{9},
		},
		&PhoneData{
			CountryData:        c.ISO3166_CountriesData[c.CountryNameZimbabwe],
			MobileBeginsWith:   []string{"71", "73", "77"},
			PhoneNumberLengths: []int{9},
		},
	}
)
