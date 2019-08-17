package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"regexp"
	"log"
)

var sql = `INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000000000', '西藏自治区公安厅', '010000000000', '', '1', '1', '540000', '42', '010000000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000030000', '西藏自治区公安厅治安管理总队', '540000000000', '', '1', '2', '540000', '3', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000050000', '西藏自治区公安厅刑事侦查总队', '540000000000', '', '1', '2', '540000', '5', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000130000', '西藏自治区公安厅监所管理总队', '540000000000', '', '1', '2', '540000', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000130100', '西藏自治区公安厅监所管理总队自治区看守所', '540000130000', '', '1', '3', '540000', '904', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000210000', '西藏自治区公安厅禁毒总队', '540000000000', '', '1', '1', '540000', '21', '010000000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000580000', '民航西藏自治区管理局公安局', '540000000000', '', '1', '2', '540000', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000B00000', '西藏自治区公安厅边防总队', '540000000000', '', '1', '2', '540000', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000B00500', '西藏自治区公安边防总队司令部刑事案件侦查队', '540000000000', '', '1', '2', '540000', '42', '540000B00000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000B00600', '西藏自治区公安边防总队情报侦查支队', '540000000000', '', '1', '2', '540000', '42', '540000B00000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000B00601', '西藏自治区公安边防总队情报侦查支队机关侦查队', '540000B00000', '', '1', '3', '540000', '42', '540000B00600');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000B00602', '西藏自治区公安边防总队情报侦查支队拉萨情报站', '540000B00000', '', '1', '3', '540000', '42', '540000B00600');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000B00603', '西藏自治区公安边防总队情报侦查支队普兰情报站', '540000B00000', '', '1', '3', '540000', '42', '540000B00600');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000B00604', '西藏自治区公安边防总队情报侦查支队樟木情报站', '540000B00000', '', '1', '3', '540000', '42', '540000B00600');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000B00605', '西藏自治区公安边防总队情报侦查支队亚东情报站', '540000B00000', '', '1', '3', '540000', '42', '540000B00600');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000B00606', '西藏自治区公安边防总队情报侦查支队仲巴情报站', '540000B00000', '', '1', '3', '540000', '42', '540000B00600');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000B00607', '西藏自治区公安边防总队情报侦查支队定日情报站', '540000B00000', '', '1', '3', '540000', '42', '540000B00600');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000B00608', '西藏自治区公安边防总队情报侦查支队吉隆情报站', '540000B00000', '', '1', '3', '540000', '42', '540000B00600');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000H00000', '西藏自治区拉萨海关缉私局', '540000000000', '', '1', '2', '540000', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000H10000', '西藏自治区聂拉木海关缉私分局', '540000000000', '', '1', '2', '540000', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000H20000', '西藏自治区日喀则海关缉私分局', '540000000000', '', '1', '2', '540000', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000H30000', '西藏自治区狮泉河海关缉私分局', '540000000000', '', '1', '2', '540000', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000M00500', '民航西藏自治区管理局公安局刑侦大队', '540000000000', '', '1', '2', '540000', '5', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000S00000', '西藏自治区森林公安局', '540000000000', '', '1', '2', '540000', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000S00500', '西藏自治区森林公安局刑侦支队', '540000000000', '', '1', '2', '540000', '804', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540000T00000', '青藏铁路公安局拉萨铁路公安处', '540000000000', '', '1', '2', '540000', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100000000', '西藏自治区拉萨市公安局', '540000000000', '', '1', '2', '540100', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100025290', '西藏自治区戒毒管理局', '540000000000', '', '1', '2', '540000', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100027527', '西藏自治区堆龙德庆县强制隔离戒毒所', '540125000000', '', '1', '4', '540125', '42', '540125000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100030000', '西藏自治区拉萨市公安局治安管理支队', '540100000000', '', '1', '3', '540100', '3', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100050000', '西藏自治区拉萨市公安局刑警支队', '540100000000', '', '1', '3', '540100', '5', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100100000', '西藏自治区拉萨市公安局铁路治安管理支队', '540100000000', '', '1', '3', '540100', '3', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100120000', '青藏铁路公安局拉萨铁路公安处', '540100000000', '', '1', '3', '540000', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100130000', '西藏自治区拉萨市公安局监所管理支队', '540100000000', '', '1', '3', '540100', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100130100', '西藏自治区拉萨市公安局监所管理科看守所', '540100000000', '', '1', '3', '540100', '904', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100130200', '西藏自治区拉萨市公安局监所管理科治安行政拘留所', '540100000000', '', '1', '3', '540100', '903', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100210000', '西藏自治区拉萨市公安局禁毒支队', '540100000000', '', '1', '2', '540100', '21', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100930000', '西藏自治区拉萨市墨竹工卡县公安局日多检查站（一级）', '540127000000', '', '1', '4', '540100', '42', '540127000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100940000', '西藏自治区拉萨市曲水县公安局聂当检查站（一级）', '540124000000', '', '1', '4', '540100', '42', '540124000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100950000', '西藏自治区拉萨市尼木县公安局卡如一级检查站（一级）', '540123000000', '', '1', '4', '540100', '42', '540123000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100960000', '西藏自治区拉萨市当雄县公安局羊八井检查站（一级）', '540122000000', '', '1', '4', '540100', '42', '540122000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100970000', '西藏自治区拉萨市当雄县公安局乌玛塘检查站（一级）', '540122000000', '', '1', '4', '540100', '42', '540122000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100980000', '西藏自治区拉萨市柳梧新区公安局柳梧高速检查站（一级）', '540104000000', '', '1', '4', '540100', '42', '540104000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100990000', '西藏自治区拉萨市公安局城关分局拉林高速公路拉萨段一级公安检查站', '540102000000', '', '1', '4', '540100', '42', '540102000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100S00000', '西藏自治区拉萨市森林公安局', '540100000000', '', '1', '3', '540100', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540100S00500', '西藏自治区拉萨市森林公安局刑侦支队', '540100000000', '', '1', '3', '540100', '804', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540102000000', '西藏自治区拉萨市公安局城关分局', '540100000000', '', '1', '3', '540102', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540102030000', '西藏自治区拉萨市公安局城关分局治安科', '540102000000', '', '1', '4', '540102', '3', '540102000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540103000000', '西藏自治区拉萨市公安局经济技术开发区公安局', '540100000000', '', '1', '3', '540103', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540104000000', '西藏自治区拉萨市柳梧新区公安局', '540100000000', '', '1', '3', '540104', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540105000000', '西藏自治区拉萨市八廓古城公安局', '540100000000', '', '1', '3', '540105', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540106000000', '西藏自治区拉萨市空港新区公安局', '540100000000', '', '1', '3', '540106', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540106030000', '西藏自治区拉萨市空港新区分局治安管理大队', '540106000000', '', '1', '4', '540106', '3', '540106000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540121000000', '西藏自治区拉萨市林周县公安局', '540100000000', '', '1', '3', '540121', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540121030000', '西藏自治区拉萨市林周县公安局治安大队', '540121000000', '', '1', '4', '540121', '3', '540121210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540121130100', '西藏自治区拉萨市林周县公安局看守所', '540121000000', '', '1', '4', '540121', '904', '540121210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540121210000', '西藏自治区拉萨市林周县禁毒大队', '540100000000', '', '1', '3', '540121', '21', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540122000000', '西藏自治区拉萨市当雄县公安局', '540100000000', '', '1', '3', '540122', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540122030000', '西藏自治区拉萨市当雄县公安局治安管理大队', '540122000000', '', '1', '4', '540122', '3', '540122000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540122100000', '西藏自治区拉萨市当雄县公安局铁路治安巡警大队', '540122000000', '', '1', '4', '540122', '3', '540122000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540122130200', '西藏自治区拉萨市当雄县公安局看守所', '540122000000', '', '1', '4', '540122', '904', '540122000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540123000000', '西藏自治区拉萨市尼木县公安局', '540100000000', '', '1', '3', '540123', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540123030000', '西藏自治区拉萨市尼木县公安局治安管理大队', '540123000000', '', '1', '4', '540123', '3', '540123210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540123130200', '西藏自治区拉萨市尼木县公安局看守所', '540123000000', '', '1', '4', '540123', '904', '540123210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540123210000', '西藏自治区拉萨市尼木县公安局禁毒大队', '540123000000', '', '1', '3', '540123', '21', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540124000000', '西藏自治区拉萨市曲水县公安局', '540100000000', '', '1', '3', '540124', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540124030000', '西藏自治区拉萨市曲水县公安局治安管理大队', '540124000000', '', '1', '4', '540124', '3', '540124000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540124130200', '西藏自治区拉萨市曲水县公安局看守所', '540124000000', '', '1', '4', '540124', '904', '540124000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540125000000', '西藏自治区拉萨市堆龙德庆县公安局', '540100000000', '', '1', '3', '540125', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540125030000', '西藏自治区拉萨市堆龙德庆县公安局治安管理大队', '540125000000', '', '1', '4', '540125', '3', '540125000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540125100000', '西藏自治区拉萨市堆龙德庆县公安局铁路段治安巡警大队', '540125000000', '', '1', '4', '540125', '3', '540125000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540125130200', '西藏自治区拉萨市堆龙德庆县公安局看守所', '540125000000', '', '1', '4', '540125', '904', '540125000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540126000000', '西藏自治区拉萨市达孜县公安局', '540100000000', '', '1', '3', '540126', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540126030000', '西藏自治区拉萨市达孜县公安局治安管理大队', '540126000000', '', '1', '4', '540126', '3', '540126000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540126130200', '西藏自治区拉萨市达孜县公安局看守所', '540126000000', '', '1', '4', '540126', '904', '540126000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540127000000', '西藏自治区拉萨市墨竹工卡县公安局', '540100000000', '', '1', '3', '540127', '42', '540100210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540127030000', '西藏自治区拉萨市墨竹工卡县公安局治安管理大队', '540127000000', '', '1', '4', '540127', '3', '540127000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540127130100', '西藏自治区拉萨市墨竹工卡县公安局看守所', '540127000000', '', '1', '4', '540127', '904', '540127000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540200000000', '西藏自治区日喀则市公安局', '540000000000', '', '1', '2', '540200', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540200030000', '西藏自治区日喀则市公安局治安管理支队', '540200000000', '', '1', '3', '540200', '3', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540200050000', '西藏自治区日喀则市公安局刑事侦查支队', '540200000000', '', '1', '3', '542300', '5', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540200130000', '西藏自治区日喀则地市公安局监所管理支队', '540200000000', '', '1', '3', '542300', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540200130100', '西藏自治区日喀则市公安局地区看守所', '540200000000', '', '1', '3', '540200', '904', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540200130200', '西藏自治区日喀则市公安局行政拘留所', '540200000000', '', '1', '3', '540200', '903', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540200210000', '西藏自治区日喀则市公安局禁毒支队', '540200000000', '', '1', '2', '542300', '21', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540200950000', '西藏自治区日喀则市萨嘎县公安局22道班一级公安检查站', '540236000000', '', '1', '4', '540200', '42', '540236050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540200960000', '西藏自治区日喀则市康马县公安局少岗一级公安检查站', '540230000000', '', '1', '4', '540200', '42', '540230050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540200970000', '西藏自治区日喀则市仁布县公安局切娃一级公安检查站', '540229000000', '', '1', '4', '540200', '42', '540229050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540200980000', '西藏自治区日喀则市拉孜县公安局查务一级公安检查站', '540225000000', '', '1', '4', '540200', '42', '540225050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540200990000', '西藏自治区日喀则市江孜县公安局东郊公安一级检查站', '540222000000', '', '1', '4', '540200', '42', '540222050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540200M00000', '西藏自治区日喀则市和平机场公安分局', '540200000000', '', '1', '3', '540200', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540200S00000', '西藏自治区日喀则市森林公安局', '540200000000', '', '1', '3', '540200', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540202000000', '西藏自治区日喀则市桑珠孜区公安局', '540200000000', '', '1', '3', '540202', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540202030000', '西藏自治区日喀则市桑珠孜区公安局治安管理大队', '540202000000', '', '1', '4', '540202', '3', '540202210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540202050000', '西藏自治区日喀则市桑珠孜区公安局刑事侦查大队', '540202000000', '', '1', '4', '540202', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540202130100', '西藏自治区日喀则市桑珠孜区公安局看守所', '540202000000', '', '1', '4', '540202', '904', '540202210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540202210000', '西藏自治区日喀则市桑珠孜区公安局禁毒大队', '540202000000', '', '1', '3', '542300', '21', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540221000000', '西藏自治区日喀则市南木林县公安局', '540200000000', '', '1', '3', '540221', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540221030000', '西藏自治区日喀则市南木林县公安局治安管理大队', '540221000000', '', '1', '4', '540221', '3', '540221050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540221050000', '西藏自治区日喀则市南木林县公安局刑事侦查大队', '540221000000', '', '1', '3', '540221', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540221130100', '西藏自治区日喀则市南木林县公安局看守所', '540221000000', '', '1', '4', '540221', '904', '540221050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540222000000', '西藏自治区日喀则市江孜县公安局', '540200000000', '', '1', '3', '540222', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540222030000', '西藏自治区日喀则市江孜县公安局治安管理大队', '540222000000', '', '1', '4', '540222', '3', '540222050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540222050000', '西藏自治区日喀则市江孜县公安局刑事侦查大队', '540222000000', '', '1', '3', '540222', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540222130100', '西藏自治区日喀则市江孜县公安局看守所', '540222000000', '', '1', '4', '540222', '904', '540222050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540223000000', '西藏自治区日喀则市定日县公安局', '540200000000', '', '1', '3', '540223', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540223030000', '西藏自治区日喀则市定日县公安局治安管理大队', '540223000000', '', '1', '4', '540223', '3', '540223050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540223050000', '西藏自治区日喀则市定日县公安局刑事侦查大队', '540223000000', '', '1', '3', '540223', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540223130100', '西藏自治区日喀则市定日县公安局看守所', '540223000000', '', '1', '4', '540223', '904', '540223050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540223B00000', '中国人民武装警察部队西藏日喀则市定日县公安边防大队', '540223000000', '', '1', '4', '540223', '42', '540223050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540224000000', '西藏自治区日喀则市萨迦县公安局', '540200000000', '', '1', '3', '540224', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540224030000', '西藏自治区日喀则市萨迦县公安局治安管理大队', '540224000000', '', '1', '4', '540224', '3', '540224050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540224050000', '西藏自治区日喀则市萨迦县公安局刑事侦查大队', '540224000000', '', '1', '3', '540224', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540224130100', '西藏自治区日喀则市萨迦县公安局看守所', '540224000000', '', '1', '4', '540224', '904', '540224050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540225000000', '西藏自治区日喀则市拉孜县公安局', '540200000000', '', '1', '3', '540225', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540225030000', '西藏自治区日喀则市拉孜县公安局治安管理大队', '540225000000', '', '1', '4', '540225', '3', '540225050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540225050000', '西藏自治区日喀则市拉孜县公安局刑事侦查大队', '540225000000', '', '1', '3', '540225', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540225130100', '西藏自治区日喀则市拉孜县公安局看守所', '540225000000', '', '1', '4', '540225', '904', '540225050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540226000000', '西藏自治区日喀则市昂仁县公安局', '540200000000', '', '1', '3', '540226', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540226030000', '西藏自治区日喀则市昂仁县公安局治安管理大队', '540226000000', '', '1', '4', '540226', '3', '540226050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540226050000', '西藏自治区日喀则市昂仁县公安局刑事侦查大队', '540226000000', '', '1', '3', '540226', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540226130100', '西藏自治区日喀则市昂仁县公安局看守所', '540226000000', '', '1', '4', '540226', '904', '540226050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540227000000', '西藏自治区日喀则市谢通门县公安局', '540200000000', '', '1', '3', '540227', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540227030000', '西藏自治区日喀则市谢通门县公安局治安管理大队', '540227000000', '', '1', '4', '540227', '3', '540227050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540227050000', '西藏自治区日喀则市谢通门县公安局刑事侦查大队', '540227000000', '', '1', '3', '540227', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540227130100', '西藏自治区日喀则市谢通门县公安局看守所', '540227000000', '', '1', '4', '540227', '904', '540227050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540228000000', '西藏自治区日喀则市白朗县公安局', '540200000000', '', '1', '3', '540228', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540228030000', '西藏自治区日喀则市白朗县公安局治安管理大队', '540228000000', '', '1', '4', '540228', '3', '540228050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540228050000', '西藏自治区日喀则市白朗县公安局刑事侦查大队', '540228000000', '', '1', '3', '540228', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540228130100', '西藏自治区日喀则市白朗县公安局看守所', '540228000000', '', '1', '4', '540228', '904', '540228050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540229000000', '西藏自治区日喀则市仁布县公安局', '540200000000', '', '1', '3', '540229', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540229030000', '西藏自治区日喀则市仁布县公安局治安管理大队', '540229000000', '', '1', '4', '540229', '3', '540229050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540229050000', '西藏自治区日喀则市仁布县公安局刑事侦查大队', '540229000000', '', '1', '3', '540229', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540229130100', '西藏自治区日喀则市仁布县公安局看守所', '540229000000', '', '1', '4', '540229', '904', '540229050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540230000000', '西藏自治区日喀则市康马县公安局', '540200000000', '', '1', '3', '540230', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540230030000', '西藏自治区日喀则市康马县公安局治安管理大队', '540230000000', '', '1', '4', '540230', '3', '540230050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540230050000', '西藏自治区日喀则市康马县公安局刑事侦查大队', '540230000000', '', '1', '3', '540230', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540230130100', '西藏自治区日喀则市康马县公安局看守所', '540230000000', '', '1', '4', '540230', '904', '540230050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540230B00000', '中国人民武装警察部队西藏日喀则市康马县公安边防大队', '540230000000', '', '1', '4', '540230', '42', '540230050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540231000000', '西藏自治区日喀则市定结县公安局', '540200000000', '', '1', '3', '540231', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540231030000', '西藏自治区日喀则市定结县公安局治安管理大队', '540231000000', '', '1', '4', '540231', '3', '540231050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540231050000', '西藏自治区日喀则市定结县公安局刑事侦查大队', '540231000000', '', '1', '3', '540231', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540231130100', '西藏自治区日喀则市定结县公安局看守所', '540231000000', '', '1', '4', '540231', '904', '540231050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540231B00000', '中国人民武装警察部队西藏日喀则市定结县公安边防大队', '540231000000', '', '1', '4', '540231', '42', '540231050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540232000000', '西藏自治区日喀则市仲巴县公安局', '540200000000', '', '1', '3', '540232', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540232030000', '西藏自治区日喀则市仲巴县公安局治安管理大队', '540232000000', '', '1', '4', '540232', '3', '540232050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540232050000', '西藏自治区日喀则市仲巴县公安局刑事侦查大队', '540232000000', '', '1', '3', '540232', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540232130100', '西藏自治区日喀则市仲巴县公安局看守所', '540232000000', '', '1', '4', '540232', '904', '540232050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540232B00000', '中国人民武装警察部队西藏日喀则市仲巴县公安边防大队', '540232000000', '', '1', '4', '540232', '42', '540232050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540233000000', '西藏自治区日喀则市亚东县公安局', '540200000000', '', '1', '3', '540233', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540233030000', '西藏自治区日喀则市亚东县公安局治安管理大队', '540233000000', '', '1', '4', '540233', '3', '540233050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540233050000', '西藏自治区日喀则市亚东县公安局刑事侦查大队', '540233000000', '', '1', '3', '540233', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540233130100', '西藏自治区日喀则市亚东县公安局看守所', '540233000000', '', '1', '4', '540233', '904', '540233050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540233B00000', '中国人民武装警察部队西藏日喀则市亚东县公安边防大队', '540233000000', '', '1', '4', '540233', '42', '540233050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540233S00000', '西藏自治区日喀则市亚东县森林公安局', '540233000000', '', '1', '4', '540233', '42', '540233050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540233S00500', '西藏自治区日喀则市亚东县森林公安局刑侦大队', '540233000000', '', '1', '4', '540233', '804', '540233050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540234000000', '西藏自治区日喀则市吉隆县公安局', '540200000000', '', '1', '3', '540234', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540234030000', '西藏自治区日喀则市吉隆县公安局治安管理大队', '540234000000', '', '1', '4', '540234', '3', '540234050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540234050000', '西藏自治区日喀则市吉隆县公安局刑事侦查大队', '540234000000', '', '1', '4', '542335', '5', '540234050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540234130100', '西藏自治区日喀则市吉隆县公安局看守所', '540234000000', '', '1', '4', '540234', '904', '540234050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540234210000', '西藏自治区日喀则市吉隆县公安局禁毒大队', '540234000000', '', '1', '4', '540234', '21', '540234050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540234510000', '西藏自治区日喀则市吉隆县公安局吉隆口岸公安分局', '540234000000', '', '1', '4', '540234', '42', '540234050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540234B00000', '中国人民武装警察部队西藏日喀则市吉隆县公安边防大队', '540234000000', '', '1', '4', '540234', '42', '540234050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540234S00000', '西藏自治区日喀则市吉隆县森林公安局', '540234000000', '', '1', '4', '540234', '42', '540234050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540234S00500', '西藏自治区日喀则市吉隆县森林公安局刑侦大队', '540234000000', '', '1', '4', '540234', '804', '540234050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540235000000', '西藏自治区日喀则市聂拉木县公安局', '540200000000', '', '1', '3', '540235', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540235030000', '西藏自治区日喀则市聂拉木县公安局治安管理大队', '540235000000', '', '1', '4', '540235', '3', '540235210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540235050000', '西藏自治区日喀则市聂拉木县公安局刑事侦查大队', '540235000000', '', '1', '4', '540235', '5', '540235210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540235130100', '西藏自治区日喀则市聂拉木县公安局看守所', '540235000000', '', '1', '4', '540235', '904', '540235210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540235210000', '西藏自治区日喀则市聂拉木县公安局禁毒大队', '540235000000', '', '1', '3', '542336', '21', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540235510000', '西藏自治区日喀则市聂拉木县公安局聂拉木口岸公安分局', '540235000000', '', '1', '4', '540235', '42', '540235210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540235S00000', '西藏自治区日喀则市聂拉木县森林公安局', '540235000000', '', '1', '4', '540235', '42', '540235210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540235S00500', '西藏自治区日喀则市聂拉木县森林公安局刑侦大队', '540235000000', '', '1', '4', '540235', '804', '540235210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540236000000', '西藏自治区日喀则市萨嘎县公安局', '540200000000', '', '1', '3', '540236', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540236030000', '西藏自治区日喀则市萨嘎县公安局治安管理大队', '540236000000', '', '1', '4', '540236', '3', '540236050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540236050000', '西藏自治区日喀则市萨嘎县公安局刑事侦查大队', '540236000000', '', '1', '3', '540236', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540236130100', '西藏自治区日喀则市萨嘎县公安局看守所', '540236000000', '', '1', '4', '540236', '904', '540236050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540236B00000', '中国人民武装警察部队西藏日喀则市萨嘎县公安边防大队', '540236000000', '', '1', '4', '540236', '42', '540236050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540237000000', '西藏自治区日喀则市岗巴县公安局', '540200000000', '', '1', '3', '540237', '42', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540237030000', '西藏自治区日喀则市岗巴县公安局治安管理大队', '540237000000', '', '1', '4', '540237', '3', '540237050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540237050000', '西藏自治区日喀则市岗巴县公安局刑事侦查大队', '540237000000', '', '1', '3', '540237', '5', '540200210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540237130100', '西藏自治区日喀则市岗巴县公安局看守所', '540237000000', '', '1', '4', '540237', '904', '540237050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540237B00000', '中国人民武装警察部队西藏日喀则市岗巴县公安边防大队', '540237000000', '', '1', '4', '540237', '42', '540237050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300000000', '西藏自治区昌都市公安局', '540000000000', '', '1', '2', '540300', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300000247', '西藏昌都强制隔离戒毒所(西藏昌都社会教育矫治所)', '540300000000', '', '1', '3', '542121', '42', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300030000', '西藏自治区昌都市公安局治安管理支队', '540300000000', '', '1', '3', '540300', '3', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300050000', '西藏自治区昌都市公安局刑警支队', '540300000000', '', '1', '3', '542100', '5', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300130100', '西藏自治区昌都市公安局看守所', '540300000000', '', '1', '3', '540300', '904', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300210000', '西藏自治区昌都市公安局禁毒支队', '540300000000', '', '1', '2', '542100', '21', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300510000', '西藏自治区昌都市公安局昌都新区公安分局', '540300000000', '', '1', '3', '540300', '42', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300890000', '西藏自治区昌都市边坝县公安局尼木一级公安检查站', '540330000000', '', '1', '4', '540300', '42', '540330210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300900000', '西藏自治区昌都市芒康县公安局盐井一级公安检查站', '540328000000', '', '1', '4', '540300', '42', '540328210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300910000', '西藏自治区昌都市芒康县公安局朱巴龙一级公安检查站', '540328000000', '', '1', '4', '540300', '42', '540328210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300920000', '西藏自治区昌都市八宿县公安局然乌一级公安检查站', '540326000000', '', '1', '4', '540300', '42', '540326210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300930000', '西藏自治区昌都市丁青县公安局布塔一级公安检查站', '540324000000', '', '1', '4', '540300', '42', '540324210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300940000', '西藏自治区昌都市丁青县公安局巴达一级公安检查站', '540324000000', '', '1', '4', '540300', '42', '540324210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300950000', '西藏自治区昌都市类乌齐县公安局甲桑卡一级公安检查站', '540323000000', '', '1', '4', '540300', '42', '540323210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300960000', '西藏自治区昌都市江达县公安局生达一级公安检查站', '540321000000', '', '1', '4', '540300', '42', '540321210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300970000', '西藏自治区昌都市江达县公安局邓柯友谊桥一级公安检查站', '540321000000', '', '1', '4', '540300', '42', '540321210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300980000', '西藏自治区昌都市江达县公安局金沙江大桥一级公安检查站', '540321000000', '', '1', '4', '540300', '42', '540321210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300990000', '西藏自治区昌都市公安局卡若区分局缅达一级公安检查站', '540302000000', '', '1', '4', '540300', '42', '540302210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300M00000', '西藏自治区昌都市邦达机场公安分局', '540300000000', '', '1', '3', '540300', '42', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300S00000', '西藏自治区昌都市森林公安局', '540300000000', '', '1', '3', '540300', '42', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540300S00500', '西藏自治区昌都市森林公安局刑侦支队', '540300000000', '', '1', '3', '540300', '804', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540302000000', '西藏自治区昌都市公安局卡若区分局', '540300000000', '', '1', '3', '540302', '42', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540302030000', '西藏自治区昌都市公安局卡若区分局治安管理大队', '540302000000', '', '1', '4', '540302', '3', '540302210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540302050000', '西藏自治区昌都市公安局卡若区分局刑事侦查大队', '540302000000', '', '1', '4', '542121', '5', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540302210000', '西藏自治区昌都市公安局卡若区分局禁毒大队', '540302000000', '', '1', '3', '542121', '21', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540302S00000', '西藏自治区昌都市公安局卡若区森林公安局', '540302000000', '', '1', '4', '540302', '42', '540302210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540302S00500', '西藏自治区昌都市公安局卡若区森林公安局刑侦大队', '540302000000', '', '1', '4', '540302', '804', '540302210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540321000000', '西藏自治区昌都市江达县公安局', '540300000000', '', '1', '3', '540321', '42', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540321030000', '西藏自治区昌都市江达县公安局治安管理大队', '540321000000', '', '1', '4', '540321', '3', '540321210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540321070000', '西藏自治区昌都市江达县公安局刑事侦查大队', '540321000000', '', '1', '4', '542122', '5', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540321130100', '西藏自治区昌都市江达县公安局看守所', '540321000000', '', '1', '4', '540321', '904', '540321210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540321210000', '西藏自治区昌都市江达县公安局禁毒大队', '540321000000', '', '1', '3', '542122', '21', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540321S00000', '西藏自治区昌都市江达县森林公安局', '540321000000', '', '1', '4', '540321', '42', '540321210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540321S00500', '西藏自治区昌都市江达县森林公安局刑侦大队', '540321000000', '', '1', '4', '540321', '804', '540321210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540322000000', '西藏自治区昌都市贡觉县公安局', '540300000000', '', '1', '3', '540322', '42', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540322030000', '西藏自治区昌都市贡觉县公安局治安管理大队', '540322000000', '', '1', '4', '540322', '3', '540322210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540322050000', '西藏自治区昌都市贡觉县公安局刑事侦查大队', '540322000000', '', '1', '4', '542123', '5', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540322130100', '西藏自治区昌都市贡觉县公安局看守所', '540322000000', '', '1', '4', '540322', '904', '540322210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540322210000', '西藏自治区昌都市贡觉县公安局禁毒大队', '540322000000', '', '1', '3', '542123', '21', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540322570000', '西藏自治区昌都市贡觉县公安局三岩片区分局', '540322000000', '', '1', '4', '540322', '42', '540322210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540322S00000', '西藏自治区昌都市贡觉县森林公安局', '540322000000', '', '1', '4', '540322', '42', '540322210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540322S00500', '西藏自治区昌都市贡觉县森林公安局刑侦大队', '540322000000', '', '1', '4', '540322', '804', '540322210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540323000000', '西藏自治区昌都市类乌齐县公安局', '540300000000', '', '1', '3', '540323', '42', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540323030000', '西藏自治区昌都市类乌齐县公安局治安管理大队', '540323000000', '', '1', '4', '540323', '3', '540323210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540323050000', '西藏自治区昌都市类乌齐县公安局刑事侦查大队', '540323000000', '', '1', '4', '542124', '5', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540323130100', '西藏自治区昌都市类乌齐县公安局看守所', '540323000000', '', '1', '4', '540323', '904', '540323210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540323210000', '西藏自治区昌都市类乌齐县公安局禁毒大队', '540323000000', '', '1', '3', '542124', '21', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540323S00000', '西藏自治区昌都市类乌齐县森林公安局', '540323000000', '', '1', '4', '540323', '42', '540323210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540323S00500', '西藏自治区昌都市类乌齐县森林公安局刑侦大队', '540323000000', '', '1', '4', '540323', '804', '540323210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540324000000', '西藏自治区昌都市丁青县公安局', '540300000000', '', '1', '3', '540324', '42', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540324030000', '西藏自治区昌都市丁青县公安局治安管理大队', '540324000000', '', '1', '4', '540324', '3', '540324210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540324050000', '西藏自治区昌都市丁青县公安局刑事侦查大队', '540324000000', '', '1', '4', '542125', '5', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540324130100', '西藏自治区昌都市丁青县公安局看守所', '540324000000', '', '1', '4', '540324', '904', '540324210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540324210000', '西藏自治区昌都市丁青县公安局禁毒大队', '540324000000', '', '1', '3', '542125', '21', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540324S00000', '西藏自治区昌都市丁青县森林公安局', '540324000000', '', '1', '4', '540324', '42', '540324210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540324S00500', '西藏自治区昌都市丁青县森林公安局刑侦大队', '540324000000', '', '1', '4', '540324', '804', '540324210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540325010000', '西藏自治区昌都市察雅县公安局', '540300000000', '', '1', '3', '540325', '42', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540325030000', '西藏自治区昌都市察雅县公安局治安管理大队', '540325010000', '', '1', '4', '540325', '3', '540325210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540325050000', '西藏自治区昌都市察雅县公安局刑事侦查大队', '540325010000', '', '1', '4', '542126', '5', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540325130100', '西藏自治区昌都市察雅县公安局看守所', '540325010000', '', '1', '4', '540325', '904', '540325210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540325210000', '西藏自治区昌都市察雅县公安局禁毒大队', '540325010000', '', '1', '3', '542126', '21', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540325S00000', '西藏自治区昌都市察雅县森林公安局', '540325010000', '', '1', '4', '540325', '42', '540325210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540325S00500', '西藏自治区昌都市察雅县森林公安局刑侦大队', '540325010000', '', '1', '4', '540325', '804', '540325210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540326000000', '西藏自治区昌都市八宿县公安局', '540300000000', '', '1', '3', '540326', '42', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540326030000', '西藏自治区昌都市八宿县公安局治安管理大队', '540326000000', '', '1', '4', '540326', '3', '540326210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540326050000', '西藏自治区昌都市八宿县公安局刑事侦查大队', '540326000000', '', '1', '4', '542127', '5', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540326130100', '西藏自治区昌都市八宿县公安局看守所', '540326000000', '', '1', '4', '540326', '904', '540326210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540326210000', '西藏自治区昌都市八宿县公安局禁毒大队', '540326000000', '', '1', '3', '542127', '21', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540326S00000', '西藏自治区昌都市八宿县森林公安局', '540326000000', '', '1', '4', '540326', '42', '540326210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540326S00500', '西藏自治区昌都市八宿县森林公安局刑侦大队', '540326000000', '', '1', '4', '540326', '804', '540326210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540327000000', '西藏自治区昌都市左贡县公安局', '540300000000', '', '1', '3', '540327', '42', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540327030000', '西藏自治区昌都市左贡县公安局治安管理大队', '540327000000', '', '1', '4', '540327', '3', '540327210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540327050000', '西藏自治区昌都市左贡县公安局刑事侦查大队', '540327000000', '', '1', '4', '542128', '5', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540327130100', '西藏自治区昌都市左贡县公安局看守所', '540327000000', '', '1', '4', '540327', '904', '540327210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540327210000', '西藏自治区昌都市左贡县公安局禁毒大队', '540327000000', '', '1', '3', '542128', '21', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540327S00000', '西藏自治区昌都市左贡县森林公安局', '540327000000', '', '1', '4', '540327', '42', '540327210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540327S00500', '西藏自治区昌都市左贡县森林公安局刑侦大队', '540327000000', '', '1', '4', '540327', '804', '540327210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540328000000', '西藏自治区昌都市芒康县公安局', '540300000000', '', '1', '3', '540328', '42', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540328030000', '西藏自治区昌都市芒康县公安局治安管理大队', '540328000000', '', '1', '4', '540328', '3', '540328210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540328050000', '西藏自治区昌都市芒康县公安局刑事侦查大队', '540328000000', '', '1', '4', '542129', '5', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540328060000', '西藏自治区昌都市芒康县公安局刑事侦查大队', '540328000000', '', '1', '4', '542129', '5', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540328070000', '西藏自治区昌都市芒康县公安局刑事侦查大队', '540328000000', '', '1', '4', '542129', '5', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540328130100', '西藏自治区昌都市芒康县公安局看守所', '540328000000', '', '1', '4', '540328', '904', '540328210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540328210000', '西藏自治区昌都市芒康县公安局禁毒大队', '540328000000', '', '1', '3', '542129', '21', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540328S00000', '西藏自治区昌都市芒康县森林公安局', '540328000000', '', '1', '4', '540328', '42', '540328210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540328S00500', '西藏自治区昌都市芒康县森林公安局刑侦大队', '540328000000', '', '1', '4', '540328', '804', '540328210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540329000000', '西藏自治区昌都市洛隆县公安局', '540300000000', '', '1', '3', '540329', '42', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540329030000', '西藏自治区昌都市洛隆县公安局治安管理大队', '540329000000', '', '1', '4', '540329', '3', '540329210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540329050000', '西藏自治区昌都市洛隆县公安局刑事侦查大队', '540329000000', '', '1', '4', '542132', '5', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540329130100', '西藏自治区昌都市洛隆县公安局看守所', '540329000000', '', '1', '4', '540329', '904', '540329210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540329210000', '西藏自治区昌都市洛隆县公安局禁毒大队', '540329000000', '', '1', '3', '542132', '21', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540329S00000', '西藏自治区昌都市洛隆县森林公安局', '540329000000', '', '1', '4', '540329', '42', '540329210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540329S00500', '西藏自治区昌都市洛隆县森林公安局刑侦大队', '540329000000', '', '1', '4', '540329', '804', '540329210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540330000000', '西藏自治区昌都市边坝县公安局', '540300000000', '', '1', '3', '540330', '42', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540330030000', '西藏自治区昌都市边坝县公安局治安管理大队', '540330000000', '', '1', '4', '540330', '3', '540330210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540330050000', '西藏自治区昌都市边坝县公安局刑事侦查大队', '540330000000', '', '1', '4', '542133', '5', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540330130100', '西藏自治区昌都市边坝县公安局看守所', '540330000000', '', '1', '4', '540330', '904', '540330210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540330210000', '西藏自治区昌都市边坝县公安局禁毒大队', '540330000000', '', '1', '3', '542133', '21', '540300210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540330S00000', '西藏自治区昌都市边坝县森林公安局', '540330000000', '', '1', '4', '540330', '42', '540330210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540330S00500', '西藏自治区昌都市边坝县森林公安局刑侦大队', '540330000000', '', '1', '4', '540330', '804', '540330210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540400000000', '西藏自治区林芝市公安局', '540000000000', '', '1', '2', '540400', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540400030000', '西藏自治区林芝市公安局治安管理支队', '540400000000', '', '1', '3', '540400', '3', '540400210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540400050000', '西藏自治区林芝市公安局刑事侦查支队', '540400000000', '', '1', '3', '542600', '5', '540400210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540400130100', '西藏自治区林芝市公安局看守所', '540400000000', '', '1', '3', '540400', '904', '540400210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540400130200', '西藏自治区林芝市公安局拘留所', '540400000000', '', '1', '3', '540400', '903', '540400210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540400210000', '西藏自治区林芝市公安局禁毒支队', '540400000000', '', '1', '2', '542600', '21', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540400980000', '西藏自治区林芝市波密县公安局玉普一级公安检查站', '540424000000', '', '1', '4', '540400', '42', '540423210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540400990000', '西藏自治区林芝市察隅县公安局察瓦龙一级公安检查站', '540425000000', '', '1', '4', '540400', '42', '540425210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540400M00000', '西藏自治区林芝市米林机场公安分局', '540400000000', '', '1', '3', '540400', '42', '540400210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540400S00000', '西藏自治区林芝市森林公安局', '540400000000', '', '1', '3', '540400', '42', '540400210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540400S00500', '西藏自治区林芝市森林公安局刑侦支队', '540400000000', '', '1', '3', '540400', '804', '540400210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540402000000', '西藏自治区林芝市巴宜区公安局', '540400000000', '', '1', '3', '540402', '42', '540400210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540402050000', '西藏自治区林芝市巴宜区公安局刑事侦查大队', '540402000000', '', '1', '4', '540402', '5', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540402130100', '西藏自治区林芝市巴宜区公安局看守所', '540402000000', '', '1', '4', '540402', '904', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540402210000', '西藏自治区林芝市巴宜区公安局禁毒大队', '540402000000', '', '1', '3', '542621', '21', '540400210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540402S00000', '西藏自治区林芝市巴宜区森林公安局', '540402000000', '', '1', '4', '540402', '42', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540402S00500', '西藏自治区林芝市巴宜区森林公安局刑侦大队', '540402000000', '', '1', '4', '540402', '804', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540403030000', '西藏自治区林芝市巴宜区公安局治安管理大队', '540402000000', '', '1', '4', '540402', '3', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540421000000', '西藏自治区林芝市工布江达县公安局', '540400000000', '', '1', '3', '540421', '42', '540400210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540421030000', '西藏自治区林芝市工布江达县公安局治安管理大队', '540421000000', '', '1', '4', '540421', '3', '540421210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540421050000', '西藏自治区林芝市工布江达县公安局刑事侦查大队', '540421000000', '', '1', '4', '540421', '5', '540421210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540421130100', '西藏自治区林芝市工布江达县公安局看守所', '540421000000', '', '1', '4', '540421', '904', '540421210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540421210000', '西藏自治区林芝市工布江达县公安局禁毒大队', '540421000000', '', '1', '3', '542622', '21', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540421S00000', '西藏自治区林芝市工布江达县森林公安局', '540421000000', '', '1', '4', '540421', '42', '540421210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540421S00500', '西藏自治区林芝市工布江达县森林公安局刑侦大队', '540421000000', '', '1', '4', '540421', '804', '540421210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540422000000', '西藏自治区林芝市米林县公安局', '540400000000', '', '1', '3', '540422', '42', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540422030000', '西藏自治区林芝市米林县公安局治安管理大队', '540422000000', '', '1', '4', '540422', '3', '540422210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540422050000', '西藏自治区林芝市米林县公安局刑事侦查大队', '540422000000', '', '1', '4', '540422', '5', '540422210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540422130100', '西藏自治区林芝市米林县公安局看守所', '540422000000', '', '1', '4', '540422', '904', '540422210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540422210000', '西藏自治区林芝市米林县公安局禁毒大队', '540422000000', '', '1', '3', '542623', '21', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540422B00000', '中国人民武装警察部队西藏自治区林芝市米林县公安边防大队', '540422000000', '', '1', '4', '540422', '42', '540422210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540422S00000', '西藏自治区林芝市米林县森林公安局', '540422000000', '', '1', '4', '540422', '42', '540422210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540422S00500', '西藏自治区林芝市米林县森林公安局刑侦大队', '540422000000', '', '1', '4', '540422', '804', '540422210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540423000000', '西藏自治区林芝市墨脱县公安局', '540400000000', '', '1', '3', '540423', '42', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540423030000', '西藏自治区林芝市墨脱县公安局治安管理大队', '540423000000', '', '1', '4', '540423', '3', '540423210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540423050000', '西藏自治区林芝市墨脱县公安局刑事侦查大队', '540423000000', '', '1', '4', '540423', '5', '540423210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540423130100', '西藏自治区林芝市墨脱县公安局看守所', '540423000000', '', '1', '4', '540423', '904', '540423210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540423210000', '西藏自治区林芝市墨脱县公安局禁毒支队', '540423000000', '', '1', '3', '542624', '21', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540423B00000', '中国人民武装警察部队西藏自治区林芝市墨脱县公安局边防大队', '540423000000', '', '1', '4', '540423', '42', '540423210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540423S00000', '西藏自治区林芝市墨脱县森林公安局', '540423000000', '', '1', '4', '540423', '42', '540423210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540423S00500', '西藏自治区林芝市墨脱县森林公安局刑侦大队', '540423000000', '', '1', '4', '540423', '804', '540423210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540424000000', '西藏自治区林芝市波密县公安局', '540400000000', '', '1', '3', '540424', '42', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540424030000', '西藏自治区林芝市波密县公安局治安管理大队', '540424000000', '', '1', '4', '540424', '3', '540424210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540424050000', '西藏自治区林芝市波密县公安局刑事侦查大队', '540424000000', '', '1', '4', '540424', '5', '540424210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540424130100', '西藏自治区林芝市波密县公安局看守所', '540424000000', '', '1', '4', '540424', '904', '540424210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540424210000', '西藏自治区林芝市波密县公安局禁毒大队', '540424000000', '', '1', '3', '542625', '21', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540424S00000', '西藏自治区林芝市波密县森林公安局', '540424000000', '', '1', '4', '540424', '42', '540424210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540424S00500', '西藏自治区林芝市波密县森林公安局刑侦大队', '540424000000', '', '1', '4', '540424', '804', '540424210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540425000000', '西藏自治区林芝市察隅县公安局', '540400000000', '', '1', '3', '540425', '42', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540425030000', '西藏自治区林芝市察隅县公安局治安管理大队', '540425000000', '', '1', '4', '540425', '3', '540425210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540425050000', '西藏自治区林芝市察隅县公安局刑事侦查大队', '540425000000', '', '1', '4', '540425', '5', '540425210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540425130100', '西藏自治区林芝市察隅县公安局看守所', '540425000000', '', '1', '4', '540425', '904', '540425210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540425210000', '西藏自治区林芝市察隅县公安局禁毒大队', '540425000000', '', '1', '3', '542626', '21', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540425B00000', '中国人民武装警察部队西藏自治区林芝市察隅县公安边防大队', '540425000000', '', '1', '4', '540425', '42', '540425210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540425S00000', '西藏自治区林芝市察隅县森林公安局', '540425000000', '', '1', '4', '540425', '42', '540425210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540425S00500', '西藏自治区林芝市察隅县森林公安局刑侦大队', '540425000000', '', '1', '4', '540425', '804', '540425210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540426000000', '西藏自治区林芝市朗县公安局', '540400000000', '', '1', '3', '540426', '42', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540426050000', '西藏自治区林芝市朗县公安局刑事侦查大队', '540426000000', '', '1', '4', '540426', '5', '540426210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540426130100', '西藏自治区林芝市朗县公安局看守所', '540426000000', '', '1', '4', '540426', '904', '540426210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540426210000', '西藏自治区林芝市朗县公安局禁毒大队', '540426000000', '', '1', '3', '542627', '21', '540402210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540426B00000', '中国人民武装警察部队西藏自治区林芝市白朗县公安边防大队', '540228000000', '', '1', '4', '540426', '42', '540228050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540426S00000', '西藏自治区林芝市朗县森林公安局', '540426000000', '', '1', '4', '540426', '42', '540426210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540426S00500', '西藏自治区林芝市朗县森林公安局刑侦大队', '540426000000', '', '1', '4', '540426', '804', '540426210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540500000000', '西藏自治区山南市公安局', '540000000000', '', '1', '2', '540500', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540500030000', '西藏自治区山南市公安局治安管理支队', '540500000000', '', '1', '3', '540500', '3', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540500050000', '西藏自治区山南市公安局刑事侦查支队', '540500000000', '', '1', '3', '542200', '5', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540500130100', '西藏自治区山南市公安局看守所', '540500000000', '', '1', '3', '540500', '904', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540500210000', '西藏自治区山南市公安局禁毒支队', '540500000000', '', '1', '2', '542200', '21', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540500990000', '西藏自治区山南市公安局迎宾亭一级公安检查站', '540500000000', '', '1', '3', '540500', '42', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540500S00000', '西藏自治区山南市森林公安局', '540500000000', '', '1', '3', '540500', '42', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540500S00500', '西藏自治区山南市森林公安局刑侦支队', '540500000000', '', '1', '3', '540500', '804', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540502000000', '西藏自治区山南市乃东区公安局', '540500000000', '', '1', '3', '540502', '42', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540502030000', '西藏自治区山南市乃东区公安局治安管理大队', '540502000000', '', '1', '4', '540502', '3', '540502050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540502050000', '西藏自治区山南市乃东区公安局刑事侦查大队', '540502000000', '', '1', '3', '540502', '5', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540502130100', '西藏自治区山南市乃东区公安局看守所', '540502000000', '', '1', '4', '540502', '904', '540502050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540521000000', '西藏自治区山南市扎囊县公安局', '540500000000', '', '1', '3', '540521', '42', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540521030000', '西藏自治区山南市扎囊县公安局治安管理大队', '540521000000', '', '1', '4', '540521', '3', '540521050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540521050000', '西藏自治区山南市扎囊县公安局刑事侦查大队', '540521000000', '', '1', '3', '540521', '5', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540521130100', '西藏自治区山南市扎囊县公安局看守所', '540521000000', '', '1', '4', '540521', '904', '540521050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540522000000', '西藏自治区山南市贡嘎县公安局', '540500000000', '', '1', '3', '540522', '42', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540522030000', '西藏自治区山南市贡嘎县公安局治安管理大队', '540522000000', '', '1', '4', '540522', '3', '540522050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540522050000', '西藏自治区山南市贡嘎县公安局刑事侦查大队', '540522000000', '', '1', '3', '540522', '5', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540522130100', '西藏自治区山南市贡嘎县公安局看守所', '540522000000', '', '1', '4', '540522', '904', '540522050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540523000000', '西藏自治区山南市桑日县公安局', '540500000000', '', '1', '3', '540523', '42', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540523030000', '西藏自治区山南市桑日县公安局治安管理大队', '540523000000', '', '1', '4', '540523', '3', '540523050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540523050000', '西藏自治区山南市桑日县公安局刑事侦查大队', '540523000000', '', '1', '3', '540523', '5', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540523130100', '西藏自治区山南市桑日县公安局看守所', '540523000000', '', '1', '4', '540523', '904', '540523050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540524000000', '西藏自治区山南市琼结县公安局', '540500000000', '', '1', '3', '540524', '42', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540524030000', '西藏自治区山南市琼结县公安局治安管理大队', '540524000000', '', '1', '4', '540524', '3', '540524050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540524050000', '西藏自治区山南市琼结县公安局刑事侦查大队', '540524000000', '', '1', '3', '540524', '5', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540524130100', '西藏自治区山南市琼结县公安局看守所', '540524000000', '', '1', '4', '540524', '904', '540524050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540525000000', '西藏自治区山南市曲松县公安局', '540500000000', '', '1', '3', '540525', '42', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540525030000', '西藏自治区山南市曲松县公安局治安管理大队', '540525000000', '', '1', '4', '540525', '3', '540525050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540525050000', '西藏自治区山南市曲松县公安局刑事侦查大队', '540525000000', '', '1', '3', '540525', '5', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540525130100', '西藏自治区山南市曲松县公安局看守所', '540525000000', '', '1', '4', '540525', '904', '540525050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540526000000', '西藏自治区山南市措美县公安局', '540500000000', '', '1', '3', '540526', '42', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540526030000', '西藏自治区山南市措美县公安局治安管理大队', '540526000000', '', '1', '4', '540526', '3', '540526050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540526050000', '西藏自治区山南市措美县公安局刑事侦查大队', '540526000000', '', '1', '3', '540526', '5', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540526130100', '西藏自治区山南市措美县公安局看守所', '540526000000', '', '1', '4', '540526', '904', '540526050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540527000000', '西藏自治区山南市洛扎县公安局', '540500000000', '', '1', '3', '540527', '42', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540527030000', '西藏自治区山南市洛扎县公安局治安管理大队', '540527000000', '', '1', '4', '540527', '3', '540527050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540527050000', '西藏自治区山南市洛扎县公安局刑事侦查大队', '540527000000', '', '1', '3', '540527', '5', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540527130100', '西藏自治区山南市洛扎县公安局看守所', '540527000000', '', '1', '4', '540527', '904', '540527050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540527B00000', '西藏自治区山南市洛扎县公安边防大队', '540527000000', '', '1', '4', '540527', '42', '540527050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540528010000', '西藏自治区山南市加查县公安局', '540500000000', '', '1', '3', '540528', '42', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540528030000', '西藏自治区山南市加查县公安局治安管理大队', '540528010000', '', '1', '4', '540528', '3', '540528050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540528050000', '西藏自治区山南市加查县公安局刑事侦查大队', '540528010000', '', '1', '3', '540528', '5', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540528130100', '西藏自治区山南市加查县公安局看守所', '540528010000', '', '1', '4', '540528', '904', '540528050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540528S00000', '西藏自治区山南市加查县森林公安局', '540528010000', '', '1', '4', '540528', '42', '540528050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540528S00500', '西藏自治区山南市加查县森林公安局刑侦大队', '540528010000', '', '1', '4', '540528', '804', '540528050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540529000000', '西藏自治区山南市隆子县公安局', '540500000000', '', '1', '3', '540529', '42', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540529050000', '西藏自治区山南市隆子县公安局刑事侦查大队', '540529000000', '', '1', '3', '540529', '5', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540529130100', '西藏自治区山南市隆子县公安局看守所', '540529000000', '', '1', '4', '540529', '904', '540529050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540529B00000', '中国人民武装警察部队西藏自治区山南市隆子县公安边防大队', '540529000000', '', '1', '4', '540529', '42', '540529050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540530000000', '西藏自治区山南市错那县公安局', '540500000000', '', '1', '3', '540530', '42', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540530030000', '西藏自治区山南市错那县公安局治安管理大队', '540530000000', '', '1', '4', '540530', '3', '540530050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540530050000', '西藏自治区山南市错那县公安局刑事侦查大队', '540530000000', '', '1', '3', '540530', '5', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540530130100', '西藏自治区山南市错那县公安局看守所', '540530000000', '', '1', '4', '540530', '904', '540530050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540530B00000', '中国人民武装警察部队西藏自治区山南市错那县公安边防大队', '540530000000', '', '1', '4', '540430', '42', '540530050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540530S00000', '西藏自治区山南市错那县森林公安局', '540530000000', '', '1', '4', '540530', '42', '540530050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540530S00500', '西藏自治区山南市错那县森林公安局刑侦大队', '540530000000', '', '1', '4', '540530', '804', '540530050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540531000000', '西藏自治区山南市浪卡子县公安局', '540500000000', '', '1', '3', '540531', '42', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540531030000', '西藏自治区山南市浪卡子县公安局治安管理大队', '540531000000', '', '1', '4', '540531', '3', '540531050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540531050000', '西藏自治区山南市浪卡子县公安局刑事侦查大队', '540531000000', '', '1', '3', '540531', '5', '540500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('540531130100', '西藏自治区山南市浪卡子县公安局看守所（拘留所）', '540531000000', '', '1', '4', '540531', '904', '540531050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542001B00400', '西藏自治区日喀则地区聂拉木县公安边防大队', '540235000000', '', '1', '4', '542326', '42', '540235050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542400000000', '西藏自治区那曲地区公安处', '540000000000', '', '1', '2', '542400', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542400030000', '西藏自治区那曲地区公安处治安管理支队', '542400000000', '', '1', '3', '542400', '3', '542400210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542400050000', '西藏自治区那曲地区公安处刑事侦查支队', '542400000000', '', '1', '3', '542400', '5', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542400080000', '西藏自治区那曲地区公安处铁路治安管理支队', '542400000000', '', '1', '3', '542400', '3', '542400210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542400130000', '西藏自治区那曲地区公安处监所管理支队', '542400000000', '', '1', '3', '542400', '42', '542400210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542400130100', '西藏自治区那曲地区看守所', '542400000000', '', '1', '3', '542400', '904', '542400210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542400130200', '西藏自治区那曲地区行政拘留所', '542400000000', '', '1', '3', '542400', '903', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542400210000', '西藏自治区那曲地区公安处禁毒支队', '542400000000', '', '1', '2', '542400', '21', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542400970000', '西藏自治区那曲地区安多县公安局雁石坪一级公安检查站', '542425000000', '', '1', '4', '542400', '42', '542425050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542400980000', '西藏自治区那曲地区比如县公安局夏曲卡一级公安检查站', '542423000000', '', '1', '4', '542400', '42', '542423050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542400990000', '西藏那曲公安处青海省格尔木市乃吉沟（南山口）一级公安检查站', '542431000000', '', '1', '4', '542400', '42', '542431050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542400S00000', '西藏自治区那曲地区森林公安局', '542400000000', '', '1', '3', '542400', '42', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542400S00500', '西藏自治区那曲地区森林公安局刑侦支队', '542400000000', '', '1', '3', '542400', '804', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542421000000', '西藏自治区那曲地区那曲县公安局', '542400000000', '', '1', '3', '542421', '42', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542421030000', '西藏自治区那曲地区那曲县公安局治安管理大队（爆炸危险品监管大队）', '542421000000', '', '1', '4', '542421', '3', '542421050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542421050000', '西藏自治区那曲地区那曲县公安局刑事侦查大队（禁毒委员会办公室）', '542421000000', '', '1', '4', '542421', '5', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542422000000', '西藏自治区那曲地区嘉黎县公安局', '542400000000', '', '1', '3', '542422', '42', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542422030000', '西藏自治区那曲地区嘉黎县公安局治安管理大队（爆炸危险品监管大队）', '542422000000', '', '1', '4', '542422', '3', '542422050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542422050000', '西藏自治区那曲地区嘉黎县公安局刑事侦查大队（禁毒委员会办公室）', '542422000000', '', '1', '4', '542422', '5', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542422130100', '西藏自治区那曲地区嘉黎县公安局看守所', '542422000000', '', '1', '4', '542422', '904', '542422050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542422S00000', '西藏自治区那曲地区嘉黎县森林公安局', '542422000000', '', '1', '4', '542422', '42', '542422050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542422S00500', '西藏自治区那曲地区嘉黎县森林公安局刑侦大队', '542422000000', '', '1', '4', '542422', '804', '542422050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542423000000', '西藏自治区那曲地区比如县公安局', '542400000000', '', '1', '3', '542423', '42', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542423030000', '西藏自治区那曲地区比如县公安局治安管理大队（爆炸危险品监管大队）', '542423000000', '', '1', '4', '542423', '3', '542423050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542423050000', '西藏自治区那曲地区比如县公安局刑事侦查大队（禁毒委员会办公室）', '542423000000', '', '1', '4', '542423', '5', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542423130100', '西藏自治区那曲地区比如县公安局看守所', '542423000000', '', '1', '4', '542423', '904', '542423050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542423S00000', '西藏自治区那曲地区比如县森林公安局', '542423000000', '', '1', '4', '542423', '42', '542423050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542423S00500', '西藏自治区那曲地区比如县森林公安局刑侦大队', '542423000000', '', '1', '4', '542423', '804', '542423050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542424000000', '西藏自治区那曲地区聂荣县公安局', '542400000000', '', '1', '3', '542424', '42', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542424030000', '西藏自治区那曲地区聂荣县公安局治安管理大队（爆炸危险品监管大队）', '542424000000', '', '1', '4', '542424', '3', '542424050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542424050000', '西藏自治区那曲地区聂荣县公安局刑事侦查大队（禁毒委员会办公室）', '542424000000', '', '1', '4', '542424', '5', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542424130100', '西藏自治区那曲地区聂荣县公安局看守所', '542424000000', '', '1', '4', '542424', '904', '542424050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542425000000', '西藏自治区那曲地区安多县公安局', '542400000000', '', '1', '3', '542425', '42', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542425030000', '西藏自治区那曲地区安多县公安局治安管理大队（爆炸危险品监管大队）', '542425000000', '', '1', '4', '542425', '3', '542425050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542425050000', '西藏自治区那曲地区安多县公安局刑事侦查大队（禁毒委员会办公室）', '542425000000', '', '1', '4', '542425', '5', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542425130100', '西藏自治区那曲地区安多县公安局看守所', '542425000000', '', '1', '4', '542425', '904', '542425050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542425S00000', '西藏自治区那曲地区安多县森林公安局', '542425000000', '', '1', '4', '542425', '42', '542425050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542425S00500', '西藏自治区那曲地区安多县森林公安局刑侦大队', '542425000000', '', '1', '4', '542425', '804', '542425050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542426000000', '西藏自治区那曲地区申扎县公安局', '542400000000', '', '1', '3', '542426', '42', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542426030000', '西藏自治区那曲地区申扎县公安局治安管理大队（爆炸危险品监管大队）', '542426000000', '', '1', '4', '542426', '3', '542426050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542426050000', '西藏自治区那曲地区申扎县公安局刑事侦查大队（禁毒委员会办公室）', '542426000000', '', '1', '4', '542426', '5', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542426130100', '西藏自治区那曲地区申扎县公安局看守所', '542426000000', '', '1', '4', '542426', '904', '542426050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542426S00000', '西藏自治区那曲地区申扎县森林公安局', '542426000000', '', '1', '4', '542426', '42', '542426050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542426S00500', '西藏自治区那曲地区申扎县森林公安局刑侦大队', '542426000000', '', '1', '4', '542426', '804', '542426050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542427000000', '西藏自治区那曲地区索县公安局', '542400000000', '', '1', '3', '542427', '42', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542427030000', '西藏自治区那曲地区索县公安局治安管理大队（爆炸危险品监管大队）', '542427000000', '', '1', '4', '542427', '3', '542427050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542427050000', '西藏自治区那曲地区索县公安局刑事侦查大队（禁毒委员会办公室）', '542427000000', '', '1', '4', '542427', '5', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542427130100', '西藏自治区那曲地区索县公安局看守所', '542427000000', '', '1', '4', '542427', '904', '542427050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542427S00000', '西藏自治区那曲地区索县森林公安局', '542427000000', '', '1', '4', '542427', '42', '542427050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542427S00500', '西藏自治区那曲地区索县森林公安局刑侦大队', '542427000000', '', '1', '4', '542427', '804', '542427050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542428000000', '西藏自治区那曲地区班戈县公安局', '542422000000', '', '1', '4', '542428', '42', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542428030000', '西藏自治区那曲地区班戈县公安局治安管理大队（爆炸危险品监管大队）', '542422000000', '', '1', '4', '542428', '3', '542428050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542428050000', '西藏自治区那曲地区班戈县公安局刑事侦查大队（禁毒委员会办公室）', '542422000000', '', '1', '4', '542428', '5', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542428130100', '西藏自治区那曲地区班戈县公安局看守所', '542422000000', '', '1', '4', '542428', '904', '542428050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542429000000', '西藏自治区那曲地区巴青县公安局', '542400000000', '', '1', '3', '542429', '42', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542429030000', '西藏自治区那曲地区巴青县公安局治安管理大队（爆炸危险品监管大队）', '542429000000', '', '1', '4', '542429', '3', '542429050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542429050000', '西藏自治区那曲地区巴青县公安局刑事侦查大队（禁毒委员会办公室）', '542429000000', '', '1', '4', '542429', '5', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542429130100', '西藏自治区那曲地区巴青县公安局看守所', '542429000000', '', '1', '4', '542429', '904', '542429050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542430000000', '西藏自治区那曲地区尼玛县公安局', '542400000000', '', '1', '3', '542430', '42', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542430030000', '西藏自治区那曲地区尼玛县公安局治安管理大队（爆炸危险品监管大队）', '542430000000', '', '1', '4', '542430', '3', '542430050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542430050000', '西藏自治区那曲地区尼玛县公安局刑事侦查大队（禁毒委员会办公室）', '542430000000', '', '1', '4', '542430', '5', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542430130100', '西藏自治区那曲地区尼玛县公安局看守所', '542430000000', '', '1', '4', '542430', '904', '542430050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542430S00000', '西藏自治区那曲地区尼玛县森林公安局', '542430000000', '', '1', '4', '542430', '42', '542430050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542430S00500', '西藏自治区那曲地区尼玛县森林公安刑侦大队', '542430000000', '', '1', '4', '542430', '804', '542430050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542431000000', '西藏自治区那曲地区双湖县公安局', '542400000000', '', '1', '3', '542431', '42', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542431030000', '西藏自治区那曲地区双湖县公安局治安管理大队（爆炸危险品监管大队）', '542431000000', '', '1', '4', '542431', '3', '542431050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542431050000', '西藏自治区那曲地区双湖县公安局刑事侦查大队（禁毒委员会办公室）', '542431000000', '', '1', '4', '542431', '5', '542400050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542431130100', '西藏自治区那曲地区双湖县公安局看守所', '542431000000', '', '1', '4', '542431', '904', '542431050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542431S00000', '西藏自治区那曲地区双湖县森林公安局', '542431000000', '', '1', '4', '542431', '42', '542431050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542431S00500', '西藏自治区那曲地区双湖县森林公安局刑侦大队', '542431000000', '', '1', '4', '542431', '804', '542431050000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542500000000', '西藏自治区阿里地区公安处', '540000000000', '', '1', '2', '542500', '42', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542500030000', '西藏自治区阿里地区公安处治安管理支队', '542500000000', '', '1', '3', '542500', '3', '542500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542500050000', '西藏自治区阿里地区公安处刑事侦查支队', '542500000000', '', '1', '3', '542500', '5', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542500130100', '西藏自治区阿里地区公安处看守所', '542500000000', '', '1', '3', '542500', '904', '542500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542500210000', '西藏自治区阿里地区公安处禁毒支队', '542500000000', '', '1', '2', '542500', '21', '540000210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542500950000', '西藏自治区阿里地区措勤县美朵公安一级检查站', '542527000000', '', '1', '4', '542500', '42', '542527000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542500960000', '西藏自治区阿里地区改则县洞措拉雄公安一级检查站', '542526000000', '', '1', '4', '542500', '42', '542526000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542500970000', '西藏自治区阿里地区革吉县革狮公安一级检查站', '542525000000', '', '1', '4', '542500', '42', '542525000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542500980000', '西藏自治区阿里地区日土县泉水湖公安一级检查站', '542524000000', '', '1', '4', '542500', '42', '542524000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542500990000', '西藏自治区阿里地区日土县多玛公安一级检查站', '542524000000', '', '1', '4', '542500', '42', '542524000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542500M00000', '西藏自治区阿里地区昆莎机场公安分局', '542500000000', '', '1', '3', '542500', '42', '542500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542500S00000', '西藏自治区阿里地区森林公安局', '542500000000', '', '1', '3', '542500', '42', '542500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542500S00500', '西藏自治区阿里地区森林公安刑侦支队', '542500000000', '', '1', '3', '542500', '804', '542500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542521000000', '西藏自治区阿里地区普兰县公安局', '542500000000', '', '1', '3', '542521', '42', '542500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542521010000', '西藏自治区阿里地区普兰县公安局刑事侦查大队', '542521000000', '', '1', '4', '542521', '42', '');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542521030000', '西藏自治区阿里地区普兰县公安局治安大队', '542521000000', '', '1', '4', '542521', '3', '542521000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542521130100', '西藏自治区阿里地区普兰县公安局看守所', '542521000000', '', '1', '4', '542521', '904', '542521000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542521B00000', '中国人民武装警察部队西藏自治区阿里地区普兰县公安边防大队', '542521000000', '', '1', '4', '542521', '42', '542521000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542522000000', '西藏自治区阿里地区札达县公安局', '542500000000', '', '1', '3', '542522', '42', '542500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542522010000', '西藏自治区阿里地区札达县公安局刑事侦查大队', '542522000000', '', '1', '4', '542522', '42', '');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542522030000', '西藏自治区阿里地区札达县公安局治安大队', '542522000000', '', '1', '4', '542522', '3', '542522000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542522130100', '西藏自治区阿里地区札达县公安局看守所', '542522000000', '', '1', '4', '542522', '904', '542522000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542522B00000', '中国人民武装警察部队西藏自治区阿里地区札达县公安边防大队', '542522000000', '', '1', '4', '542522', '42', '542522000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542523000000', '西藏自治区阿里地区噶尔县公安局', '542500000000', '', '1', '3', '542523', '42', '542500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542523010000', '西藏自治区阿里地区噶尔县公安局刑事侦查大队', '542523000000', '', '1', '4', '542523', '42', '');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542523030000', '西藏自治区阿里地区噶尔县公安局治安大队', '542523000000', '', '1', '4', '542523', '3', '542523000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542523B00000', '中国人民武装警察部队西藏自治区阿里地区噶尔县公安边防大队', '542523000000', '', '1', '4', '542523', '42', '542523000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542524000000', '西藏自治区阿里地区日土县公安局', '542500000000', '', '1', '3', '542524', '42', '542500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542524010000', '西藏自治区阿里地区日土县公安局刑事侦查大队', '542524000000', '', '1', '4', '542524', '42', '');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542524030000', '西藏自治区阿里地区日土县公安局治安大队', '542524000000', '', '1', '4', '542524', '3', '542524000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542524130100', '西藏自治区阿里地区日土县公安局看守所', '542524000000', '', '1', '4', '542524', '904', '542524000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542524B00000', '中国人民武装警察部队西藏自治区阿里地区日土县公安边防大队', '542524000000', '', '1', '4', '542524', '42', '542524000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542524S00000', '西藏自治区阿里地区日土县森林公安局', '542524000000', '', '1', '4', '542524', '42', '542524000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542524S00500', '西藏自治区阿里地区日土县森林公安局刑侦大队', '542524000000', '', '1', '4', '542524', '804', '542524000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542525000000', '西藏自治区阿里地区革吉县公安局', '542500000000', '', '1', '3', '542525', '42', '542500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542525010000', '西藏自治区阿里地区革吉县公安局刑事侦查大队', '542525000000', '', '1', '4', '542525', '42', '');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542525030000', '西藏自治区阿里地区革吉县公安局治安大队', '542525000000', '', '1', '4', '542525', '3', '542525000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542525130100', '西藏自治区阿里地区革吉县公安局看守所', '542525000000', '', '1', '4', '542525', '904', '542525000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542526000000', '西藏自治区阿里地区改则县公安局', '542500000000', '', '1', '3', '542526', '42', '542500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542526010000', '西藏自治区阿里地区改则县公安局刑事侦查大队', '542526000000', '', '1', '4', '542526', '42', '');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542526030000', '西藏自治区阿里地区改则县公安局治安大队', '542526000000', '', '1', '4', '542526', '3', '542526000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542526310100', '西藏自治区阿里地区改则县公安局看守所', '542526000000', '', '1', '4', '542526', '904', '542526000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542526S00000', '西藏自治区阿里地区改则县森林公安局', '542526000000', '', '1', '4', '542526', '42', '542526000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542526S00500', '西藏自治区阿里地区改则县森林公安局刑侦大队', '542526000000', '', '1', '4', '542526', '804', '542526000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542527000000', '西藏自治区阿里地区措勤县公安局', '542500000000', '', '1', '3', '542527', '42', '542500210000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542527010000', '西藏自治区阿里地区措勤县公安局刑事侦查大队', '542527000000', '', '1', '4', '542527', '42', '');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542527030000', '西藏自治区阿里地区措勤县公安局治安大队', '542527000000', '', '1', '4', '542527', '3', '542527000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('542527130100', '西藏自治区阿里地区措勤县公安局看守所', '542527000000', '', '1', '4', '542527', '904', '542527000000');
INSERT INTO sys_organisations_2 (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('544001B10100', '西藏自治区山南地区错那县公安边防大队部', '540530000000', '', '1', '4', '544001', '42', '540530050000');



INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410200000000', '河南省开封市公安局', '410000000000', '开封市公安局', '1', '2', '410200', '42', '410000210000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410200030000', '河南省开封市公安局治安支队', '410200000000', '开封市公安局治安管理处', '1', '3', '410200', '3', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410200050000', '河南省开封市公安局刑事侦查支队', '410200000000', '开封市公安局刑事侦查支队', '1', '3', '410200', '5', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410200090000', '河南省开封市公安局治安和出入境管理支队', '410200000000', '治安和出入境管理支队', '1', '3', '410200', '3', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410200100000', '河南省开封市公安局犯罪侦查支队', '410200000000', '犯罪侦查支队', '1', '3', '410200', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410200101000', '河南省开封市公安局犯罪侦查支队禁毒大队', '410200100000', '', '1', '4', '410200', '21', '410000210000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410200130000', '河南省开封市公安局监所管理支队', '410200000000', '开封市公安局监所管理支队', '1', '3', '410211', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410200210000', '河南省开封市公安局缉毒支队', '410200000000', '', '1', '3', '410200', '21', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410200410000', '河南省开封市公安局看守所', '410200000000', '', '1', '3', '410211', '904', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410200440000', '河南省开封市公安局行政拘留所', '410200000000', '', '1', '3', '410211', '903', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410200S00000', '河南省开封市森林公安局', '410200000000', '开封市公安局森林警察支队', '1', '3', '410200', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410202000000', '河南省开封市龙亭分局', '410200000000', '开封市龙亭分局', '1', '3', '410200', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410202030000', '河南省开封市龙亭分局治安管理大队', '410202000000', '开封市龙亭分局治安管理大队', '1', '4', '410202', '3', '410202050000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410202050000', '河南省开封市龙亭分局刑事侦查大队', '410202000000', '开封市龙亭分局刑事侦查大队', '1', '3', '410202', '5', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410203000000', '河南省开封市顺河分局', '410200000000', '开封市顺河分局', '1', '3', '410200', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410203030000', '河南省开封市顺河分局治安大队', '410203000000', '开封市顺河分局治安大队', '1', '4', '410203', '3', '410203050000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410203050000', '河南省开封市顺河分局刑事侦查大队', '410203000000', '开封市顺河分局刑事侦查大队', '1', '3', '410203', '5', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410204000000', '河南省开封市鼓楼分局', '410200000000', '开封市鼓楼分局', '1', '3', '410200', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410204030000', '河南省开封市鼓楼分局治安大队', '410204000000', '开封市鼓楼分局治安大队', '1', '4', '410204', '3', '410204050000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410204050000', '河南省开封市鼓楼分局刑事侦查大队', '410204000000', '开封市鼓楼分局刑事侦查大队', '1', '3', '410204', '5', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410205000000', '河南省开封市禹王台分局', '410200000000', '开封市禹王台分局', '1', '3', '410200', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410205030000', '河南省开封市禹王台分局治安大队', '410205000000', '开封市禹王台分局治安大队', '1', '4', '410205', '3', '410205050000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410205050000', '河南省开封市禹王台分局刑事侦查大队', '410205000000', '开封市禹王台分局刑事侦查大队', '1', '3', '410205', '5', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410211000000', '河南省开封市金明分局', '410200000000', '开封市金明分局', '1', '3', '410200', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410211030000', '河南省开封市金明分局治安大队', '410211000000', '开封市金明分局治安大队', '1', '4', '410215', '3', '410211050000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410211050000', '河南省开封市金明分局刑事侦查大队', '410211000000', '开封市金明分局刑事侦查大队', '1', '3', '410211', '5', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410215000000', '河南省开封市公安局金明池分局', '410200000000', '', '1', '3', '410215', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410221000000', '河南省开封市杞县公安局', '410200000000', '开封市杞县公安局', '1', '3', '410221', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410221030000', '河南省开封市杞县公安局治安科', '410221000000', '开封市杞县公安局治安科', '1', '4', '410221', '3', '410221000000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410221410000', '河南省开封市杞县公安局看守所', '410221000000', '开封市杞县公安局看守所', '1', '4', '410221', '904', '410221000000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410221440000', '河南省开封市杞县公安局行政拘留所', '410221000000', '开封市杞县公安局行政拘留所', '1', '4', '410221', '903', '410221000000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410221S00000', '河南省开封市杞县森林公安局', '410221000000', '开封市杞县公安局森林警察大队', '1', '4', '410221', '42', '410221000000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410222000000', '河南省开封市通许县公安局', '410200000000', '开封市通许县公安局', '1', '3', '410222', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410222030000', '河南省开封市通许县公安局治安大队', '410222000000', '开封市通许县公安局治安大队', '1', '4', '410222', '3', '410222000000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410222410000', '河南省开封市通许县公安局看守所', '410222000000', '开封市通许县公安局看守所', '1', '4', '410222', '904', '410222000000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410222440000', '河南省开封市通许县公安局行政拘留所', '410222000000', '开封市通许县公安局行政拘留所', '1', '4', '410222', '903', '410222000000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410222S00000', '河南省开封市通许县森林公安局', '410222000000', '开封市通许县公安局森林警察大队', '1', '4', '410222', '42', '410222000000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410223000000', '河南省开封市尉氏县公安局', '410200000000', '开封市尉氏县公安局', '1', '3', '410223', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410223030000', '河南省开封市尉氏县公安局治安大队', '410223000000', '开封市尉氏县公安局治安大队', '1', '4', '410223', '3', '410223210000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410223210000', '河南省开封市尉氏县公安局禁毒大队', '410223000000', '开封市尉氏县公安局禁毒大队', '1', '3', '410223', '21', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410223410000', '河南省开封市尉氏县公安局看守所', '410223000000', '开封市尉氏县公安局看守所', '1', '4', '410223', '904', '410223210000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410223440000', '河南省开封市尉氏县公安局拘留所', '410223000000', '开封市尉氏县公安局拘留所', '1', '4', '410223', '903', '410223210000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410223S00000', '河南省开封市尉氏县森林公安局', '410223000000', '开封市尉氏县公安局森林警察大队', '1', '4', '410223', '42', '410223210000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410224000000', '河南省开封市祥符区公安分局', '410200000000', '开封市开封县公安局', '1', '3', '410224', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410224010000', '河南省开封市祥符区公安分局刑警大队', '410224000000', '', '1', '4', '410224', '42', '');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410224030000', '河南省开封市祥符区公安分局治安大队', '410224000000', '开封市开封县公安局治安大队', '1', '4', '410224', '3', '410224000000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410224410000', '河南省开封市祥符区公安分局看守所', '410224000000', '开封市开封县公安局看守所', '1', '4', '410224', '904', '410224000000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410224440000', '河南省开封市祥符区公安分局行政拘留所', '410224000000', '开封市开封县公安局行政拘留所', '1', '4', '410224', '903', '410224000000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410224S00000', '河南省开封市祥符区森林公安局', '410224000000', '开封市开封县公安局森林警察大队', '1', '4', '410224', '42', '410224000000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410225000000', '河南省兰考县公安局', '410200000000', '开封市兰考县公安局', '1', '2', '413100', '42', '410000000000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410225030000', '河南省兰考县公安局治安大队', '410225000000', '开封市兰考县公安局治安大队', '1', '3', '413100', '3', '410225050000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410225050000', '河南省兰考县公安局刑事侦查大队', '410000000000', '开封市兰考县公安局刑事侦查大队', '1', '3', '413100', '5', '410000000000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410225080000', '河南省兰考县公安局禁毒大队', '410000000000', '兰考县公安局禁毒大队', '1', '2', '413100', '21', '410000000000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410225410000', '河南省兰考县看守所', '410225000000', '开封市兰考县公安局看守所', '1', '3', '413100', '904', '410225050000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410225440000', '河南省兰考县拘留所', '410225000000', '开封市兰考县公安局拘留所', '1', '3', '413100', '903', '410225050000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410225S00000', '河南省兰考县森林公安局', '410225000000', '开封市兰考县公安局森林警察大队', '1', '3', '413100', '42', '410225050000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410269000000', '河南省开封市公安局开发区第一分局', '410200000000', '', '1', '3', '410215', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410270000000', '河南省开封市公安局梁苑分局', '410200000000', '', '1', '3', '410211', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410271000000', '河南省开封市公安局杏花营分局', '410200000000', '', '1', '3', '410211', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410272000000', '河南省开封市公安局金明第一分局', '410200000000', '', '1', '3', '410211', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410273000000', '河南省开封市公安局金明第三分局', '410200000000', '', '1', '3', '410211', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410274000000', '河南省开封市公安局金耀分局', '410200000000', '', '1', '3', '410211', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410275000000', '河南省开封市公安局金明第二分局', '410200000000', '', '1', '3', '410211', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410276000000', '河南省开封市公安局禹王台第二分局', '410200000000', '', '1', '3', '410205', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410277000000', '河南省开封市公安局禹王台第一分局', '410200000000', '', '1', '3', '410205', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410278000000', '河南省开封市公安局禹王台第三分局', '410200000000', '', '1', '3', '410205', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410279000000', '河南省开封市公安局繁塔分局', '410200000000', '', '1', '3', '410205', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410280000000', '河南省开封市公安局机场分局', '410200000000', '', '1', '3', '410205', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410281000000', '河南省开封市公安局红洋楼分局', '410200000000', '', '1', '3', '410205', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410282000000', '河南省开封市公安局南苑分局', '410200000000', '', '1', '3', '410204', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410283000000', '河南省开封市公安局相国寺分局', '410200000000', '', '1', '3', '410204', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410284000000', '河南省开封市公安局鼓楼第三分局', '410200000000', '', '1', '3', '410204', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410285000000', '河南省开封市公安局鼓楼第二分局', '410200000000', '', '1', '3', '410204', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410286000000', '河南省开封市公安局鼓楼第一分局', '410200000000', '', '1', '3', '410204', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410287000000', '河南省开封市公安局州桥分局', '410200000000', '', '1', '3', '410204', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410289000000', '河南省开封市公安局顺河第二分局', '410200000000', '', '1', '3', '410203', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410290000000', '河南省开封市公安局顺河第一分局', '410200000000', '', '1', '3', '410203', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410291000000', '河南省开封市公安局顺河第三分局', '410200000000', '', '1', '3', '410203', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410292000000', '河南省开封市公安局宋门分局', '410200000000', '', '1', '3', '410203', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410293000000', '河南省开封市公安局土柏岗分局', '410200000000', '', '1', '3', '410203', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410294000000', '河南省开封市公安局铁塔分局', '410200000000', '', '1', '3', '410203', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410295000000', '河南省开封市公安局柳园口分局', '410200000000', '', '1', '3', '410202', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410296000000', '河南省开封市公安局龙亭第一分局', '410200000000', '', '1', '3', '410202', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410297000000', '河南省开封市公安局龙亭第二分局', '410200000000', '', '1', '3', '410202', '42', '410200101000');
INSERT INTO sys_organisations (dwdm, dwmc, sjdm, dwjc, yxx, dwjb, xzqh, dwjz, ywsjdm) VALUES ('410298000000', '河南省开封市公安局午朝门分局', '410200000000', '', '1', '3', '410202', '42', '410200101000');
`

type SysOrganisation struct {
	Dwdm   string `json:"dwdm" xorm:"pk VARCHAR(12)" colName:"单位代码"`    //单位代码
	Dwmc   string `json:"dwmc" xorm:"VARCHAR(200)" colName:"单位名称"`      //单位名称
	Sjdm   string `json:"sjdm" xorm:"VARCHAR(12) index" colName:"上级单位"` //上级单位
	Dwjc   string `json:"dwjc" xorm:"VARCHAR(200)" colName:"单位简称"`      //单位简称
	Yxx    string `json:"yxx" xorm:"VARCHAR(1)" colName:"有效性"`          //有效性
	Dwjb   string `json:"dwjb" xorm:"VARCHAR(1)" colName:"单位级别"`        //单位级别
	Xzqh   string `json:"xzqh" xorm:"VARCHAR(6)" colName:"所属地区"`        //所属地区
	Dwjz   string `json:"dwjz" xorm:"VARCHAR(4)" colName:"单位警种"`        //单位警种
	Ywsjdm string `json:"ywsjdm" xorm:"VARCHAR(12)" colName:"业务上级单位"`   //业务上级单位
}

func (s *SysOrganisation) TableName() string {
	return "sys_organisations"
}

var total int
var success int
var dataReg = regexp.MustCompile(`VALUES \((.*)\);`)

func main() {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "123456.abcd", "localhost:3306", "tqry") + "&loc=Asia%2FChongqing"
	Engine, err := xorm.NewEngine("mysql", params)
	if err != nil {
		fmt.Println(err)
		return
	}
	del := []string{"54", "4102"}
	for _, v := range del {
		_, err := Engine.Exec("delete from sys_organisations where dwdm like '" + v + "%';")
		if err != nil {
			log.Fatal(err)
		}
	}
	var dbs []*SysOrganisation
	insert := strings.Split(sql, "\n")
	for _, line := range insert {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		aa := dataReg.FindAllStringSubmatch(line, -1)
		if len(aa) > 0 && len(aa[0]) > 1 {
			line = aa[0][1]
			line = strings.Replace(line, "'", "", -1)
			arr := strings.Split(line, ",")
			if len(arr) >= 9 {
				//去掉不要的警种
				jz := strings.TrimLeft(strings.TrimSpace(arr[7]), "0")
				dbs = append(dbs, &SysOrganisation{Dwdm: arr[0],
					Dwmc: strings.TrimSpace(arr[1]),
					Sjdm: strings.TrimSpace(arr[2]),
					Dwjc: strings.TrimSpace(arr[3]),
					Yxx: strings.TrimSpace(arr[4]),
					Dwjb: strings.TrimSpace(arr[5]),
					Xzqh: strings.TrimSpace(arr[6]),
					Dwjz: jz,
					Ywsjdm: strings.TrimSpace(arr[8])})
				if len(dbs) >= 200 {
					BatchInsert(Engine, dbs)
					dbs = []*SysOrganisation{}
				}
			}
		}
	}
	if len(dbs) > 0 {
		BatchInsert(Engine, dbs)
	}

}

func Batch(Engine *xorm.Engine, datas []*SysOrganisation) error {
	session := Engine.NewSession()
	defer session.Close()
	session.Begin()
	_, err := session.Insert(datas)
	if err != nil {
		session.Rollback()
		return err
	}
	return session.Commit()
}

func BatchInsert(Engine *xorm.Engine, datas []*SysOrganisation) {
	defer func() {
		fmt.Println(total, success)
	}()
	total += len(datas)
	var failure int
	err := Batch(Engine, datas)
	if err != nil {
		for i, _ := range datas {
			_, err := Engine.InsertOne(datas[i])
			if err != nil {
				failure++
				fmt.Println(err)
			}
		}
	}
	success += len(datas) - failure
}
