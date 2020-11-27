fpm -s dir -t rpm -n gobackup --rpm-os linux -v v1.0-centos7 \
  ./gobackup=/usr/bin/ \
  ./gobackup.yml=/etc/gobackup/gobackup.yml  \
  ./gobackup.service=/usr/lib/systemd/system/gobackup.service