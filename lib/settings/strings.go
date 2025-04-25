package settings

import (
	"html/template"
	"maps"
)

type Strings map[string]string

var DefaultStrings = make(Strings).set(map[string]string{
	"title_challenge": "Checking you are not a bot",
	"title_error":     "Oh no!",

	"noscript_warning": "<p>Sadly, you may need to enable JavaScript to get past this challenge. This is required because AI companies have changed the social contract around how website hosting works.</p>",

	"details_title": "Why am I seeing this?",
	"details_text": `
<p>
	You are seeing this because the administrator of this website has set up <a href="https://git.gammaspectra.live/git/go-away">go-away</a> 
	to protect the server against the scourge of <a href="https://thelibre.news/foss-infrastructure-is-under-attack-by-ai-companies/">AI companies aggressively scraping websites</a>.
</p>
<p>
	Mass scraping can and does cause downtime for the websites, which makes their resources inaccessible for everyone.
</p>
<p>
	Please note that some challenges requires the use of modern JavaScript features and some plugins may disable these.
	Disable such plugins for this domain (for example, JShelter) if you encounter any issues.
</p>
`,
	"details_contact_admin_with_request_id": "If you have any issues contact the site administrator and provide the following Request Id",

	"button_refresh_page": "Refresh page",

	"status_loading_challenge":   "Loading challenge",
	"status_starting_challenge":  "Starting challenge",
	"status_loading":             "Loading...",
	"status_calculating":         "Calculating...",
	"status_challenge_success":   "Challenge success!",
	"status_challenge_done_took": "Done! Took",
	"status_error":               "Error:",
})

func (s Strings) set(v map[string]string) Strings {
	maps.Copy(s, v)
	return s
}

func (s Strings) Get(value string) template.HTML {
	v, ok := (s)[value]
	if !ok {
		// fallback
		return template.HTML("string:" + value)
	}
	return template.HTML(v)
}
