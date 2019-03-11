# Glow - Git-Flow (for Gitlab)

![glow logo](./assets/glow-logo.svg)

You can use this tool in any git project of course. But there are some commands which are bound to the Gitlab api. So therefore I promote this as a tool tailored for Gitlab.

## Installation

Here you can find all [available Binaries](https://github.com/meinto/glow/releases). Download the binary and run the install command:

```bash
<name-of-binary> install
```

## Workflow

> Important!  
> Some commands need additional information like git author or Gitlab namespace and project name.  
> These informations can be stored in a config file or can be passed through flags.
> To configure the glow.json config file run the "init" command

![glow workflow](./assets/glow.jpg?raw=true)

### Feature Development

The following command will create a new feature branch in the following form: `features/dvader/death-star`. The name of the author (`dvader`) is grabbed from the config file.

```bash
# author grabbed from config
glow feature death-star
```

After you created the feature branch it is automatically checked out.  
When you finish your feature you can create a merge request in Gitlab:

```bash
# Gitlab information grabbed from config
glow close
```

### Create a release

I recommend to use [Semver](https://semver.org/) for versioning. The following command will create a release branch with the following format: `release/v1.2.3`.

```bash
glow release 1.2.3
```

### Publish a release

When you decide that the release is stable and you want to publish it, the following command will create a merge request on the `master` branch in Gitlab.

```bash
glow publish
```

### Close a release

After publishing the release, you have to merge all changes made on the release branch back into `develop`. The following command creates a merge request of the release branch into `develop`.

```bash
glow close
```

## Config

For some commands you must provide information like the url of your Gitlab instance or your Gitlab ci token. These informations can be put in a `glow.json` file. Glow will lookup this json in the directory where its executed.

You can create this json with the `init` command. The json will be automatically added to the `.gitignore`:

```bash
glow init
```

**List of all config params**

```json
{
  "author": "dvadar",
  "gitlabEndpoint": "https://gitlab.com",
  "projectNamespace": "my-namespace",
  "projectName": "my-project",
  "gitlabCIToken": "abc",
  "gitPath": "/usr/local/bin/git",
  "useBuiltInGitBindings": false,
}
```

## Built in git bindings

glow uses the native git installation per default. The default configured path to git is `/usr/local/bin/git`. You can change the path with the flag `--gitPath` or the property `gitPath` in the config file.

You also can use the built in git client by setting the property "useBuiltInGitBindings" in the config file to `true`. The built in git client uses the library [go-git](https://github.com/src-d/go-git) which unforunately has some perfomance problems. Therefore this tool tries to use a native git installation first.