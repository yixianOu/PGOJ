-- MySQL dump 10.13  Distrib 8.0.36, for Win64 (x86_64)
--
-- Host: localhost    Database: oj_micro
-- ------------------------------------------------------
-- Server version	8.0.36

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `judgestatus`
--

DROP TABLE IF EXISTS `judgestatus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `judgestatus` (
  `judge_id` int NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `problem_id` int NOT NULL,
  `oj` varchar(255) NOT NULL,
  `result` varchar(127) NOT NULL,
  `time` int NOT NULL,
  `memory` int NOT NULL,
  `length` int NOT NULL,
  `language` varchar(255) NOT NULL,
  `submittime` datetime DEFAULT CURRENT_TIMESTAMP,
  `judger` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `contest` int DEFAULT '0',
  `contestproblem` int DEFAULT '0',
  `code` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `testcase` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `message` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `problemtitle` varchar(255) NOT NULL,
  `rating` int DEFAULT '0',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `input_data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `sample_out_put` text,
  `user_out_put` text,
  PRIMARY KEY (`judge_id`),
  KEY `judgestatus_FK` (`problem_id`),
  KEY `judgestatus_FK_1` (`user_id`),
  CONSTRAINT `judgestatus_FK` FOREIGN KEY (`problem_id`) REFERENCES `problem` (`problem_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `judgestatus_FK_1` FOREIGN KEY (`user_id`) REFERENCES `user_login` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `judgestatus`
--

LOCK TABLES `judgestatus` WRITE;
/*!40000 ALTER TABLE `judgestatus` DISABLE KEYS */;
/*!40000 ALTER TABLE `judgestatus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `problem`
--

DROP TABLE IF EXISTS `problem`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `problem` (
  `problem_id` int NOT NULL AUTO_INCREMENT,
  `author` varchar(255) NOT NULL,
  `addtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `oj` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `des` text NOT NULL,
  `input` text NOT NULL,
  `output` text NOT NULL,
  `sinput` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `soutput` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `source` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `time` int NOT NULL,
  `memory` int NOT NULL,
  `hint` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `auth` tinyint(1) NOT NULL DEFAULT '0',
  `level` int NOT NULL,
  `problem_code` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`problem_id`),
  UNIQUE KEY `problem_un` (`title`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `problem`
--

LOCK TABLES `problem` WRITE;
/*!40000 ALTER TABLE `problem` DISABLE KEYS */;
INSERT INTO `problem` VALUES (1,'orician','2024-08-31 09:33:35','hnoj','problem_test','a test for create a problem','input_test','output_test','','','nil',5,10,'between 1 and 10',1,1,'p0001'),(2,'orician','2024-08-31 09:38:08','','problem_test_again','a test for create a problem','input_test','output_test','','','',5,10,'',1,1,'P0001');
/*!40000 ALTER TABLE `problem` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `problem_tag`
--

DROP TABLE IF EXISTS `problem_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `problem_tag` (
  `id` int NOT NULL AUTO_INCREMENT,
  `problem_id` int NOT NULL,
  `tag_id` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `problem` (`problem_id`,`tag_id`),
  KEY `tag_id` (`tag_id`),
  KEY `problem_id` (`problem_id`),
  CONSTRAINT `problem_tag_ibfk_1` FOREIGN KEY (`problem_id`) REFERENCES `problem` (`problem_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `problem_tag_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `tag` (`tag_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `problem_tag`
--

LOCK TABLES `problem_tag` WRITE;
/*!40000 ALTER TABLE `problem_tag` DISABLE KEYS */;
INSERT INTO `problem_tag` VALUES (1,1,1),(2,1,4),(8,1,5),(11,2,1),(3,2,5),(4,2,6);
/*!40000 ALTER TABLE `problem_tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `problemdata`
--

DROP TABLE IF EXISTS `problemdata`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `problemdata` (
  `problemdata_id` int NOT NULL AUTO_INCREMENT,
  `problem_id` int NOT NULL,
  `submission` int NOT NULL,
  `ac` int NOT NULL,
  `mle` int NOT NULL,
  `tle` int NOT NULL,
  `rte` int NOT NULL,
  `ole` int NOT NULL,
  `ce` int NOT NULL,
  `wa` int NOT NULL,
  `ue` int NOT NULL,
  `score` int NOT NULL,
  `auth` tinyint(1) NOT NULL,
  `sf` int NOT NULL,
  `fe` int NOT NULL,
  PRIMARY KEY (`problemdata_id`),
  UNIQUE KEY `problemdata_un` (`problem_id`),
  CONSTRAINT `problemdata_ibfk_1` FOREIGN KEY (`problem_id`) REFERENCES `problem` (`problem_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `problemdata`
--

LOCK TABLES `problemdata` WRITE;
/*!40000 ALTER TABLE `problemdata` DISABLE KEYS */;
INSERT INTO `problemdata` VALUES (1,1,0,0,0,0,0,0,0,0,0,10,1,0,0),(2,2,0,0,0,0,0,0,0,0,0,15,2,0,0);
/*!40000 ALTER TABLE `problemdata` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tag`
--

DROP TABLE IF EXISTS `tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tag` (
  `tag_id` int NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(255) NOT NULL,
  PRIMARY KEY (`tag_id`),
  UNIQUE KEY `tag_name` (`tag_name`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tag`
--

LOCK TABLES `tag` WRITE;
/*!40000 ALTER TABLE `tag` DISABLE KEYS */;
INSERT INTO `tag` VALUES (5,'hello world'),(1,'test'),(4,'test_again'),(6,'world');
/*!40000 ALTER TABLE `tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `testcases`
--

DROP TABLE IF EXISTS `testcases`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `testcases` (
  `test_id` int NOT NULL AUTO_INCREMENT,
  `problem_id` int NOT NULL,
  `test_group` int NOT NULL,
  `inputFileName` varchar(255) NOT NULL,
  `outputFileName` varchar(255) NOT NULL,
  PRIMARY KEY (`test_id`),
  UNIQUE KEY `testcases_uni` (`problem_id`,`test_group`),
  CONSTRAINT `testcases_ibfk_1` FOREIGN KEY (`problem_id`) REFERENCES `problem` (`problem_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `testcases`
--

LOCK TABLES `testcases` WRITE;
/*!40000 ALTER TABLE `testcases` DISABLE KEYS */;
INSERT INTO `testcases` VALUES (2,1,2,'1/input_file/P1607_2.in','1/output_file/P1607_2.out'),(3,1,1,'1/input_file/P1607_1.in','1/output_file/P1607_1.out');
/*!40000 ALTER TABLE `testcases` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_login`
--

DROP TABLE IF EXISTS `user_login`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_login` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `password` varchar(127) NOT NULL,
  `is_superuser` tinyint NOT NULL,
  `username` varchar(127) NOT NULL,
  `email` varchar(255) NOT NULL,
  `cover_image_url` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户认证授权';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_login`
--

LOCK TABLES `user_login` WRITE;
/*!40000 ALTER TABLE `user_login` DISABLE KEYS */;
INSERT INTO `user_login` VALUES (30,'$2a$10$3tY..p1aArsuiq5JXsr5Qeal7xTo4wWQlR5ReewzIBmZaLJp6/FMa',2,'oyx','2018783812@qq.com','https://oj-file.oss-cn-shenzhen.aliyuncs.com/1user_cover.jpg'),(33,'$2a$10$NfELzonJIiYn4tjRDYkxfu3U0wX4b8yCoMbVnMwmbnPlhc9J2Py.y',2,'orician','20223003671@hainanu.edu.cn','https://DOMAINNAME/oj-file/user_cover/33.png');
/*!40000 ALTER TABLE `user_login` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_profile`
--

DROP TABLE IF EXISTS `user_profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_profile` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `phone` varchar(127) DEFAULT NULL,
  `name` varchar(127) DEFAULT NULL,
  `AC_count` int NOT NULL DEFAULT '0',
  `submit_count` int NOT NULL DEFAULT '0',
  `score` int NOT NULL DEFAULT '0',
  `description` text,
  `rating` int unsigned DEFAULT '0',
  `AC_problem` text COMMENT '中间用竖线隔开',
  `school` varchar(127) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`),
  CONSTRAINT `user_profile` FOREIGN KEY (`user_id`) REFERENCES `user_login` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_profile`
--

LOCK TABLES `user_profile` WRITE;
/*!40000 ALTER TABLE `user_profile` DISABLE KEYS */;
INSERT INTO `user_profile` VALUES (12,30,'123456789','oyx',0,0,0,'just a test',0,NULL,'lzjzx'),(14,33,NULL,NULL,0,0,0,NULL,0,NULL,NULL);
/*!40000 ALTER TABLE `user_profile` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'oj_micro'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-09-22 11:44:57
