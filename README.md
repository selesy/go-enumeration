# go-enumeration

This project provides a code generator that will allow the easy creation
of Java-style enumerations.

## Installing

The following command will install the go-enumeration executable:

```shell
go get -u github.com/selesy/go-enumeration
```

The executable name is ``go-enumeration`` and it can be found in the
``$GOPATH/bin`` directory.  You should make sure that this directory
is on your shell's path.

## Usage

The go-enumeration code generator uses a private data table annotated
with directives to determine how the table's structure should be
converted into a (effectively) constant enumeration with cargo data.
Instructions detailing the required format of the code are provided
by extensive ``go docs`` which can be viewed using the following
command:

```shell
go docs $GOPATH/src/github.com/selesy/go-enumeration/pkg
```

Certain facets of the code generation process should be handled by
the build pipeline and are therefore documented using the
``go-enumeration`` command's help system.  Execute the following
command to view the executable's help:

```shell
go-enumeration -help
```
