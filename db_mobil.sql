-- Adminer 4.7.5 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `tb_jenis_mobil`;
CREATE TABLE `tb_jenis_mobil` (
  `id_jenis` int(11) NOT NULL AUTO_INCREMENT,
  `nm_jenis` varchar(100) NOT NULL,
  PRIMARY KEY (`id_jenis`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `tb_jenis_mobil` (`id_jenis`, `nm_jenis`) VALUES
(1,	'Low Cost Green Car'),
(2,	'Low Sport Utility Vehicle'),
(3,	'Sport Utility Vehicle'),
(4,	'Multi Purpose Vehicle'),
(5,	'Sedan'),
(6,	'Truck'),
(7,	'Sport'),
(8,	'City Car');

DROP TABLE IF EXISTS `tb_mobil`;
CREATE TABLE `tb_mobil` (
  `id_mobil` int(11) NOT NULL AUTO_INCREMENT,
  `id_perusahaan` int(11) NOT NULL,
  `nm_mobil` varchar(100) NOT NULL,
  `jenis_penggerak` varchar(5) NOT NULL,
  `banyak_roda` int(11) NOT NULL,
  `id_jenis` int(11) NOT NULL,
  `harga` int(11) NOT NULL,
  PRIMARY KEY (`id_mobil`),
  KEY `id_perusahaan` (`id_perusahaan`),
  KEY `id_jenis` (`id_jenis`),
  CONSTRAINT `tb_mobil_ibfk_1` FOREIGN KEY (`id_perusahaan`) REFERENCES `tb_perusahaan` (`id_perusahaan`) ON DELETE CASCADE,
  CONSTRAINT `tb_mobil_ibfk_2` FOREIGN KEY (`id_jenis`) REFERENCES `tb_jenis_mobil` (`id_jenis`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `tb_mobil` (`id_mobil`, `id_perusahaan`, `nm_mobil`, `jenis_penggerak`, `banyak_roda`, `id_jenis`, `harga`) VALUES
(3,	1,	'Kijang Innova Venturer',	'FR',	4,	3,	675000000);

DROP TABLE IF EXISTS `tb_perusahaan`;
CREATE TABLE `tb_perusahaan` (
  `id_perusahaan` int(11) NOT NULL AUTO_INCREMENT,
  `nm_perusahaan` varchar(50) NOT NULL,
  `alamat` varchar(255) NOT NULL,
  PRIMARY KEY (`id_perusahaan`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `tb_perusahaan` (`id_perusahaan`, `nm_perusahaan`, `alamat`) VALUES
(1,	'Toyota',	''),
(2,	'Suzuki',	''),
(3,	'Honda',	'');

-- 2020-04-16 07:04:04
