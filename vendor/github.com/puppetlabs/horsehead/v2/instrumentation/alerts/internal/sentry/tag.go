package sentry

import "github.com/puppetlabs/horsehead/v2/instrumentation/alerts/trackers"

func tagsToSentryTags(tags []trackers.Tag) map[string]string {
	tm := make(map[string]string, len(tags))
	for _, tag := range tags {
		tm[tag.Key] = tag.Value
	}

	return tm
}
