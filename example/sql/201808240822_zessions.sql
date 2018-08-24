-- セッション
CREATE TABLE `zessions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `session_id` VARCHAR(255),
  `data` TEXT,
  `created_at` DATETIME,
  `updated_at` DATETIME,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
