-- Adminer 4.7.5 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `tb_mobil`;
CREATE TABLE `tb_mobil` (
  `id_mobil` int(11) NOT NULL AUTO_INCREMENT,
  `kode_mobil` varchar(10) NOT NULL,
  `merk` varchar(50) NOT NULL,
  `tipe` varchar(100) NOT NULL,
  `harga` int(11) NOT NULL,
  `warna` varchar(20) NOT NULL,
  `penggerak` varchar(4) NOT NULL,
  `banyak_roda` int(11) NOT NULL,
  `banyak_bangku` int(11) NOT NULL,
  `jenis_mesin` varchar(20) NOT NULL,
  `mesin` varchar(20) NOT NULL,
  PRIMARY KEY (`id_mobil`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `tb_mobil` (`id_mobil`, `kode_mobil`, `merk`, `tipe`, `harga`, `warna`, `penggerak`, `banyak_roda`, `banyak_bangku`, `jenis_mesin`, `mesin`) VALUES
(1,	'T001',	'Toyota',	'Kijang Innova Venturer',	675000000,	'Merah',	'FR',	4,	6,	'Diesel',	'D4D');

-- 2020-04-13 09:53:35
