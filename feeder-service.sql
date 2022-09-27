-- MySQL dump 10.13  Distrib 8.0.30, for macos12.4 (arm64)
--
-- Host: 127.0.0.1    Database: feeder-service
-- ------------------------------------------------------
-- Server version	8.0.30

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
-- Table structure for table `feeders`
--

DROP TABLE IF EXISTS `feeders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `feeders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL,
  `barcode` varchar(45) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `feeders`
--

LOCK TABLES `feeders` WRITE;
/*!40000 ALTER TABLE `feeders` DISABLE KEYS */;
INSERT INTO `feeders` VALUES (1,'968ec500-c7f1-4527-a416-fb50bf567674','00001-AL03005090R-SMIT','eFishery_0001','2022-09-19 07:44:49','2022-09-19 07:46:32'),(2,'f179b6bc-1efd-428b-b692-8758ca8c21c9','00002-AL03005090R-F3ot','eFishery_0002','2022-09-19 07:45:01','2022-09-19 07:46:42'),(5,'de82a743-aae5-4e43-89a7-a8fc87402060','00005-AL15005090R-dxbm','eFishery_0003','2022-09-19 07:45:54','2022-09-19 07:46:49');
/*!40000 ALTER TABLE `feeders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pond_feeders`
--

DROP TABLE IF EXISTS `pond_feeders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pond_feeders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `pond_uuid` varchar(36) NOT NULL,
  `feeder_uuid` varchar(36) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_pond_feeders_ponds` (`pond_uuid`),
  KEY `fk_pond_feeders_feeders` (`feeder_uuid`),
  CONSTRAINT `fk_pond_feeders_feeders` FOREIGN KEY (`feeder_uuid`) REFERENCES `feeders` (`uuid`),
  CONSTRAINT `fk_pond_feeders_ponds` FOREIGN KEY (`pond_uuid`) REFERENCES `ponds` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pond_feeders`
--

LOCK TABLES `pond_feeders` WRITE;
/*!40000 ALTER TABLE `pond_feeders` DISABLE KEYS */;
INSERT INTO `pond_feeders` VALUES (1,'26f1b9ee-65d9-4c7d-afb1-f7137fefa784','968ec500-c7f1-4527-a416-fb50bf567674'),(2,'26f1b9ee-65d9-4c7d-afb1-f7137fefa784','f179b6bc-1efd-428b-b692-8758ca8c21c9'),(3,'06691a99-f10a-4d79-8298-75b9e344c0f6','de82a743-aae5-4e43-89a7-a8fc87402060');
/*!40000 ALTER TABLE `pond_feeders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ponds`
--

DROP TABLE IF EXISTS `ponds`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ponds` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ponds`
--

LOCK TABLES `ponds` WRITE;
/*!40000 ALTER TABLE `ponds` DISABLE KEYS */;
INSERT INTO `ponds` VALUES (1,'26f1b9ee-65d9-4c7d-afb1-f7137fefa784','Kolam Gurame','2022-09-19 07:50:48','2022-09-19 07:50:48'),(2,'06691a99-f10a-4d79-8298-75b9e344c0f6','Kolam Nila','2022-09-19 07:51:18','2022-09-19 07:51:18');
/*!40000 ALTER TABLE `ponds` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schedules`
--

DROP TABLE IF EXISTS `schedules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schedules` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` varchar(36) NOT NULL,
  `pond_uuid` varchar(36) NOT NULL,
  `time_start` time DEFAULT NULL,
  `time_end` time DEFAULT NULL,
  `duration_run` int DEFAULT NULL,
  `duration_pause` int DEFAULT NULL,
  `schedule_type` enum('basic','continues','advance') DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid` (`uuid`),
  KEY `fk_schedules_ponds` (`pond_uuid`),
  CONSTRAINT `fk_schedules_ponds` FOREIGN KEY (`pond_uuid`) REFERENCES `ponds` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schedules`
--

LOCK TABLES `schedules` WRITE;
/*!40000 ALTER TABLE `schedules` DISABLE KEYS */;
INSERT INTO `schedules` VALUES (1,'77e25f2f-d497-475d-8553-1346fccbf1ef','26f1b9ee-65d9-4c7d-afb1-f7137fefa784','09:24:00','09:24:40',NULL,NULL,'basic','2022-09-19 07:54:41'),(4,'8f8ef701-3967-45a6-981d-c4b15ffd7614','26f1b9ee-65d9-4c7d-afb1-f7137fefa784','12:00:00','12:00:01',NULL,NULL,'basic','2022-09-19 08:01:50'),(5,'39b0fa0c-a566-4611-8e56-08ffb7f37f48','06691a99-f10a-4d79-8298-75b9e344c0f6','09:24:05','09:24:45',NULL,NULL,'basic','2022-09-19 08:03:07');
/*!40000 ALTER TABLE `schedules` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-09-27 17:36:43
