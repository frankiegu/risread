-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.7.22 - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  10.1.0.5464
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- 导出 ris_read 的数据库结构
CREATE DATABASE IF NOT EXISTS `ris_read` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `ris_read`;

-- 导出  表 ris_read.book_attainment 结构
CREATE TABLE IF NOT EXISTS `book_attainment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_info_id` int(11) DEFAULT NULL,
  `book_info_id` int(11) DEFAULT NULL,
  `publish_time` datetime DEFAULT NULL,
  `content` varchar(500) DEFAULT NULL,
  `scan_times` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='读书心得';

-- 数据导出被取消选择。
-- 导出  表 ris_read.book_info 结构
CREATE TABLE IF NOT EXISTS `book_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_info_id` int(11) DEFAULT NULL,
  `link` varchar(300) DEFAULT NULL,
  `name` varchar(300) DEFAULT NULL,
  `copyright` varchar(100) DEFAULT NULL,
  `cover` varchar(100) DEFAULT NULL,
  `content_legal` tinyint(4) DEFAULT NULL,
  `publish_time` datetime DEFAULT NULL,
  `audit_time` datetime DEFAULT NULL,
  `reward` tinyint(4) DEFAULT NULL,
  `download_times` int(11) DEFAULT NULL,
  `author` varchar(50) DEFAULT NULL,
  `introduction` varchar(300) DEFAULT NULL,
  `book_type_id` int(11) DEFAULT NULL,
  `save_name` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8;

-- 数据导出被取消选择。
-- 导出  表 ris_read.book_info2_book_list 结构
CREATE TABLE IF NOT EXISTS `book_info2_book_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `book_info_id` int(11) DEFAULT NULL,
  `book_list_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- 数据导出被取消选择。
-- 导出  表 ris_read.book_info_comment 结构
CREATE TABLE IF NOT EXISTS `book_info_comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `book_info_id` int(11) DEFAULT NULL,
  `user_info_id` int(11) DEFAULT NULL,
  `content` varchar(500) DEFAULT NULL,
  `publish_time` datetime DEFAULT NULL,
  `scan_times` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- 数据导出被取消选择。
-- 导出  表 ris_read.book_list 结构
CREATE TABLE IF NOT EXISTS `book_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_info_id` int(11) DEFAULT NULL,
  `book_list_type_id` int(11) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `instruction` varchar(500) DEFAULT NULL,
  `publish_time` datetime DEFAULT NULL,
  `publish` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=66 DEFAULT CHARSET=utf8;

-- 数据导出被取消选择。
-- 导出  表 ris_read.book_list_book_infos 结构
CREATE TABLE IF NOT EXISTS `book_list_book_infos` (
  `book_list_id` int(11) NOT NULL,
  `book_info_id` int(11) NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- 数据导出被取消选择。
-- 导出  表 ris_read.book_list_comment 结构
CREATE TABLE IF NOT EXISTS `book_list_comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `book_list_id` int(11) DEFAULT NULL,
  `user_info_id` int(11) DEFAULT NULL,
  `content` varchar(500) DEFAULT NULL,
  `publish_time` datetime DEFAULT NULL,
  `mark` int(11) NOT NULL DEFAULT '0' COMMENT '评论的分数',
  `scan_times` int(11) NOT NULL DEFAULT '0' COMMENT '浏览的次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='书单的评论';

-- 数据导出被取消选择。
-- 导出  表 ris_read.book_list_favorite 结构
CREATE TABLE IF NOT EXISTS `book_list_favorite` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `book_list_id` int(11) DEFAULT NULL,
  `user_info_id` int(11) DEFAULT NULL,
  `favorite_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='收藏书单';

-- 数据导出被取消选择。
-- 导出  表 ris_read.book_list_type 结构
CREATE TABLE IF NOT EXISTS `book_list_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- 数据导出被取消选择。
-- 导出  表 ris_read.book_type 结构
CREATE TABLE IF NOT EXISTS `book_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 数据导出被取消选择。
-- 导出  表 ris_read.upload_book 结构
CREATE TABLE IF NOT EXISTS `upload_book` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_info_id` int(11) NOT NULL,
  `book_name` varchar(250) DEFAULT NULL,
  `upload_time` datetime DEFAULT NULL,
  `save_name` varchar(250) DEFAULT NULL,
  `size` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- 数据导出被取消选择。
-- 导出  表 ris_read.user_info 结构
CREATE TABLE IF NOT EXISTS `user_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT 'ris_reader',
  `riches` int(11) NOT NULL DEFAULT '0',
  `phone` varchar(13) NOT NULL DEFAULT 'ris_reader',
  `password` varchar(20) NOT NULL,
  `gender` varchar(10) NOT NULL DEFAULT '不公开',
  `signature` varchar(250) DEFAULT NULL,
  `email` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`id`,`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- 数据导出被取消选择。
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
