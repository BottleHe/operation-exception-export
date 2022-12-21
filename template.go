package main

var (
	expTmp = `// {{ .Note }}
package {{ .Package }};

import work.bottle.plugin.exception.OperationException;

public final class {{ .Name }}Exception extends OperationException {

    public static final {{ .Name }}Exception Default = new {{ .Name }}Exception();

    public {{ .Name }}Exception() {
        super({{ .Code }}, "{{ .Message }}");
    }
}
`
)
