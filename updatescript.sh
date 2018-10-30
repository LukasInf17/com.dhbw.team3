#!/bin/bash

cd /root/go/src/github.com/invitation/
source /etc/profile

git fetch origin
reslog=$(git log HEAD..origin/master --oneline)
if [[ "${reslog}" != "" ]] ; then
	git merge origin/master
	buffalo build
	cd poster
	pdflatex conference_poster_4.tex
	cp conference_poster_4.pdf /var/www/html/
fi
