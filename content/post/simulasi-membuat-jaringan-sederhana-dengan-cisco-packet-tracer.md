---
title: "Simulasi Membuat Jaringan Sederhana Dengan Cisco Packet Tracer"
slug: simulasi-membuat-jaringan-sederhana-dengan-cisco-packet-tracer
date: 2020-06-27T12:11:30+07:00
draft: false

type: post

tags:
    - cisco-packet-tracer
    - jaringan
    - computer-network

image: ""
description: "Simulasi Membuat Jaringan Sederhana dengan Cisco Packet Tracer"
---

Sebelum membuat jaringan komputer yang kompleks, ada baiknya apabila memahami terlebih dahulu untuk membuat jaringan sederhana. Di dalam membuat jaringan sederhana ini, menggunakan simulasi pada Cisco Packet Tracer. Dengan demikian, mahasiswa dapat memahami fundamental jaringan komputer.

Di dalam membuat jaringan sederhana, kita mulai dengan simulasi 2 PC dan satu switch. Kedua PC tersebut nantinya akan saling berkomunikasi satu dengan yang lainnya. Sebagai perantara antar PC, menggunakan switch yang berfungsi sebagai penghubung antara beberapa perangkat di dalam suatu jaringan komputer. Dengan demikian, switch akan menerima suatu data atau informasi dan akan diteruskan ke perangkat lainnya.

## Alat yang Digunakan

Alat yang digunakan di dalam melakukan praktikum jaringan komputer, khususnya pada modul satu ini diantaranya sebagai berikut.
- Komputer atau laptop.
- Cisco Packet Tracer.

## Langkah-langkah

- Buka Cisco Packet Tracer.

- Login terlebih dahulu menggunaan akun [Netacad.com](https://www.netacad.com/). Apabila belum memiliki akun, dapat membuat akun di website tersebut.

![Login Cisco Packet Tracer](/img/simulasi-membuat-jaringan-sederhana-dengan-cisco-packet-tracer/login-cisco-packet-tracer.png)

- Setelah login, akan muncul workspace Cisco Packet Tracer seperti berikut.

![Cisco Packet Tracer Workspace](/img/simulasi-membuat-jaringan-sederhana-dengan-cisco-packet-tracer/cisco-packet-tracer-workspace.png)

- Buatlah topologi jaringan seperti pada gambar di bawah.

![Network Topology](/img/simulasi-membuat-jaringan-sederhana-dengan-cisco-packet-tracer/network-topology.png)

- Lakukan konfigurasi IP pada PC-1 dengan cara klik pada PC-1, pilih tab Desktop, klik IP Coniguration. Pada field IP address, isikan nilai 192.168.15.1, sedangkan pada field subnet mask isikan nilai 255.255.255.0.

![PC-1 IP Configuration](/img/simulasi-membuat-jaringan-sederhana-dengan-cisco-packet-tracer/pc-1-ip-configuration.png)

- Lakukan konfigurasi yang sama pada PC-2.
![PC-2 IP Configuration](/img/simulasi-membuat-jaringan-sederhana-dengan-cisco-packet-tracer/pc-2-ip-configuration.png)

- Terakhir, lakukan tes dengan melakukan ping pada masing-masing PC. Klik pada PC, pilih desktop, lalu klik command prompt. PC-1 akan melakukan ping pada PC-2 dan sebaliknya.

![PC-1 Test Ping](/img/simulasi-membuat-jaringan-sederhana-dengan-cisco-packet-tracer/pc-1-ping.png)

![PC-2 Test Ping](/img/simulasi-membuat-jaringan-sederhana-dengan-cisco-packet-tracer/pc-2-ping.png)

## Pembahasan dan Analisa

Dari percobaan di atas, pertama-tama membuka Cisco Packet Tracer dan melakukan login. Selanjutnya, membuat topologi jaringan berupa 2 PC dan 1 switch. Lalu, memberikan ip address dan subnet mask pada masing-masing PC. Untuk mengetes konfigurasi, dapat dilakukan dengan melakukan ping pada masing-masing PC, jika ada reply atau jawaban dari PC tujuan, maka konfigurasi telah berhasil.

## Kesimpulan

Dari pembahasan dan analisa yang telah dilakukan, kita dapat menyimpulkan bahwa untuk dapat menghubungkan 2 PC agar dapat saling berkomunikasi, kita harus mengkonfigurasikan IP address dan subnet mask pada perangkat terkait.
