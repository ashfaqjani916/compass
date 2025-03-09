
#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

# Define CLI name and installation directory
BIN_NAME="compass"
INSTALL_DIR="/usr/local/bin"

echo "🚀 Building Compass CLI in the root directory..."
go build -o "$BIN_NAME"

echo "📂 Moving $BIN_NAME to $INSTALL_DIR..."
sudo mv "$BIN_NAME" "$INSTALL_DIR/$BIN_NAME"

echo "✅ Installation complete! Run 'compass --help' to get started."
