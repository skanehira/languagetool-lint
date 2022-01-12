![GitHub Repo stars](https://img.shields.io/github/stars/skanehira/languagetool-lint?style=social)
![GitHub](https://img.shields.io/github/license/skanehira/languagetool-lint)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/skanehira/languagetool-lint)
![GitHub all releases](https://img.shields.io/github/downloads/skanehira/languagetool-lint/total)
![GitHub CI Status](https://img.shields.io/github/workflow/status/skanehira/languagetool-lint/ci?label=CI)
![GitHub Release Status](https://img.shields.io/github/workflow/status/skanehira/languagetool-lint/Release?label=release)

# languagetool-lint
Lint CLI for [languagetool](https://github.com/languagetool-org/languagetool).

## Requirements
- [languagetool](https://github.com/languagetool-org/languagetool).

## Installation

```sh
$ go install github.com/skanehira/languagetool-lint@latest
```

## Use as a lint tool
1. Run your `languagetool server` in local.  
   NOTE: You can also use [docker-languagetool](https://github.com/Erikvl87/docker-languagetool) to run `languagetool server`.
2. Execute `languagetool-lint` like bellow.
   ```sh
   $ languagetool-lint -a http://localhost:8081 -l "en-US" your_text_file
   your_text_file:2:27: The abbreviation “e.g.” (= for example) requires two periods.
   ```

## Use as a LSP Server
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
3. Add your LSP client settings.  
   e.g. coc.nvim
   ```
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
