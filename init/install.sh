#!/bin/sh
set -e;

printf -- '\nThis script will download the latest `semver` binary from https://github.com/usvc/semver/releases.\n';
printf -- '  Enter any key followed by an enter to continue.\n';
printf -- '  Or hit ctrl+c to exit.\n';

read x;

# check for write permissions
touch ./__do_we_have_write_permissions;
rm -rf ./__do_we_have_write_permissions;

# retrieve latest version
curl https://api.github.com/repos/usvc/semver/releases > ./.version_info;
cat ./.version_info | jq '.[].tag_name' -r | egrep "^v[0-9]+\.[0-9]+\.[0-9]+$" | sed -e 's/v//g' | sort -V | tail -n 1 > ./.version_latest;

# determine system environment
export SYSTEM="$(uname -s)";
export ARCHITECTURE="$(uname -p)";
case $SYSTEM in
  *inux)  export OS=linux; ;;
  *arwin) export OS=darwin; ;;
  *)      export OS=windows;
          export BIN_EXT=.exe; ;;
esac
case $ARCHITECTURE in
  x86_64) export ARCH=amd64; ;;
  *86)    export ARCH=386; ;;
  arm*)   export ARCH=arm; ;;
esac

# download it
curl -Lo "./semver${BIN_EXT}" "https://github.com/usvc/semver/releases/download/v$(cat ./.version_latest)/semver-${OS}-${ARCH}${BIN_EXT}";
chmod +x "./semver${BIN_EXT}";

# clean up
rm -rf ./.verison_info;
rm -rf ./.version_latest;

# done
printf -- "run ./semver --help for more information\n";
printf -- "to include it in your path, shift the binary to /bin\n\n";