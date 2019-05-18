# GO-SYNTAX-STUDY

# Setting GOPATH variable dynamically

1. Adding this function in your ~/.bashrc - 

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
 2. Creating an empty file .gopath in your root workspace path
``` touch .gotpath```

 Good suggestion found [here](http://hgfischer.org/article/managing-multiple-gopaths/)

# Good Practies in GoLang
[Good Practies](https://github.com/golang-standards/project-layout)