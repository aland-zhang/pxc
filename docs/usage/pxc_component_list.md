## pxc component list

List all visible component executables

### Synopsis


		List all available component files on a user's PATH.

		Available component files are those that are:
		- executable
		- anywhere on the user's PATH
		- begin with "pxc-"

```
pxc component list [flags]
```

### Options

```
  -h, --help        help for list
      --name-only   If true, display only the binary name of each component, rather than its full path
```

### Options inherited from parent commands

```
      --pxc.config string       Config file (default is $HOME/.pxc/config.yml)
      --pxc.config-dir string   Config directory (default "/home/lpabon/.pxc")
      --pxc.context string      Force context name for the command
      --pxc.token string        Portworx authentication token
      --pxc.v int32             [0-3] Log level verbosity
```

### SEE ALSO

* [pxc component](pxc_component.md)	 - Provides utilities for interacting with components

###### Auto generated by spf13/cobra on 2-Jul-2020