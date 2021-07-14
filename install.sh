#!/usr/bin/env sh
# Shamelessly copied from https://github.com/technosophos/helm-template
# https://github.com/c-bata/kube-prompt
PROJECT_NAME="kube-prompt"
PROJECT_GH="c-bata/kube-prompt"
export GREP_COLOR="never"
ZIP_TYPE='.zip'
SPLIT_CHAT='_'
# SHELL_FOLDER=$(cd "$(dirname "$0")";pwd)
PLUGIN_TMP_FILE="/tmp/${PROJECT_NAME}${ZIP_TYPE}"
: ${VERSION="$1"}

# initArch discovers the architecture for this system.
initArch() {
  ARCH=$(uname -m)
  case $ARCH in
  armv5*) ARCH="armv5" ;;
  armv6*) ARCH="armv6" ;;
  armv7*) ARCH="armv7" ;;
  aarch64) ARCH="arm64" ;;
  x86) ARCH="386" ;;
  x86_64) ARCH="amd64" ;;
  i686) ARCH="386" ;;
  i386) ARCH="386" ;;
  esac
}

# initOS discovers the operating system for this system.
initOS() {
  OS=$(uname | tr '[:upper:]' '[:lower:]')

  case "$OS" in
  # Msys support
  msys*) OS='windows' ;;
  # Minimalist GNU for Windows
  mingw*) OS='windows' ;;
  darwin) OS='darwin' ;;
  esac
}

# verifySupported checks that the os/arch combination is supported for
# binary builds.
verifySupported() {
  supported="linux-amd64\nlinux-arm64\nlinux-386\ndarwin-amd64"
  if ! echo "${supported}" | grep -q "${OS}-${ARCH}"; then
    echo "No prebuild binary for ${OS}${SPLIT_CHAT}${ARCH}."
    exit 1
  fi

  if ! type "curl" >/dev/null && ! type "wget" >/dev/null; then
    echo "Either curl or wget is required"
    exit 1
  fi
}

# getDownloadURL checks the latest available version.
getDownloadURL() {
  version=$VERSION
#   https://github.com/c-bata/kube-prompt/releases/download/v0.9.38/kube-prompt-v0.9.38-darwin-amd64.zip
  if [ -n "$version" ]; then
    DOWNLOAD_URL="https://github.com/$PROJECT_GH/releases/download/$version/kube-prompt-$version-${OS}${SPLIT_CHAR}${ARCH}${ZIP_TYPE}"
  else
    # Use the GitHub API to find the download url for this project.
    #    https://api.github.com/repos/c-bata/kube-prompt/releases/latest
    #    https://api.github.com/repos/c-bata/kube-prompt/releases
    url="https://api.github.com/repos/$PROJECT_GH/releases"
    if type "curl" >/dev/null; then
      # reponse=$(curl -s $url)
      echo "DOWNLOAD_URLs=\$(curl -s $url |grep ${OS}${SPLIT_CHAT}${ARCH} |grep ${ZIP_TYPE}| awk '/\"browser_download_url\":/{gsub( /[,\"]/,\"\", \$2); print \$2}')"
      echo "DOWNLOAD_URL=\`echo \$DOWNLOAD_URLs|head -1|awk '{print $1}'\`"
      DOWNLOAD_URLs=$(curl -s $url |grep ${OS}${SPLIT_CHAT}${ARCH} |grep ${ZIP_TYPE}| awk '/"browser_download_url":/{gsub( /[,"]/,"", $2); print $2}')
      DOWNLOAD_URL=`echo $DOWNLOAD_URLs|head -1|awk '{print $1}'`
      if [ -z "$DOWNLOAD_URL" ]; then
        echo "no url get, try manually"
        exit 1
      fi
    elif type "wget" >/dev/null; then
      DOWNLOAD_URLs=$(wget -q -O - $url  |grep ${OS}${SPLIT_CHAT}${ARCH} |grep ${ZIP_TYPE}| awk '/"browser_download_url":/{gsub( /[,"]/,"", $2); print $2}')
      DOWNLOAD_URL=`echo $DOWNLOAD_URLs|head -1|awk '{print $1}'`
    fi
  fi
}

# downloadFile downloads the latest binary package and also the checksum
# for that binary.
downloadFile() {
  echo "Downloading $DOWNLOAD_URL"
  if type "curl" >/dev/null; then
    curl -L "$DOWNLOAD_URL" -o "$PLUGIN_TMP_FILE"
  elif type "wget" >/dev/null; then
    wget -q -O "$PLUGIN_TMP_FILE" "$DOWNLOAD_URL"
  fi
}

# installFile verifies the SHA256 for the file, then unpacks and
# installs it.
installFile() {
  FILE_TMP="/tmp/$PROJECT_NAME"
  mkdir -p "$FILE_TMP"
  if [ "${ZIP_TYPE}" = '.zip' ]; then
    if ! command -v unzip; then
        echo 'please install unzip'
        command -v apt && sudo -S apt install -y unzip
    fi
    if command -v unzip; then
        echo "Preparing to install into ~"
        unzip -o "$PLUGIN_TMP_FILE" -d "$FILE_TMP"
        chmod +x $FILE_TMP/kube-prompt
        sudo -S mv $FILE_TMP/kube-prompt /usr/local/bin/kube-prompt
    fi
#   else
#     tar xf "$PLUGIN_TMP_FILE" -C "$FILE_TMP"
#     FILE_TMP_BIN="$FILE_TMP/$PROJECT_NAME"
#     echo "Preparing to install into ~"
#     # mkdir -p "$SHELL_FOLDER/cqhttp"
#     cp "$FILE_TMP_BIN" ~  
  fi
}

# fail_trap is executed if an error occurs.
fail_trap() {
  result=$?
  if [ "$result" != "0" ]; then
    echo "Failed to install $PROJECT_NAME"
    printf '\tFor support, go to https://github.com/c-bata/kube-prompt or 17682318150 \n'
  fi
  exit $result
}

# Execution

#Stop execution on any error
trap "fail_trap" EXIT
set -e
initArch
initOS
verifySupported
getDownloadURL
downloadFile
installFile
# testVersion