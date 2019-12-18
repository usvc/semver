#!/bin/sh
set -x;
if [ "${GIT_EMAIL}" = "" ]; then
  printf -- 'the $GIT_EMAIL variable was not defined. exiting with status 1...\n';
  exit 1;
elif [ "${GIT_NAME}" = "" ]; then
  printf -- 'the $GIT_NAME variable was not defined. exiting with status 2...\n';
  exit 2;
elif [ "${REPO_URL}" = "" ]; then
  printf -- 'the $REPO_URL variable was not defined. exiting with status 3...\n';
  exit 3;
elif 
git config --global user.email "${GIT_EMAIL}";
git config --global user.name "${GIT_NAME}";
git remote set-url --add --push origin ${REPO_URL};
git checkout master;
git push -u origin master --tags --force;
