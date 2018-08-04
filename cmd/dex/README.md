# dex entry point

This is our custom dex entry point. It is here so we can add custom connectors without having to fully fork the repo, rather we can treat the majority of dex as a dependency.

Unfortunately, we can't import the main like that. Because of this, it is basically copied from https://github.com/coreos/dex/tree/master/cmd/dex verbatim here.

Some modifications will be likely to add our own custom auth endpoints and services. When this is done, we should record them here to allow easier reconciliation when we bring in changes from upstream. Ideally our code should be in it's own file/packages, with the very bare minimum to inject it in the upstream files.

## Modifications