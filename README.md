# d2herogrid

![logo](assets/logo/logo.png)

d2herogrid is a simple application that creates custom Dota 2 hero grids based on hero winrates. 

The application is able to generate many different hero grid layouts, as well as sorting heroes by winrate in any skill bracket (+ pro matches).

## Usage
```
> ./d2herogrid [OPTION]... [SKILL BRACKETS]...


DESCRIPTION: 
The SKILL BRACKETS argument is one or more Dota 2 skill brackets that grids should be generated for.

If no skill bracket argument is specified, it defaults to Immortal.

PARAMETERS:
  -l string
        Grid layout (default "mainstat")
            Layouts:
                Main Stat:    "mainstat" / "ms"
                Single:       "single" / "s"
                Attack Type:  "attacktype" / "a"
                Role:         "role" / "r"
                Legs:         "legs" / "l"
                Modify:       "modify" / "m"
  -n string
        Grid name (default "d2hg")
  -p string
        Path to Dota 2 userdata directory (default ".")
  -s    Sort ascending (low-high) [default: high-low]
```

### Examples

#### New Immortal Main Stat Grid (Default)

```bash
> ./d2herogrid
```

#### New Divine Main Stat Grid

```bash
> ./d2herogrid divine
```

#### New Divine Attack Type Grid

```bash
> ./d2herogrid -layout a divine
```

#### New Herald, Guardian & Crusader Role Grids

```bash
> ./d2herogrid -layout role herald guardian crusader
```

## Installation

The easiest way to use d2herogrid is to simply place it in `Steam/userdata/<your_steamID3>/570/cfg` and launch it from within that directory. The application will then modify the existing `hero_grid_config.json` or create a new one if it does not exist. Launching the application with no command-line arguments will result in the application creating a new Immortal Mainstat grid.

Alternatively, the program can be launched from anywhere using the `-p <path>` option, where path is the absolute path to the `[...]/570/cfg` directory.

### Example

```bash
> ./d2herogrid -p "C:\Program Files (x86)\Steam\userdata\<your_steamID3>\570\remote\cfg" [OPTIONS] [ARGS]
```
