var e1 = eveCables.save({"_key":"420527197812010018", "name":"董爱华", "tel":"13511111111", "landline":"02888888888","address":"成都市科环北路", "city":"成都"})._id;

var id1 = entIdCards.save({"_key":"420527197812010018"})._id;
var p1 = entPersons.save({"name":"董爱华", "id":"420527197812010018", "tels":["13511111111"], "landlines":["02888888888"],"addresses":["成都市科环北路"], "cities":["成都"]})._id;
var tel1 = entTels.save({"_key":"13511111111"})._id;
var landline1 = entLandlines.save({"_key":"02888888888"})._id;

relCable.save(id1, p1, {"type":"idperson", "event":e1});
relCable.save(id1, tel1, {"type":"idtel","event":e1});
relCable.save(id1, landline1, {"type":"idlandline", "event":e1});


var e2 = eveCables.save({"_key":"520527197901010018", "name":"喻波", "tel":"13611111111", "landline":"01000000000","address":"北京市东城区", "city":"北京"})._id;
var id2 = entIdCards.save({"_key":"520527197901010018"})._id;
var p2 = entPersons.save({"name":"喻波", "id":"520527197901010018", "tels":["13611111111"], "landlines":["01000000000"], "addresses":["北京市东城区"], "cities":["北京"]})._id;
var tel2 = entTels.save({"_key":"13611111111"})._id;
var landline2 = entLandlines.save({"_key":"01000000000"})._id;

relCable.save(id2, p2, {"type":"idperson", "event":e2});
relCable.save(id2, tel2, {"type":"idtel","event":e2});
relCable.save(id2, landline2, {"type":"idlandline", "event":e2});


var e3 = eveCables.save({"_key":"620527196901010018", "name":"王志海", "tel":"13711111111", "landline":"01011111111","address":"北京市西城区", "city":"北京"})._id;
var id3 = entIdCards.save({"_key":"620527196901010018"})._id;
var p3 = entPersons.save({"name":"王志海", "id":"620527196901010018", "tels":["13711111111"], "landlines":["01011111111"], "addresses":["北京市西城区"], "cities":["北京"]})._id;
var tel3 = entTels.save({"_key":"13711111111"})._id;
var landline3 = entLandlines.save({"_key":"01011111111"})._id;

relCable.save(id3, p3, {"type":"idperson", "event":e3});
relCable.save(id3, tel3, {"type":"idtel","event":e3});
relCable.save(id3, landline3, {"type":"idlandline", "event":e3});


var e4 = eveCables.save({"_key":"720527197003010018", "name":"薛晓文", "tel":"13811111111", "landline":"01022222222","address":"北京市海淀区", "city":"北京"})._id;
var id4 = entIdCards.save({"_key":"720527197003010018"})._id;
var p4 = entPersons.save({"name":"薛晓文", "id":"720527197003010018", "tels":["13811111111"], "landlines":["01022222222"], "addresses":["北京市海淀区"], "cities":["北京"]})._id;
var tel4 = entTels.save({"_key":"13811111111"})._id;
var landline4 = entLandlines.save({"_key":"01022222222"})._id;

relCable.save(id4, p4, {"type":"idperson", "event":e4});
relCable.save(id4, tel4, {"type":"idtel","event":e4});
relCable.save(id4, landline4, {"type":"idlandline", "event":e4});

var e5 = eveStudies.save({"_key":"420527200905010018", "name":"董懿萱", "sex":"女", "address":"成都市科环北路", "tel":"13568984989", "school":"贵溪小学"})._id;
var s5 = entSchools.save({"name":"贵溪小学"})._id;
var p5 = entPersons.save({"name":"董懿萱", "sex":"女", "id":"420527200905010018", "tels":["13568984989"], "addresses":["成都市科环北路"], "schools":["贵溪小学"]})._id;
var tel5 = entTels.save({"_key":"13568984989"})._id;
var p6 = entPersons.save({"name":"王一", "id":"420527198903010018"})._id;
var id5 = entIdCards.save({"_key":"420527200905010018"})._id;
var id6 = entIdCards.save({"_key":"420527198903010018"})._id;


relStudy.save(id5, s5, {"type":"studyin", "event":e5})._id;
relStudy.save(id5, id1, {"type":"family", "subtype":"father", "event":e5})._id;
relStudy.save(id5, id6, {"type":"family", "subtype":"mother", "event":e5})._id;
relStudy.save(id5, p5, {"type":"idperson", "event":e5})._id;
relStudy.save(id6, p6, {"type":"idperson", "event":e5})._id;
relStudy.save(id5, tel5, {"type":"idtel", "event":e5})._id;


var e7 = eveStudies.save({"_key":"520527199904010018", "name":"喻小小", "sex":"男", "address":"北京市东城区", "tel":"13611111111", "school":"北京中学"})._id;
var s7 = entSchools.save({"name":"北京中学"})._id;
var p7 = entPersons.save({"name":"喻小小", "sex":"男", "id":"520527199904010018", "tels":["13611111111"], "addresses":["北京市东城区"], "schools":["北京中学"]})._id;
var p8 = entPersons.save({"name":"李一", "id":"520527198802010018"})._id;
var id7 = entIdCards.save({"_key":"520527199904010018"})._id;
var id8 = entIdCards.save({"_key":"520527198802010018"})._id;

var e8 = eveKuaidis.save({"_key":"12345678",
    "from":{"address":"成都", "name":"董爱华", "tel":"13568984989"},
    "to":{"address":"北京", "name":"喻波", "tel":"13568984988"},
    "receiver":{"name":"王小二", "tel":"13534567891", "address":"成都", "company":"顺丰"},
    "sender":{"name":"李小二", "tel":"13423456781", "address":"北京"},
    "__time":1493668984000,
    "updatedAt":1493678984000})._id;
var e9 = eveKuaidis.save({"_key":"22345678",
    "from":{"address":"北京", "name":"喻波", "tel":"13568984988"},
    "to":{"address":"北京", "name":"王志海", "tel":"13568984987"},
    "receiver":{"name":"王小三", "tel":"14534567891", "address":"北京", "company":"顺丰"},
    "sender":{"name":"李小三", "tel":"13523456781", "address":"北京"},
    "__time":1493678984000,
    "updatedAt":1493778984000})._id;
var e10 = eveKuaidis.save({"_key":"32345678",
    "from":{"address":"北京", "name":"王志海", "tel":"13568984987"},
    "to":{"address":"北京", "name":"薛晓文", "tel":"13568984986"},
    "receiver":{"name":"王小三", "tel":"14534567891", "address":"北京", "company":"顺丰"},
    "sender":{"name":"李小三", "tel":"13523456781", "address":"北京"},
    "__time":1493778984000,
    "updatedAt":1493878984000})._id;

var tel9 = entTels.save({"_key":"13568984988"})._id;
var tel10 = entTels.save({"_key":"13568984987"})._id;
var tel11 = entTels.save({"_key":"13568984986"})._id;