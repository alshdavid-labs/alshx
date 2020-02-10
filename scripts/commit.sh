readonly STAGE_UNSET=0
readonly STAGE_PENDING=1
readonly STAGE_SET=2

WANTS_MESSAGE=$STAGE_SET
MESSAGE="Update"

WANTS_BRANCH=$STAGE_UNSET
BRANCH_NAME=""

WANTS_AMEND=$STAGE_UNSET

WANTS_PUSH=$STAGE_SET

for var in "$@"
do
    if [ $WANTS_MESSAGE = $STAGE_PENDING ]; then
        WANTS_MESSAGE=$STAGE_SET
        MESSAGE=$var
        continue
    fi
    if [ $WANTS_BRANCH = $STAGE_PENDING ]; then
        WANTS_BRANCH=$STAGE_SET
        BRANCH_NAME=$var
        continue
    fi
    if [ "$var" = "-ne" ]; then
        WANTS_MESSAGE=$STAGE_SET
        continue
    fi
    if [ "$var" = "-m" ]; then
        WANTS_MESSAGE=$STAGE_PENDING
        continue
    fi        
    if [ "$var" = "-b" ]; then
        WANTS_BRANCH=$STAGE_PENDING
        continue
    fi
    if [ "$var" = "-np" ]; then
        WANTS_PUSH=$STAGE_UNSET
        continue
    fi
    if [ "$var" = "-a" ]; then
        WANTS_AMEND=$STAGE_SET
        continue
    fi
done

if [ $WANTS_BRANCH = $STAGE_SET ]; then
    git checkout -b $BRANCH_NAME
fi

git add .

if [ $WANTS_AMEND = $STAGE_SET ]; then
    if [ $WANTS_MESSAGE = $STAGE_SET ]; then
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

if [ $WANTS_PUSH = $STAGE_SET ]; then
    if [ $WANTS_AMEND = $STAGE_SET ]; then
        git push -u origin HEAD -f
    else
        git push -u origin HEAD
    fi
fi
