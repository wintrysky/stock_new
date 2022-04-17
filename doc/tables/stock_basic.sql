use stockv2;
CREATE TABLE `stock_basic` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `symbol` varchar(100) NOT NULL COMMENT '代码',
  `name` varchar(200) NOT NULL COMMENT '名称',
  `industry` varchar(100) NOT NULL DEFAULT '-' COMMENT '所属行业',
  `is_china` varchar(1) NOT NULL DEFAULT '-' COMMENT '中概股',
  `is_hk` varchar(1) NOT NULL DEFAULT '' COMMENT '香港',
  `is_hot` varchar(1) NOT NULL DEFAULT '-' COMMENT '热门',
  `is_etf` varchar(1) NOT NULL DEFAULT '-' COMMENT 'ETF',
  `is_risk` varchar(1) NOT NULL DEFAULT '-' COMMENT '政策风险',
  `is_option` varchar(1) NOT NULL COMMENT '是否期权',
  `is_too_high` varchar(1) NOT NULL COMMENT '高位可空',
  `blacklist_flag` varchar(1) NOT NULL DEFAULT '-' COMMENT '是否黑名单',
  `description` text COMMENT '描述',
  `open_price` decimal(12,3) NOT NULL COMMENT '开盘',
  `low_price` decimal(12,3) NOT NULL COMMENT '最低',
  `high_price` decimal(12,3) NOT NULL COMMENT '最高',
  `current_price` decimal(12,3) NOT NULL COMMENT '最新价',
  `yesterday_price` decimal(12,3) NOT NULL COMMENT '昨收',
  `pe` decimal(12,3) NOT NULL COMMENT '市盈率(静)',
  `total_market_cap` decimal(12,3) NOT NULL COMMENT '总市值',
  `increase_rate_yesterday` decimal(12,3) DEFAULT NULL COMMENT '涨跌幅',
  `increase_rate_5day` decimal(12,3) DEFAULT NULL COMMENT '5日涨跌幅',
  `increase_rate_10day` decimal(12,3) DEFAULT NULL COMMENT '10日涨跌幅',
  `increase_rate_20day` decimal(12,3) DEFAULT NULL COMMENT '20日涨跌幅',
  `increase_rate_60day` decimal(12,3) DEFAULT NULL COMMENT '60日涨跌幅',
  `increase_rate_120day` decimal(12,3) DEFAULT NULL COMMENT '120日涨跌幅',
  `increase_rate_250day` decimal(12,3) DEFAULT NULL COMMENT '250日涨跌幅',
  `increase_rate_form_year` decimal(12,3) DEFAULT NULL COMMENT '年初至今涨跌幅',
  `trade_date` datetime DEFAULT NULL COMMENT '股票交易日',
  `calculate_date` datetime DEFAULT NULL COMMENT '计算周涨跌幅完成时间,如果无效时间,表示不可交易',
  `last_week_rate` decimal(12,3) DEFAULT NULL COMMENT '上周涨跌率,0.1表示涨了10%,-0.1表示跌了10%',
  `compare_yes_last_week` decimal(12,3) DEFAULT NULL COMMENT '昨日收盘对比上周最高点涨跌幅0.1表示涨了10%,-0.1表示跌了10%',
  `create_time` datetime NOT NULL,
  `con_id` bigint(20) DEFAULT NULL COMMENT 'IB股票标记值',
  `sec_type` varchar(20) DEFAULT NULL COMMENT 'STK/OPT',
  `prim_exchange` varchar(150) DEFAULT NULL COMMENT 'NASDAQ.NMS/NYSE',
  `exchange` varchar(20) DEFAULT NULL COMMENT 'ISLAND/SMART',
  `currency` varchar(20) NOT NULL COMMENT 'USD',
  `enabled` varchar(1) NOT NULL DEFAULT 'Y' COMMENT '是否可用',
  `wt_moat` varchar(1) NOT NULL DEFAULT '-' COMMENT '护城河',
  `wt_tenfold` varchar(1) NOT NULL DEFAULT '-' COMMENT '十倍股',
  `wt_good_financial` varchar(1) NOT NULL DEFAULT '-' COMMENT '优质财报',
  `wt_weight` varchar(1) NOT NULL DEFAULT '-' COMMENT '权重股',
  `wt_low_pe` varchar(1) NOT NULL DEFAULT '-' COMMENT '低市盈率',
  `wt_high_occupy` varchar(1) NOT NULL DEFAULT '-' COMMENT '高占有率',
  `wt_special_business` varchar(1) NOT NULL DEFAULT '-' COMMENT '独特业务',
  `wt_frontier_technology` varchar(1) NOT NULL DEFAULT '-' COMMENT '前沿科技,如VR、AI',
  `wt_good_industry` varchar(1) NOT NULL DEFAULT '-' COMMENT '前景行业，如电动车，新能源',
  `bt_little_market` varchar(1) NOT NULL DEFAULT '-' COMMENT '小市值',
  `bt_debt` varchar(1) NOT NULL DEFAULT '-' COMMENT '亏损',
  `bt_bad_industry` varchar(1) NOT NULL DEFAULT '-' COMMENT '夕阳行业',
  `update_time` datetime NOT NULL,
  `increase_rate_curr_day` decimal(12,3) DEFAULT NULL COMMENT '涨跌幅',
  `ikey` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`,`enabled`),
  UNIQUE KEY `uidx_stock_basic_symbol` (`symbol`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;