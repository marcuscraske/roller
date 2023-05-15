# Docs

## Files
- `roller.yaml` - template configuration
- `.roller.state.yaml` - tracks state changes (tracked files, used to detect deleted files)


## Configuration
### roller.yaml

````
template:
  repo: [url to git repo, recommend to use ssh url]
  vars:
    key:
      value: value
      description: An optional short description of what this variable does
    <free-form fields>
  replace:
    key: value
    <free-form fields>
  ignore:
    - <pattern of files to ignore>
  actions:
    pre: # optional, executed before applying templating
    - shell: [shell command here]
    post: # optional, executed after applying templating but before diff
    - shell: [shell command here]

action:
  [unqiue action name]:
    shell: [shell command here]
    working_dir: [working dir, optional, defaults to root dir]
````


### Templating Engine
- <https://pkg.go.dev/text/template>
- <https://github.com/flosch/pongo2>
