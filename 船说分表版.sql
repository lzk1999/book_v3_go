/*
ShipSayCMS QQGroup: 249310348 user: admin password: shipsay
*/

CREATE TABLE IF NOT EXISTS `shipsay_article_chapter_1` (
  `chapterid` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `articleid` int(11) unsigned NOT NULL DEFAULT '0',
  `articlename` varchar(50) NOT NULL DEFAULT '',
  `posterid` smallint(6) unsigned NOT NULL DEFAULT '0',
  `poster` varchar(30) NOT NULL DEFAULT '',
  `postdate` int(11) unsigned NOT NULL DEFAULT '0',
  `lastupdate` int(11) unsigned NOT NULL DEFAULT '0',
  `chaptername` varchar(100) NOT NULL DEFAULT '',
  `chapterorder` smallint(6) unsigned NOT NULL DEFAULT '0',
  `words` int(11) unsigned NOT NULL DEFAULT '0',
  `chaptertype` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `attachment` text,
  `summary` text,
  `isimage` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `volumeid` int(11) unsigned NOT NULL DEFAULT '0',
  `pushed` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`chapterid`),
  KEY `lastupdate` (`lastupdate`),
  KEY `articleid` (`articleid`,`chapterorder`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `shipsay_article_article` (
  `articleid` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `siteid` int(11) unsigned NOT NULL DEFAULT '0',
  `postdate` int(11) unsigned NOT NULL DEFAULT '0',
  `lastupdate` int(11) unsigned NOT NULL DEFAULT '0',
  `infoupdate` int(11) unsigned NOT NULL DEFAULT '0',
  `articlename` varchar(50) NOT NULL DEFAULT '',
  `articlecode` varchar(200) NOT NULL DEFAULT '',
  `backupname` varchar(100) NOT NULL DEFAULT '',
  `keywords` varchar(100) NOT NULL DEFAULT '',
  `roles` varchar(200) NOT NULL DEFAULT '',
  `initial` char(1) NOT NULL DEFAULT '',
  `authorid` int(11) unsigned NOT NULL DEFAULT '0',
  `author` varchar(30) NOT NULL DEFAULT '',
  `posterid` int(11) unsigned NOT NULL DEFAULT '0',
  `poster` varchar(30) NOT NULL DEFAULT '',
  `sortid` smallint(3) unsigned NOT NULL DEFAULT '0',
  `typeid` smallint(3) unsigned NOT NULL DEFAULT '0',
  `intro` text,
  `lastchapterid` int(11) unsigned NOT NULL DEFAULT '0',
  `lastchapter` varchar(100) NOT NULL DEFAULT '',
  `lastsummary` text,
  `chapters` smallint(6) unsigned NOT NULL DEFAULT '0',
  `words` int(11) unsigned NOT NULL DEFAULT '0',
  `lastvisit` int(11) unsigned NOT NULL DEFAULT '0',
  `dayvisit` int(11) unsigned NOT NULL DEFAULT '0',
  `weekvisit` int(11) unsigned NOT NULL DEFAULT '0',
  `monthvisit` int(11) unsigned NOT NULL DEFAULT '0',
  `allvisit` int(11) unsigned NOT NULL DEFAULT '0',
  `lastvote` int(11) unsigned NOT NULL DEFAULT '0',
  `dayvote` int(11) unsigned NOT NULL DEFAULT '0',
  `weekvote` int(11) unsigned NOT NULL DEFAULT '0',
  `monthvote` int(11) unsigned NOT NULL DEFAULT '0',
  `allvote` int(11) unsigned NOT NULL DEFAULT '0',
  `goodnum` int(11) unsigned NOT NULL DEFAULT '0', 
  `fullflag` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `imgflag` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `display` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `rgroup` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `lastvolume` varchar(100) NOT NULL DEFAULT '',
  `lastvolumeid` int(11) unsigned NOT NULL DEFAULT '0',
  `pushed` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `ratenum` int(11) unsigned NOT NULL DEFAULT '0',
  `ratesum` int(11) unsigned NOT NULL DEFAULT '0', 
  PRIMARY KEY (`articleid`),
  KEY `articlename` (`articlename`),
  KEY `lastupdate` (`lastupdate`),
  KEY `allvisit` (`allvisit`) USING BTREE,
  KEY `monthvisit` (`monthvisit`) USING BTREE,
  KEY `weekvisit` (`dayvisit`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `shipsay_article_bookcase` (
  `caseid` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `articleid` int(11) unsigned NOT NULL DEFAULT '0',
  `articlename` varchar(50) NOT NULL DEFAULT '',
  `userid` int(11) unsigned NOT NULL DEFAULT '0',
  `username` varchar(30) NOT NULL DEFAULT '',
  `chapterid` int(11) unsigned NOT NULL DEFAULT '0',
  `chaptername` varchar(100) NOT NULL DEFAULT '',
  PRIMARY KEY (`caseid`),
  KEY `articleid` (`articleid`),
  KEY `userid` (`userid`),
  KEY `chapterid` (`chapterid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS  `shipsay_article_report` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `articleid` int(11) NOT NULL,
  `chapterid` int(11) NOT NULL,
  `articlename` varchar(100) NOT NULL,
  `chaptername` varchar(200) NOT NULL,
  `repurl` varchar(100) NOT NULL,
  `ip` varchar(25) NOT NULL,
  `reptime` int(11) NOT NULL,
  `content` varchar(200) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `chapterid` (`chapterid`),
  KEY `reptime` (`reptime`),
  KEY `articlename` (`articlename`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS  `shipsay_article_search` (
  `searchid` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `searchtime` int(11) NOT NULL,
  `keywords` varchar(100) NOT NULL,
  `results` smallint(6) NOT NULL,
  `searchsite` varchar(100) NOT NULL,
  PRIMARY KEY (`searchid`),
  KEY `searchtime` (`searchtime`),
  KEY `keywords` (`keywords`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `shipsay_system_users` (
  `uid` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uname` varchar(30) NOT NULL DEFAULT '',
  `name` varchar(60) NOT NULL DEFAULT '',
  `pass` varchar(32) NOT NULL DEFAULT '',
  `salt` varchar(32) NOT NULL DEFAULT '',
  `groupid` tinyint(3) NOT NULL DEFAULT '0',
  `regdate` int(11) unsigned NOT NULL DEFAULT '0',
  `lastlogin` int(11) unsigned NOT NULL DEFAULT '0',
  `email` varchar(60) NOT NULL,
  PRIMARY KEY (`uid`),
  UNIQUE KEY `uname` (`uname`),
  KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `shipsay_universal_push` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `url` varchar(255) CHARACTER SET utf8mb4 NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `url` (`url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS  `shipsay_article_langtail` (
  `langid` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `sourceid` int(11) NOT NULL,
  `langname` varchar(50) NOT NULL DEFAULT '',
  `sourcename` varchar(50) NOT NULL DEFAULT '',
  `uptime` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`langid`),
  KEY `sourceid` (`sourceid`,`langid`),
  UNIQUE KEY `langname`(`langname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `shipsay_system_users` VALUES ('1', 'admin', 'admin', 'e3b561a8265d317f58a046bc07ebe755', 'ffbcbba18762184e', '2', '1588248146', '1588248146', 'admin@shipsay.com');