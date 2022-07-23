package errors

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
	"testing"
)

func TestFormatNew(t *testing.T) {
	tests := []struct {
		error
		format string
		want   string
	}{{
		New("error"),
		"%s",
		"error",
	}, {
		New("error"),
		"%v",
		"error",
	}, {
		New("error"),
		"%+v",
		"error\nMy-IAM/errors.TestFormatNew\n\t" +
			"F:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:26\ntesting.tRunner\n\t" +
			"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
	}, {
		New("error"),
		"%q",
		`"error"`,
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.error, tt.format, tt.want)
	}
}
func testFormatRegexp(t *testing.T, n int, arg interface{}, format, want string) {
	t.Helper()
	got := fmt.Sprintf(format, arg)
	gotLines := strings.SplitN(got, "\n", -1)
	wantLines := strings.SplitN(want, "\n", -1)

	if len(wantLines) > len(gotLines) {
		t.Errorf("test %d: wantLines(%d) > gotLines(%d):\n got: %q\n want: %q", n+1, len(wantLines), len(gotLines), got, want)
		return
	}

	for i, w := range wantLines {
		match, err := regexp.MatchString(w, gotLines[i])
		if err != nil {
			t.Fatal(err)
		}
		if !match {
			t.Errorf("test %d: line %d: fmt.Sprintf(%q, err):\n got: %q\n want: %q", n+1, i+1, format, got, want)
			//t.Errorf("got:%q\n want:%q\n", got, want)
		}
	}
}

func TestFormatErrorf(t *testing.T) {
	tests := []struct {
		error
		format string
		want   string
	}{{
		Errorf("%s", "error"),
		"%s",
		"error",
	}, {
		Errorf("%s", "error"),
		"%v",
		"error",
	}, {
		Errorf("%s", "error"),
		"%+v",
		"error\nMy-IAM/errors.TestFormatErrorf\n\t" +
			"F:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:78\ntesting.tRunner\n\t" +
			"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
	}}
	for i, tt := range tests {
		testFormatRegexp(t, i, tt.error, tt.format, tt.want)
	}
}

func TestFormatWrap(t *testing.T) {
	tests := []struct {
		error
		format string
		want   string
	}{{
		Wrap(New("error"), "error2"),
		"%s",
		"error2",
	}, {
		Wrap(New("error"), "error2"),
		"%v",
		"error2",
	}, {
		Wrap(New("error"), "error2"),
		"%+v",
		"error\nMy-IAM/errors.TestFormatWrap\n\t" +
			"F:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:103\ntesting.tRunner\n\t" +
			"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\t" +
			"D:/Go/src/runtime/asm_amd64.s:1581\nerror2\nMy-IAM/errors.TestFormatWrap\n\t" +
			"F:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:103\ntesting.tRunner\n\t" +
			"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
	}, {
		Wrap(io.EOF, "error"),
		"%s",
		"error",
	}, {
		Wrap(io.EOF, "error"),
		"%v",
		"error",
	}, {
		Wrap(io.EOF, "error"),
		" %+v",
		" EOF\nerror\nMy-IAM/errors.TestFormatWrap\n\t" +
			"F:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:120\ntesting.tRunner\n\t" +
			"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
	}, {
		Wrap(Wrap(io.EOF, "error1"), "error2"),
		"%+v",
		"EOF\nerror1\nMy-IAM/errors.TestFormatWrap\n\t" +
			"F:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:126\ntesting.tRunner\n\t" +
			"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\t" +
			"D:/Go/src/runtime/asm_amd64.s:1581\nerror2\nMy-IAM/errors.TestFormatWrap\n\t" +
			"F:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:126\ntesting.tRunner\n\t" +
			"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
	}, {
		Wrap(New("error with space"), "context"),
		"%q",
		`"context"`,
	}}
	for i, tt := range tests {
		testFormatRegexp(t, i, tt.error, tt.format, tt.want)
	}
}

func TestFormatWrapf(t *testing.T) {
	tests := []struct {
		error
		format string
		want   string
	}{{
		Wrapf(io.EOF, "error%d", 2),
		"%s",
		"error2",
	}, {
		Wrapf(io.EOF, "error%d", 2),
		"%v",
		"error2",
	}, {
		Wrapf(io.EOF, "error%d", 2),
		"%+v",
		"EOF\nerror2\nMy-IAM/errors.TestFormatWrapf\n\t" +
			"F:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:158\ntesting.tRunner\n\t" +
			"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
	}, {
		Wrapf(New("error"), "error%d", 2),
		"%s",
		"error2",
	}, {
		Wrapf(New("error"), "error%d", 2),
		"%v",
		"error2",
	}, {
		Wrapf(New("error"), "error%d", 2),
		"%+v",
		"error\nMy-IAM/errors.TestFormatWrapf\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:163\ntesting.tRunner\n\t" +
			"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581\nerror2\nMy-IAM/errors.TestFormatWrapf\n\t" +
			"F:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:163\ntesting.tRunner\n\tD:/Go/src/testing/testing.go:1259\nruntime.goexit\n\t" +
			"D:/Go/src/runtime/asm_amd64.s:1581",
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.error, tt.format, tt.want)
	}
}

func TestFormatWithMessage(t *testing.T) {
	tests := []struct {
		error
		format string
		want   []string
	}{{
		WithMessage(New("error"), "error2"),
		"%s",
		[]string{"error2"},
	}, {
		WithMessage(New("error"), "error2"),
		"%v",
		[]string{"error2"},
	}, {
		WithMessage(New("error"), "error2"),
		"%+v",
		[]string{
			"error",
			"My-IAM/errors.TestFormatWithMessage\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:191\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
			"error2"},
	}, {
		WithMessage(io.EOF, "addition1"),
		"%s",
		[]string{"addition1"},
	}, {
		WithMessage(io.EOF, "addition1"),
		"%v",
		[]string{"addition1"},
	}, {
		WithMessage(io.EOF, "addition1"),
		"%+v",
		[]string{"EOF", "addition1"},
	}, {
		WithMessage(WithMessage(io.EOF, "addition1"), "addition2"),
		"%v",
		[]string{"addition2"},
	}, {
		WithMessage(WithMessage(io.EOF, "addition1"), "addition2"),
		"%+v",
		[]string{"EOF", "addition1", "addition2"},
	}, {
		Wrap(WithMessage(io.EOF, "error1"), "error2"),
		"%+v",
		[]string{"EOF", "error1", "error2",
			"My-IAM/errors.TestFormatWithMessage\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:219\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
		},
	}, {
		WithMessage(Errorf("error%d", 1), "error2"),
		"%+v",
		[]string{"error1",
			"My-IAM/errors.TestFormatWithMessage\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:226\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
			"error2"},
	}, {
		WithMessage(WithStack(io.EOF), "error"),
		"%+v",
		[]string{
			"EOF",
			"My-IAM/errors.TestFormatWithMessage\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:233\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
			"error"},
	}, {
		WithMessage(Wrap(WithStack(io.EOF), "inside-error"), "outside-error"),
		"%+v",
		[]string{
			"EOF",
			"My-IAM/errors.TestFormatWithMessage\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:241\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
			"inside-error",
			"My-IAM/errors.TestFormatWithMessage\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:241\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
			"outside-error"},
	}}

	for i, tt := range tests {
		testFormatCompleteCompare(t, i, tt.error, tt.format, tt.want, true)
	}
}
func TestFormatGeneric(t *testing.T) {
	starts := []struct {
		err  error
		want []string
	}{
		{New("new-error"), []string{
			"new-error",
			"My-IAM/errors.TestFormatGeneric\n\t" +
				"F:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:262\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
		},
		}, {Errorf("errorf-error"), []string{
			"errorf-error",
			"My-IAM/errors.TestFormatGeneric\n\t" +
				"F:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:268\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
		},
		}, {errors.New("errors-new-error"), []string{
			"errors-new-error"},
		},
	}

	wrappers := []wrapper{
		{
			func(err error) error { return WithMessage(err, "with-message") },
			[]string{"with-message"},
		}, {
			func(err error) error { return WithStack(err) },
			[]string{
				"My-IAM/errors.(func·002|TestFormatGeneric.func2)\n\t" +
					"F:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:284",
			},
		}, {
			func(err error) error { return Wrap(err, "wrap-error") },
			[]string{
				"wrap-error",
				"My-IAM/errors.(func·003|TestFormatGeneric.func3)\n\t" +
					"F:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:290",
			},
		}, {
			func(err error) error { return Wrapf(err, "wrapf-error%d", 1) },
			[]string{
				"wrapf-error1",
				"My-IAM/errors.(func·004|TestFormatGeneric.func4)\n\t" +
					"F:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:297",
			},
		},
	}

	for s := range starts {
		err := starts[s].err
		want := starts[s].want
		testFormatCompleteCompare(t, s, err, "%+v", want, false)
		testGenericRecursive(t, err, want, wrappers, 3)
	}
}

func TestFormatWithStack(t *testing.T) {
	tests := []struct {
		error
		format string
		want   []string
	}{{
		WithStack(io.EOF),
		"%s",
		[]string{"EOF"},
	}, {
		WithStack(io.EOF),
		"%v",
		[]string{"EOF"},
	}, {
		WithStack(io.EOF),
		"%+v",
		[]string{"EOF",
			"My-IAM/errors.TestFormatWithStack\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:191\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
		},
	}, {
		WithStack(New("error")),
		"%s",
		[]string{"error"},
	}, {
		WithStack(New("error")),
		"%v",
		[]string{"error"},
	}, {
		WithStack(New("error")),
		"%+v",
		[]string{"error",
			"My-IAM/errors.TestFormatWithStack\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:206\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
			"My-IAM/errors.TestFormatWithStack\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:206\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
		},
	}, {
		WithStack(WithStack(io.EOF)),
		"%+v",
		[]string{"EOF",
			"My-IAM/errors.TestFormatWithStack\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:215\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
			"My-IAM/errors.TestFormatWithStack\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:215\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
		},
	}, {
		WithStack(WithStack(Wrapf(io.EOF, "message"))),
		"%+v",
		[]string{"EOF",
			"message",
			"My-IAM/errors.TestFormatWithStack\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:224\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
			"My-IAM/errors.TestFormatWithStack\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:224\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
			"My-IAM/errors.TestFormatWithStack\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:224\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
		},
	}, {
		WithStack(Errorf("error%d", 1)),
		"%+v",
		[]string{"error1",
			"My-IAM/errors.TestFormatWithStack\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:236\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
			"My-IAM/errors.TestFormatWithStack\n\tF:/workspace/VScode/项目/IAM/My-IAM/errors/format_test.go:236\ntesting.tRunner\n\t" +
				"D:/Go/src/testing/testing.go:1259\nruntime.goexit\n\tD:/Go/src/runtime/asm_amd64.s:1581",
		},
	}}

	for i, tt := range tests {
		testFormatCompleteCompare(t, i, tt.error, tt.format, tt.want, true)
	}
}

func wrappedNew(message string) error { // This function will be mid-stack inlined in go 1.12+
	return New(message)
}
func TestFormatWrappedNew(t *testing.T) {
	tests := []struct {
		error
		format string
		want   string
	}{{
		wrappedNew("error"),
		"%+v",
		"error\n" +
			"github.com/marmotedu/errors.wrappedNew\n" +
			"\t.+/github.com/marmotedu/errors/format_test.go:364\n" +
			"github.com/marmotedu/errors.TestFormatWrappedNew\n" +
			"\t.+/github.com/marmotedu/errors/format_test.go:373",
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.error, tt.format, tt.want)
	}
}

var stackLineR = regexp.MustCompile(`\.`)

// parseBlocks parses input into a slice, where:
//  - incase entry contains a newline, its a stacktrace
//  - incase entry contains no newline, its a solo line.
//
// Detecting stack boundaries only works incase the WithStack-calls are
// to be found on the same line, thats why it is optionally here.
//
// Example use:
//
// for _, e := range blocks {
//   if strings.ContainsAny(e, "\n") {
//     // Match as stack
//   } else {
//     // Match as line
//   }
// }
//

func parseBlocks(input string, detectStackboundaries bool) ([]string, error) {
	var blocks []string

	stack := ""
	wasSatck := false
	lines := map[string]bool{}
	for _, l := range strings.Split(input, "\n") {
		isStackLine := stackLineR.MatchString(l)

		switch {
		case !isStackLine && wasSatck:
			blocks = append(blocks, stack, l)
			stack = ""
			lines = map[string]bool{}
		case isStackLine:
			if wasSatck {
				// Detecting two stacks after another, possible cause lines match in
				// our tests due to WithStack(WithStack(io.EOF)) on same line.
				if detectStackboundaries {
					if lines[l] {
						if len(stack) == 0 {
							return nil, errors.New("len of block must not be zero here")
						}

						blocks = append(blocks, stack)
						stack = l
						lines = map[string]bool{l: true}
						continue
					}
				}
				stack = stack + "\n" + l
			} else {
				stack = l
			}
			lines[l] = true
		case !isStackLine && !wasSatck:
			blocks = append(blocks, l)
		default:
			return nil, errors.New("must not happen")
		}
		wasSatck = isStackLine
	}
	// Use up stack
	if stack != "" {
		blocks = append(blocks, stack)
	}
	return blocks, nil
}

func testFormatCompleteCompare(t *testing.T, n int, arg interface{}, format string, want []string, detectStackBoundaries bool) {
	gotStr := fmt.Sprintf(format, arg)

	got, err := parseBlocks(gotStr, detectStackBoundaries)
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != len(want) {
		t.Fatalf("test %d: fmt.Sprintf(%s, err) -> wrong number of blocks: got(%d) want(%d)\n got: %s\nwant: %s\ngotStr: %q",
			n+1, format, len(got), len(want), prettyBlocks(got), prettyBlocks(want), gotStr)
	}

	for i := range got {
		if strings.ContainsAny(want[i], "\n") {
			// Match as stack
			match, err := regexp.MatchString(want[i], got[i])
			if err != nil {
				t.Fatal(err)
			}
			if !match {
				t.Fatalf("test %d: block %d: fmt.Sprintf(%q, err):\n got:\n%q\n want:\n%q\n all-got:\n%s\n all-want:\n%s\n",
					n+1, i+1, format, got[i], want[i], prettyBlocks(got), prettyBlocks(want))
			}
		} else {
			// Match as message
			if got[i] != want[i] {
				t.Fatalf("test %d: fmt.Sprintf(%s, err) at block %d got != want:\n got: %q\nwant: %q", n+1, format, i+1, got[i], want[i])
			}
		}
	}
}

type wrapper struct {
	wrap func(err error) error
	want []string
}

func prettyBlocks(blocks []string) string {
	var out []string

	for _, b := range blocks {
		out = append(out, fmt.Sprintf("%v", b))
	}

	return "   " + strings.Join(out, "\n   ")
}

func testGenericRecursive(t *testing.T, beforeErr error, beforeWant []string, list []wrapper, maxDepth int) {
	if len(beforeWant) == 0 {
		panic("beforeWant must not be empty")
	}
	for _, w := range list {
		if len(w.want) == 0 {
			panic("want must not be empty")
		}

		err := w.wrap(beforeErr)

		// Copy required cause append(beforeWant, ..) modified beforeWant subtly.
		beforeCopy := make([]string, len(beforeWant))
		copy(beforeCopy, beforeWant)

		beforeWant := beforeCopy
		last := len(beforeWant) - 1
		var want []string

		// Merge two stacks behind each other.
		if strings.ContainsAny(beforeWant[last], "\n") && strings.ContainsAny(w.want[0], "\n") {
			want = append(beforeWant[:last], append([]string{beforeWant[last] + "((?s).*)" + w.want[0]}, w.want[1:]...)...)
		} else {
			want = append(beforeWant, w.want...)
		}

		testFormatCompleteCompare(t, maxDepth, err, "%+v", want, false)
		if maxDepth > 0 {
			testGenericRecursive(t, err, want, list, maxDepth-1)
		}
	}
}
