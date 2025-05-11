# MMonster Hunter API Library for Golang

Using the [mhdb-wilds](https://github.com/LartTyler/mhdb) API. Currently, only Monster Hunter Wilds is supported by this Library

See the official [mhdb-wilds documentation](https://docs.wilds.mhdb.io/#introduction) for more information

## Usage

```golang
import wildsapi "github.com/Lofter1/mhapi-go"

func main() {
    client := wildsapi.GetDefaultClient()
    skills, err := client.FetchSkills(
        wildsapi.QueryOptions{Page: 1},
    )
}
```