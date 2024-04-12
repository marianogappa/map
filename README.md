# map

CLI tool that takes `(lat,long)` markers from STDIN and opens a World Map in the browser with those markers.

<img width="1874" alt="Screenshot 2024-04-11 at 18 23 31" src="https://github.com/marianogappa/wm/assets/1078546/98b9e8f9-321b-4429-aec7-ed42fc81ce92">

## Installation

```bash
go install github.com/marianogappa/map@latest
```

## Usage

```bash
echo "52.5170365,13.3888599,Berlin" | map -separator comma
```

If you don't have the coordinates, use [locator](https://github.com/marianogappa/locator)!

```bash
$ echo "Berlin" | locator | map
```

## Notes

- By default it reads tab-separated "{lat} {long} {label}" lines from STDIN
- Uses [Leaflet.js](https://leafletjs.com/) for drawing the map
  
## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/marianogappa/wm/blob/main/LICENSE) file for details.
