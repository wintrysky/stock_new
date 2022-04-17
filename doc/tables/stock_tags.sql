USE stockv2;
create table tag_dictionary
(
	id                   	bigint not null AUTO_INCREMENT,
	tag_name            	varchar(64) not null comment 'tag name', 
	tag_desc				varchar(2000) not null comment 'description', 
	primary key (id)
);


create unique index uidx_tag_name on tag_dictionary(tag_name);


create table stock_tags
(
	id                   	bigint not null AUTO_INCREMENT,
	symbol            		varchar(100) not null comment 'symbol', 
	tag_name            	varchar(64) not null comment 'tag name', 
	primary key (id)
);

create unique index uidx_symbol_tag_name on stock_tags(symbol,tag_name);