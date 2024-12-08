#!/bin/bash

set -e

# Define the CLI name and installation directory
CLI_NAME="dev_cli"
INSTALL_DIR="/usr/local/bin"

# Define base URL for downloads
BASE_URL="https://github.com/amar-jay/dev_cli/releases/download"

# Fetch the latest version if VERSION is not explicitly set
VERSION=$(curl -s https://api.github.com/repos/amar-jay/dev_cli/releases/latest | jq -r '.tag_name')
if [[ -z "$VERSION" ]]; then
    echo "Unable to fetch the latest version. Aborting."
    exit 1
fi

# Detect the platform and architecture
OS=$(uname -s)
ARCH=$(uname -m)

case "$OS" in
    Linux*)
        PLATFORM_OS="linux"
        ;;
    Darwin*)
        PLATFORM_OS="darwin"
        ;;
    *)
        echo "Unsupported operating system: $OS"
        exit 1
        ;;
esac

case "$ARCH" in
    x86_64)
        PLATFORM_ARCH="386"
        ;;
    arm64|aarch64)
        PLATFORM_ARCH="arm64"
        ;;
    armv7l)
        PLATFORM_ARCH="armv7"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# Construct the download URL
TAR_NAME="${CLI_NAME}_${VERSION}_${PLATFORM_OS}_${PLATFORM_ARCH}.tar.gz"
DOWNLOAD_URL="${BASE_URL}/${VERSION}/${TAR_NAME}"

# Function to clean up previous installations
cleanup_previous_installation() {
    echo "Removing previous versions of ${CLI_NAME}..."
    sudo rm -f "${INSTALL_DIR}/${CLI_NAME}" || true
}

# Function to download and install the CLI
install_new_version() {
    echo "Downloading ${CLI_NAME} version ${VERSION} from ${DOWNLOAD_URL}..."
    curl -L "${DOWNLOAD_URL}" -o "/tmp/${TAR_NAME}"

    echo "Extracting the binary..."
    tar -xzf "/tmp/${TAR_NAME}" -C "/tmp"

    echo "Installing ${CLI_NAME} to ${INSTALL_DIR}..."
    sudo mv "/tmp/${CLI_NAME}" "${INSTALL_DIR}/"
    sudo chmod +x "${INSTALL_DIR}/${CLI_NAME}"

    echo "${CLI_NAME} version ${VERSION} installed successfully!"
}

# Execute the functions
cleanup_previous_installation
install_new_version