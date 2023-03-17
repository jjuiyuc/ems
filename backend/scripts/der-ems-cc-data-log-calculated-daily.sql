INSERT INTO der_ems.`cc_data_log_calculated_daily`
	(`gw_uuid`,`latest_log_date`,`pv_produced_lifetime_energy_ac_diff`,`load_consumed_lifetime_energy_ac_diff`,`battery_lifetime_energy_ac_diff`,`grid_lifetime_energy_ac_diff`,`load_self_consumed_energy_percent_ac`,`off_peak_period_pre_ubiik_cost`,`off_peak_period_post_ubiik_cost`,`on_peak_period_pre_ubiik_cost`,`on_peak_period_post_ubiik_cost`,`mid_peak_period_pre_ubiik_cost`,`mid_peak_period_post_ubiik_cost`,`created_at`,`updated_at`)
VALUES
	('0E0BA27A8175AF978C49396BDE9D7A1E','2022-07-31 15:59:00',5,10,15,5,10,302.03,265.48,0,0,0,0,'2022-07-31 16:15:00','2022-07-31 16:15:00'),
	('0E0BA27A8175AF978C49396BDE9D7A1E','2022-08-01 15:59:15',10,20,30,15,15,114.84,114.81,296.6,226.11,0,0,'2022-08-01 16:15:00','2022-08-01 16:15:00'),
	('0E0BA27A8175AF978C49396BDE9D7A1E','2022-08-02 15:59:30',15,30,45,25,20,87.43,87.4,299.95,209.36,0,0,'2022-08-02 16:15:00','2022-08-02 16:15:00');
