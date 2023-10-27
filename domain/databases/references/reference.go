package references

type reference struct {
	commits     Commits
	contentKeys ContentKeys
}

func createReference(
	commits Commits,
) Reference {
	return createReferenceInternally(commits, nil)
}

func createReferenceWithContentKeys(
	commits Commits,
	contentKeys ContentKeys,
) Reference {
	return createReferenceInternally(commits, contentKeys)
}

func createReferenceInternally(
	commits Commits,
	contentKeys ContentKeys,
) Reference {
	out := reference{
		contentKeys: contentKeys,
		commits:     commits,
	}

	return &out
}

// Commits returns the commits, if any
func (obj *reference) Commits() Commits {
	return obj.commits
}

// HasContentKeys returns true if there is contentKeys, false otherwise
func (obj *reference) HasContentKeys() bool {
	return obj.contentKeys != nil
}

// ContentKeys returns the contentKeys
func (obj *reference) ContentKeys() ContentKeys {
	return obj.contentKeys
}
