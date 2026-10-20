-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 20 Okt 2026 pada 09.18
-- Versi server: 10.4.28-MariaDB
-- Versi PHP: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `article`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `posts`
--

CREATE TABLE `posts` (
  `id` int(11) NOT NULL,
  `title` varchar(200) DEFAULT NULL,
  `content` text DEFAULT NULL,
  `category` varchar(100) DEFAULT NULL,
  `created_date` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `status` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `posts`
--

INSERT INTO `posts` (`id`, `title`, `content`, `category`, `created_date`, `updated_date`, `status`) VALUES
(8, 'Transformasi Digital di Sektor Perbankan Indonesia', 'Perkembangan teknologi mengubah cara masyarakat bertransaksi. Layanan mobile banking, dompet digital, dan pembayaran nirsentuh kini menjadi bagian dari kebiasaan sehari-hari. Bagi perbankan, kondisi ini menuntut pembaruan sistem inti, peningkatan keamanan data, serta pengalaman pengguna yang lebih sederhana. Institusi yang mampu beradaptasi dengan cepat akan lebih siap menghadapi persaingan di era digital.', 'Teknologi', '2026-10-20 07:01:35', '2026-10-20 07:08:23', 'publish'),
(10, 'Pemanfaatan Cloud Computing bagi Pelaku UMKM', 'Komputasi awan memungkinkan pelaku usaha kecil menggunakan infrastruktur teknologi tanpa harus membeli perangkat mahal. Penyimpanan data, aplikasi kasir, hingga layanan pelanggan dapat diakses dari mana saja dengan biaya yang mengikuti kebutuhan. Langkah awal yang bijak adalah memilih penyedia layanan yang andal, memahami kebijakan keamanan data, dan memulai dari kebutuhan yang paling mendesak.', 'Teknologi', '2026-10-20 07:02:51', '2026-10-20 07:04:54', 'thrash'),
(11, 'Transformasi Digital di Sektor Perbankan Indonesia', 'Perkembangan teknologi mengubah cara masyarakat bertransaksi. Layanan mobile banking, dompet digital, dan pembayaran nirsentuh kini menjadi bagian dari kebiasaan sehari-hari. Bagi perbankan, kondisi ini menuntut pembaruan sistem inti, peningkatan keamanan data, serta pengalaman pengguna yang lebih sederhana. Institusi yang mampu beradaptasi dengan cepat akan lebih siap menghadapi persaingan di era digital.', 'Teknologi', '2026-10-20 07:08:10', '2026-10-20 07:08:10', 'publish'),
(12, 'Transformasi Digital di Sektor Perbankan Indonesia', 'Perkembangan teknologi mengubah cara masyarakat bertransaksi. Layanan mobile banking, dompet digital, dan pembayaran nirsentuh kini menjadi bagian dari kebiasaan sehari-hari. Bagi perbankan, kondisi ini menuntut pembaruan sistem inti, peningkatan keamanan data, serta pengalaman pengguna yang lebih sederhana. Institusi yang mampu beradaptasi dengan cepat akan lebih siap menghadapi persaingan di era digital.', 'Teknologi', '2026-10-20 07:18:06', '2026-10-20 07:18:06', 'publish');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `posts`
--
ALTER TABLE `posts`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `posts`
--
ALTER TABLE `posts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
