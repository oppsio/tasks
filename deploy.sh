#! /bin/bash

USER=ricardo-rossi
cd $HOME
git clone git@github.com:oppsio/releases.git
mkdir ~/releases/$CIRCLE_BRANCH/
cp ~/$CIRCLE_PROJECT_REPONAME/$CIRCLE_PROJECT_REPONAME ~/releases/$CIRCLE_BRANCH/
cd ~/releases/
git add .
git status
COMMITMSG="Updating /$CIRCLE_BRANCH/$CIRCLE_PROJECT_REPONAME"
echo $COMMITMSG
git commit -am $COMMITMSG
git push
