#!/bin/bash

cd /root/go/src/github.com/invitation/
source /etc/profile

git fetch origin
reslog=$(git log HEAD..origin/master --oneline)
if [[ "${reslog}" != "" ]] ; then
	git merge origin/master
	buffalo build
fi
