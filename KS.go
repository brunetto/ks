package main 

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
// 	"reflect"
)

func main() {
	tGlob0 := time.Now()
	
	var distro1file, distro2file string
	var distro1, distro2 Distro
	var d, prob float64
	
	
	if len(os.Args) < 3 {
		log.Fatal("Please, you need to provide two files to be compared...")
	}
	
	distro1file = os.Args[1]
	distro2file = os.Args[2]
	
	distro1.Populate(distro1file)
	distro2.Populate(distro2file)
	
	fmt.Println(distro2)
	
	d, prob = KSTest(&distro1, &distro2)

	fmt.Println(d, prob)
	
	tGlob1 := time.Now()
	fmt.Println()
	log.Println("Wall time for all ", tGlob1.Sub(tGlob0))
}

type Distro struct {
	Values []float64
	
}

func (distro *Distro) Populate (inFile string) () {
	var fileObj *os.File
	var nReader *bufio.Reader
	var readLine string
	var err error
	var nLines, voids int
// 	var reg = regexp.MustCompile(`(\d\.{0,1}\d*)e{0,1}(\D{0,1}\d*)`)
	var reg = regexp.MustCompile(`([+-]*\d*\.*\d*e*[+-]\d*)`) // should match any number
	var res []string
	var num float64
	var exp int64
	var temp float64
	
	log.Println("Opening data file...")
	if fileObj, err = os.Open(inFile); err != nil {
		log.Fatal("Can't open %s: error: %s\n", inFile, err)
	}
	defer fileObj.Close()
	
	nLines, voids = LinesCount(fileObj)
	distro.Values = make([]float64, nLines-voids)
	
// 	log.Println("Made slice with ", nLines, " elements")
	
	nReader = bufio.NewReader(fileObj)
	line := 0
	for {
		if readLine, err = Readln(nReader); err != nil {
			log.Println("Done reading ", line, " lines from file " , inFile, " with ", nLines, " lines with err ", err)
			break
		}
		if len(readLine) == 0 {
			break
		}
		if res = reg.FindStringSubmatch(readLine); len(res) == 0 {
			log.Fatal("Regexp found nothing, problem reading file.")
		} else if len(res) == 3 {
			num, _ = strconv.ParseFloat(res[1], 64)
			exp, _ = strconv.ParseInt(res[2], 10, 64)
			temp = float64(num * math.Pow10(int(exp)))
		} else if len(res) == 2 {
			temp, _ = strconv.ParseFloat(res[1], 64)
		} else {
			log.Fatal("Wrong res ", res)
		}
// 		num, _ = strconv.ParseFloat(res[1], 64)
// 		exp, _ = strconv.ParseInt(res[1], 10, 64)
// 		temp = int64(num * math.Pow10(int(exp)))
		distro.Values[line] = temp
		fmt.Println("temp ", distro.Values[line])
		line++
	}
	sort.Sort(float64arr(distro.Values))
}

func KSTest(distro1, distro2 *Distro) (d, prob float64) {
	var en1, en2, en, dt, fn1,fn2 float64
	var j1, j2 int
	d = 0
	prob = 0
	en1 = float64(len(distro1.Values))
	en2 = float64(len(distro2.Values))
	j1 = 1
	j2 = 1
	fn1 = 0
	fn2 = 0
	
	log.Println("KS test")
	log.Println("Dimensions ", len(distro1.Values), len(distro2.Values))
	for {
		if (j1 > len(distro1.Values)-1 || j2 > len(distro2.Values)-1) {
			break
		}
		fmt.Println(j1, distro1.Values[j1], j2, distro2.Values[j2])
		time.Sleep(1 * time.Millisecond)
		if distro1.Values[j1] <= distro2.Values[j2] {
			fmt.Println("Case 1")
			j1++
			fn1 = float64(j1)/en1
		} else if distro1.Values[j1] >= distro2.Values[j2] {
			fmt.Println("Case 2")
			j2++
			fn2 = float64(j2)/en2
		}
		
		dt = math.Abs(fn2-fn1)
		
		if dt > d {
			fmt.Println("Case 3")
			d = dt
		}
// 		os.Exit(0)
	}
	log.Println("Done KS start prob")
	
	en = math.Sqrt(en1*en2 / (en1+en2))
	prob = ProbKs((en+0.12+0.11/en)*d)
	
	return d, prob
}

func ProbKs (alam float64) (float64) {
	const EPS1 = 0.001
	const EPS2 = 1.0e-8
	var a2, fac, sum, term, termbf float64
	
	log.Println("Calculating probability")
	
	fac = 2.0
	sum = 0.0
	termbf = 0.0
	
	a2 = -2.0*alam*alam
	for j:=1; j<=100;j++ {
		term=fac*math.Exp(a2*float64(j*j))
		sum += term
		if ((math.Abs(term) <= EPS1*termbf) || (math.Abs(term) <= EPS2*sum)) {
			return sum
		}
		fac = -fac
		termbf = math.Abs(term)
	}
	return 1.0
}


func LinesCount(fileObj *os.File) (totCount int, voidCount int) {
	var ( 
		nReader *bufio.Reader
		readLine string
		err error
	)
	
	log.Println("Counting lines")
	
	nReader = bufio.NewReader(fileObj)
	
	totCount, voidCount = 0, 0
	for {
		if readLine, err = Readln(nReader); err != nil {break}
		if len(readLine) == 0 {voidCount++}
		totCount++
	}
	
	// 	log.Println("Rewind file")
	if _, err := fileObj.Seek(0,0); err != nil {
		log.Fatal("Failed seek of file with error ", err)
	}
// 	log.Println("Done")
	return totCount, voidCount
}

func Readln(r *bufio.Reader) (string, error) {
	/*from http://stackoverflow.com/questions/6141604/go-readline-string*/
  var (isPrefix bool = true
       err error = nil
       line, ln []byte
      )
  for isPrefix && err == nil {
      line, isPrefix, err = r.ReadLine()
      ln = append(ln, line...)
  }
  return string(ln),err
}

type float64arr []float64
func (a float64arr) Len() int { return len(a) }
func (a float64arr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a float64arr) Less(i, j int) bool { return a[i] < a[j] }














