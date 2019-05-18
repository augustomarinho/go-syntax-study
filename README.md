#GO-SYNTAX-STUDY

# Setting GOPATH variable dynamically

### Adding this function in your ~/.bashrc (Good suggestion found on http://hgfischer.org/article/managing-multiple-gopaths/)
```cd () {
    builtin cd "$@"
    cdir=$PWD
    while [ "$cdir" != "/" ]; do
        if [ -e "$cdir/.gopath" ]; then
            export GOPATH=$cdir
            break
        fi
        cdir=$(dirname "$cdir")
    done
}
```

### Creating an empty file .gopath in your root workspace path
``` touch .gotpath```