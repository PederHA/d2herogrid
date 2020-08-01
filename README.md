# d2herogrid

![logo](assets/logo/logo.png)

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

## Examples

### New Immortal Main Stat Grid (Default)

```bash
> ./d2herogrid
```

### New Divine Main Stat Grid

```bash
> ./d2herogrid divine
```

### New Divine Attack Type Grid

```bash
> ./d2herogrid -layout a divine
```

### New Herald, Guardian & Crusader Role Grids

```bash
> ./d2herogrid -layout r herald guardian crusader
```



## Notes

I am so awful at structuring my Go projects.