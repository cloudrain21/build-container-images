Simple framework for creating Dockerfile of various version
===========

You can use this simple framework for creating various versions of Dockerfiles
which include different versions of operating systems and solutios.

You don't have to create simillar Dockerfiles repeatedly.

Just give values in config file for parameters of template.

## Usage

Install python3

```
$ brew install python3
```

Set /usr/local/bin/python3 to your main python script

```
$ ln -s /usr/local/bin/python3 /usr/local/bin/python
$ export PATH=/usr/local/bin:$PATH
```

Create dockerfile template && config docker.ini

Examples

```
$ vi template/Dockerfile.client.template
$ vi conf/docker.ini
```

Run script for creating Dockerfile
(docker filename = Dockerfile.config_section_name)

```
$ python mk_dockerfile.py conf/docker.ini
```

Now you can see created Dockerfile in dockerfile directory.

```
$ ls -al dockerfile/*
```

Build you docker image using newly created Dockerfile.

```
$ docker build -f dockerfile/Docker.xxx.xx --tag image_name:tag .
```
