
# catsay

`catsay` is a fun and customizable command-line program written in Go that prints a random or custom message inside a cat-themed speech bubble. The program features random ASCII cat faces, random colors, and an optional typewriter effect for the message.

## Features

- **Random Cats**: Displays one of many possible ASCII cat faces.
- **Random Colors**: The cat and its speech bubble are printed in a random color.
- **Custom Quotes**: Displays a random quote from your custom list (`myQuotes`) if no message is provided.
- **Custom Messages**: Use the `-m` flag to pass a custom message to the cat.
- **Typewriter Effect**: The message is printed with a typewriter-style effect, with customizable speed or the option to disable it.
- **Flags**:
  - `-m`: Pass a custom message.
  - `-speed`: Adjust the speed of the typewriter effect (in milliseconds per character).
  - `-disable-typewriter`: Disable the typewriter effect and print the message instantly.

## Installation

To run this program, you'll need [Go](https://golang.org/dl/) installed on your machine.

1. Clone or download the repository:

   ```bash
   git clone <repository-url>
   cd <repository-folder>
   ```

2. Run the program using:

   ```bash
   go run catsay.go
   ```

## Usage

You can run `catsay` with a variety of options using flags.

### Run with a Custom Message

To display a custom message:

```bash
go run catsay.go -m "Hello, I'm a cool cat!"
```

Example output:

```
----------------------
| Hello, I'm a cool cat! |
----------------------
   /\_/\
  ( @_@ )
   > ^ <
```

### Run with a Random Quote

If you run the program without the `-m` flag, it will print a random quote from the custom `myQuotes` list:

```bash
go run catsay.go
```

Example output:

```
----------------------------------------------------
| Hello, World! |
----------------------------------------------------
   /\_/\
  ( o.o )
   > ^ <
```

### Adjust Typewriter Speed

To change the speed of the typewriter effect, use the `-speed` flag with the desired milliseconds per character:

```bash
go run catsay.go -m "Fast typing cat!" -speed 10
```

This prints the message with each character appearing every 10 milliseconds.

### Disable Typewriter Effect

To disable the typewriter effect and print the message instantly, use the `-disable-typewriter` flag:

```bash
go run catsay.go -m "Instant cat!" -disable-typewriter
```

### Example with All Options

```bash
go run catsay.go -m "I am a speedy cat!" -speed 20
```

This command prints a custom message with a typewriter speed of 20 milliseconds per character.

## Customization

### Adding More Cats

You can extend the program by adding more ASCII cat variations to the `cats` array in the code:

```go
var cats = []string{
	"  /\\_/\\\n ( o.o )\n  > ^ <",
	"  /\\_/\\\n ( -.- )\n  > ^ <",
	"  /\\_/\\\n ( O.O )\n  > ^ <",
  //add custom cats here
}
```

### Adding More Quotes

To add more custom quotes to the `myQuotes` array:

```go
var myQuotes = []string{
    "Hello, World!",
    "The cake is a lie",
    // Add your quotes here
}
```

### Adjusting Typewriter Speed

You can change the default typewriter speed by adjusting the default value in the `-speed` flag:

```go
speedFlag := flag.Int("speed", 50, "Set typewriter speed in milliseconds per character")
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
