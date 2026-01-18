
#!/bin/bash
#
# Freight Installer
# Usage: curl -fsSL https://raw.githubusercontent.com/devbytes-cloud/freight/main/curl.sh | bash
#

set -e

# Configuration
REPO="devbytes-cloud/freight"
BINARY_NAME="freight"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

info() {
    printf "${GREEN}[INFO]${NC} %s\n" "$1"
}

warn() {
    printf "${YELLOW}[WARN]${NC} %s\n" "$1"
}

error() {
    printf "${RED}[ERROR]${NC} %s\n" "$1"
    exit 1
}

# Detect OS
detect_os() {
    case "$(uname -s)" in
        Linux*)  OS="Linux";;
        Darwin*) OS="Darwin";;
        MINGW*|MSYS*|CYGWIN*) OS="Windows";;
        *)       error "Unsupported operating system: $(uname -s)";;
    esac
}

# Detect architecture
detect_arch() {
    case "$(uname -m)" in
        x86_64|amd64)  ARCH="x86_64";;
        arm64|aarch64) ARCH="arm64";;
        armv6l|armv7l) ARCH="armv6";;
        *)             error "Unsupported architecture: $(uname -m)";;
    esac
}

# Get latest version from GitHub releases
get_latest_version() {
    VERSION=$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    if [ -z "$VERSION" ]; then
        error "Failed to fetch latest version. Check if releases exist at https://github.com/${REPO}/releases"
    fi
}

# Download and install
install() {
    detect_os
    detect_arch
    get_latest_version

    info "Detected: OS=$OS, ARCH=$ARCH"
    info "Installing ${BINARY_NAME} ${VERSION}..."

    # Build download URL
    ARCHIVE_NAME="freight_${OS}_${ARCH}.tar.gz"
    DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/${ARCHIVE_NAME}"

    # Create temp directory
    TMP_DIR=$(mktemp -d)
    trap "rm -rf ${TMP_DIR}" EXIT

    info "Downloading from ${DOWNLOAD_URL}..."
    curl -fsSL "${DOWNLOAD_URL}" -o "${TMP_DIR}/${ARCHIVE_NAME}" || error "Download failed. Archive may not exist for your platform."

    # Extract archive
    info "Extracting..."
    tar -xzf "${TMP_DIR}/${ARCHIVE_NAME}" -C "${TMP_DIR}"

    # Set binary extension for Windows
    EXT=""
    if [ "$OS" = "Windows" ]; then
        EXT=".exe"
    fi

    # Install binary
    BINARY_PATH="${TMP_DIR}/${BINARY_NAME}${EXT}"
    if [ ! -f "$BINARY_PATH" ]; then
        error "Binary not found in archive"
    fi

    info "Installing to ${INSTALL_DIR}/${BINARY_NAME}${EXT}..."

    # Check if we need sudo
    if [ -w "$INSTALL_DIR" ]; then
        mv "$BINARY_PATH" "${INSTALL_DIR}/${BINARY_NAME}${EXT}"
        chmod +x "${INSTALL_DIR}/${BINARY_NAME}${EXT}"
    else
        warn "Elevated permissions required for ${INSTALL_DIR}"
        sudo mv "$BINARY_PATH" "${INSTALL_DIR}/${BINARY_NAME}${EXT}"
        sudo chmod +x "${INSTALL_DIR}/${BINARY_NAME}${EXT}"
    fi

    echo ""
    info "✔ Successfully installed ${BINARY_NAME} ${VERSION}"
    info "Run '${BINARY_NAME} --help' to get started"
}

# Run installation
install