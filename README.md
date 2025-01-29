# Tempfetch

A simple golang CLI app to fetch or forecast weather

## Install

1. Clone the repo
 
```bash
git clone https://github.com/hoangthanhtien/tempfetch.git
cd tempfetch

```
2. Install dependencies

```bash
go mod tidy
```

3. Build

```sh
go build .

```

4. Run

```sh
./tempfetch forecast 

```

Run options: 
- `forecast`: Default get current user location, fetch the weather forecast and print ASCII chart 


Example:
```
 31.90 ┤                                                                                                                                    ╭─╮
 30.83 ┤                                                            ╭╮                    ╭─────╮                  ╭────╮                 ╭─╯ ╰╮
 29.76 ┤                                                           ╭╯╰───╮               ╭╯     ╰╮                ╭╯    ╰╮               ╭╯    ╰─╮                ╭─────╮
 28.69 ┤                                        ╭╮                ╭╯     ╰╮              │       ╰╮              ╭╯      ╰╮             ╭╯       ╰╮              ╭╯     ╰─╮
 27.61 ┤         ╭────╮                    ╭────╯╰╮              ╭╯       ╰╮            ╭╯        ╰─╮           ╭╯        ╰─╮           │         ╰─╮           ╭╯        ╰─╮
 26.54 ┤        ╭╯    ╰───╮              ╭─╯      ╰─╮          ╭─╯         ╰──╮        ╭╯           ╰───╮       │           ╰───╮      ╭╯           ╰─────╮    ╭╯           ╰──
 25.47 ┼╮      ╭╯         ╰──────╮      ╭╯          ╰──╮      ╭╯              ╰───╮   ╭╯                ╰──╮   ╭╯               ╰──────╯                  ╰────╯
 24.40 ┤╰──────╯                 ╰──────╯              ╰──────╯                   ╰───╯                    ╰───╯
                  2025-01-29              2025-01-30              2025-01-31              2025-02-01              2025-02-02              2025-02-03              2025-02-04
```

