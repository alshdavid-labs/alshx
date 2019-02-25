ls $2

if [ $? -eq 1 ]; then
    echo Making directory $2
    mkdir $2
fi

sudo mount -t ntfs -o rw,auto,nobrowse $1 $2
