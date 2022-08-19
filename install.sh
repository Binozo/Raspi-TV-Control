#!/bin/sh

if [ "$(id -u)" -ne 0 ]; then
        echo 'This script must be run by root' >&2
        exit 1
fi

read -p "Do you want to install Raspi-TV-Control (y/n)?" choice
case "$choice" in
  y|Y ) echo "Installing...";;
  n|N ) exit;;
  * ) echo "invalid" && exit;;
esac

arch=$(dpkg --print-architecture)
realuser="${SUDO_USER:-${USER}}"
echo "Searching executable for architecture $arch..."

releasesJson=$(curl -s https://api.github.com/repos/Binozo/Raspi-TV-Control/releases/latest)
# Getting Tag name
tagName=$(echo $releasesJson | grep -o -P '(?<="tag_name": ").*(?=", "target_commitish)')
echo "Latest release is $tagName"

downloadUrl="https://github.com/Binozo/Raspi-TV-Control/releases/download/$tagName/RaspiTVControl_$arch"
path="/opt/raspi-tv-control/"
cd $path && cd ..
chown "$realuser" raspi-tv-control

echo "Downloading $downloadUrl to $path..."
filename="raspitvcontrol"
mkdir -p $path
wget "$downloadUrl" -O "$path$filename"
chmod +x "$path$filename"

echo "Installing systemd service..."
echo "
[Unit]
Description=Raspi-TV-Control
ConditionPathExists=$path$filename
After=network.target

[Service]
User=$realuser

WorkingDirectory=$path
ExecStart=$path$filename
Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
" > /etc/systemd/system/raspitvcontrol.service

systemctl daemon-reload
systemctl enable raspitvcontrol.service
systemctl start raspitvcontrol.service
systemctl status raspitvcontrol

echo "Done!"