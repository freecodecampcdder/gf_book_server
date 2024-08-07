/*
 Navicat Premium Data Transfer

 Source Server         : 本机Mysql
 Source Server Type    : MySQL
 Source Server Version : 50740
 Source Host           : 127.0.0.1:3306
 Source Schema         : gxs_book

 Target Server Type    : MySQL
 Target Server Version : 50740
 File Encoding         : 65001

 Date: 17/07/2023 18:29:58
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int(10) UNSIGNED NOT NULL,
  `receiver_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '接收人名称',
  `receiver_phone` varchar(11) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '接收人手机号',
  `address_content` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '地址',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '1正常 2默认地址',
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `id`(`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of address
-- ----------------------------
INSERT INTO `address` VALUES (1, 4, '朱淑祺', '18254705629', '山东省济宁市鱼台县老砦乡后六屯村133号', 1, '2023-06-20 15:55:30', '2023-06-26 02:27:55', NULL);
INSERT INTO `address` VALUES (2, 4, '郭星铄', '18888888888', '山东省临沂市郯城县老郭村', 1, '2023-06-20 08:45:10', '2023-06-26 02:27:55', NULL);
INSERT INTO `address` VALUES (3, 4, '仔仔', '18888888881', '山东省济宁市汶上县仔仔村', 1, '2023-06-20 08:45:35', '2023-06-26 02:27:55', NULL);
INSERT INTO `address` VALUES (4, 4, '杨翔宇', '18888088888', '山东省济南市历下区县老杨村', 2, '2023-06-20 08:46:07', '2023-06-26 02:27:55', NULL);
INSERT INTO `address` VALUES (8, 4, '朱老师', '18888888888', '山东省济南市临沂市鱼台县', 1, '2023-06-26 02:24:13', '2023-06-26 02:27:55', '2023-06-26 02:29:09');

-- ----------------------------
-- Table structure for book
-- ----------------------------
DROP TABLE IF EXISTS `book`;
CREATE TABLE `book`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `language_id` int(11) UNSIGNED NOT NULL COMMENT '语种ID',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '书名',
  `author` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '作者名称',
  `cover` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '图书缩略图',
  `translator` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '译者',
  `description` text CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '简介',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1可借 3下架',
  `isbn` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '国际标准图书编号',
  `press` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '出版社',
  `press_time` int(11) NOT NULL COMMENT '出版时间',
  `page_num` int(11) NOT NULL COMMENT '页数',
  `wish_num` int(11) NOT NULL DEFAULT 0 COMMENT '心愿数量',
  `collect_num` int(11) NOT NULL COMMENT '收藏数量',
  `borrow_num` int(11) NOT NULL COMMENT '被借数量',
  `price` int(11) NOT NULL COMMENT '租售价格*100',
  `buy_price` int(11) NOT NULL COMMENT '售卖价格*100',
  `inventory_num` int(11) NOT NULL DEFAULT 0 COMMENT '库存数量',
  `user_id` int(11) NOT NULL COMMENT '添加人ID',
  `recommended` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '推荐系数 越大越靠前',
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of book
-- ----------------------------
INSERT INTO `book` VALUES (1, 1, '什么是科学', '吴国盛', 'http://rwj48creq.hb-bkt.clouddn.com/book/242dd42a2834349b033b895da0bc02ce36d3d4391096.webp', '', '在本书中，吴国盛教授先是梳理当代中国人的科学观念，然后阐述现代科学起源于古希腊的理性科学，并对比中西文化的差异。接着又从两个方面讨论现代科学的产生：第一、宗教直接助力现代科学产生；第二、数理实验科学塑造了现代科学的标准。此外，吴国盛教授还讨论另一科学传统——博物学（自然志），并由此回到中国的传统文化，坦言虽然在数理实验科学的意义上中国没有科学，但是在博物学的意义上，中国有科学，并由思考中国文化的当代命运。 [2]\r\n《什么是科学》第二版增添了“对批评的答复”和“地方性知识的哲学和社会学”两个附录，增加了索引，正文重新做了润色，统一了个别术语，改正了个别说法。', 1, '9787100207904', '商务印书馆', 1672416000, 309, 85, 64, 35, 99, 9900, 159, 1, 0, '2023-06-26 17:25:47', '2023-06-26 17:25:51', NULL);
INSERT INTO `book` VALUES (2, 1, '软件测试', '闫川川', 'http://rwj48creq.hb-bkt.clouddn.com/book/de3ccfb07110241c.jpg', '', '该教材以软件测试技术为主要研究对象，讲解了软件测试的基本原理、基本方法、基本技术、基本标准和规范、基本理论及基本软件测试工具。全书共8章，主要内容包括软件测试概论、软件测试基本知识、黑盒测试、白盒测试、软件测试流程、性能测试、软件测试自动化和软件测试管理。', 1, '9787302473299', '清华大学出版社', 1501516800, 146, 0, 0, 0, 9, 99, 100, 0, 1, '2023-07-10 06:06:46', '2023-07-10 06:35:55', NULL);
INSERT INTO `book` VALUES (4, 1, '测试001', '测试001', 'http://rwj48creq.hb-bkt.clouddn.com/book/1689239602924282.jpg', '', '测试001测试001测试001', 1, '37082719980409', '大明湖出版社', 1689177600, 159, 0, 0, 0, 15, 59, 99, 0, 35, '2023-07-13 07:58:06', '2023-07-13 09:13:25', NULL);
INSERT INTO `book` VALUES (5, 1, '测试0021', '测试0021', 'http://rwj48creq.hb-bkt.clouddn.com/book/16892361733355.jpg', '', '测试0021测试0021测试0021', 1, '1354789244', '444', 1689177600, 59, 0, 0, 0, 15, 59, 99, 0, 10, '2023-07-13 08:16:45', '2023-07-13 09:11:54', NULL);
INSERT INTO `book` VALUES (6, 1, '测试005', '测试005', 'http://rwj48creq.hb-bkt.clouddn.com/book/168923695800581.png', '', '测试005测试005测试005', 1, '测试005测试005', '测试005', 1689177600, 59, 1, 1, 0, 1, 15, 99, 0, 35, '2023-07-13 08:29:42', '2023-07-17 08:58:31', NULL);

-- ----------------------------
-- Table structure for book_sort
-- ----------------------------
DROP TABLE IF EXISTS `book_sort`;
CREATE TABLE `book_sort`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `book_id` int(10) UNSIGNED NOT NULL COMMENT '图书ID',
  `sort_id` int(10) UNSIGNED NOT NULL COMMENT '分类ID',
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of book_sort
-- ----------------------------
INSERT INTO `book_sort` VALUES (4, 1, 2, '2023-06-26 17:27:47', '2023-06-26 17:27:49', NULL);
INSERT INTO `book_sort` VALUES (5, 1, 1, '2023-06-28 16:07:59', '2023-06-28 16:08:27', NULL);
INSERT INTO `book_sort` VALUES (6, 2, 1, '2023-07-10 06:06:46', '2023-07-10 14:35:55', '2023-07-10 06:35:55');
INSERT INTO `book_sort` VALUES (7, 2, 2, '2023-07-10 06:06:46', '2023-07-10 14:35:55', '2023-07-10 06:35:55');
INSERT INTO `book_sort` VALUES (8, 2, 1, '2023-07-10 06:35:55', '2023-07-10 06:35:55', NULL);
INSERT INTO `book_sort` VALUES (9, 2, 2, '2023-07-10 06:35:55', '2023-07-10 06:35:55', NULL);
INSERT INTO `book_sort` VALUES (12, 6, 1, '2023-07-13 08:29:42', '2023-07-13 08:29:42', NULL);
INSERT INTO `book_sort` VALUES (13, 6, 2, '2023-07-13 08:29:42', '2023-07-13 08:29:42', NULL);
INSERT INTO `book_sort` VALUES (14, 5, 2, '2023-07-13 09:11:54', '2023-07-13 09:11:54', NULL);
INSERT INTO `book_sort` VALUES (15, 4, 1, '2023-07-13 09:13:25', '2023-07-13 09:13:25', NULL);

-- ----------------------------
-- Table structure for book_tag
-- ----------------------------
DROP TABLE IF EXISTS `book_tag`;
CREATE TABLE `book_tag`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `book_id` int(10) UNSIGNED NOT NULL COMMENT '图书ID',
  `tag_id` int(10) UNSIGNED NOT NULL COMMENT '标签ID',
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of book_tag
-- ----------------------------
INSERT INTO `book_tag` VALUES (4, 1, 4, '2023-06-26 17:28:00', '2023-06-26 17:28:02', NULL);
INSERT INTO `book_tag` VALUES (5, 2, 4, '2023-07-10 06:06:46', '2023-07-10 14:35:55', '2023-07-10 06:35:55');
INSERT INTO `book_tag` VALUES (6, 2, 4, '2023-07-10 06:35:55', '2023-07-10 06:35:55', NULL);
INSERT INTO `book_tag` VALUES (9, 6, 2, '2023-07-13 08:29:42', '2023-07-13 08:29:42', NULL);
INSERT INTO `book_tag` VALUES (10, 6, 3, '2023-07-13 08:29:42', '2023-07-13 08:29:42', NULL);
INSERT INTO `book_tag` VALUES (11, 6, 4, '2023-07-13 08:29:42', '2023-07-13 08:29:42', NULL);
INSERT INTO `book_tag` VALUES (12, 5, 2, '2023-07-13 09:11:54', '2023-07-13 09:11:54', NULL);
INSERT INTO `book_tag` VALUES (13, 5, 3, '2023-07-13 09:11:54', '2023-07-13 09:11:54', NULL);
INSERT INTO `book_tag` VALUES (14, 5, 4, '2023-07-13 09:11:54', '2023-07-13 09:11:54', NULL);

-- ----------------------------
-- Table structure for collect
-- ----------------------------
DROP TABLE IF EXISTS `collect`;
CREATE TABLE `collect`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `book_id` int(11) NOT NULL,
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of collect
-- ----------------------------
INSERT INTO `collect` VALUES (1, 4, 6, '2023-07-17 03:06:32', '2023-07-17 03:06:32', '2023-07-17 03:07:59');
INSERT INTO `collect` VALUES (2, 4, 6, '2023-07-17 03:08:11', '2023-07-17 03:08:11', '2023-07-17 08:58:06');
INSERT INTO `collect` VALUES (3, 4, 6, '2023-07-17 08:58:28', '2023-07-17 08:58:28', '2023-07-17 08:58:28');
INSERT INTO `collect` VALUES (4, 4, 6, '2023-07-17 08:58:29', '2023-07-17 08:58:29', '2023-07-17 08:58:30');
INSERT INTO `collect` VALUES (5, 4, 6, '2023-07-17 08:58:31', '2023-07-17 08:58:31', NULL);

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `user_id` int(11) NOT NULL,
  `book_id` int(11) NOT NULL,
  `context` text CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (6, '2023-07-17 10:22:56', '2023-07-17 10:22:56', NULL, 4, 6, '你好呀');

-- ----------------------------
-- Table structure for copy
-- ----------------------------
DROP TABLE IF EXISTS `copy`;
CREATE TABLE `copy`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of copy
-- ----------------------------

-- ----------------------------
-- Table structure for language
-- ----------------------------
DROP TABLE IF EXISTS `language`;
CREATE TABLE `language`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `language` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '语种',
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of language
-- ----------------------------
INSERT INTO `language` VALUES (1, '中文简体', '2023-06-26 11:36:31', '2023-06-27 14:33:50', NULL);
INSERT INTO `language` VALUES (2, '中文繁体', '2023-06-26 11:36:44', '2023-06-27 14:34:02', NULL);
INSERT INTO `language` VALUES (3, '英语', '2023-06-26 17:27:30', '2023-06-27 14:34:08', NULL);

-- ----------------------------
-- Table structure for like
-- ----------------------------
DROP TABLE IF EXISTS `like`;
CREATE TABLE `like`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `book_id` int(11) NOT NULL,
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of like
-- ----------------------------
INSERT INTO `like` VALUES (1, 4, 6, '2023-07-17 02:48:12', '2023-07-17 02:48:12', '2023-07-17 02:48:22');
INSERT INTO `like` VALUES (2, 4, 6, '2023-07-17 05:59:41', '2023-07-17 05:59:41', '2023-07-17 05:59:42');
INSERT INTO `like` VALUES (3, 4, 6, '2023-07-17 05:59:45', '2023-07-17 05:59:45', '2023-07-17 05:59:46');
INSERT INTO `like` VALUES (4, 4, 6, '2023-07-17 05:59:49', '2023-07-17 05:59:49', '2023-07-17 06:00:12');
INSERT INTO `like` VALUES (5, 4, 6, '2023-07-17 06:00:13', '2023-07-17 06:00:13', '2023-07-17 06:02:04');
INSERT INTO `like` VALUES (6, 4, 6, '2023-07-17 06:13:44', '2023-07-17 06:13:44', '2023-07-17 06:15:45');
INSERT INTO `like` VALUES (7, 4, 6, '2023-07-17 06:22:49', '2023-07-17 06:22:49', '2023-07-17 06:22:54');
INSERT INTO `like` VALUES (8, 4, 6, '2023-07-17 06:22:55', '2023-07-17 06:22:55', '2023-07-17 06:25:11');
INSERT INTO `like` VALUES (9, 4, 6, '2023-07-17 06:43:00', '2023-07-17 06:43:00', '2023-07-17 07:15:54');
INSERT INTO `like` VALUES (10, 4, 6, '2023-07-17 07:16:00', '2023-07-17 07:16:00', '2023-07-17 08:52:45');
INSERT INTO `like` VALUES (11, 4, 6, '2023-07-17 08:52:46', '2023-07-17 08:52:46', '2023-07-17 08:52:48');
INSERT INTO `like` VALUES (12, 4, 6, '2023-07-17 08:52:49', '2023-07-17 08:52:49', '2023-07-17 08:58:30');
INSERT INTO `like` VALUES (13, 4, 6, '2023-07-17 08:58:31', '2023-07-17 08:58:31', NULL);

-- ----------------------------
-- Table structure for sort
-- ----------------------------
DROP TABLE IF EXISTS `sort`;
CREATE TABLE `sort`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sort
-- ----------------------------
INSERT INTO `sort` VALUES (-1, 0, '全部', '2023-06-26 11:22:52', '2023-06-26 11:22:56', NULL);
INSERT INTO `sort` VALUES (1, 0, '生活', '2023-06-26 11:23:18', '2023-06-26 11:23:19', NULL);
INSERT INTO `sort` VALUES (2, 0, '科学', '2023-06-26 17:26:34', '2023-06-26 17:26:36', NULL);
INSERT INTO `sort` VALUES (3, 0, '测试001', '2023-07-11 07:47:32', '2023-07-11 15:55:31', '2023-07-11 07:55:31');
INSERT INTO `sort` VALUES (4, 0, '测试02', '2023-07-14 01:24:27', '2023-07-14 09:28:33', '2023-07-14 01:28:33');

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_at` datetime(0) NOT NULL,
  `updated_at` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0),
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tag
-- ----------------------------
INSERT INTO `tag` VALUES (2, '电影', '2023-06-26 11:36:31', '2023-06-26 11:36:33', NULL);
INSERT INTO `tag` VALUES (3, '音乐', '2023-06-26 11:36:44', '2023-06-26 11:36:46', NULL);
INSERT INTO `tag` VALUES (4, '教育', '2023-06-26 17:27:30', '2023-06-26 17:27:34', NULL);
INSERT INTO `tag` VALUES (5, '测试001', '2023-07-10 09:10:47', '2023-07-10 17:21:16', '2023-07-10 09:21:16');
INSERT INTO `tag` VALUES (6, '111', '2023-07-14 07:43:11', '2023-07-14 15:44:00', '2023-07-14 07:44:01');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `role` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '角色',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态 0正常 1关闭',
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '账户名称',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '密码',
  `nickname` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '昵称',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '头像',
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '描述',
  `wish` int(11) NOT NULL COMMENT '希望？',
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '邮箱',
  `mobile` varchar(11) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '手机号',
  `score` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '分数？',
  `balance` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '余额',
  `user_salt` varchar(20) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '加密盐',
  `push_switch` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0关闭 1开启',
  `created_at` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  `updated_at` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '修改时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `id`(`id`) USING BTREE,
  INDEX `email`(`email`) USING BTREE,
  INDEX `user_name`(`user_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (4, 'user', 0, '363301318@qq.com', '0369d0aca4d4811a350d8a407d6fc520', '郭星铄', 'http://rwj48creq.hb-bkt.clouddn.com/avatar/1687331569167274.jpg', '祝我早日完成.太好了', 0, '363301318@qq.com', '1825705629', '', 0, '7tCjVWVRwg', 0, '2023-06-21 15:52:40', '2023-06-21 07:52:41', NULL);
INSERT INTO `user` VALUES (5, 'admin', 0, '363301318@qq.com', '0369d0aca4d4811a350d8a407d6fc520', '郭星铄', 'http://rwj48creq.hb-bkt.clouddn.com/avatar/1687331569167274.jpg', '祝我早日完成.太好了', 0, '363301318@qq.com', '1825705629', '', 0, '7tCjVWVRwg', 0, '2023-06-21 15:52:40', '2023-06-21 07:52:41', NULL);
INSERT INTO `user` VALUES (6, 'admin', 1, 'admin', '079994d099689d4f5a03fd84c300dea6', 'Admin', '', '', 0, '123456789@qq.com', '18254705629', '', 0, 'ZXpjxf6yMI', 0, '2023-07-10 16:13:51', '2023-07-10 08:13:51', NULL);
INSERT INTO `user` VALUES (7, 'admin', 0, 'ZhuTest001', '182b540ef87c69962016fa2be574cabd', '测试朱001', '', '', 0, '363301318@qq.com', '18254705629', '', 0, 'cmWXjpXDRr', 0, '2023-07-14 16:38:13', '2023-07-14 16:38:13', '2023-07-14 08:38:14');

SET FOREIGN_KEY_CHECKS = 1;
