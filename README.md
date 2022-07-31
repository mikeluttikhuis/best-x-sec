## Best-x-Sec
Loads data from GPX files and calculates your best power performance over the supplied timeframe.

### Usage

```
$ go install github.com/mikeluttikhuis/best-x-sec@latest
$ ./best-x-sec <path_to_gpx_file> <seconds to calculate max power for>
```

### Example

```
$ ./best-x-sec ride.gpx 30
Best 30s power: 544w
```