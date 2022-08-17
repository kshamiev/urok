
apt update
apt install systemctl
apt install systemd
apt install wget
apt install mysql-client
apt install sudo
apt install mc

wget https://repo.manticoresearch.com/repository/manticoresearch_jammy/dists/manticore_5.0.2-220530-348514c86_amd64.tgz https://repo.manticoresearch.com/repository/manticoresearch_jammy/dists/jammy/main/binary-amd64/manticore-columnar-lib_1.15.4-220522-2fef34e_amd64.deb
tar -xvf manticore_5.0.2-220530-348514c86_amd64.tgz
sudo dpkg -i manticore-*5.0.2-220530-348514c86_amd64.deb manticore-columnar-lib_1.15.4-220522-2fef34e_amd64.deb

при запуске севриса вылезает вот такая ошибка:

systemctl start manticore

System has not been booted with systemd as init system (PID 1). Can't operate.
Failed to connect to bus: Host is down