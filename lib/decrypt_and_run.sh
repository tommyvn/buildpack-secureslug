#!/bin/bash

KEY_PROVIDER=${KEY_PROVIDER:=dummy}

# move the encrypted tar and key to the work directory
shopt -s dotglob nullglob
ENCR_APP=$(mktemp -d)
mv ./* ${ENCR_APP}/

PASSPHRASE=$(cat ${ENCR_APP}/encr_key | ${ENCR_APP}/lib/providers/${KEY_PROVIDER} decrypting)

gpg --batch --no-tty --passphrase="${PASSPHRASE}" --no-use-agent -d ${ENCR_APP}/slug.tgz.gpg | tar -xzf - --strip 2 -C ./ ./app/
PROC=${DYNO%.*}
PROC_CMD=$(cat /app/Procfile | grep ^$PROC | cut -f 2- -d ':')
source .profile.d/*
rm -fr ${ENCR_APP}
exec $PROC_CMD
