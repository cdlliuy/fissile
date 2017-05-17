## fissile build layer stemcell

Builds a Docker layer that is the base for all images

### Synopsis



This command creates a docker image to be used as a base layer for all role images,
similar to BOSH 'stemcells'.

Fissile will create a Dockerfile and a directory structure with all dependencies in 
`<work-dir>/base_dockerfile`. After that, it will build an image named 
`<repository>-role-base:<FISSILE_VERSION>`.


```
fissile build layer stemcell
```

### Options inherited from parent commands

```
  -c, --cache-dir string         Local BOSH cache directory. (default "~/.bosh/cache")
      --config string            config file (default is $HOME/.fissile.yaml)
  -d, --dark-opinions string     Path to a BOSH deployment manifest file that contains properties that should not have opinionated defaults.
  -F, --from string              Docker image used as a base for the layers (default "ubuntu:14.04")
  -l, --light-opinions string    Path to a BOSH deployment manifest file that contains properties to be used as defaults.
  -M, --metrics string           Path to a CSV file to store timing metrics into.
  -N, --no-build                 If specified, the Dockerfile and assets will be created, but the image won't be built.
  -o, --output string            Choose output format, one of human, json, or yaml (currently only for 'show properties') (default "human")
  -r, --release string           Path to dev BOSH release(s).
  -n, --release-name string      Name of a dev BOSH release; if empty, default configured dev release name will be used
  -v, --release-version string   Version of a dev BOSH release; if empty, the latest dev release will be used
  -p, --repository string        Repository name prefix used to create image names. (default "fissile")
  -m, --role-manifest string     Path to a yaml file that details which jobs are used for each role.
  -w, --work-dir string          Path to the location of the work directory. (default "/var/fissile")
  -W, --workers int              Number of workers to use. (default 2)
```

### SEE ALSO
* [fissile build layer](fissile_build_layer.md)	 - Has subcommands for building Docker layers used during the creation of your images.

###### Auto generated by spf13/cobra on 15-May-2017