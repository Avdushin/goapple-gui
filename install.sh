echo "INSTALLING COMPONENTS..."
sudo pacman -S st --noconfirm
clear
sudo dnf install st -y
clear
sudo eopkg it st -y
clear
sudo apt install st -y
clear
chmod +x *.desktop
sudo mkdir /etc/goapple-gui/
sudo cp -rf language.txt /etc/goapple-gui
echo "DONE..."