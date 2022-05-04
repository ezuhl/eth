package main

import (
	"context"
	"fmt"
	"github.com/open-policy-agent/opa/rego"
	"github.com/sanity-io/litter"
)

var module = `
package example.authz

default allow = false

allow {
    input.method == "GET"
    input.path == ["salary", input.subject.user]
}

allow {
    is_admin
}

is_admin {
    input.subject.groups[_] = "admin"
}
`

func main(){


	ctx := context.TODO()
	input := map[string]interface{}{
		"method": "GET",
		"path": []interface{}{"salary", "bob"},
		"subject": map[string]interface{}{
			"user": "bob",
			"groups": []interface{}{"sales", "marketing"},
		},
	}

	query, err := eval(ctx)


	results, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		litter.Dump(err)
		return
	}
	//if !results.Allowed() {
	//	fmt.Println("not allowed")
	//	litter.Dump(results)
	//	return
	//}

	fmt.Println("allowed")
	if err != nil {
		// Handle evaluation error.
		litter.Dump(results)
		return
	} else if len(results) == 0 {
		// Handle undefined result.
		litter.Dump(results)
		//	return
	} else if result, ok := results[0].Bindings["x"].(bool); !ok {
		// Handle unexpected result type.
		litter.Dump(result)
	} else {
		// Handle result/decision.
		fmt.Printf("%+v", results) // => [{Expressions:[true] Bindings:map[x:true]}]
		litter.Dump(results)
		//	return
	}

	litter.Dump(results)
}



func eval(ctx context.Context) (rego.PreparedEvalQuery,error){



	query  := rego.New(
		rego.Query("x = data.example.authz.allow"),
		rego.Module("example.rego", module),
	)

	return query.PrepareForEval(ctx)
}