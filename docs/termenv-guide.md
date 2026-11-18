## `termenv`

The `termenv` package provides terminal helpers for projects using
`github.com/nsymx/scoopbucket`.

Import the package in your Go code:

```go
import "github.com/nsymx/scoopbucket/termenv"
```

### OS detection and validation
Since we aim to make this project work with linux and nothing else (for now, at least), we call the following function that will make sure the requirements are met. If not, the program will exit with an error message.

```go
func main() {
	termenv.DetectOS()

	// Continue with platform-specific application setup.
}
```
### Get the user's home directory

`GetUserHomePath` returns the current user's home directory after checking
that the path exists and is a directory. Handle errors before using the path.

```go
homePath, err := termenv.GetUserHomePath()
if err != nil {
	log.Fatalf("could not find a valid home directory: %v", err)
}

fmt.Println("Home directory:", homePath)
```

### Clearscreen
A very simple function to clear the terminal.

```go
termenv.ClearScreen()
fmt.Println("Screen cleared")
```

### Simple yes or no prompt.

`Prompt` displays a prompt and reads an answer from the `io.Reader` you pass
in. It returns `true` for `y` and `false` for `n` or a blank answer.
Input is trimmed and case-insensitive; other answers cause the prompt to be
shown again.

For interactive terminal input, pass `os.Stdin`:

```go
confirmed := termenv.Prompt(os.Stdin, "Continue? [y/N] ")
if !confirmed {
	fmt.Println("Cancelled")
	return
}

fmt.Println("Continuing")
```

You can also pass a reader in tests or other non-interactive code:

```go
confirmed := termenv.Prompt(strings.NewReader("y\n"), "Continue? [y/N] ")
```

### Use the predefined colors

The package exports four Lip Gloss colors:

| Variable | Default | Intended use |
| --- | --- | --- |
| `ColorText` | `#FFFFFF` | General text |
| `ColorAlpha` | `#1016CF` | Accent text |
| `ColorSuccess` | `#1BE160` | Success messages |
| `ColorError` | `#E20F0F` | Error messages |

For example, build a style with the success color:

```go
import "charm.land/lipgloss/v2"

successStyle := lipgloss.NewStyle().
	Bold(true).
	Foreground(termenv.ColorSuccess)

fmt.Println(successStyle.Render("Operation complete"))
```

The predefined style variables are created using the color values available
when the package initializes. If you change a color variable later, an
already-created style is not automatically rebuilt.

### Render text with predefined styles using [lipgloss](charm.land/lipgloss)

- `BasicTextStyle`: bold text using `ColorText`.
- `AlphaTextStyle`: bold text using `ColorAlpha`.
- `GreeterStyle`: bold text with the text color, hidden border settings,
  and top and left margins.
- `EnumeratorStyle`: bold text using `ColorAlpha`.
- `ErrorStyle`: bold text using `ColorError`.

Use `Render` to style a string:

```go
fmt.Println(termenv.BasicTextStyle.Render("Hello"))
fmt.Println(termenv.AlphaTextStyle.Render("Section 1"))
fmt.Println(termenv.GreeterStyle.Render("Welcome"))
fmt.Println(termenv.EnumeratorStyle.Render("1. First item"))
fmt.Println(termenv.ErrorStyle.Render("Something went wrong"))
```

The style variables are exported, so applications can derive their own
styles or assign new ones. For example:

```go
customErrorStyle := termenv.ErrorStyle.Italic(true)
fmt.Println(customErrorStyle.Render("Please try again"))
```

### Logger

`StandardLogger` and `CompactLogger` are package-level Charmbracelet loggers.
The standard logger includes timestamps; the compact logger omits them.
Both write to standard error by default.

```go
termenv.StandardLogger.Info("application started")
termenv.StandardLogger.Warn("configuration file not found")
termenv.StandardLogger.Error("could not load configuration")
```

Use `CompactLogger` when a timestamp-free message is useful:

```go
termenv.CompactLogger.Print("Working...")
```
