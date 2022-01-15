package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "languagetool-lint",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func exitError(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func init() {
	rootCmd.PersistentFlags().StringP("addr", "a", "http://localhost:8010", "languagetool server host")
	rootCmd.PersistentFlags().StringP("lang", "l", "en-US", "language")
}

type Rule struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

type Context struct {
	Text   string `json:"text,omitempty"`
	Offset int64  `json:"offset,omitempty"`
	Length int64  `json:"length,omitempty"`
}

type Match struct {
	Message      string  `json:"message,omitempty"`
	ShortMessage string  `json:"short_message,omitempty"`
	Context      Context `json:"context,omitempty"`
	Offset       int64   `json:"offset,omitempty"`
	Length       int64   `json:"length,omitempty"`
	Sentence     string  `json:"sentence,omitempty"`
	Rule         Rule    `json:"rule,omitempty"`
}

type Response struct {
	Matches []Match `json:"matches,omitempty"`
}

type Annotation struct {
	Text string `json:"text"`
}

type Request struct {
	Annotation []Annotation `json:"annotation,omitempty"`
}

type Line struct {
	Text string
	Num  int64
}

var skipWords = map[string]struct{}{}

func getpos(lines map[string]Line, word string) [][]int {
	lnumCol := [][]int{}
	// Found word from all lines.
	for text, line := range lines {
		if strings.Contains(text, word) {
			col := strings.Index(line.Text, word)
			lnumCol = append(lnumCol, []int{int(line.Num), col})
		}
	}
	return lnumCol
}

// format is format languagetool response to lint format
func format(fname string, lines map[string]Line, resp Response) ([]string, error) {
	errors := []string{}
	for _, m := range resp.Matches {
		word := m.Context.Text[m.Context.Offset : m.Context.Offset+m.Context.Length]
		if _, skip := skipWords[word]; skip {
			continue
		}
		skipWords[word] = struct{}{}
		lumCol := getpos(lines, word)
		for _, l := range lumCol {
			errors = append(errors, fmt.Sprintf("%s:%d:%d: %s", fname, l[0], l[1], m.Message))
		}
	}
	return errors, nil
}

func Execute() {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		lang, err := rootCmd.PersistentFlags().GetString("lang")
		if err != nil {
			exitError(err)
		}
		addr, err := rootCmd.PersistentFlags().GetString("addr")
		if err != nil {
			exitError(err)
		}

		if lang == "" || addr == "" {
			_ = rootCmd.Help()
			return
		}

		var (
			fname string
			sc    *bufio.Scanner
		)

		// create request
		if len(args) == 0 {
			fname = "stdin"
			sc = bufio.NewScanner(os.Stdin)
		} else {
			fname = args[0]
			f, err := os.Open(fname)
			if err != nil {
				exitError(err)
			}
			defer f.Close()
			sc = bufio.NewScanner(f)
		}
		var lnum int64 = 1
		lines := make(map[string]Line)
		var req Request
		for sc.Scan() {
			t := sc.Text()
			lines[strings.TrimSpace(t)] = Line{
				Text: t,
				Num:  lnum,
			}
			req.Annotation = append(req.Annotation, Annotation{
				Text: t + "\n",
			})
			lnum++
		}

		data, err := json.Marshal(req)
		if err != nil {
			exitError(err)
		}
		v := url.Values{}
		v.Add("data", string(data))
		v.Add("language", lang)

		resp, err := http.DefaultClient.PostForm(addr+"/v2/check", v)
		if err != nil {
			exitError(err)
		}
		defer resp.Body.Close()
		var langResp Response
		if err := json.NewDecoder(resp.Body).Decode(&langResp); err != nil {
			exitError(err)
		}
		result, err := format(fname, lines, langResp)
		if err != nil {
			exitError(err)
		}
		for _, l := range result {
			println(l)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		exitError(err)
	}
}
