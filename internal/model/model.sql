CREATE TABLE `1min_trade_data` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `time` bigint unsigned COMMENT '交易时间',
  `open` float(20,10) COMMENT '开盘价',
  `close` float(20,10)  COMMENT '收盘价',
  `highest_price` float(20,10)  COMMENT '最高价',
  `lowest_price` float(20,10)  COMMENT '最低价',
  `volume` float(10,5) unsigned COMMENT '交易量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='1min的k线信息';

CREATE TABLE `5min_trade_data` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `time` bigint unsigned COMMENT '交易时间',
  `open` float(20,10) COMMENT '开盘价',
  `close` float(20,10)  COMMENT '收盘价',
  `highest_price` float(20,10)  COMMENT '最高价',
  `lowest_price` float(20,10)  COMMENT '最低价',
  `volume` float(10,5) unsigned COMMENT '交易量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='5min的k线信息';