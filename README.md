#### commands : 

        get version : go version  
        running golang file :  go run main.go  
        
        creating project
        lenovo@ubuntu:~/workspace/go/examples/string$ go mod init stringex
        go: creating new go.mod: module stringex

        building project
        lenovo@ubuntu:~/workspace/go/examples/string$ go build main.go
                
        running project
        lenovo@ubuntu:~/workspace/go/examples/string$ ./main 

        or building and running main file
        go run main.go  

### practice

        https://pkg.go.dev/builtin
        https://pkg.go.dev/std


        
#### installing
        
        https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04

        check release here : 
        curl -OL https://golang.org/dl/go1.20.5.linux-amd64.tar.gz
        sha256sum go1.20.5.linux-amd64.tar.gz
        sudo tar -C /usr/local -xvf go1.20.5.linux-amd64.tar.gz
        sudo nano ~/.profile or sudo vi ~/.profile
        source ~/.profile
        set : export PATH=$PATH:/usr/local/go/bin


#### official documentation for installing

        https://go.dev/doc/install
        
        
#### UninstallGo

        https://stackoverflow.com/questions/42186003/how-to-uninstall-go
        https://github.com/golang/go/issues/60983
        

#### creating and running project

        mkdir demo
        demo>go mod init <github.com/username/modname>
        demo>main.go create and edit
        demo>go build
        demo>./modname

#### go packages : 

        https://godoc.org/?q=quote
        https://golangbot.com/go-packages/

#### GitHub

        github.com/emirpasic/gods :data structures


#### todo : 
        
                exporting components  
                roadmap : 
                https://www.educative.io/courses/the-way-to-go
                https://www.callicoder.com/categories/golang/
                https://tutorialedge.net/golang/
                https://blog.logrocket.com/functional-programming-in-go/
                https://flaviocopes.com/golang-data-structures/
                set in golang : 
                https://stackoverflow.com/questions/34018908/golang-why-dont-we-have-a-set-datastructure
                http://www.golangbootcamp.com/book
                https://yourbasic.org/golang/
                https://golang.org/pkg/builtin/ : explore builtin funcitons
                https://golang.org/doc/effective_go.html
                https://quii.gitbook.io/learn-go-with-tests
                zetcode.com
                github.com/mercadolibre/golang-tutorial-public
                https://www.topfreebooks.org/best-go-programming-books/
                https://www.ardanlabs.com/categories/go-programing/
                https://golangbyexample.com/golang-comprehensive-tutorial/
                https://blog.logrocket.com/rest-api-golang-gin-gorm/
                https://github.com/bitfield/ftl-code
                https://bitfieldconsulting.com/golang


#### must read : 

https://go101.org/article/101.html  
            
            
#### practice : 
        
                https://github.com/quii/learn-go-with-tests
                github.com/striversity/glft
                github.com/striversity/gotr
                yourbasic.org/algorithms
                github.com/inancgumus/learngo
                golangdocs.com
                https://gosamples.dev/
                https://www.ardanlabs.com/categories/go-programing/


#### exporting function : declare function with uppercase letter

        if you want to import func from other package you have to export first 
        to export function, declare function with uppercase letter

