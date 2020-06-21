#### structure

          + module folder
             + package1
                + file1
                + file2
                + package2
                  file1
                  file2
           + file with main func
           go.mod               // helps to "go build" and to generate modName.exe file to execute  

     
#### exporting function 

        to export function , declare function starting with uppercase letter

#### files exist in same package no need to use import statement to use functions

      + packer                      // module
        + numbers                   // package
          + addition.go             // source file
          + helperAddition.go
      + mainfunc.go                 // staring point to execute
      
      In the above example, addition.go files can use functions from helperAddition.go
      ** make sure function declared with uppercase letter to export
         
         
#### declaring module

      declare top folder as module      
      use command : you can run this command as soon as you create top folder(to be converted to module)
        
        directory-inside-module-name> go mod init github.com/<username-or-orgname>/<module-name>
      
        example : 
        $ cd packer
        $ go mod init github.com/callicoder/packer
        
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
      
      
      
         