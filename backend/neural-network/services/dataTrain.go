package services

import (
	"container/list"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type DataTrainService struct {
	classes map[string]float64
}

const separatorSentence = "#"
const separatorClassification = "("

func (d *DataTrainService) GetData(path string) string {
	resp, err := http.Get(path)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	bodyString = strings.Replace(bodyString, "\u0000", " ", -1)
	return bodyString
}

func (d *DataTrainService) SeparateData(data string) (map[string][]string, map[string]float64) {
	sentences := strings.Split(data, separatorSentence)
	classifiedSentences := make(map[string][]string, len(sentences))

	d.classes = make(map[string]float64)

	re := regexp.MustCompile(`\(([^)]+)\)`)
	for i := 0; i < len(sentences); i++ {
		text := sentences[i]
		if len(text) == 0 {
			continue
		}

		sentence := strings.TrimSpace(strings.Split(text, separatorClassification)[0])
		classesFind := string(re.Find([]byte(text)))
		classesText := classesFind[1 : len(classesFind)-1]
		classesArray := strings.Split(classesText, ",")

		d.insertClasses(classesArray)

		classifiedSentences[sentence] = classesArray
	}

	ClassesQuantity = len(d.classes)
	return classifiedSentences, d.classes
}

func (d *DataTrainService) insertClasses(words []string) {
	for _, word := range words {
		d.insertClass(word)
	}
}

func (d *DataTrainService) insertClass(word string) {
	_, ok := d.classes[word]
	if ok {
		return
	}

	classes := make(map[string]float64, len(d.classes)+1)

	for w, id := range d.classes {
		classes[w] = id
	}

	classes[word] = float64(len(d.classes))

	d.classes = classes
}

func (d *DataTrainService) GetOrdenedArrayClasses(classes []string) []string {
	// ordenedClasses := make([]string, len(d.classes))
	// for _, class := range classes {
	// 	ordenedClasses[d.classes[class]] = class
	// }
	// return ordenedClasses
	return nil
}

func (d *DataTrainService) GetWordsArray(mapSentences map[string][]string) *list.List {
	listWords := list.New()
	sentences := getMapIndexers(mapSentences)
	for i := 0; i < len(sentences); i++ {
		words := strings.Split(sentences[i], " ")
		for j := 0; j < len(words); j++ {
			listWords.PushBack(words[j])
		}
	}
	return listWords
}

func (d *DataTrainService) GetClassIdentity(class string) float64 {
	identity, ok := d.classes[class]
	if ok {
		return -1
	}
	return identity
}

func (d *DataTrainService) GetClassVector(classes []string) []float64 {
	indexClasses := make(map[int]int, len(classes))
	for i, class := range classes {
		indexClasses[i] = int(d.classes[class])
	}

	classesVector := make([]float64, len(d.classes))
	for _, value := range indexClasses {
		classesVector[value] = 1
	}

	return classesVector
}

// func (d *DataTrainService) GetClassVector(classes []string) []float64 {
// 	// indexClasses := make([]int, len(classes))
// 	classesVector := make([]float64, len(d.classes))
// 	for i, classText := range classes {
// 		for classKey, value := range d.classes {
// 			if classKey == classText {
// 				classesVector[i] = value
// 				break
// 			}
// 		}
// 		// indexClasses[i] = int(d.classes[class])
// 	}

// 	return classesVector
// }

func (d *DataTrainService) GetClassesString(classes []float64) []string {
	classesText := make([]string, len(classes))

	// index := 0
	for class, i := range d.classes {
		if classes[int(i)] > 0 {
			classesText[int(i)] = class
			// break
		}
	}

	return classesText
}

func getMapIndexers(m map[string][]string) []string {
	keys := make([]string, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func GetMapValues(m map[string]float64) []float64 {
	values := make([]float64, len(m))

	i := 0
	for k := range m {
		values[i] = m[k]
		i++
	}

	return values
}
