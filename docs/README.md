# Docs

## Files
- `roller.yaml` - template configuration
- `.roller.state.yaml` - tracks state changes (tracked files, used to detect deleted files)


## Configuration
### roller.yaml

````
template:
  repo: [url to git repo]
  vars:
    key: value
    <free-form fields>
  replace:
    key: value
    <free-form fields>
  ignore:
    <pattern of files to ignore>

action:
  [action name: cant be a reserved keyword]:
    shell: [shell command here]
    working_dir: [working dir, optional, defaults to root dir]
````


### Templating Engine
- <https://pkg.go.dev/text/template>
- <https://github.com/flosch/pongo2>
