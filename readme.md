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