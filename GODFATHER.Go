#!/bin/bash

_get_zlink () {
    local regex
    regex='(https?)://github.com/.+/.+'
    if [[ $UPSTREAM_REPO == "godfather" ]]
    then
        echo "aHR0cHM6Ly9naXRodWIuY29tL2dvZGZhdGhlcmFra2lpL0dPREZBVEhFUlVTRVJCT1QvYXJjaGl2ZS9tYXN0ZXIuemlw" | base64 -d
    elif [[ $UPSTREAM_REPO == "powerfull" ]]
    then
        echo "aHR0cHM6Ly9naXRodWIuY29tL2dvZGZhdGhlcmFra2lpL0dPREZBVEhFUlVTRVJCT1QvYXJjaGl2ZS9tYXN0ZXIuemlw" | base64 -d
    elif [[ $UPSTREAM_REPO =~ $regex ]]
    then
        if [[ $UPSTREAM_REPO_BRANCH ]]
        then
            echo "${UPSTREAM_REPO}/archive/${UPSTREAM_REPO_BRANCH}.zip"
        else
            echo "${UPSTREAM_REPO}/archive/master.zip"
        fi
    else
        echo "aHR0cHM6Ly9naXRodWIuY29tL2dvZGZhdGhlcmFra2lpL0dPREZBVEhFUlVTRVJCT1QvYXJjaGl2ZS9tYXN0ZXIuemlw" | base64 -d
    fi
}

_get_repolink () {
    local regex
    local rlink
    regex='(https?)://github.com/.+/.+'
    if [[ $UPSTREAM_REPO == "godfahter" ]]
    then
        rlink=`echo "aHR0cHM6Ly9naXRodWIuY29tL2dvZGZhdGhlcmFra2lpL0dPREZBVEhFUlVTRVJCT1Q=" | base64 -d`
    elif [[ $UPSTREAM_REPO == "powerfull" ]]
    then
        rlink=`echo "aHR0cHM6Ly9naXRodWIuY29tL2dvZGZhdGhlcmFra2lpL0dPREZBVEhFUlVTRVJCT1Q=" | base64 -d`
    elif [[ $UPSTREAM_REPO =~ $regex ]]
    then
        rlink=`echo "${UPSTREAM_REPO}"`
    else
        rlink=`echo "aHR0cHM6Ly9naXRodWIuY29tL2dvZGZhdGhlcmFra2lpL0dPREZBVEhFUlVTRVJCT1Q=" | base64 -d`
    fi
    echo "$rlink"
}


_run_python_code() {
    python3${pVer%.*} -c "$1"
}

_run_godfatherpackgit() {
    $(_run_python_code 'from git import Repo
import sys
OFFICIAL_UPSTREAM_REPO = "https://github.com/godfahterakkii/GODFAHTERBOT"
ACTIVE_BRANCH_NAME = "master"
repo = Repo.init()
origin = repo.create_remote("temponame", OFFICIAL_UPSTREAM_REPO)
origin.fetch()
repo.create_head(ACTIVE_BRANCH_NAME, origin.refs[ACTIVE_BRANCH_NAME])
repo.heads[ACTIVE_BRANCH_NAME].checkout(True) ')
}

_run_godfathergit() {
    local repolink=$(_get_repolink)
    $(_run_python_code 'from git import Repo
import sys
OFFICIAL_UPSTREAM_REPO="'$repolink'"
ACTIVE_BRANCH_NAME = "'$UPSTREAM_REPO_BRANCH'" or "master"
repo = Repo.init()
origin = repo.create_remote("temponame", OFFICIAL_UPSTREAM_REPO)
origin.fetch()
repo.create_head(ACTIVE_BRANCH_NAME, origin.refs[ACTIVE_BRANCH_NAME])
repo.heads[ACTIVE_BRANCH_NAME].checkout(True) ')
}

_start_bot () {
    local zippath
    zippath="GODFATHERBOT.zip"
    echo "  Downloading source code ..."
    wget -q $(_get_zlink) -O "$zippath"
    echo "  Unpacking Data ..."
    GODFATHERPATH=$(zipinfo -1 "$zippath" | grep -v "/.");
    unzip -qq "$zippath"
    echo "Done"
    echo "  Cleaning ..."
    rm -rf "$zippath"
    _run_godfatherpackgit
    cd $GODFATHERPATH
    _run_godfahtergit
    python3 ../setup/updater.py ../requirements.txt requirements.txt
    chmod -R 755 bin
    echo "    Starting GodFatherUserBot    "
    python3 -m GodFather
}

_start_bot
