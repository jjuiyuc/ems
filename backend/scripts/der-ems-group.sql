INSERT INTO der_ems.`group`
	(`id`,`name`,`type_id`,`parent_id`,`created_at`,`updated_at`)
VALUES
    (1,'Admin',1,null,UTC_TIMESTAMP(),UTC_TIMESTAMP()),
	(2,'AreaOwner_TW',2,1,UTC_TIMESTAMP(),UTC_TIMESTAMP()),
	(3,'AreaMaintainer_TW',3,2,UTC_TIMESTAMP(),UTC_TIMESTAMP()),
	(4,'Serenegray',4,2,UTC_TIMESTAMP(),UTC_TIMESTAMP()),
	(5,'CHT_Yuanli',4,2,UTC_TIMESTAMP(),UTC_TIMESTAMP()),
	(6,'Huayu',4,2,UTC_TIMESTAMP(),UTC_TIMESTAMP());