/*
Navicat MySQL Data Transfer

Source Server         : 192.168.130.240
Source Server Version : 50542
Source Host           : 192.168.130.240:3306
Source Database       : szga

Target Server Type    : MYSQL
Target Server Version : 50542
File Encoding         : 65001

Date: 2016-03-30 13:59:34
*/
SET character_set_client='utf8';
SET character_set_connection='utf8';
SET character_set_database='utf8';
SET character_set_results='utf8';
SET character_set_server='utf8';
SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for auth_menus
-- ----------------------------
DROP TABLE IF EXISTS `auth_menus`;
CREATE TABLE `auth_menus` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `contains_url` varchar(1024) DEFAULT NULL,
  `sort_number` int(11) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `enable` int(11) DEFAULT NULL,
  `parent` bigint(20) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of auth_menus
-- ----------------------------
INSERT INTO `auth_menus` VALUES ('5', '系统管理', '/szga/', '', '7', '', '1', '0', '2016-02-16 23:45:10', '2016-03-28 17:08:44');
INSERT INTO `auth_menus` VALUES ('6', '系统用户管理', '/szga/sysusers', '/organizations/tree,/users/list,/users/read,/users/update,/users/create,/users/delete', '1', '', '1', '5', '2016-02-16 23:45:44', '2016-03-28 14:37:34');
INSERT INTO `auth_menus` VALUES ('12', '角色管理', '/szga/roles', '/roles/list,/users/tree,/menus/tree,/roles/getusers,/roles/getmenus,/roles/add2user,/roles/add2menu,/roles/update,/roles/create,/roles/delete', '3', '', '1', '5', '2016-02-17 21:31:51', '2016-03-04 10:44:24');
INSERT INTO `auth_menus` VALUES ('13', '菜单管理', '/szga/menus', '/menus/list,/roles/tree,/menus/getroles,/menus/add2role,/menus/update,/menus/create,/menus/delete,/menus/read,/menus/tree,/menus/dropdown', '4', '菜单管理', '1', '5', '2016-02-17 21:33:56', '2016-03-04 10:48:30');
INSERT INTO `auth_menus` VALUES ('14', '授权管理', '/szga/licensing', '/manager/license', '5', '', '1', '5', '2016-02-17 21:40:03', '2016-03-04 10:52:37');
INSERT INTO `auth_menus` VALUES ('17', '用户管理', '/szga/users', '/organizations/tree,/szga_users/list,/szga_users/delete,/szga_users/create,/szga_users/update', '3', '用户管理', '1', '0', '2016-02-17 21:48:18', '2016-03-28 14:32:39');
INSERT INTO `auth_menus` VALUES ('18', '终端管理', '/szga/terminals', '/szga_terminals/list,/szga_terminals/update,/szga_terminals/create,/szga_terminals/delete', '4', '终端管理', '1', '0', '2016-02-17 21:48:32', '2016-03-23 13:28:20');
INSERT INTO `auth_menus` VALUES ('19', '设备管理', '/szga/devices', '/szga_devices/list,/szga_devices/update,/szga_devices/delete,/szga_devices/create', '5', '设备管理', '1', '0', '2016-02-17 21:48:44', '2016-03-28 14:37:12');
INSERT INTO `auth_menus` VALUES ('20', '统计报表', '/szga/', '', '1', '统计报表', '1', '0', '2016-02-17 21:48:59', '2016-03-01 09:26:54');
INSERT INTO `auth_menus` VALUES ('21', '用户信息展示', '/szga/userinfoReports', '/search/aggcount,/szga_users/countstatus,/szga_users/getavgonofftime,/szga_users/gettimespacestatus,/szga_devices/countstatus,/szgaonlines/firstlogin,/organizations/list,/homepagecharts/create', '1', '用户信息展示', '1', '20', '2016-02-17 21:49:20', '2016-03-30 11:23:44');
INSERT INTO `auth_menus` VALUES ('22', '终端信息展示', '/szga/terminalReports', '/search/aggcount,/szga_terminals/countstatus,/szgaonlines/online4map,/szgaonlines/firstlogin,/homepagecharts/create', '2', '终端信息展示', '1', '20', '2016-02-17 21:49:39', '2016-03-30 11:23:49');
INSERT INTO `auth_menus` VALUES ('23', '网络信息展示', '/szga/netReports', '/search/aggcount,/homepagecharts/create', '3', '网络信息展示', '1', '20', '2016-02-17 21:50:04', '2016-03-30 11:23:53');
INSERT INTO `auth_menus` VALUES ('24', '数据信息展示', '/szga/dataReports', '/search/aggcount,/homepagecharts/create', '4', '数据信息展示', '1', '20', '2016-02-17 21:50:19', '2016-03-30 11:23:58');
INSERT INTO `auth_menus` VALUES ('25', '应用信息展示', '/szga/appReports', '/search/aggcount,/szga_uploads/getserverinfo,/szga_uploads/getonoffstatus,/homepagecharts/create', '5', '应用信息展示', '1', '20', '2016-02-17 21:50:39', '2016-03-30 11:24:05');
INSERT INTO `auth_menus` VALUES ('26', '安全事件信息展示', '/szga/securityinfoReports', '/search/aggcount,,/szga_uploads/setflagsecevent,/szga_uploads/setflagdlpeve,/szga_uploads/getdlpeve,/szga_uploads/getsecevent,/homepagecharts/create', '6', '安全事件信息展示', '1', '20', '2016-02-17 21:51:01', '2016-03-30 11:24:47');
INSERT INTO `auth_menus` VALUES ('28', '策略管理', '/szga/plans', '/szga_strategies/get', '7', '策略管理', '1', '0', '2016-02-17 21:51:36', '2016-03-23 13:28:42');
INSERT INTO `auth_menus` VALUES ('29', '我的', '/szga/me', '/users/read,/users/modify', '8', '我的', '1', '0', '2016-02-17 21:51:50', '2016-03-23 13:28:47');
INSERT INTO `auth_menus` VALUES ('45', '组织机构', '/szga/organizations', '/organizations/list,/organizations/tree,/organizations/read,/organizations/delete,/organizations/update,/organizations/create', '2', '', '1', '5', '2016-02-19 19:03:16', '2016-03-28 13:55:43');
INSERT INTO `auth_menus` VALUES ('46', '主页', '/szga/system', '/homepagecharts/save,/homepagecharts/delete,/homepagecharts/list,/search/aggcount,/szga_users/countstatus,/szga_users/getavgonofftime,/szga_users/gettimespacestatus,/szga_devices/countstatus,/szgaonlines/firstlogin,/organizations/list,/homepagecharts/create,/szga_terminals/countstatus,/szgaonlines/online4map,/szga_uploads/getserverinfo,/szga_uploads/getonoffstatus,/search/search,/szga_uploads/setflagsecevent,/szga_uploads/setflagdlpeve,/szga_uploads/getdlpeve,/szga_uploads/getsecevent', '0', '', '1', '0', '2016-02-21 22:15:01', '2016-03-30 13:07:21');
INSERT INTO `auth_menus` VALUES ('47', '信息查询', '/szga/reports', '/organizations/list,/search/search', '6', '安全事件信息查询', '1', '0', '2016-02-24 16:31:10', '2016-03-28 17:08:31');
INSERT INTO `auth_menus` VALUES ('55', '基础信息导入', '/szga/importdata', '/szga_uploads/handle', '6', '基础信息导入', '1', '5', '2016-03-30 11:37:40', '2016-03-30 13:59:53');
