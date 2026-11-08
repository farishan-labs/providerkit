package providerkit

import (
	"regexp"
	"strings"
)

func Slugify(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "_", "-")
	s = strings.ReplaceAll(s, ".", "-")
	re := regexp.MustCompile(`[^a-z0-9-]+`)
	s = re.ReplaceAllString(s, "-")
	re2 := regexp.MustCompile(`-+`)
	s = re2.ReplaceAllString(s, "-")
	return strings.Trim(s, "-")
}

var prettifyMap = map[string]string{
	"openai":      "OpenAI",
	"anthropic":   "Anthropic",
	"google":      "Google",
	"meta":        "Meta",
	"mistral":     "Mistral",
	"cohere":      "Cohere",
	"nvidia":      "NVIDIA",
	"deepseek":    "DeepSeek",
	"z-ai":        "Z.ai",
	"groq":        "Groq",
	"cerebras":    "Cerebras",
	"cloudflare":  "Cloudflare",
	"together":    "Together AI",
	"fireworks":   "Fireworks AI",
	"novita":      "Novita AI",
	"perplexity":  "Perplexity",
	"sambanova":   "SambaNova",
	"replicate":   "Replicate",
	"huggingface": "Hugging Face",
	"x-ai":        "xAI",
	"openrouter":  "OpenRouter",
}

func Prettify(slug string) string {
	slug = Slugify(slug)
	if pretty, ok := prettifyMap[slug]; ok {
		return pretty
	}
	words := strings.Split(slug, "-")
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}

func HasTag(tags []string, tag string) bool {
	for _, t := range tags {
		if strings.EqualFold(strings.TrimSpace(t), strings.TrimSpace(tag)) {
			return true
		}
	}
	return false
}

func SplitModelID(id string) (string, string, bool) {
	parts := strings.SplitN(strings.TrimSpace(id), "/", 2)
	if len(parts) < 2 || parts[0] == "" || parts[1] == "" {
		return "", "", false
	}
	return parts[0], parts[1], true
}

func BuildURL(base, path string) string {
	base = strings.TrimRight(strings.TrimSpace(base), "/")
	path = strings.TrimSpace(path)
	if path == "" {
		return base
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return base + path
}

func Redact(s, secret string) string {
	if secret == "" {
		return s
	}
	return strings.ReplaceAll(s, secret, "***")
}

func JoinTags(tags []string) string {
	out := make([]string, 0, len(tags))
	for _, t := range tags {
		t = strings.TrimSpace(t)
		if t != "" {
			out = append(out, t)
		}
	}
	return strings.Join(out, ",")
}

func SplitTags(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
