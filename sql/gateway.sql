-- MySQL dump 10.13  Distrib 5.7.23, for Linux (x86_64)
--
-- Host: localhost    Database: gateway
-- ------------------------------------------------------
-- Server version	5.7.23

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `service_api`
--

DROP TABLE IF EXISTS `service_api`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `service_api` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `service_url_id` int(11) unsigned NOT NULL COMMENT '服务地址',
  `method` varchar(50) COLLATE utf8_unicode_ci NOT NULL COMMENT '请求方式',
  `api_alias` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '网关别名',
  `api_path` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '网关配置转发地址',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_api_alias` (`method`,`api_alias`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_api`
--

LOCK TABLES `service_api` WRITE;
/*!40000 ALTER TABLE `service_api` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_api` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `service_url`
--

DROP TABLE IF EXISTS `service_url`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `service_url` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `service_name` varchar(50) COLLATE utf8_unicode_ci NOT NULL COMMENT '服务名称',
  `service_url` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '服务地址',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`service_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `service_url`
--

LOCK TABLES `service_url` WRITE;
/*!40000 ALTER TABLE `service_url` DISABLE KEYS */;
/*!40000 ALTER TABLE `service_url` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '用户名称',
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '邮箱',
  `password` varchar(60) COLLATE utf8_unicode_ci NOT NULL COMMENT '密码',
  `phone` char(11) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '手机号',
  `real_name` char(10) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '真实姓名',
  `last_login_ip` char(16) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '上次登陆ip',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `last_login_time` int(11) NOT NULL COMMENT '上一次登陆的时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (3,'admin','123456@qq.com','e10adc3949ba59abbe56e057f20f883e','','','0.0.0.0','2018-12-02 15:07:45','2020-03-27 08:14:45',1585296885);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-03-27  8:16:43
