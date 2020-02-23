
# CLI - Random numbers

The program prints numbers from 1 to 10 in random order. It makes sure that consecutive numbers are not printed

# Programming Language
Golang is used as programming language

# Installing Go
## Installing on Linux

### Download
Download the archive from https://golang.org/dl/  the archive file appropriate for the architecture. For instance, if the system architecture is 64-bit x86 on Linux, the archive to be downloaded will be  `go1.13.8.linux-amd64.tar.gz`.

### Extract
Once downloaded, extract it into  `${HOME}/opt`, creating a Go tree in `${HOME}/opt/go`. For example:

create the directory

    mkdir ${HOME}/opt
    
extract the downloaded archive (replacing the archive name with the downloaded one)

    tar -C ${HOME}/opt -xzf go1.13.8.linux-amd64.tar.gz

### Path
Add  `${HOME}/opt/go/bin`  to the  `PATH`  environment variable.  The following command can be executed within the shell as it is:

    export PATH=$PATH:${HOME}/opt/go/bin

### Setting GOPATH
`GOPATH`  can be any directory on your system **other than** the Go installation path `${HOME}/opt/go`. Default will be set to  `$HOME/go`.  You may also use home directory  `GOPATH=$HOME` or the directory where the code is downloaded.

    mkdir ${HOME}/go && export GOPATH=$HOME/go

## Installing on macOS (package installer)

### Download
Download the archive from https://golang.org/dl/  appropriate for the architecture. For instance, if the system architecture is 64-bit, the archive to be downloaded will be  `go1.13.8.darwin-amd64.pkg`.

### Install
Open the file and follow the prompts to install the Go tools. The package installs the Go distribution to /usr/local/go

The package should put the  `/usr/local/go/bin`  directory in your  `PATH`  environment variable. You may need to restart any open Terminal sessions for the change to take effect.

# Running the program
Following instructions are for both Linux and macOS systems.

Download the program from github by pointing the browser at [https://github.com/VanagaS/Adjust](https://github.com/VanagaS/Adjust) and clicking on the green button `Clone or download` and selecting `Download ZIP` option from the dropdown. Alternatively it can be downloaded directly using the URL

    https://github.com/VanagaS/Adjust/archive/master.zip

Unzip the downloaded zip file. Open a terminal and change to the location where the downloaded zip file is extracted. 

Change to the source directory, assuming, we are currently in the extracted folder, 

    cd Adjust-master/go-cli/

## Run the program

Execute the program using the command

    go run go-cli.go


## Uninstalling Go
Remove the Go installation that we did in the above steps.

### Linux 
Delete the Go installed directory 

    rm -fR ${HOME}/opt/go
    
The environment variables created in the installation step above `(PATH & GOPATH)` persist only till the terminal session exits. 

### macOS
Delete the Go installed directory from `/usr/local/go`. You'll need sudo permissions

    sudo rm -fR /usr/local/go
    
The Go bin directory from PATH environment must also be removed. The /etc/paths.d/go file must be removed.

    sudo rm -f /etc/paths.d/go
    
