#! /bin/bash

git config --global user.email "ricardo@endata.com"
git config --global user.name "Ricardo Rossi"
cd $HOME
git clone git@github.com:oppsio/releases.git
mkdir -p ~/releases/$CIRCLE_BRANCH
cp ~/$CIRCLE_PROJECT_REPONAME/$CIRCLE_PROJECT_REPONAME ~/releases/$CIRCLE_BRANCH
cd ~/releases
git add .
git status
COMMITMSG="Updating /$CIRCLE_BRANCH/$CIRCLE_PROJECT_REPONAME"
echo $COMMITMSG
git commit -v -m "$COMMITMSG"
git push -v --progress
