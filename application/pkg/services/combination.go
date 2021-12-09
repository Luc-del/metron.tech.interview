package services

//
//func makeCombinations(items ...int) <-chan combination {
//	c := make(chan combination)
//	defer close(c)
//
//	sets := makeUniqueSets(items)
//
//	for minLength := 1; minLength <len(items); minLength++ {
//		//TODO
//	}
//
//	return c
//}
//
//
//func rCombinations(p int, n []int, c []int, ccc [][][]int) [][][]int {
//	if len(n) == 0 || p <= 0 {
//		return ccc
//	}
//	if len(ccc) == 0 {
//		ccc = make([][][]int, p)
//	}
//	p--
//	for i := range n {
//		cc := make([]int, len(c)+1)
//		copy(cc, c)
//		cc[len(cc)-1] = n[i]
//		ccc[len(cc)-1] = append(ccc[len(cc)-1], cc)
//		ccc = rCombinations(p, n[i+1:], cc, ccc)
//	}
//	return ccc
//}
//
//func makeUniqueSets(n []int) [][][]int {
//	sets := rCombinations(len(n), n, nil, nil)
//	for i, sizedSets := range sets {
//		for j, set := range sizedSets {
//			if !checkUniqueness(set) {
//				sets[i] = append(sets[i][:j-1], sets[i][j+1:]...)
//			}
//		}
//	}
//
//	return sets
//}
//
//func checkUniqueness(set []int) bool {
//	unique := make(map[int]struct{})
//	for _, x := range set {
//		unique[x] = struct{}{}
//	}
//
//	return len(unique) == len(set)
//}
//
