INSERT INTO der_ems.`device`
	(`id`,`modbusid`,`module_id`,`model_id`,`gw_id`,`power_capacity`,`extra_info`,`created_at`,`updated_at`)
VALUES
	(1,1,1,1,1,24,null,'2022-07-01 00:00:00','2022-07-01 00:00:00'),
	(2,253,1,3,1,22,null,'2022-07-01 00:00:00','2022-07-01 00:00:00'),
	(3,252,1,4,1,19.8,null,'2022-07-01 00:00:00','2022-07-01 00:00:00'),
	(4,251,1,5,1,20,'{"voltage": 51.2, "energyCapacity": 30, "chargingSources": "Solar + Grid", "reservedForGridOutagePercent": 20}','2022-07-01 00:00:00','2022-07-01 00:00:00'),
	(5,1,2,1,1,24,null,'2022-07-01 00:00:00','2022-07-01 00:00:00');