CREATE TABLE `trade_pair` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `price` float(20,10) COMMENT '价格',
  `amt` float(20,10)  COMMENT '待成交',
  `total` float(20,10)  COMMENT '待成交+已成交',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='交易对信息';