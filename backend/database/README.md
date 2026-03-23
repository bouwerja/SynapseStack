- Installing on new machine

```bash
# Update local index and install wget
sudo apt update && sudo apt install -y wget

# Download the MySQL APT config helper
wget https://dev.mysql.com/get/mysql-apt-config_0.8.32-1_all.deb

# Run the config tool (A menu will appear)
# 1. Select 'MySQL Server & Cluster'
# 2. Select 'mysql-9.0-innovation' (or latest 9.x)
# 3. Select 'OK' at the bottom
sudo dpkg -i mysql-apt-config_0.8.32-1_all.deb

# Update the package list again to include the new repository
sudo apt update

# Install the server
sudo apt install -y mysql-server

# Check version (should say 9.x.x)
mysql -V

# Secure the installation (set root password, etc.)
sudo mysql_secure_installation
```

- Upgrading existing version

```bash
#!/bin/bash

# 1. Update and install prerequisites
sudo apt update && sudo apt install -y wget gnupg2 lsB-release

# 2. Download the MySQL APT Config tool
# Note: Version 0.8.32 covers the latest 2025/2026 releases
WGET_URL="https://dev.mysql.com/get/mysql-apt-config_0.8.32-1_all.deb"
wget $WGET_URL

# 3. Pre-configure the selection to MySQL 9.0 Innovation
# This prevents the interactive pop-up menu from stopping the script
export DEBIAN_FRONTEND=noninteractive
echo "mysql-apt-config mysql-apt-config/select-server select mysql-9.0-innovation" | sudo debconf-set-selections

# 4. Install the config package
sudo dpkg -i mysql-apt-config_0.8.32-1_all.deb

# 5. Update repositories to include MySQL 9.x
sudo apt update

# 6. Install MySQL Server and Client
sudo apt install -y mysql-server

# 7. Check the version to confirm success
mysql --version
```
