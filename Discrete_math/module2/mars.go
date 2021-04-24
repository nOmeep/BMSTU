package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	//"github.com/skorobogatov/input"
)

type Member struct { // удобнее сортировать и делать проверки людей имея вспомогательную структуру
	number int64
	undesirable []int64
}

type Group struct {
	first []bool
	second []bool
	firstCount int64
	secondCount int64
}

func getScanned(crew *[]Member, nel int64) {
	for i := 0; int64(i) < nel; i++ { 
		var (
			curMember Member = setMember(int64(i)) // текущий рассматриваемый член экспедиции 
			preferences []byte = make([]byte, 0) // предпочтения каждого члена экипажа
		)

		for j := 0; int64(j) < nel * 2; j++ { // считываем предпочтения
			var (
				curSymb byte // + или - 
			)
			//fmt.Printf("IM HERE")
			fmt.Scanf("%c", &curSymb);

			if j % 2 == 0 { // иначе он считывает энтеры и пробелы 
				preferences = append(preferences, curSymb) // добавляем символ в массив предпочтений
			}
		}

		for j := 0; int64(j) < nel; j++ {
			if strings.Compare(string(preferences[j]), "+") == 0 {
				curMember.undesirable = append(curMember.undesirable, int64(j)); // обновляем список нежелательных "тиммейтов" человека
			}
		}

		*crew = append(*crew, curMember); // добавляем члена экипажа после того, как все с ним сделали и проверили
	}
}

func setMember(index int64) Member {
	var tmp Member
	tmp.number = index
	tmp.undesirable = make([]int64, 0)
	return tmp;
} 

func setGroup(nel int64) Group {
	var tmp Group
	tmp.first = make([]bool, nel)
	tmp.second = make([]bool, nel)
	return tmp
}

func sortComb( b *[]Group) {
	sort.Slice(*b, func(i, j int) bool {
		if math.Abs(float64((*b)[i].firstCount - (*b)[i].secondCount)) == math.Abs(float64((*b)[j].firstCount - (*b)[j].secondCount)) {
			if (*b)[i].firstCount == (*b)[j].firstCount {
				var (
					i1 int64
					j1 int64
				)
				for k := 0; k < len((*b)[i].first); k++ {
					if (*b)[i].first[k] {
						i1 = i1 * int64(10) + int64(k)
					}
					if (*b)[j].first[k] {
						j1 = j1 * int64(10) + int64(k)
					}
				}
				return i1 < j1
			} else {
				return (*b)[i].firstCount < (*b)[j].firstCount
			}
		} else {
			return math.Abs(float64(int64((*b)[i].firstCount - (*b)[i].secondCount))) < math.Abs(float64(int64((*b)[j].firstCount - (*b)[j].secondCount)))
		}
	})
}

func sortMemb(current int64, fCount *int64, sCount *int64, isFirst bool, f *[]bool, s *[]bool, using *[]bool, members []Member) (ok bool) {
	if isFirst == false { // меняем второй

		if (*using)[current] {
			return !(*f)[current]
		}
		if !(*s)[current] {
			*sCount++
		}
		(*s)[current] = true

	} else if isFirst == true { // меняем первый 

		if (*using)[current] {
			return !(*s)[current]
		}
		if !(*f)[current] {
			*fCount++
		}
		(*f)[current] = true

	}

	(*using)[current] = true
	for i := 0; i < len(members[current].undesirable); i++ {
		if !sortMemb(members[current].undesirable[i], fCount, sCount, !isFirst, f, s, using, members) {
			return false
		}
	}

	return true
}

func main() {
	var (
		nel int64
	)

	fmt.Scanf("%d", &nel);

	var (
		crew []Member = make([]Member, 0) // массив с членами экипажа
	)

	getScanned(&crew, nel)

	// с заполением вроде все ОК
	
	// Делим людей по группам
	var (
		using []bool = make([]bool, nel);
		grPr []Group = make([]Group, 0)
		flag bool = true
	)	

	for i := 0; int64(i) < nel; i++ {
		if !using[i] {
			var tmp Group
			tmp = setGroup(nel)
			if flag {
				flag = sortMemb(int64(i), &tmp.firstCount, &tmp.secondCount, true, &tmp.first, &tmp.second, &using, crew)
			} else {
				flag = flag && sortMemb(int64(i), &tmp.firstCount, &tmp.secondCount, true, &tmp.first, &tmp.second, &using, crew)
			}
			grPr = append(grPr, tmp)
		}
	}

	// следующий этап 

	var (
		destribution []Group = make([]Group, 1 << uint64(len(grPr)))
	)

	for i := 0; i < len(destribution); i++ {

		destribution[i].first = make([]bool, nel)
		destribution[i].second = make([]bool, nel)

		for j := 0; j < len(grPr); j++ {

			if uint64(i) & (1 << uint64(j)) != 0 {

				for k := 0; int64(k) < nel; k++ {
					if grPr[j].first[k] {

						destribution[i].first[k] = true
						destribution[i].firstCount++

					}
					if grPr[j].second[k] {

						destribution[i].second[k] = true
						destribution[i].secondCount++
						
					}
				}
			} else {

				for k := 0; int64(k) < nel; k++ {

					if grPr[j].second[k] {

						destribution[i].first[k] = true
						destribution[i].firstCount++

					}
					if grPr[j].first[k] {

						destribution[i].second[k] = true
						destribution[i].secondCount++

					}
				}
			}
		}
	}

	if !flag {
		fmt.Printf("No solution")
		os.Exit(1);
	}

	// Сортируем комбинации группировок
	sortComb(&destribution)

	for i := 0; int64(i) < nel; i++ {
		if destribution[0].first[i] == true {
			fmt.Printf("%d ", i + 1)
		}
	}

	fmt.Println()
}