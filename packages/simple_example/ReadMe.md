creating module : module name : zoo

        // best practice need to be confirmed yet
        your folder name (its good if it is zoo )>
        go mod init github.com/zoo
    

folder structure

        zoo
          animal
               tiger.go
               clinic.go
          main.go ( declare package main : folder name "main" optional )
              

run

        go build  (run this command where there is main.go file)
        zoo.exe
