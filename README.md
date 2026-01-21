# madc
MAnage Docker Compose easily.

# why
I built this project because I run much selfhosted applications in a low hardware, and need to toggle then all time.

# use
put the your containers in madc.conf:
**WARNING**: is automatic added your home directory before the path
```conf
container_name:/path/to/file.yml
```

list the containers
```console
$ ./madc
```

up container:
```console
$ madc container_name u
```

down container:
```console
$ madc container_name d
```

# todo
- change it to a tui
