package core

func AppendIfNotExisted(targets []string, item string) []string {
	for _, s := range targets {
		if s == item {
			return targets
		}
	}

	return append(targets, item)
}
