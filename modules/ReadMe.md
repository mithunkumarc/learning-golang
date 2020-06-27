#### structure

          + module folder
             + package1
                + file1.go                     // declare : package <package1> in the first line
                + file2.go                     // declare : package <package1> in the first line     
                + package2
                  file1.go                     // declare : package <package2> in the first line      
                  file2.go                     // declare : package <package2> in the first line      
           + file with main func
           go.mod      
           // helps to "go build" and to generate modName.exe file to execute  
           
           
#### declaring package and sourcefile inside package similar to java , but to run project module is required


#### exporting function 

        to export function , declare function starting with uppercase letter

#### files exist in same package no need to use import statement to use functions

      + packer                      // module
        + numbers                   // package
          + addition.go             // source file : declare : "package numbers" in the first line
          + helperAddition.go       // declare : "package numbers" in the first line              
      + mainfunc.go                 // staring point to execute, "package main"
      
      In the above example, addition.go files can use functions from helperAddition.go
      ** make sure function declared with uppercase letter to export
         
         
#### declaring module

      declare top folder as module      
      use command : you can run this command as soon as you create top folder(to be converted to module)
        
        directory-inside-module-name> go mod init github.com/<username-or-orgname>/<module-name>
      
        example : 
        $ cd packer
        $ go mod init github.com/callicoder/packer
        
#### importing package : local

          import "github.com/<user-or-orgname>/<modulename>/<package1>"

          import (
                    "github.com/callicoder/packer/numbers"
          )
        
#### build and run : windows

        go build
        moduleName.exe
        
        
####  read text file in command prompt

      type filename.go
      
      
#### go.sum file

      sum containing the expected cryptographic checksums of the content of specific module versions
      
#### go.mod file

      required dependencies will be listed in this file
      similar to package.json file in nodejs
      
      
      
         
