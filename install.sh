echo "INSTALLING COMPONENTS..."
sudo pacman -S st --no-confirm
clear
sudo dnf install st -y
clear
sudo eopkg it st -y
clear
sudo apt install st -y
clear
chmod +x *.desktop
echo "DONE..."