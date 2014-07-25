// Paste the code from here
// https://github.com/AfterShip/node-phone/blob/master/lib/index.js#L1

var iso3166_data = [];

var res = [];
var stringify = function(obj) {
	var wrap = function(a) {
	 	var b = [];
	 	for (var i = 0; i < a.length; i++) {
	 		b[i] = "\"" + a[i] + "\"";
	 	}
	 	return b;
	};
	return "&PhoneData{\n\
	Alpha2:\"" + obj.alpha2 + "\",\n\
	Alpha3:\"" + obj.alpha3 + "\",\n\
	CountryCode:\"" + obj.country_code + "\",\n\
	CountryName:\"" + obj.country_name + "\",\n\
	MobileBeginsWith: []string{" + wrap(obj.mobile_begin_with).join(", ") + "},\n\
	PhoneNumberLengths: []int{" + obj.phone_number_lengths.join(",") + "},\n\
},\n";
};

for(var i = 0; i < iso3166_data.length; i++) {
	res.push(stringify(iso3166_data[i]))
}

console.log(res.join(""))
