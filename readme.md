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
| `--debug`           | Prints a list of the directories and/or files expected to be ignored. |

## Todo
- [ ] Add colors based on extension.
- [ ] Replace the output extension names with their filetype names.
- [ ] Make the formatted output less garbage.
- [ ] Allow the usage of an optional ignore file separate from `.gitignore`
- [ ] GitHub repo support.
- [ ] Output to markdown.
- [x] Sort the final output.


## Installation
### Windows
- Clone and build the repo or download the executable for your system.
- Move the executable somewhere, or leave it there.
- Add a the path to the executable to your `PATH` environment variable.
  - If you want to be the only user with access to the executable, add it to the user variable `PATH`.
  - If you want to have goloc available system-wide, add it to the system variable `PATH`.
- Restart your terminal.

### Linux
- Clone and build the repo or download the binary for your system.
- Copy or move the binary to `/usr/bin`.
- Update your paths accordingly.

### MacOS
- Same as Linux, but on MacOS.
