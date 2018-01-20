package path

import "testing"

func TestFollow(t *testing.T) {
	useCases := map[string]struct {
		diagram string
		path    string
		steps   int
	}{
		"AOC test case": {
			diagram: `     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 
`,
			path:  "ABCDEF",
			steps: 38,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			path, steps := Follow(uc.diagram)
			if path != uc.path {
				t.Fatalf("Expected path was `%s`, got `%s`", uc.path, path)
			}

			if steps != uc.steps {
				t.Fatalf("Expected number of  was `%d`, got `%d`", uc.steps, steps)
			}
		})
	}
}
