package main

type slice []int

func arrayMax(w []int) int {
	var max int
	for i, v := range w {
		if i == 0 || max < v {
			max = v
		}
	}
	return max
}

func arrayContains(w []int, x int) bool {
	for _, v := range w {
		if v == x {
			return true
		}
	}
	return false
}


func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func buildSolution(g slice) slice {
	aux := make([]int, len(g))
	for i := 0; i < len(g); i++ {
		if i > 0 {
			aux[i] = g[i] - g[i-1]
		}
	}
	return aux
}

func mergeSliceOfSlices(a, b []slice) []slice {
	aux := make([]slice, len(a) + len(b))
	for i, v := range a {
		aux[i] = v
	}
	for i, v := range b {
		aux[i + len(a)] = v
	}
	return aux
}

func sliceIncluded(V, W slice) bool {
	for _, v := range V {
		if !arrayContains(W, v) {
			return false
		}
	}
	return true
}

func sumSlice(s slice) int {
	aux := 0
	for _, v := range s {
		aux += v
	}
	return aux
}

func initialSolution(W []int, b int) []slice {
	max := arrayMax(W)

	C := make([]int, 1)
	C[0] = max
	for _, v := range W {
		if arrayContains(W, max - v) {
			C = append(C, v)
		}
	}

	G := make([][]slice, b-1)

	G[0] = make([]slice, 0)
	for _, v := range C {
		if v != max {
			G[0] = append(G[0], []int{v, max})
		}
	}
	
	for i := 1; i <= b-2; i++ {
		G[i] = make([]slice, 0)
		for _, g := range G[i-1] {
			for _, v := range C {
				for _, h := range g {
					if arrayContains(W, absInt(h - v)) {
						G[i] = append(G[i], slice(append(g, v)))
					}
				}
			}
		}  
	}

	S := make([]slice, 0)
	for _, g := range G[len(G) - 1] {
		s := buildSolution(g)
		S = append(S, s)
	}

	return S
}

func extendLeft(Si []slice, W []int, b int) []slice {
	Saux := make([]slice, 0)
	for _, s := range Si {
		for _, w := range W {
			aux := make([]int, b)
			for k := 0; k < b; k++ {
				aux[k] = w + sumSlice(s[0:k])
			}

			if sliceIncluded(aux, W) {
				Saux = append(Saux, append(s, w))
			}
		}
	}
	return Saux
}

func extendRight(Si []slice, W []int, b int) []slice {
	Saux := make([]slice, 0)
	for _, s := range Si {
		for _, w := range W {
			aux := make([]int, b)
			for k := 0; k < b; k++ {
				aux[k] = w + sumSlice(s[len(s) - 1 - k:])
			}

			if sliceIncluded(aux, W) {
				Saux = append(Saux, append(s, w))
			}
		}
	}
	return Saux
}

func attack(W []int, b, N int) []slice {
	Sb := initialSolution(W, b)
	S := make([][]slice, N - b + 1)
	S[0] = Sb

	for i := b + 1; i <= N ; i++ {
		S[i - b] = mergeSliceOfSlices(extendLeft(S[i - b - 1], W, b), extendRight(S[i - b - 1], W, b))
	}

	Sn := S[len(S) - 1]
	// Sn = {s | s ∈ Sn ∧ Lb(Sn) = W}
	return Sn
}