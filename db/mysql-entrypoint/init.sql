use monitor_db;

CREATE TABLE IF NOT EXISTS `resources`(
  `id` VARCHAR(36) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `http_method` VARCHAR(255) NOT NULL, 
  `url` VARCHAR(2048) NOT NULL, 
  `status_code` INT NOT NULL, 
  PRIMARY KEY(`id`)
);

CREATE TABLE IF NOT EXISTS `resources_logs`(
  `id` INT NOT NULL AUTO_INCREMENT, 
  `resource_id` VARCHAR(36) NOT NULL, 
  `message` VARCHAR(999) NOT NULL, 
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  PRIMARY KEY(`id`),
  KEY `fk_logs_resources_idx` (`resource_id`)
);

INSERT INTO `monitor_db`.`resources`(
  `id`,
  `name`,
  `url`,
  `http_method`,
  `status_code`
) VALUES (
  "6b4c28f4-6831-495a-9444-19c93452faa3",
  "Random OK",
  "https://httpstat.us/200",
  "GET",
  200
);

INSERT INTO `monitor_db`.`resources`(
  `id`,
  `name`,
  `url`,
  `http_method`,
  `status_code`
) VALUES (
  "dad6f8fb-3149-4d0b-861e-03d29c6accf0",
  "Random Maybe OK",
  "httpstat.us/random/200,200,201,400-404",
  "GET",
  200
);

INSERT INTO `monitor_db`.`resources`(
  `id`,
  `name`,
  `url`,
  `http_method`,
  `status_code`
) VALUES (
  "66805003-f9a2-4772-b577-d5ff66207707",
  "Random UNHEALTHY",
  "httpstat.us/random/400-422,500-504",
  "GET",
  200
);