#!/bin/bash

# Function to display error messages and exit
function error_exit() {
    echo "$1" 1>&2
    exit 1
}

# Check if neovim is already installed
if ! command -v nvim &> /dev/null; then
    echo "Neovim is not installed. Installing..."
    # Check the package manager and install neovim accordingly
    if command -v apt-get &> /dev/null; then
        sudo apt-get install neovim -y || error_exit "Failed to install neovim"
    elif command -v pacman &> /dev/null; then
        sudo pacman -S neovim --noconfirm || error_exit "Failed to install neovim"
    elif command -v brew &> /dev/null; then
        brew install neovim || error_exit "Failed to install neovim"
    else
        error_exit "Unable to determine package manager. Please install Neovim manually."
    fi
else
    echo "Neovim is already installed."
fi

# Clone the nvim-config repository
echo "Cloning nvim-config repository..."
mkdir -p ~/.config
git clone http://github.com/amarjay/nvim-config ~/.config/nvim || error_exit "Failed to clone nvim-config repository"

# Install dependencies
echo "Installing dependencies..."
# Example: Install Python dependencies using pip
pip install pynvim || error_exit "Failed to install Python dependencies"

# Beautiful description of the nvim-config
echo ""
echo "=================================================================="
echo "          Welcome to amarjay's Neovim Configuration!"
echo "=================================================================="
echo ""
echo "This is amarjay's nvim-config, built with Neovim, LazyVim, Mason,"
echo "and other awesome plugins, meticulously crafted for an efficient and"
echo "enjoyable editing experience."
echo ""
echo "Neovim is a powerful text editor that provides an intuitive and extensible"
echo "editing environment. With LazyVim, we bring laziness to the next level,"
echo "automating repetitive tasks and reducing cognitive load. Mason enhances"
echo "Neovim's capabilities, allowing seamless integration of various tools"
echo "and workflows."
echo ""
echo "With this configuration, you'll unleash your productivity, effortlessly"
echo "navigating through code, writing prose, or tinkering with configurations."
echo "Say goodbye to mundane tasks and hello to a world of efficient editing!"
echo ""
echo "Happy editing!"
echo ""
echo "=================================================================="
echo ""

# Additional dependencies may go here, depending on the configuration

echo "Setup completed successfully."

nvim --version #+PlugInstall +qall

