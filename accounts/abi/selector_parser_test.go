name, rest, err := parseIdentifier(unescapedSelector)
	if err != nil {
		return SelectorMarshaling{}, fmt.Errorf("failed to parse selector '%s': %v", unescapedSelector, err)
	}
	args := []interface{}{}
	if len(rest) >= 2 && rest[0] == '(' && rest[1] == ')' {
		rest = rest[2:]
	} else {
		args, rest, err = parseCompositeType(rest)
		if err != nil {
			return SelectorMarshaling{}, fmt.Errorf("failed to parse selector '%s': %v", unescapedSelector, err)
		}
	}
	if len(rest) > 0 {
		return SelectorMarshaling{}, fmt.Errorf("failed to parse selector '%s': unexpected string '%s'", unescapedSelector, rest)
	}

	// Reassemble the fake ABI and constuct the JSON
	fakeArgs, err := assembleArgs(args)
	if err != nil {
		return SelectorMarshaling{}, fmt.Errorf("failed to parse selector: %v", err)
	}

	return SelectorMarshaling{name, "function", fakeArgs}, nil
}
