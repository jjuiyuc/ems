INSERT INTO der_ems.`tou`
	(`tou_location_id`,`voltage_type`,`tou_type`,`period_type`,`peak_type`,`is_summer`,`period_stime`,`period_etime`,`basic_charge`,`basic_rate`,`flow_rate`,`enable_at`,`disable_at`,`created_at`,`updated_at`)
VALUES
	(1,'Low voltage','Two-section','Weekdays','On-peak',1,'07:30:00','22:30:00',262.5,236.2,3.42,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'Low voltage','Two-section','Weekdays','Off-peak',1,'00:00:00','07:30:00',262.5,47.2,1.46,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'Low voltage','Two-section','Weekdays','Off-peak',1,'22:30:00','00:00:00',262.5,47.2,1.46,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'Low voltage','Two-section','Saturday','Mid-peak',1,'07:30:00','22:30:00',262.5,47.2,2.14,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'Low voltage','Two-section','Saturday','Off-peak',1,'00:00:00','07:30:00',262.5,47.2,1.46,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'Low voltage','Two-section','Saturday','Off-peak',1,'22:30:00','00:00:00',262.5,47.2,1.46,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'Low voltage','Two-section','Sunday & Holiday','Off-peak',1,'00:00:00','00:00:00',262.5,47.2,1.46,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'Low voltage','Two-section','Weekdays','On-peak',0,'07:30:00','22:30:00',262.5,173.2,3.33,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'Low voltage','Two-section','Weekdays','Off-peak',0,'22:30:00','00:00:00',262.5,34.6,1.39,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24');

INSERT INTO der_ems.`tou`
	(`tou_location_id`,`voltage_type`,`tou_type`,`period_type`,`peak_type`,`is_summer`,`period_stime`,`period_etime`,`basic_charge`,`basic_rate`,`flow_rate`,`enable_at`,`disable_at`,`created_at`,`updated_at`)
VALUES
	(1,'Low voltage','Two-section','Saturday','Mid-peak',0,'07:30:00','22:30:00',262.5,34.6,2.06,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'Low voltage','Two-section','Saturday','Off-peak',0,'00:00:00','07:30:00',262.5,34.6,1.39,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'Low voltage','Two-section','Saturday','Off-peak',0,'22:30:00','00:00:00',262.5,34.6,1.39,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'Low voltage','Two-section','Sunday & Holiday','Off-peak',0,'00:00:00','00:00:00',262.5,34.6,1.39,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'High voltage','Two-section','Weekdays','On-peak',1,'07:30:00','22:30:00',0.0,223.6,3.29,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'High voltage','Two-section','Weekdays','Off-peak',1,'00:00:00','07:30:00',0.0,44.7,1.41,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'High voltage','Two-section','Weekdays','Off-peak',1,'22:30:00','00:00:00',0.0,44.7,1.41,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'High voltage','Two-section','Saturday','Mid-peak',1,'07:30:00','22:30:00',0.0,44.7,1.97,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'High voltage','Two-section','Saturday','Off-peak',1,'00:00:00','07:30:00',0.0,44.7,1.41,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'High voltage','Two-section','Saturday','Off-peak',1,'22:30:00','00:00:00',0.0,44.7,1.41,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24');

INSERT INTO der_ems.`tou`
	(`tou_location_id`,`voltage_type`,`tou_type`,`period_type`,`peak_type`,`is_summer`,`period_stime`,`period_etime`,`basic_charge`,`basic_rate`,`flow_rate`,`enable_at`,`disable_at`,`created_at`,`updated_at`)
VALUES
	(1,'High voltage','Two-section','Sunday & Holiday','Off-peak',1,'00:00:00','00:00:00',0.0,44.7,1.41,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'High voltage','Two-section','Weekdays','On-peak',0,'07:30:00','22:30:00',0.0,166.9,3.17,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'High voltage','Two-section','Weekdays','Off-peak',0,'00:00:00','07:30:00',0.0,33.3,1.31,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'High voltage','Two-section','Weekdays','Off-peak',0,'22:30:00','00:00:00',0.0,33.3,1.31,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'High voltage','Two-section','Saturday','Mid-peak',0,'07:30:00','22:30:00',0.0,33.3,1.87,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'High voltage','Two-section','Saturday','Off-peak',0,'00:00:00','07:30:00',0.0,33.3,1.31,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'High voltage','Two-section','Saturday','Off-peak',0,'22:30:00','00:00:00',0.0,33.3,1.31,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24'),
	(1,'High voltage','Two-section','Sunday & Holiday','Off-peak',0,'00:00:00','00:00:00',0.0,33.3,1.31,'2022-01-01','2023-12-31','2022-07-05 15:51:24','2022-07-05 15:51:24');
