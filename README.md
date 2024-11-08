## homebrew-template

Template for managing homebrew taps

---

### environment

- Docker
- Golang

---

### node package

- allow `node-formulas` folder to be checked in
- add a new node package description under `node-formulas`

```json
{
  "name": "",
  "description": "",
  "homepage": "",
  "license": ""
}
```

- run the command to upgrade your node package

`./cli.sh upgrade-node-package <org> <name>`