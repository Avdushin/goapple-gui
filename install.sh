echo "INSTALLING COMPONENTS..."
sudo pacman -S st --no-confirm
sudo dnf install st -y
sudo eopkg it st -y
sudo apt install st -y
chmod +x *.desktop
echo "DONE..."