# OPA Authorization strategy

The current folder contains the OPA authorization policy definition.
For information about OPA, see:

- [OPA - Github](https://github.com/open-policy-agent/opa/)
- [OPA - website](https://www.openpolicyagent.org/)

## Using REPL (interactive interpreter)

[Download OPA](https://www.openpolicyagent.org/docs/get-started.html#prerequisites)

Run the following command:

```shell
opa run -w authz.rego common.rego policies:../example/policies.json
```

Enter the interactive interpreter:

```shell
OPA 0.9.2 (commit 9fbff4c3, built at 2018-09-24T16:12:26Z)

> data.authz.authorized
false
# This matches against an action/resource from a statement in a policy.
> data.authz.authorized with input as { "subjects": [ "team:local:admins" ], "action": "iam:teams:create", "resource": "iam:teams" }
true
> data.authz.authorized with input as { "subjects": [ "team:local:admins" ], "action": "iam:teams:create", "resource": "iam:users" }
false
# This matches against an action/resource from a statement in a policy.
> data.authz.authorized with input as { "subjects": [ "team:local:admins" ], "action": "infra:nodes:delete", "resource": "infra:nodes" }
true
>
```

## Run OPA unit tests

Run the test and get only the final result:

```shell
opa test authz.rego common.rego authz_test.rego
```

Run the test and print the details:

```shell
opa test -v authz.rego common.rego authz_test.rego
```