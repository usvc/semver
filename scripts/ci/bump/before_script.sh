#!/bin/sh
set -x;
if [ "${BASE64_DEPLOY_KEY}" = "" ]; then
  printf -- 'the $BASE64_DEPLOY_KEY variable was not defined. exiting with status 1...\n';
  exit 1;
elif [ "${REPO_HOSTNAME}" = "" ]; then
  printf -- 'the $REPO_HOSTNAME variable was not defined. exiting with status 2...\n';
  exit 2;
else
  mkdir -p ~/.ssh;
  printf -- "${BASE64_DEPLOY_KEY}" | base64 -d > ~/.ssh/id_rsa;
  chmod 600 -R ~/.ssh/id_rsa;
  ssh-keyscan -t rsa ${REPO_HOSTNAME} >> ~/.ssh/known_hosts;
  exit 0;
fi;
