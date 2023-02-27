-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.7.3-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Dumping structure for table test.transaction
CREATE TABLE IF NOT EXISTS `transaction` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `trans_no` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `type` varchar(1) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Deposit/Withdraw',
  `account_no` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `amount` int(11) DEFAULT NULL,
  `trans_dt` datetime DEFAULT NULL,
  `created_dt` datetime DEFAULT NULL,
  `created_by` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `updated_dt` datetime DEFAULT NULL,
  `updated_by` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `account_no` (`account_no`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table test.transaction: ~15 rows (approximately)
/*!40000 ALTER TABLE `transaction` DISABLE KEYS */;
INSERT INTO `transaction` (`id`, `trans_no`, `type`, `account_no`, `amount`, `trans_dt`, `created_dt`, `created_by`, `updated_dt`, `updated_by`) VALUES
	(1, '', 'D', '1234567890', 10, '2023-02-27 19:19:34', '2023-02-27 19:19:34', 'admin', '0000-00-00 00:00:00', ''),
	(2, '', 'D', '1234567890', 10, '2023-02-27 19:44:47', '2023-02-27 19:44:47', 'admin', '0000-00-00 00:00:00', ''),
	(3, '', 'D', '1234567890', 10, '2023-02-27 19:45:11', '2023-02-27 19:45:11', 'admin', '0000-00-00 00:00:00', ''),
	(4, '', 'D', '1234567890', 10, '2023-02-27 19:46:25', '2023-02-27 19:46:25', 'admin', '0000-00-00 00:00:00', ''),
	(5, '', 'D', '1234567890', 10, '2023-02-27 19:49:04', '2023-02-27 19:49:04', 'admin', '0000-00-00 00:00:00', ''),
	(6, '', 'D', '1234567890', 10, '2023-02-27 19:50:11', '2023-02-27 19:50:11', 'admin', '0000-00-00 00:00:00', ''),
	(7, '', 'D', '4444444888', 10, '2023-02-27 19:51:14', '2023-02-27 19:51:14', 'admin', '0000-00-00 00:00:00', ''),
	(8, '', 'D', '4444444888', 10, '2023-02-27 19:51:57', '2023-02-27 19:51:57', 'admin', '0000-00-00 00:00:00', ''),
	(9, '', 'D', '4444444888', 10, '2023-02-27 19:52:56', '2023-02-27 19:52:56', 'admin', '0000-00-00 00:00:00', ''),
	(10, '', 'D', '4444444888', 10, '2023-02-27 19:55:56', '2023-02-27 19:55:56', 'admin', '0000-00-00 00:00:00', ''),
	(11, '', 'D', '4444444888', 10, '2023-02-27 19:57:38', '2023-02-27 19:57:38', 'admin', '0000-00-00 00:00:00', ''),
	(12, '', 'D', '4444444888', 10, '2023-02-27 19:58:32', '2023-02-27 19:58:32', 'admin', '0000-00-00 00:00:00', ''),
	(13, '', 'D', '4444444888', 10, '2023-02-27 19:59:08', '2023-02-27 19:59:08', 'admin', '0000-00-00 00:00:00', ''),
	(14, '', 'D', '4444444888', 10, '2023-02-27 19:59:29', '2023-02-27 19:59:29', 'admin', '0000-00-00 00:00:00', ''),
	(15, '', 'D', '4444444888', 10, '2023-02-27 19:59:42', '2023-02-27 19:59:42', 'admin', '0000-00-00 00:00:00', ''),
	(16, '', 'D', '4444444888', 10, '2023-02-27 21:08:17', '2023-02-27 21:08:17', 'admin', '0000-00-00 00:00:00', ''),
	(17, '', 'D', '4444444888', 10, '2023-02-27 21:08:17', '2023-02-27 21:08:17', 'admin', '0000-00-00 00:00:00', ''),
	(18, '', 'D', '4444444888', 10, '2023-02-27 21:09:51', '2023-02-27 21:09:51', 'admin', '0000-00-00 00:00:00', ''),
	(19, '', 'D', '4444444888', 10, '2023-02-27 21:10:12', '2023-02-27 21:10:12', 'admin', '0000-00-00 00:00:00', ''),
	(20, '', 'W', '1234444444', 20, '2023-02-27 21:14:45', '2023-02-27 21:14:45', 'admin', '0000-00-00 00:00:00', ''),
	(21, '', 'W', '1234444444', 20, '2023-02-27 21:16:00', '2023-02-27 21:16:00', 'admin', '0000-00-00 00:00:00', ''),
	(22, '', 'W', '1234444444', 20, '2023-02-27 21:18:50', '2023-02-27 21:18:50', 'admin', '0000-00-00 00:00:00', ''),
	(23, '', 'D', '1111110000', 500, '2023-02-27 21:46:40', '2023-02-27 21:46:40', 'admin', '0000-00-00 00:00:00', ''),
	(24, '', 'D', '1111112200', 800, '2023-02-27 21:51:55', '2023-02-27 21:51:55', 'admin', '0000-00-00 00:00:00', '');
/*!40000 ALTER TABLE `transaction` ENABLE KEYS */;

-- Dumping structure for table test.wallet
CREATE TABLE IF NOT EXISTS `wallet` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `citizen_id` varchar(13) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `title` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `firstname` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `lastname` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `address` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `subdistrict` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `district` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `province` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `zipcode` varchar(5) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `tel` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_dt` datetime DEFAULT NULL,
  `created_by` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `updated_dt` datetime DEFAULT NULL,
  `updated_by` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `citizen_id` (`citizen_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table test.wallet: ~1 rows (approximately)
/*!40000 ALTER TABLE `wallet` DISABLE KEYS */;
INSERT INTO `wallet` (`id`, `citizen_id`, `title`, `firstname`, `lastname`, `address`, `subdistrict`, `district`, `province`, `zipcode`, `tel`, `created_dt`, `created_by`, `updated_dt`, `updated_by`) VALUES
	(1, '2209141648126', 'นาย', 'มานะ', 'สีดำ', 'ถนนบางนา', 'บางนา', 'บางนา', 'กรุงเทพ', '54120', '104444', '2023-02-27 17:19:14', 'admin', '0000-00-00 00:00:00', ''),
	(3, '2209141648125', 'นาย', 'สมสุข', 'สีดำ', 'ถนนบางนา', 'บางนา', 'บางนา', 'กรุงเทพ', '54120', '104444', '2023-02-27 17:20:49', 'admin', '0000-00-00 00:00:00', ''),
	(7, '2209141648100', 'นาง', 'งามตา', 'แก้วมา', 'ถนนบางนา', 'บางนา', 'บางนา', 'กรุงเทพ', '54120', '104444', '2023-02-27 21:46:40', 'admin', '0000-00-00 00:00:00', ''),
	(9, '2209141648200', 'นางสาว', 'สมหญิง', 'ใจจันทร์', 'ถนนสุขุมวิท', 'พระโขนง', 'บางนา', 'กรุงเทพ', '54120', '104444', '2023-02-27 21:51:55', 'admin', '0000-00-00 00:00:00', '');
/*!40000 ALTER TABLE `wallet` ENABLE KEYS */;

-- Dumping structure for table test.wallet_account
CREATE TABLE IF NOT EXISTS `wallet_account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `citizen_id` varchar(13) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `account_no` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `balance` decimal(10,2) DEFAULT NULL,
  `status` varchar(1) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Activate/Deactivate',
  `remark` varchar(250) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_dt` datetime DEFAULT NULL,
  `created_by` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `updated_dt` datetime DEFAULT NULL,
  `updated_by` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `cid` (`citizen_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Dumping data for table test.wallet_account: ~3 rows (approximately)
/*!40000 ALTER TABLE `wallet_account` DISABLE KEYS */;
INSERT INTO `wallet_account` (`id`, `citizen_id`, `account_no`, `balance`, `status`, `remark`, `created_dt`, `created_by`, `updated_dt`, `updated_by`) VALUES
	(1, '2209141648126', '1234567890', 1060.00, 'A', 'a', '2023-02-27 18:11:13', 'admin', '2023-02-27 19:50:11', 'admin'),
	(2, '2209141648126', '1234444444', 24500.00, 'A', 'b', '2023-02-27 18:11:14', 'admin', '2023-02-27 22:10:00', 'admin'),
	(3, '2209141648120', '4444444888', 1151.60, 'A', 'c', '2023-02-27 18:11:15', 'admin', '2023-02-27 21:10:12', 'admin'),
	(4, '2209141648100', '1111110000', 500.00, 'A', '', '2023-02-27 21:46:40', 'admin', '2023-02-27 21:46:40', 'admin'),
	(5, '2209141648200', '1111112200', 800.00, 'A', '', '2023-02-27 21:51:55', 'admin', '2023-02-27 21:51:55', 'admin');
/*!40000 ALTER TABLE `wallet_account` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
