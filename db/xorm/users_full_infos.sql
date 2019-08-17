DROP VIEW IF EXISTS szga_users_full_infos ;

CREATE VIEW szga_users_full_infos AS 
SELECT
  ui.id AS id,
  th.terminal_id AS terminal_id,
  th.terminal_hw_id AS terminal_hardware_id,
  ui.card_type AS card_type,
  ui.card_name AS card_name,
  ui.card_version AS card_version,
  ui.card_sn AS card_sn,
  ui.cert_sn AS cert_sn,
  ui.user_name AS user_name,
  ui.user_id AS user_id,
  ui.police_number AS police_number,
  ui.police_type AS police_type,
  ui.user_depart AS user_depart,
  ui.user_zone AS user_zone,
  ui.user_org AS user_org,
  ui.digit_data_sn AS digit_data_sn,
  ui.digit_data_alg AS digit_data_alg,
  ui.digit_data_valid_time AS digit_data_valid_time,
  ui.reg_time AS reg_time,
  ui.deleted AS deleted,
  ui.deleted_time AS deleted_time,
  ui.telephone AS telephone,
  th.address_city AS address_city,
  th.address_district AS address_district,
  th.address_province AS address_province,
  th.heartbeat_time AS heartbeat_time,
  th.latitude AS latitude,
  th.longitude AS longitude,
  th.offline_time AS offline_time,
  th.online AS online,
  th.online_time AS online_time,
  th.state_change_time AS state_change_time,
  th.imsi AS imsi,
  th.imei AS imei,
  sog.name AS user_org_name 
FROM
  szga_users_infos ui 
  LEFT JOIN szga_terminals_infos ti 
    ON ti.police_number = ui.police_number 
  LEFT JOIN szga_terminal_heartbeat th 
    ON th.`terminal_hw_id` = ti.`terminal_hardware_id` 
  LEFT JOIN sys_organization sog 
    ON ui.user_org = sog.id 
ORDER BY ui.police_number ;

