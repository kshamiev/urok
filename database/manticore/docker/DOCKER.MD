docker run --name test --rm -it ubuntu:22.04

apt update
apt install wget
apt install mysql-client
mkdir -p /var/run/manticore
mkdir -p /usr/local/lib/manticore

cd tmp

wget https://repo.manticoresearch.com/repository/manticoresearch_jammy/dists/manticore_5.0.2-220530-348514c86_amd64.tgz https://repo.manticoresearch.com/repository/manticoresearch_jammy/dists/jammy/main/binary-amd64/manticore-columnar-lib_1.15.4-220522-2fef34e_amd64.deb
tar -xvf manticore_5.0.2-220530-348514c86_amd64.tgz
dpkg -i manticore-*5.0.2-220530-348514c86_amd64.deb manticore-columnar-lib_1.15.4-220522-2fef34e_amd64.deb

wget https://repo.manticoresearch.com/repository/morphology/en.pak.tgz
tar -xvf en.pak.tgz
mv en.pak /usr/share/manticore

wget https://repo.manticoresearch.com/repository/morphology/ru.pak.tgz
tar -xvf ru.pak.tgz
mv ru.pak /usr/share/manticore

wget https://go.dev/dl/go1.17.13.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.13.linux-amd64.tar.gz
// ////

192.168.0.103

apt install sudo
apt install mc
apt install systemctl
apt install systemd
