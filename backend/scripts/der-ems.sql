-- MySQL dump 10.13  Distrib 8.0.29, for Linux (x86_64)
--
-- Host: localhost    Database: der_ems
-- ------------------------------------------------------
-- Server version	8.0.27

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
-- Table structure for table `ai_data`
--

DROP TABLE IF EXISTS `ai_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ai_data` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `gw_uuid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `log_date` datetime NOT NULL,
  `gw_id` bigint DEFAULT NULL,
  `location_id` bigint DEFAULT NULL,
  `local_ai_data` json DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `gw_uuid_log_date_UNIQUE` (`gw_uuid`,`log_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cc_data`
--

DROP TABLE IF EXISTS `cc_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cc_data` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `gw_uuid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `log_date` datetime NOT NULL,
  `gw_id` bigint DEFAULT NULL,
  `location_id` bigint DEFAULT NULL,
  `local_cc_data` json DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `gw_uuid_log_date_UNIQUE` (`gw_uuid`,`log_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cc_data_log`
--

DROP TABLE IF EXISTS `cc_data_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cc_data_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `gw_uuid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `log_date` datetime NOT NULL,
  `gw_id` bigint DEFAULT NULL,
  `location_id` bigint DEFAULT NULL,
  `grid_is_peak_shaving` int DEFAULT NULL,
  `load_grid_average_power_ac` float DEFAULT NULL,
  `battery_grid_average_power_ac` float DEFAULT NULL,
  `grid_contract_power_ac` float DEFAULT NULL,
  `load_pv_average_power_ac` float DEFAULT NULL,
  `load_battery_average_power_ac` float DEFAULT NULL,
  `battery_soc` float DEFAULT NULL,
  `battery_produced_average_power_ac` float DEFAULT NULL,
  `battery_consumed_average_power_ac` float DEFAULT NULL,
  `battery_charging_from` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `battery_discharging_to` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `pv_average_power_ac` float DEFAULT NULL,
  `load_average_power_ac` float DEFAULT NULL,
  `load_links` json DEFAULT NULL,
  `grid_links` json DEFAULT NULL,
  `pv_links` json DEFAULT NULL,
  `battery_links` json DEFAULT NULL,
  `battery_pv_average_power_ac` float DEFAULT NULL,
  `grid_pv_average_power_ac` float DEFAULT NULL,
  `grid_produced_average_power_ac` float DEFAULT NULL,
  `grid_consumed_average_power_ac` float DEFAULT NULL,
  `battery_lifetime_operation_cycles` float DEFAULT NULL,
  `battery_produced_lifetime_energy_ac` float DEFAULT NULL,
  `battery_consumed_lifetime_energy_ac` float DEFAULT NULL,
  `battery_average_power_ac` float DEFAULT NULL,
  `battery_voltage` float DEFAULT NULL,
  `all_produced_lifetime_energy_ac` float DEFAULT NULL,
  `pv_produced_lifetime_energy_ac` float DEFAULT NULL,
  `grid_produced_lifetime_energy_ac` float DEFAULT NULL,
  `all_consumed_lifetime_energy_ac` float DEFAULT NULL,
  `load_consumed_lifetime_energy_ac` float DEFAULT NULL,
  `grid_consumed_lifetime_energy_ac` float DEFAULT NULL,
  `grid_average_power_ac` float DEFAULT NULL,
  `battery_lifetime_energy_ac` float DEFAULT NULL,
  `grid_lifetime_energy_ac` float DEFAULT NULL,
  `load_self_consumed_lifetime_energy_ac` float DEFAULT NULL,
  `grid_power_cost` float DEFAULT NULL,
  `grid_power_cost_savings` float DEFAULT NULL,
  `load_pv_consumed_lifetime_energy_ac` float DEFAULT NULL,
  `battery_pv_consumed_lifetime_energy_ac` float DEFAULT NULL,
  `grid_pv_consumed_lifetime_energy_ac` float DEFAULT NULL,
  `pv_energy_cost_savings` float DEFAULT NULL,
  `pv_co2_savings` float DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `gw_uuid_log_date_UNIQUE` (`gw_uuid`,`log_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cc_data_log_calculated_daily`
--

DROP TABLE IF EXISTS `cc_data_log_calculated_daily`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cc_data_log_calculated_daily` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `gw_uuid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `latest_log_date` datetime NOT NULL,
  `gw_id` bigint DEFAULT NULL,
  `location_id` bigint DEFAULT NULL,
  `pv_produced_lifetime_energy_ac_diff` float DEFAULT NULL,
  `load_consumed_lifetime_energy_ac_diff` float DEFAULT NULL,
  `battery_lifetime_energy_ac_diff` float DEFAULT NULL,
  `grid_lifetime_energy_ac_diff` float DEFAULT NULL,
  `load_self_consumed_energy_percent_ac` float DEFAULT NULL,
  `off_peak_period_pre_ubiik_cost` float DEFAULT NULL,
  `off_peak_period_post_ubiik_cost` float DEFAULT NULL,
  `on_peak_period_pre_ubiik_cost` float DEFAULT NULL,
  `on_peak_period_post_ubiik_cost` float DEFAULT NULL,
  `mid_peak_period_pre_ubiik_cost` float DEFAULT NULL,
  `mid_peak_period_post_ubiik_cost` float DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `gw_uuid_latest_log_date_UNIQUE` (`gw_uuid`,`latest_log_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `cc_data_log_calculated_monthly`
--

DROP TABLE IF EXISTS `cc_data_log_calculated_monthly`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cc_data_log_calculated_monthly` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `gw_uuid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `latest_log_date` datetime NOT NULL,
  `gw_id` bigint DEFAULT NULL,
  `location_id` bigint DEFAULT NULL,
  `pv_produced_lifetime_energy_ac_diff` float DEFAULT NULL,
  `load_consumed_lifetime_energy_ac_diff` float DEFAULT NULL,
  `battery_lifetime_energy_ac_diff` float DEFAULT NULL,
  `grid_lifetime_energy_ac_diff` float DEFAULT NULL,
  `load_self_consumed_energy_percent_ac` float DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `gw_uuid_latest_log_date_UNIQUE` (`gw_uuid`,`latest_log_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `location`
--

DROP TABLE IF EXISTS `location`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `location` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `customer_number` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `field_number` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `lat` double DEFAULT NULL,
  `lng` double DEFAULT NULL,
  `weather_lat` float DEFAULT NULL,
  `weather_lng` float DEFAULT NULL,
  `timezone` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `tou_location_id` bigint DEFAULT NULL,
  `voltage_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `tou_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `customer_number_field_number_UNIQUE` (`customer_number`,`field_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `device`
--

DROP TABLE IF EXISTS `device`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `device` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `modbusid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `uueid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `model_id` bigint NOT NULL,
  `gw_uuid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `remark` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `enable` tinyint(1) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `modbusid_uueid_UNIQUE` (`modbusid`,`uueid`),
  KEY `device_model_id_device_model_id_foreign` (`model_id`),
  KEY `device_gw_uuid_gateway_uuid_foreign` (`gw_uuid`),
  CONSTRAINT `device_gw_uuid_gateway_uuid_foreign` FOREIGN KEY (`gw_uuid`) REFERENCES `gateway` (`uuid`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `device_model_id_device_model_id_foreign` FOREIGN KEY (`model_id`) REFERENCES `device_model` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `device_model`
--

DROP TABLE IF EXISTS `device_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `device_model` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `device_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `model_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `capacity` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `gateway`
--

DROP TABLE IF EXISTS `gateway`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `gateway` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uuid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `location_id` bigint DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uuid_UNIQUE` (`uuid`),
  KEY `gateway_location_id_location_id_foreign` (`location_id`),
  CONSTRAINT `gateway_location_id_location_id_foreign` FOREIGN KEY (`location_id`) REFERENCES `location` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `login_log`
--

DROP TABLE IF EXISTS `login_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `login_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tou`
--

DROP TABLE IF EXISTS `tou`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tou` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `tou_location_id` int DEFAULT NULL,
  `voltage_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `tou_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `period_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `peak_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `is_summer` tinyint(1) DEFAULT NULL,
  `period_stime` time DEFAULT NULL,
  `period_etime` time DEFAULT NULL,
  `basic_charge` float DEFAULT NULL,
  `basic_rate` float DEFAULT NULL,
  `flow_rate` float DEFAULT NULL,
  `enable_at` date DEFAULT NULL,
  `disable_at` date DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tou_holiday`
--

DROP TABLE IF EXISTS `tou_holiday`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tou_holiday` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `tou_location_id` bigint DEFAULT NULL,
  `year` year DEFAULT NULL,
  `day` date DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tou_location`
--

DROP TABLE IF EXISTS `tou_location`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tou_location` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `power_company` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `location` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password_last_changed` datetime DEFAULT NULL,
  `password_retry_count` int DEFAULT '0',
  `reset_pwd_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `pwd_token_expiry` datetime DEFAULT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `locked_at` datetime DEFAULT NULL,
  `expiration_date` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user_gateway_right`
--

DROP TABLE IF EXISTS `user_gateway_right`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_gateway_right` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint DEFAULT NULL,
  `gw_id` bigint DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id_gw_id_UNIQUE` (`user_id`,`gw_id`),
  KEY `user_gateway_right_user_id_user_id_foreign` (`user_id`),
  KEY `user_gateway_right_gw_id_gateway_id_foreign` (`gw_id`),
  CONSTRAINT `user_gateway_right_gw_id_gateway_id_foreign` FOREIGN KEY (`gw_id`) REFERENCES `gateway` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `user_gateway_right_user_id_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `weather_forecast`
--

DROP TABLE IF EXISTS `weather_forecast`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `weather_forecast` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `lat` float NOT NULL,
  `lng` float NOT NULL,
  `alt` float DEFAULT NULL,
  `valid_date` datetime NOT NULL,
  `data` json DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `lat_lng_valid_date_UNIQUE` (`lat`,`lng`,`valid_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-06-02 11:03:15
