#!/bin/bash

echo "Memulai backup log...."

mkdir -p "backup-log"

if [ ! -f syslog.txt ]; then
    echo "[ERROR] File syslog tidak ditemukan di $(pwd)"
    exit 1
fi

if cp "syslog.txt" "backup-log/"; then
    echo "[SUKSES] Berhasil dibackup"
else
    echo "[ERROR] Gagal dibackup"
    exit 1
fi