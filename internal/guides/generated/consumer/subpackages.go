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
package consumer

var SubpackagesGuide = `
{{% pageinfo color="warning" %}}

#### Notice: Subpackages support is available with kpt version v0.34.0+ for [cfg] commands only

{{% /pageinfo %}}

This guide walks you through an example to get, view, set and apply contents of an
example kpt package with a [subpackage] in its directory tree.

## Steps

1. [Fetch a remote package](#fetch-a-remote-package)
2. [View the package contents](#view-the-package-contents)
3. [Provide the setter values](#provide-the-setter-values)
4. [Apply the package](#apply-the-package)

## Fetch a remote package

Fetch an example kpt package with [subpackage] using [kpt pkg get].

### get command

  kpt pkg get https://github.com/GoogleContainerTools/kpt.git/package-examples/wordpress \
  wordpress

## View the package contents

The primary package artifacts are Kubernetes resource configuration
(e.g. YAML files).

### List package contents in a tree structure

Once you fetch the package onto local list its contents using [tree] command.

  kpt cfg tree wordpress/

Output:

  wordpress
  ├── [Kptfile]  Kptfile wordpress
  ├── [wordpress-deployment.yaml]  Deployment wordpress
  ├── [wordpress-deployment.yaml]  Service wordpress
  ├── [wordpress-deployment.yaml]  PersistentVolumeClaim wp-pv-claim
  └── Pkg: mysql
      ├── [Kptfile]  Kptfile mysql
      ├── [mysql-deployment.yaml]  PersistentVolumeClaim mysql-pv-claim
      ├── [mysql-deployment.yaml]  Deployment wordpress-mysql
      └── [mysql-deployment.yaml]  Service wordpress-mysql

There are two kpt packages in the output:

1. wordpress
2. mysql

` + "`" + `mysql` + "`" + ` package is a [subpackage] of ` + "`" + `wordpress` + "`" + ` package

Optionally, users may use other commands like [count], [grep], [cat] to
further view and understand the package contents.

### List setters in the package

The fetched package contains [setters]. Invoke [list-setters] command to list
the [setters] recursively in all the packages.

  kpt cfg list-setters wordpress/

Output:

  wordpress/
           NAME             VALUE      SET BY   DESCRIPTION   COUNT   REQUIRED
    gcloud.core.project   PROJECT_ID                          3       No
    image                 wordpress                           1       No
    tag                   4.8                                 1       No
    teamname              YOURTEAM                            3       Yes
  
  wordpress/mysql/
           NAME             VALUE      SET BY   DESCRIPTION   COUNT   REQUIRED
    gcloud.core.project   PROJECT_ID                          3       No
    image                 wordpress                           1       No
    tag                   4.8                                 1       No
    teamname              YOURTEAM                            3       Yes

You may notice that the [auto-setter] ` + "`" + `gcloud.core.project` + "`" + ` is already set if you
have ` + "`" + `gcloud` + "`" + ` configured on your local.

## Provide the setter values

Provide the values for all the [required setters]. By default, [set] 
command is performed only on the resource files of provided package and not its 
subpackages. ` + "`" + `--recurse-subpackages(-R)` + "`" + ` can be leveraged to run the command on 
subpackages recursively.

  kpt cfg set wordpress/ teamname myteam -R

Output:

  wordpress/
  set 3 field(s)
  
  wordpress/mysql/
  set 3 field(s)

## Apply the package

Now that you have configured the package, apply it to the cluster

  kubectl apply -f wordpress/ -R

Output:

  service/wordpress-mysql created
  persistentvolumeclaim/mysql-pv-claim created
  deployment.apps/wordpress-mysql created
  service/wordpress created
  persistentvolumeclaim/wp-pv-claim created
  deployment.apps/wordpress created
`