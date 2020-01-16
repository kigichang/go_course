set names utf8;

grant all on mytest.* to 'abc'@'172.%' identified by '1234test';
grant all on mytest.* to 'abc'@'localhost' identified by '1234test';
flush privileges;

create database if not exists mytest charset utf8;
use mytest;

CREATE TABLE `member` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `info` text,
  `birthday` date DEFAULT NULL,
  `register` datetime NOT NULL,
  `login` datetime NOT NULL,
  `vip` char(1) NOT NULL DEFAULT 'A',
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8