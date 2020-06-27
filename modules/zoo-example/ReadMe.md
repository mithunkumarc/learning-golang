1. create zoo folder

2. create module : go mod init <github.com/<username>/<modulename>

        zoo > go mod init github.com/mithun/zoo

3. folder structure : package and source file structure same as java, but there must be top level folder module "zoo"

	      zoo : declared as module
		  animals : package
		      tiger.go : package animals
		      lion.go : package animals
		  birds : package
		      parrot.go : package birds
		      emu.go : package birds
	       clinic.go : package main
	
4. zoo > go build

5. zoo > zoo.exe	
