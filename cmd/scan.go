package cmd

import (
	"fmt"

	"github.com/akmanon/k-ray/internal/k8s"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var nameSpace string
var output string

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan kubernetes cluster for health issues",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := k8s.NewClient()
		if err != nil {
			fmt.Println("Failed to connect to k8s cluster", err)
			return
		}
		ns := nameSpace
		if ns == "" {
			ns = metav1.NamespaceAll
		}
		findings, err := k8s.ScanPods(client, nameSpace)

		if len(findings) == 0 {
			fmt.Println("No crashing pods found")
			return
		}

		fmt.Println(findings)
		_ = client
	},
}

func init() {
	scanCmd.Flags().StringVarP(&nameSpace, "namespace", "n", "", "Namespace to scan")
	scanCmd.Flags().StringVarP(&output, "output", "o", "table", "Output format (table|json)")
	rootCmd.AddCommand(scanCmd)
}
