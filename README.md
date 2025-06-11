# goEEW

This repository demonstrates basic Modified Mercalli Intensity (MMI) calculations in Go.
The `mmi` package implements several empirical intensity prediction equations (IPEs)
based on published studies, and `main.go` provides a small example that prints the
predicted intensity for a given magnitude and distance.

## Implemented IPEs

The [`mmi`](mmi/) package contains Go implementations of the following MMI models:

- **Bakun & Wentworth (1997)** – calibrated for Californian events
- **Atkinson & Wald (2007)** – recommended for near‑source distances
- **Allen & Wald (2012)** – global active‑crust model

A helper function `BestEstimate` chooses between these equations based on simple
distance heuristics.

## Building and running

Go modules are used, so Go automatically resolves dependencies. To build the
example program run:

```bash
go build
```

This creates an executable named `goeew` in the current directory. You can then
run it directly:

```bash
./goeew
```

Alternatively, you can run the example without building a binary using:

```bash
go run .
```

The program prints the predicted MMI values for each model and a best estimate
for the magnitude and distance hard‑coded in `main.go`.
