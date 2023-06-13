# cleanDesktop

cleanDesktop is a command-line tool written in Go that helps you organize files in a folder based on their extensions. It creates subfolders for each unique file extension and moves the corresponding files into their respective folders.

## Installation

1. Make sure you have Go installed on your system. If not, you can download and install it from the official Go website: [https://golang.org](https://golang.org)

2. Clone this repository to your local machine using the following command:

   ```shell
   git clone https://github.com/iamcool0090/cleanDesktop.git
   ```
3. Change into the cloned directory:

    ```shell
    cd cleanDesktop
    ```
4. Build the executable using the Go compiler:

    ```shell
    go build
    ```
    
5. The file-organizer executable should now be available in the current directory. You can use it from the command line.

## Usage

   Run the executable with the desired command-line arguments:

    
	 ./cleanDesktop.exe -dir /path/to/search/directory -root organized_files
    
  Replace `/path/to/search/directory` with the path to the directory you want to organize. By default, the current directory will be used if no directory is specified.

The `-root` flag is optional and allows you to specify the root file name (name of the folder to store segregated files). By default, the root file name is set to `organized_files`.

Check the directory to see the organized files in their respective subfolders based on their extensions.
