WANTS_MESSAGE=0
MESSAGE="Update"

WANTS_BRANCH=0
BRANCH_NAME=""

WANTS_AMEND=0

WANTS_PUSH=1
for var in "$@"
do
    if [ $WANTS_MESSAGE = 1 ]; then
        MESSAGE=$var
        continue
    fi
    if [ $WANTS_BRANCH = 1 ]; then
        BRANCH_NAME=$var
        continue
    fi
    if [ "$var" = "-b" ]; then
        WANTS_BRANCH=1
        continue
    fi
    if [ "$var" = "-np" ]; then
        WANTS_PUSH=0
        continue
    fi
    if [ "$var" = "-a" ]; then
        WANTS_AMEND=1
        continue
    fi
    if [ "$var" = "-ne" ]; then
        WANTS_MESSAGE=1
        continue
    fi
    if [ "$var" = "-m" ]; then
        WANTS_MESSAGE=1
        continue
    fi
done

if [ $WANTS_BRANCH = 1 ]; then
    git checkout -b $BRANCH_NAME
fi

git add .

if [ $WANTS_AMEND = 1 ]; then
    if [ $WANTS_MESSAGE = 1 ]; then
        git commit --amend -m $MESSAGE
    else
        git commit --amend --no-edit
    fi
else
    git commit -m $MESSAGE
fi

if [ $? != 0 ]; then
    return
fi

if [ $WANTS_PUSH = 1 ]; then
    if [ $WANTS_AMEND = 1 ]; then
        git push -u origin HEAD -f
    else
        git push -u origin HEAD
    fi
fi
