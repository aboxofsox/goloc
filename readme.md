# GoLoc
A simple CLI tool to calculate the total LoC of a given directory.

## Usage
```bash
goloc
```

## Flags
| Flag                | Description                                                           |
|---------------------|-----------------------------------------------------------------------|
| `--use-gitignore`   | Use your `.gitignore` file as the basis of directory exclusion.         |
| `--ignore <dir>...` | Define your own list of directories to exclude.                       |
| `--no-format`       | Prints the output without the bad formatting.                         |
| `--out-file` | Save the output as a markdown file.

## Todo
- [x] Replace the output extension names with their filetype names.
- [x] Make the formatted output less garbage.
- [ ] Allow the usage of an optional ignore file separate from `.gitignore`
- [x] GitHub repo support.
- [x] Output to markdown.
- [x] Sort the final output.
- [x] Add file extenions to ignore list.


## Installation
### Recommended
If you have Go installed, simply `go install github.com/aboxofsox/goloc`

### Windows
If you want to add the executable as a global command in Windows:

- Download the executable
- Put the executable somewhere on your C: drive. For example: `C:\Program Files\goloc`.
- Add the path to wherever you put that executable to your PATH System Variable.
- Restart any terminal windows you have open.

### Linux & MacOS
- Download the correct binary for your system.
- Copy the binary to `/usr/bin`
