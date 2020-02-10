NO_MESSAGE=0
AMEND=0
PUSH=1
MESSAGE="Update"
for var in "$@"
do
    if [ "$MESSAGE" = "<obtaining>" ]; then
        MESSAGE=$var
        continue
    fi
    if [ "$var" = "-np" ]; then
        PUSH=0
        continue
    fi
    if [ "$var" = "-a" ]; then
        AMEND=1
        continue
    fi
    if [ "$var" = "-ne" ]; then
        NO_MESSAGE=1
        continue
    fi
    if [ "$var" = "-m" ]; then
        MESSAGE="<obtaining>"
        continue
    fi
done

git add .

if [ $AMEND = 1 ]; then
    if [ $NO_MESSAGE = 1 ]; then
        git commit --amend --no-edit
    else
        git commit --amend -m \"$MESSAGE\"
    fi
else
    git commit -m \"$MESSAGE\"
fi

if [ $? != 0 ]; then
    return
fi

if [ $PUSH = 1 ]; then
    if [ $AMEND = 1 ]; then
        git push -u origin HEAD -f
    else
        git push -u origin HEAD
    fi
fi
