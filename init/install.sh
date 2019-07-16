#!/bin/sh
set -e;
printf -- '\n';

# get user confirmation
printf -- 'this script downloads the latest `semver` binary from https://github.com/usvc/semver/releases.\n';
printf -- "  > hit ctrl+c to exit within 5 seconds if that's not what you want...\n";
N=5; while [ $N -gt 0 ]; do
  sleep 1; printf -- ".";
  N=$(($N-1))
done;
printf -- '\n\n';

# check for write permissions
printf -- "checking for write permissions at $(pwd)...\n"
touch ./__do_we_have_write_permissions;
rm -rf ./__do_we_have_write_permissions;
printf -- 'we have permission\n\n';

# determine system environment
printf -- 'getting system information...\n';
SYSTEM="$(uname -s)";
ARCHITECTURE="$(uname -p)";
case $SYSTEM in
  *inux)  OS=linux; ;;
  *arwin) OS=darwin; ;;
  *)      OS=windows;
          BIN_EXT=.exe; ;;
esac
case $ARCHITECTURE in
  x86_64) ARCH=amd64; ;;
  *86)    ARCH=386; ;;
  arm*)   ARCH=arm; ;;
esac
printf -- "os: ${OS}, architecture: ${ARCH}\n\n";

# retrieve latest version
printf -- 'retrieving latest version info...\n'
curl -s https://api.github.com/repos/usvc/semver/releases > ./.version_info;
cat ./.version_info | jq '.[].tag_name' -r | egrep "^v[0-9]+\.[0-9]+\.[0-9]+$" | sed -e 's/v//g' | sort -V | tail -n 1 > ./.version_latest;
BIN_URL="https://github.com/usvc/semver/releases/download/v$(cat ./.version_latest)/semver-${OS}-${ARCH}${BIN_EXT}";
BIN_PATH="./semver-v$(cat ./.version_latest)${BIN_EXT}";
printf -- "latest version: $(cat ./.version_latest)\n\n";
rm -rf ./.verison_info;
rm -rf ./.version_latest;

# download it
printf -- "retrieving binary from ${BIN_URL}\n"
printf -- "  > downloading to ${BIN_PATH}...\n"
curl -Lo "${BIN_PATH}" "${BIN_URL}";
chmod +x "${BIN_PATH}";
printf -- "done.\n\n";

# done
printf -- "run '${BIN_PATH} --help' for more information\n";
printf -- '- to include it in your path:\n';
printf -- "  1. sudo mv '${BIN_PATH}' /bin/semver\n";
printf -- '  2. sudo ln -s $(pwd)';
printf --                       "/${BIN_PATH} /bin/semver\n";

printf -- '\nthanks for using usvc/semver (:\n';
