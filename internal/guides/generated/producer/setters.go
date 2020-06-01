// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "mdtogo"; DO NOT EDIT.
package producer

var SettersGuide = `Setters provide a solution for template-free setting or substitution of field
values through package metadata (OpenAPI).  They are a safer alternative to
other substitution techniques which do not have the context of the
structured data -- e.g. using ` + "`" + `sed` + "`" + ` to replace values.

The OpenAPI definitions for setters are defined in a Kptfile and referenced by
a fields through comments on the fields.

Setters may be invoked to programmatically modify the configuration
using ` + "`" + `kpt cfg set` + "`" + ` to set and/or substitute values.

{{% pageinfo color="primary" %}}
Creating a setter requires that the package has a Kptfile.  If one does
not exist for the package, run ` + "`" + `kpt pkg init DIR/` + "`" + ` to create one.
{{% /pageinfo %}}

## Setters explained

Following is a short explanation of the command that will be demonstrated
in this guide.

### Data model

- Fields reference setters through OpenAPI definitions specified as
  line comments -- e.g. ` + "`" + `# { "$ref": "#/definitions/..." }` + "`" + `
- OpenAPI definitions are provided through the Kptfile

### Command control flow

1. Read the package Kptfile and resources.
2. Change the setter OpenAPI value in the Kptfile
3. Locate all fields which reference the setter and change their values.
4. Write both the modified Kptfile and resources back to the package.

{{< svg src="images/set-command" >}}

#### Creating a Setter

Setters may be created either manually (by editing the Kptfile directly), or
programmatically (through the ` + "`" + `create-setter` + "`" + ` command).  The ` + "`" + `create-setter` + "`" + `
command will:

1. create a new OpenAPI definition for a setter in the Kptfile
2. create references to the setter definition on the resource fields

  # Kptfile -- original
  openAPI:
    definitions: {}

  # deployment.yaml -- original
  kind: Deployment
  metadata:
    name: foo
  spec:
    replicas: 3

  # create or update a setter named "replicas"
  # match fields with the value "3"
  kpt cfg create-setter hello-world/ replicas 3

  # Kptfile -- updated
  openAPI:
    definitions:
      io.k8s.cli.setters.replicas:
        x-k8s-cli:
          setter:
            name: "replicas"
            value: "3"

  # deployment.yaml -- updated
  kind: Deployment
  metadata:
    name: foo
  spec:
    replicas: 3 # {"$ref":"#/definitions/io.k8s.cli.setters.replicas"}

#### Invoking a Setter

  # deployment.yaml -- original
  kind: Deployment
  metadata:
   name: helloworld-gke
   labels:
     app: hello
  spec:
   replicas: 3 # {"$ref":"#/definitions/io.k8s.cli.setters.replicas"}

  # set the replicas field to 5
  kpt cfg set DIR/ replicas 5

  # deployment.yaml -- updated
  kind: Deployment
  metadata:
   name: helloworld-gke
   labels:
     app: hello
  spec:
   replicas: 5 # {"$ref":"#/definitions/io.k8s.cli.setters.replicas"}

#### Types

Setters may have types specified which ensure that the configuration is always
serialized correctly as yaml 1.1 -- e.g. if a string field such as an
annotation or arg has the value "on", then it would need to be quoted otherwise
it will be parsed as a bool by yaml 1.1.

This may be done by modifying the Kptfile OpenAPI definitions as shown here:

  openAPI:
    definitions:
      io.k8s.cli.setters.version:
        x-k8s-cli:
          setter:
            name: "version"
            value: "3"
        type: string

Set would change the configuration like this:

  kind: Deployment
  metadata:
    name: foo
    annotations:
      version: "3" # {"$ref":"#/definitions/io.k8s.cli.setters.version"}

#### Setting Lists

It is possible to create setters for fields which are a list of strings.
The setter type must be ` + "`" + `array` + "`" + `, and the reference must be on the list field.
The list setter will take variable args for its value rather than a single value.

**Note:** Currently ` + "`" + `create-setter` + "`" + ` will not directly create a setter reference for a
list field.  The simplest way to create a list setter is to create a setter for one of
the elements, and then move the reference to the list field.


  # example.yaml
  apiVersion: example.com/v1beta1
  kind: Example
  spec:
    list:
    - "a"

  # Kptfile
  kind: Kptfile


` + "`" + `$ kpt cfg create-setter --type array . list a` + "`" + `

**Note:** Move the setter reference from the element (` + "`" + `- "a"` + "`" + `) to the list (` + "`" + `list: ` + "`" + `)

  # example.yaml
  apiVersion: example.com/v1beta1
  kind: Example
  spec:
    list: # {"$ref":"#/definitions/io.k8s.cli.setters.list"}
    - "a"

  # Kptfile
  kind: Kptfile
  openAPI:
    definitions:
      io.k8s.cli.setters.list:
        type: array
        x-k8s-cli:
          setter:
            name: list
            listValues:
            - "a"

` + "`" + `$ kpt cfg set . list a b c` + "`" + `

  # example.yaml
  apiVersion: example.com/v1beta1
  kind: Example
  spec:
    list: # {"$ref":"#/definitions/io.k8s.cli.setters.list"}
    - "a"
    - "b"
    - "c"

  # Kptfile
  kind: Kptfile
  openAPI:
    definitions:
      io.k8s.cli.setters.list:
        type: array
        x-k8s-cli:
          setter:
            name: list
            listValues:
            - "a"
            - "b"
            - "c"

#### Enumerations

Setters may be configured to map an enum input to a different value set
in the configuration.

e.g. users set ` + "`" + `small` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `large` + "`" + ` cpu sizes, and these are mapped
to numeric values set in the configuration.

This may be done by modifying the Kptfile OpenAPI definitions as shown here:

  openAPI:
    definitions:
      io.k8s.cli.setters.cpu:
        x-k8s-cli:
          setter:
            name: "cpu"
            value: "small"
            # enumValues will replace the user provided key with the
            # map value when setting fields.
            enumValues:
              small: "0.5"
              medium: "2"
              large: "4"

Set would change the configuration like this:

  kind: Deployment
  metadata:
    name: foo
  spec:
    template:
      spec:
        containers:
        - name: foo
      resources:
        requests:
          cpu: "0.5" # {"$ref":"#/definitions/io.k8s.cli.setters.cpu"}`