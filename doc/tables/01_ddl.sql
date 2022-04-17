use stockv2;
alter table stock_basic add `is_block` varchar(1) NOT NULL default '-' COMMENT '板块' after is_option;
alter table stock_basic add `is_star` varchar(1) NOT NULL  COMMENT 'STAR' after is_option;
alter table stock_basic add `is_yestoday_hot` varchar(1) NOT NULL  COMMENT '昨日强势股' after is_option;
alter table stock_basic add `is_yestoday_hot_date` datetime  COMMENT '昨日强势股' after is_option;
