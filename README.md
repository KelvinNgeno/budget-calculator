# Budget Calculator (KES)

A simple monthly budget calculator. Enter your income and expenses and
it tells you whether you can save, broke even, or are over budget.

## Download and run

1. Go to the **Releases** page of this repo and download the file for your system:
   - Windows: `budget-windows.exe`
   - Mac (M1/M2/M3): `budget-mac-apple-silicon`
   - Mac (Intel): `budget-mac-intel`
   - Linux: `budget-linux`
2. Run it. Your browser opens at http://localhost:8080.
3. Keep the window open while you use it. Close it to stop.

**Windows:** if you see "Windows protected your PC", click *More info*, then *Run anyway*.
**Mac/Linux:** run `chmod +x <filename>` first, then `./<filename>`.

## Run from source

Requires Go: `go run .`

## Command-line version

`go run ./cli`