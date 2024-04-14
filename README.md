# map

CLI tool that takes `(lat,long)` markers from STDIN and opens a World Map in the browser with those markers.

<img width="1884" alt="Screenshot 2024-04-14 at 13 30 14" src="https://github.com/marianogappa/map/assets/1078546/7dc4f4a2-91ba-42c9-b38d-0d4cec795111">


## Installation

```bash
go install github.com/marianogappa/map@latest
```

## Usage

```bash
echo "52.5170365,13.3888599,Berlin" | map -separator comma
```

If you don't have the coordinates, use [gps](https://github.com/marianogappa/gps)!

```bash
$ echo "Berlin" | gps | map
```

## Notes

- By default it reads tab-separated "{lat} {long} {label}" lines from STDIN
- Uses [Leaflet.js](https://leafletjs.com/) for drawing the map
  
## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/marianogappa/wm/blob/main/LICENSE) file for details.
