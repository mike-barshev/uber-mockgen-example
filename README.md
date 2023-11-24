# uber-mockgen-example

This repository provides examples of unique `go generate` directives that can be inserted into any Go file to generate standardized mock filenames in the format `filename_mock.go`. The directive doesn't need to be changed for each file. Additionally, an example of configuring the GoLand IDE is included, demonstrating nested files, which conveniently displays the project's content, and file watchers, allowing mock generation after each file save, providing a seamless experience.

## Using `go generate` with mockgen in files

1. In the `example_var.go` file, mockgen is utilized through go generate variables (more details can be found in `go help generate`). This is a straightforward approach without using bash, but it creates files in the format `filename.go_mock.go`.

2. The `example_bash.go` file provides an example of mockgen using a command that separates the basename, allowing the creation of a file in the format `filename_mock.go`. To use this method, bash must be correctly set up with GOPATH. An example for macOS is provided:

    - `nano $HOME/.bash_profile` (if the file does not exist - create it)
    - Insert into the file:
      ```bash
      export GOPATH=$HOME/go
      export GOROOT=$HOME/sdk/go1.21.4
      export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
      ```
      *Specify your SDK; the example is for GoLand with Settings -> GOROOT -> go1.21.4 installed. You can try using the `go env` command if this does not work*  
      <br/>
    - Save the file with `Control + X` -> `Y`
    - The command `cat $HOME/.bash_profile` should display the file's content.
    - The command `echo $PATH | tr ':' '\n'` should also display changes.
    - Verify that `go version` and `mockgen -version` return the expected versions.

## Configuring file nesting

![Link to JetBrains Docks](https://www.jetbrains.com/help/webstorm/file-nesting-dialog.html)  

Example:
![File Nesting Example Image](https://resources.jetbrains.com/help/img/idea/2023.2/ws_coffee_example_default_fw_output_dark.png)

To configure a compact file display:
- Click on the three dots to the right of the project menu (next to "select opened files" and "collapse files").
- Choose "file nesting."
- Find the `.go` pattern in the left column and add to the right column, alongside existing values:
  ```
  .go_mock.go; _easyjson.go; _mock.go; _reform.go; _string.go; _test.go
  ```
- Save and observe the changes.

## Configuring file watchers
![Link to File Watchers JetBrains Docs](https://www.jetbrains.com/help/idea/using-file-watchers.html)

File watchers allow actions to be performed upon saving a file, including auto-saving. To set up mock generation when a file changes:
- (Optional step) Ensure that the file watchers plugin is enabled in GoLand -> Settings -> Plugins -> Installed. Most likely, this plugin is already enabled.
- Go to GoLand -> Settings -> Tools -> File Watchers.
- Click the plus sign, then "<custom>."
- Fill in the fields:
    - **Program**: `go`
    - **Arguments**: `generate`
- Optionally, set any additional checkboxes you need, save, and enable the file watcher.

Example how file watcher can be configured:
![File Watchers Configuration Demo Image](https://resources.jetbrains.com/help/img/idea/2023.2/ws_scss_example_file_watcher_settings_dark.png)