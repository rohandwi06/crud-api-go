create table `mahasiswas` (
    `id` int unsigned auto_increment primary key,
    `nama` varchar(255),
    `nim` varchar(12) unique,
    `prodi` varchar(255),
    `kelas` varchar(255),
    `created_at` DATETIME(3) NULL,
    `updated_at` DATETIME(3) NULL,
    `deleted_at` DATETIME(3) NULL

);