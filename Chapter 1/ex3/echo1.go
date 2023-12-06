package main

func Echo1(args []string) string {
	s, sep := "", ""
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}
