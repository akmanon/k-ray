package output

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/akmanon/k-ray/pkg/models"
)

func PrintTable(findings []models.Findings) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "SEVERITY\tNAMESPACE\tRESOURCE\tREASON\tRESTARTS")

	for _, f := range findings {
		fmt.Fprintf(
			w,
			"%s\t%s\t%s\t%s\t%d\n",
			f.Severity,
			f.Namespace,
			f.Resource,
			f.Reason,
			f.Restarts,
		)
	}
	w.Flush()
}
