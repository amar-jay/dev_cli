#!/bin/bash

set -e

# Define the CLI name and installation directory
CLI_NAME="dev_cli"
INSTALL_DIR="/usr/local/bin"

# Define base URL for downloads
BASE_URL="https://github.com/amar-jay/dev_cli/releases/download"

# Prompt the user for the version
read -p "Enter the version of ${CLI_NAME} to install (e.g., 1.0.12): " VERSION
if [[ -z "$VERSION" ]]; then
    echo "Version cannot be empty. Aborting."
    exit 1
fi

# Detect the platform
case "$(uname -s)" in
    Linux*)     PLATFORM="linux_amd64" ;;
    Darwin*)    PLATFORM="darwin_amd64" ;;
    *)          echo "Unsupported platform: $(uname -s)" && exit 1 ;;
esac

# Construct the download URL
TAR_NAME="${CLI_NAME}_${VERSION}_${PLATFORM}.tar.gz"
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

