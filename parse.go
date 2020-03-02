package matfile_go

type Match struct {
	Metadata
}

//func ParseMatFile(path string) (*Match, error) {
//
//	f, err := os.Open(path)
//	if err != nil {
//		return nil, fmt.Errorf("unable to open mat file: %w", err)
//	}
//	defer f.Close()
//
//	metadata := &Metadata{}
//
//	scanner := bufio.NewScanner(f)
//	for scanner.Scan() {
//		if strings.HasPrefix(`;`, scanner.Text()) {
//			parseMetadataString(scanner.Text(), metadata)
//		}
//	}
//
//	return nil, nil
//}
