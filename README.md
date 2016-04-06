# Dot

How to put some dots on a SVG canvas with Go?

## Usage

Clone the repository.

On OS X:

    $ go run *.go > dot.svg
    $ open -a “Google Chrome” dot.svg

Everywhere else:

    $ go run *.go > dot.svg

You can build the project with:

    $ go build -o dot *.go

Then you can use the generated binary:

    $ ./dot -h
    Usage of ./dot:
      -bc string
            The background color. (default "#234d51")
      -dc string
            The dots colors. (default "#ff513f,#ffce00")
      -generator string
            Choose your generator (simple, centered, grid, sinusoidal). (default "simple")
      -height int
            The height of the canvas. (default 720)
      -n int
            The number of dots to draw. (default 1000)
      -radius int
            The dots radius. (default 5)
      -width int
            The width of the canvas. (default 1280)

    $ ./dot -generate centered > dot.svg

## Notes

I’m using this scheme: https://color.adobe.com/Gold-Coral-Teals-color-theme-7174629

## License

This MIT License
