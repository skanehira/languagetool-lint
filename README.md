![GitHub Repo stars](https://img.shields.io/github/stars/skanehira/languagetool-lint?style=social)
![GitHub](https://img.shields.io/github/license/skanehira/languagetool-lint)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/skanehira/languagetool-lint)
![GitHub all releases](https://img.shields.io/github/downloads/skanehira/languagetool-lint/total)
![GitHub CI Status](https://img.shields.io/github/workflow/status/skanehira/languagetool-lint/ci?label=CI)
![GitHub Release Status](https://img.shields.io/github/workflow/status/skanehira/languagetool-lint/Release?label=release)

# languagetool-lint
Lint CLI for [languagetool](https://github.com/languagetool-org/languagetool).

![](https://i.gyazo.com/a27405b2d85b5b44e57fd505e2b43333.gif)

## Requirements
- [languagetool](https://github.com/languagetool-org/languagetool).

## Installation

```sh
$ go install github.com/skanehira/languagetool-lint@latest
```

## Use as a lint tool
1. Run your `languagetool server` in local.  
   NOTE: You can use [docker-languagetool](https://github.com/Erikvl87/docker-languagetool) to run `languagetool server`.
   ```sh
   $ docker run -d -p 8010:8010 erikvl87/languagetool
   ```
2. Execute `languagetool-lint` like bellow.
   ```sh
   $ cat your_text_file
   When you type |:write|, then it would be executed in terminal, and
   you can choose any options(e.g add labels).

   $ languagetool-lint -a http://localhost:8081 -l "en-US" your_text_file
   your_text_file:2:27: The abbreviation “e.g.” (= for example) requires two periods.
   ```
3. You can also use from stdin.
   ```sh
   $ echo "this is a pen." | languagetool-lint
   stdin:1:0: This sentence does not start with an uppercase letter.
   ```

## Use as a Language Server
1. Install [efm-langserver](https://github.com/mattn/efm-langserver)
2. Add config as bellow.
   ```yaml
   version: 2
   tools:
     languagetool-lint: &languagetool-lint
       lint-command: 'languagetool-lint'
       lint-ignore-exit-code: true
       lint-stdin: true
       lint-formats:
         - '%f:%l:%c: %m'
   languages:
     markdown:
       - <<: *languagetool-lint
   ```
3. Add your Language Client settings.  
   e.g. coc.nvim
   ```vim
   call coc#config('languageserver', {
         \ 'efm': {
           \ 'command': 'efm-langserver',
           \ 'args': [],
           \ 'trace.server': 'verbose',
           \ 'filetypes': ['markdown']
           \ }
         \})
   ```

## Author
skanehira

## Thanks
- [languagetool](https://github.com/languagetool-org/languagetool).
- [docker-languagetool](https://github.com/Erikvl87/docker-languagetool)
