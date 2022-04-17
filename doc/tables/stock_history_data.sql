use stockv2;
CREATE TABLE `stock_history_data` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `symbol` varchar(100) DEFAULT NULL,
  `symbol_id` bigint(20) NOT NULL,
  `open_price` decimal(12,3) DEFAULT NULL,
  `low_price` decimal(12,3) DEFAULT NULL,
  `high_price` decimal(12,3) DEFAULT NULL,
  `close_price` decimal(12,3) DEFAULT NULL,
  `trade_time` varchar(20) DEFAULT NULL COMMENT 'yyyyMMdd',
  `trade_time_d` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uidx_history_data_2` (`symbol`,`trade_time`),
  KEY `idx_history_data_1` (`symbol_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
