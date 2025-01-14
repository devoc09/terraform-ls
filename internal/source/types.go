package source

type Lines []Line

func (l Lines) Copy() Lines {
	newLines := make(Lines, len(l))

	for i, line := range l {
		newLines[i] = line.Copy()
	}

	return newLines
}
