## pxc cluster list-by-context

Show Portworx and Kubernetes information for every context in your kubeconfig

### Synopsis

Show Portworx and Kubernetes information for every context in your kubeconfig

```
pxc cluster list-by-context [flags]
```

### Examples

```

  # Scan through all the contexts in your kubeconfig and display cluster information
  pxc cluster list-by-context

  # Show cluster information for contexts that match contain 'east-coast' in the name
  pxc cluster list-by-context --context-match='*east-coast*,*south*'

  # Output all Kubernetes and Portworx cluster information as json for your contexts that match '*on-prem*'
  pxc cluster list-by-context -o json --context-match='*on-prem*'
```

### Options

```
      --context-match string   Comma separated list of expressions match the appropriate context
  -h, --help                   help for list-by-context
  -o, --output string          Output in yaml|json|wide
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

* [pxc cluster](pxc_cluster.md)	 - Manage Portworx cluster

###### Auto generated by spf13/cobra on 18-Sep-2020