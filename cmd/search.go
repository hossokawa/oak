/*
Copyright © 2026 Leonardo Hossokawa <leonardoholiver@gmail.com>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/hossokawa/oak/pkg/pokeapi"
	"github.com/hossokawa/oak/pkg/textutil"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [name_or_id]",
	Short: "Look up a pokémon by name or id",
	RunE: func(cmd *cobra.Command, args []string) error {
		q := args[0]

		client := pokeapi.NewClient()

		p, err := client.GetPokemon(q)
		if err != nil {
			return err
		}

		fmt.Printf("ID: %d\n", p.Id)
		fmt.Printf("Name: %s\n", textutil.Capitalize(textutil.Normalize(p.Name)))

		fmt.Print("Types: ")
		if len(p.Types) == 2 {
			fmt.Printf(
				"%s/%s",
				textutil.Capitalize(textutil.Normalize(p.Types[0].Type.Name)),
				textutil.Capitalize(textutil.Normalize(p.Types[1].Type.Name)),
			)
		} else {
			fmt.Printf("%s", textutil.Capitalize(p.Types[0].Type.Name))
		}

		fmt.Print("\nStats:\n")
		for _, s := range p.Stats {
			if s.Stat.Name == "hp" {
				fmt.Printf(
					"\t%s: %d\n",
					strings.ToUpper(s.Stat.Name),
					s.BaseStat,
				)
				continue
			}
			fmt.Printf(
				"\t%s: %d\n",
				textutil.Capitalize(textutil.Normalize(s.Stat.Name)),
				s.BaseStat,
			)
		}

		fmt.Print("Abilities:")
		for _, a := range p.Abilities {
			fmt.Printf("\n\t%s", textutil.Capitalize(textutil.Normalize(a.Ability.Name)))
			if a.Hidden {
				fmt.Printf(" (Hidden)")
			}
		}
		fmt.Println()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
