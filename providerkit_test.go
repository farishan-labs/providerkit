package providerkit

import "testing"

func TestSlugify(t *testing.T) {
	cases := []struct{ in, want string }{
		{"OpenAI", "openai"},
		{"Together AI", "together-ai"},
		{"z.ai", "z-ai"},
		{"Foo  Bar", "foo-bar"},
		{"a__b..c", "a-b-c"},
		{"", ""},
	}
	for _, c := range cases {
		if got := Slugify(c.in); got != c.want {
			t.Errorf("Slugify(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestPrettify(t *testing.T) {
	if got := Prettify("openai"); got != "OpenAI" {
		t.Errorf("got %q", got)
	}
	if got := Prettify("together"); got != "Together AI" {
		t.Errorf("got %q", got)
	}
	if got := Prettify("foo-bar"); got != "Foo Bar" {
		t.Errorf("got %q", got)
	}
}

func TestHasTag(t *testing.T) {
	tags := []string{"LLM", " free "}
	if !HasTag(tags, "llm") {
		t.Error("expected true")
	}
	if !HasTag(tags, " free ") {
		t.Error("expected true")
	}
	if HasTag(tags, "vision") {
		t.Error("expected false")
	}
}

func TestSplitModelID(t *testing.T) {
	p, m, ok := SplitModelID("openai/gpt-4o")
	if !ok || p != "openai" || m != "gpt-4o" {
		t.Fatalf("got %q %q %v", p, m, ok)
	}
	if _, _, ok := SplitModelID("nope"); ok {
		t.Error("expected false")
	}
	if _, _, ok := SplitModelID("/x"); ok {
		t.Error("expected false")
	}
}

func TestBuildURL(t *testing.T) {
	if got := BuildURL("https://example.com/", "/models"); got != "https://example.com/models" {
		t.Errorf("got %q", got)
	}
	if got := BuildURL("https://example.com", "models"); got != "https://example.com/models" {
		t.Errorf("got %q", got)
	}
	if got := BuildURL("https://example.com/", ""); got != "https://example.com" {
		t.Errorf("got %q", got)
	}
}

func TestRedact(t *testing.T) {
	if got := Redact("key=secret123", "secret123"); got != "key=***" {
		t.Errorf("got %q", got)
	}
	if got := Redact("x", ""); got != "x" {
		t.Errorf("got %q", got)
	}
}

func TestVersion(t *testing.T) {
	if Version == "" {
		t.Fatal("Version must not be empty")
	}
}

func TestJoinSplitTags(t *testing.T) {
	if got := JoinTags([]string{" a ", "", "b"}); got != "a,b" {
		t.Errorf("got %q", got)
	}
	if got := len(SplitTags(" a , , b ")); got != 2 {
		t.Errorf("got %d", got)
	}
}
