package pir

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"testing"
)

const LOGQ = uint64(32)
const SEC_PARAM = uint64(1 << 10)

// Test that DB packing methods are correct, when each database entry is ~ 1 Z_p elem.
func TestDBMediumEntries(t *testing.T) {
	N := uint64(4)
	d := uint64(9)
	pir := SimplePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	vals := []uint64{1, 2, 3, 4}
	DB := MakeDB(N, d, &p, vals)
	if DB.info.packing != 1 || DB.info.ne != 1 {
		panic("Should not happen.")
	}

	for i := uint64(0); i < N; i++ {
		if DB.GetElem(i) != (i + 1) {
			panic("Failure")
		}
	}
}

// Test that DB packing methods are correct, when multiple database entries fit in 1 Z_p elem.
func TestDBSmallEntries(t *testing.T) {
	N := uint64(4)
	d := uint64(3)
	pir := SimplePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	vals := []uint64{1, 2, 3, 4}
	DB := MakeDB(N, d, &p, vals)
	if DB.info.packing <= 1 || DB.info.ne != 1 {
		panic("Should not happen.")
	}

	for i := uint64(0); i < N; i++ {
		if DB.GetElem(i) != (i + 1) {
			panic("Failure")
		}
	}
}

// Test that DB packing methods are correct, when each database entry requires multiple Z_p elems.
func TestDBLargeEntries(t *testing.T) {
	N := uint64(4)
	d := uint64(12)
	pir := SimplePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	vals := []uint64{1, 2, 3, 4}
	DB := MakeDB(N, d, &p, vals)
	if DB.info.packing != 0 || DB.info.ne <= 1 {
		panic("Should not happen.")
	}

	for i := uint64(0); i < N; i++ {
		if DB.GetElem(i) != (i + 1) {
			panic("Failure")
		}
	}
}

// Print the BW used by SimplePIR
func TestSimplePirBW(t *testing.T) {
	N := uint64(1 << 20)
	d := uint64(2048)

	log_N, _ := strconv.Atoi(os.Getenv("LOG_N"))
	D, _ := strconv.Atoi(os.Getenv("D"))
	if log_N != 0 {
		N = uint64(1 << log_N)
	}
	if D != 0 {
		d = uint64(D)
	}

	pir := SimplePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)
	DB := SetupDB(N, d, &p)

	fmt.Printf("Executing with entries consisting of %d (>= 1) bits; p is %d; packing factor is %d; number of DB elems per entry is %d.\n",
		d, p.p, DB.info.packing, DB.info.ne)

	pir.GetBW(DB.info, p)
}

// Print the BW used by DoublePIR
func TestDoublePirBW(t *testing.T) {
	N := uint64(1 << 20)
	d := uint64(2048)

	log_N, _ := strconv.Atoi(os.Getenv("LOG_N"))
	D, _ := strconv.Atoi(os.Getenv("D"))
	if log_N != 0 {
		N = uint64(1 << log_N)
	}
	if D != 0 {
		d = uint64(D)
	}

	pir := DoublePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)
	DB := SetupDB(N, d, &p)

	fmt.Printf("Executing with entries consisting of %d (>= 1) bits; p is %d; packing factor is %d; number of DB elems per entry is %d.\n",
		d, p.p, DB.info.packing, DB.info.ne)

	pir.GetBW(DB.info, p)
}

// Test SimplePIR correctness on DB with short entries.
func TestSimplePir(t *testing.T) {
	N := uint64(1 << 20)
	d := uint64(8)
	pir := SimplePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	DB := MakeRandomDB(N, d, &p)
	RunPIR(&pir, DB, p, []uint64{262144})
}

// Test SimplePIR correctness on DB with long entries
func TestSimplePirLongRow(t *testing.T) {
	N := uint64(1 << 20)
	d := uint64(32)
	pir := SimplePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	DB := MakeRandomDB(N, d, &p)
	RunPIR(&pir, DB, p, []uint64{1})
}

// Test SimplePIR correctness on big DB
func TestSimplePirBigDB(t *testing.T) {
	N := uint64(1 << 25)
	d := uint64(7)
	pir := SimplePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	DB := MakeRandomDB(N, d, &p)
	RunPIR(&pir, DB, p, []uint64{0})
}

// Test SimplePIR correctness on DB with short entries, and batching.
func TestSimplePirBatch(t *testing.T) {
	N := uint64(1 << 20)
	d := uint64(8)
	pir := SimplePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	DB := MakeRandomDB(N, d, &p)
	RunPIR(&pir, DB, p, []uint64{0, 0, 0, 0})
}

// Test SimplePIR correctness on DB with long entries, and batching.
func TestSimplePirLongRowBatch(t *testing.T) {
	N := uint64(1 << 20)
	d := uint64(32)
	pir := SimplePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	DB := MakeRandomDB(N, d, &p)
	RunPIR(&pir, DB, p, []uint64{0, 0, 0, 0})
}

// Test DoublePIR correctness on DB with short entries.
func TestDoublePir(t *testing.T) {
	N := uint64(1 << 28)
	d := uint64(3)
	pir := DoublePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	DB := MakeRandomDB(N, d, &p)
	RunPIR(&pir, DB, p, []uint64{0})
}

// Test DoublePIR correctness on DB with long entries.
func TestDoublePirLongRow(t *testing.T) {
	N := uint64(1 << 20)
	d := uint64(32)
	pir := DoublePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	DB := MakeRandomDB(N, d, &p)

	fmt.Printf("Executing with entries consisting of %d (>= 1) bits; p is %d; packing factor is %d; number of DB elems per entry is %d.\n",
		d, p.p, DB.info.packing, DB.info.ne)

	RunPIR(&pir, DB, p, []uint64{1 << 19})
}

// Test DoublePIR correctness on big DB
func TestDoublePirBigDB(t *testing.T) {
	N := uint64(1 << 25)
	d := uint64(7)
	pir := DoublePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	DB := MakeRandomDB(N, d, &p)
	RunPIR(&pir, DB, p, []uint64{0})
}

// Test DoublePIR correctness on DB with short entries, and batching.
func TestDoublePirBatch(t *testing.T) {
	N := uint64(1 << 20)
	d := uint64(8)
	pir := DoublePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	DB := MakeRandomDB(N, d, &p)
	RunPIR(&pir, DB, p, []uint64{0, 0, 0, 0})
}

// Test DoublePIR correctness on DB with long entries, and batching.
func TestDoublePirLongRowBatch(t *testing.T) {
	N := uint64(1 << 20)
	d := uint64(32)
	pir := DoublePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	DB := MakeRandomDB(N, d, &p)
	RunPIR(&pir, DB, p, []uint64{0, 0, 0, 0})
}

// Benchmark SimplePIR performance.
func BenchmarkSimplePirSingle(b *testing.B) {
	_, err := os.Create("simple-cpu.out")
	if err != nil {
		panic("Error creating file")
	}

	N := uint64(1 << 20)
	d := uint64(2048)

	log_N, _ := strconv.Atoi(os.Getenv("LOG_N"))
	D, _ := strconv.Atoi(os.Getenv("D"))
	if log_N != 0 {
		N = uint64(1 << log_N)
	}
	if D != 0 {
		d = uint64(D)
	}

	pir := SimplePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	i := uint64(0) // index to query
	if i >= p.l*p.m {
		panic("Index out of dimensions")
	}

	DB := MakeRandomDB(N, d, &p)
	var tputs []float64

	var processing []float64
	var gener []float64
	var hint []float64
	var dec []float64
	var hint_size []float64
	var query_size float64
	var answer_size float64
	var tput float64
	var time_hint int64
	var time_query_generation int64
	var time_query_processing int64
	var time_decoding int64

	for j := 0; j < 4; j++ {
		tput, _, hint_size, query_size, answer_size, time_hint, time_query_generation, time_query_processing, time_decoding = RunPIR(&pir, DB, p, []uint64{i})
		tputs = append(tputs, tput)
		gener = append(gener, float64(time_query_generation))
		processing = append(processing, float64(time_query_processing))
		hint = append(hint, float64(time_hint))
		dec = append(dec, float64(time_decoding))

	}
	// fmt.Printf("Avg SimplePIR tput, except for first run: %f MB/s\n", avg(tputs))
	// fmt.Printf("Std dev of SimplePIR tput, except for first run: %f MB/s\n", stddev(tputs))

	// read the file
	var path = "sample.csv"
	f_, err_ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err_ != nil {
		log.Fatal(err_)
	}
	defer f_.Close()
	var data [][]string

	data = append(data, []string{"Simple", fmt.Sprintf("%f", math.Pow(2, float64(log_N))), fmt.Sprintf("%f", float64(d)/(8*1024.0)), fmt.Sprintf("%f", math.Pow(2, float64(log_N))*float64(d)/(8*1024.0)/1024), fmt.Sprintf("%f", avg(hint)/1000.0), fmt.Sprintf("%f", avg(gener)/1000.0), fmt.Sprintf("%f", avg(processing)/1000.0), fmt.Sprintf("%f", avg(dec)/1000.0), fmt.Sprintf("%f", hint_size[0]/1024), fmt.Sprintf("%f", query_size), fmt.Sprintf("%f", answer_size)})

	w := csv.NewWriter(f_)
	w.WriteAll(data)

	if err_ != nil {
		log.Fatal(err_)
	}

}

// Benchmark DoublePIR performance.
func BenchmarkDoublePirSingle(b *testing.B) {
	_, err := os.Create("double-cpu.out")
	if err != nil {
		panic("Error creating file")
	}

	N := uint64(1 << 20)
	d := uint64(2048)

	log_N, _ := strconv.Atoi(os.Getenv("LOG_N"))
	D, _ := strconv.Atoi(os.Getenv("D"))
	if log_N != 0 {
		N = uint64(1 << log_N)
	}
	if D != 0 {
		d = uint64(D)
	}

	pir := DoublePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	i := uint64(0) // index to query
	if i >= p.l*p.m {
		panic("Index out of dimensions")
	}

	DB := MakeRandomDB(N, d, &p)
	var tputs []float64

	var processing []float64
	var gener []float64
	var hint []float64
	var dec []float64
	var hint_size []float64
	var query_size float64
	var answer_size float64
	var tput float64
	var time_hint int64
	var time_query_generation int64
	var time_query_processing int64
	var time_decoding int64

	for j := 0; j < 4; j++ {
		tput, _, hint_size, query_size, answer_size, time_hint, time_query_generation, time_query_processing, time_decoding = RunPIR(&pir, DB, p, []uint64{i})
		tputs = append(tputs, tput)
		gener = append(gener, float64(time_query_generation))
		processing = append(processing, float64(time_query_processing))
		hint = append(hint, float64(time_hint))
		dec = append(dec, float64(time_decoding))

	}
	// fmt.Printf("Avg SimplePIR tput, except for first run: %f MB/s\n", avg(tputs))
	// fmt.Printf("Std dev of SimplePIR tput, except for first run: %f MB/s\n", stddev(tputs))
	fmt.Printf("Avg Doubl timee query generation : %f micros\n", avg(gener)/1000.0)
	fmt.Printf("Avg Doubl time query processing : %f micros\n", avg(processing)/1000.0)

	// read the file
	var path = "sample.csv"
	f_, err_ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err_ != nil {
		log.Fatal(err_)
	}
	defer f_.Close()
	var data [][]string
	data = append(data, []string{"Double", fmt.Sprintf("%f", math.Pow(2, float64(log_N))), fmt.Sprintf("%f", float64(d)/(8*1024.0)), fmt.Sprintf("%f", math.Pow(2, float64(log_N))*(float64(d)/(8*1024.0))/1024), fmt.Sprintf("%f", avg(hint)/1000.0), fmt.Sprintf("%f", avg(gener)/1000.0), fmt.Sprintf("%f", avg(processing)/1000.0), fmt.Sprintf("%f", avg(dec)/1000.0), fmt.Sprintf("%f", hint_size[1]/(1024)), fmt.Sprintf("%f", hint_size[0]/(1024*1024)), fmt.Sprintf("%f", query_size), fmt.Sprintf("%f", answer_size)})

	w := csv.NewWriter(f_)
	w.WriteAll(data)

	if err_ != nil {
		log.Fatal(err_)
	}

}

// Benchmark SimplePIR performance, on 1GB databases with increasing row length.
func BenchmarkSimplePirVaryingDB(b *testing.B) {
	flog, err := os.OpenFile("simple-comm.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Error creating log file")
	}
	defer flog.Close()

	writer := csv.NewWriter(flog)
	defer writer.Flush()

	records := []string{"N", "d", "tput", "tput_stddev", "offline_comm", "online_comm"}
	writer.Write(records)

	pir := SimplePIR{}

	// Set N, D
	total_sz := 33

	for d := uint64(1); d <= 32768; d *= 2 {
		N := uint64(1<<total_sz) / d
		p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

		i := uint64(0) // index to query
		if i >= p.l*p.m {
			panic("Index out of dimensions")
		}

		DB := MakeRandomDB(N, d, &p)
		var tputs []float64
		var offline_cs []float64
		var online_cs []float64

		var processing []float64
		var gener []float64

		for j := 0; j < 5; j++ {
			tput, _, _, _, time_query_generation, time_query_processing := RunFakePIR(&pir, DB, p, []uint64{i}, nil, false)
			tputs = append(tputs, tput)
			gener = append(gener, float64(time_query_generation))
			processing = append(processing, float64(time_query_processing))
		}
		// fmt.Printf("Avg SimplePIR tput, except for first run: %f MB/s\n", avg(tputs))
		// fmt.Printf("Std dev of SimplePIR tput, except for first run: %f MB/s\n", stddev(tputs))
		fmt.Printf("Avg SimplePIR time query generation nanoseconds: %f s\n", avg(gener))
		fmt.Printf("Avg SimplePIR time query processing nanoseconds: %f s\n", avg(processing))

		if (stddev(offline_cs) != 0) || (stddev(online_cs) != 0) {
			fmt.Printf("%f %f SHOULD NOT HAPPEN\n", stddev(offline_cs), stddev(online_cs))
			//panic("Should not happen!")
		}
		writer.Write([]string{strconv.FormatUint(N, 10),
			strconv.FormatUint(d, 10),
			strconv.FormatFloat(avg(tputs), 'f', 4, 64),
			strconv.FormatFloat(stddev(tputs), 'f', 4, 64),
			strconv.FormatFloat(avg(offline_cs), 'f', 4, 64),
			strconv.FormatFloat(avg(online_cs), 'f', 4, 64)})
	}
}

// Benchmark DoublePIR performance, on 1 GB databases with increasing row length.
func BenchmarkDoublePirVaryingDB(b *testing.B) {
	flog, err := os.OpenFile("double-comm.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Error creating log file")
	}
	defer flog.Close()

	writer := csv.NewWriter(flog)
	defer writer.Flush()

	records := []string{"N", "d", "tput", "tput_stddev", "offline_comm", "online_comm"}
	writer.Write(records)

	pir := DoublePIR{}

	// Set N, D
	total_sz := 33
	for d := uint64(1); d <= 32768; d *= 2 {
		N := uint64(1<<total_sz) / d
		p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

		i := uint64(0) // index to query
		if i >= p.l*p.m {
			panic("Index out of dimensions")
		}

		DB := MakeRandomDB(N, d, &p)
		var tputs []float64
		var offline_cs []float64
		var online_cs []float64

		var processing []float64
		var gener []float64

		for j := 0; j < 5; j++ {
			tput, _, _, _, time_query_generation, time_query_processing := RunFakePIR(&pir, DB, p, []uint64{i}, nil, false)
			tputs = append(tputs, tput)
			gener = append(gener, float64(time_query_generation))
			processing = append(processing, float64(time_query_processing))
		}
		// fmt.Printf("Avg SimplePIR tput, except for first run: %f MB/s\n", avg(tputs))
		// fmt.Printf("Std dev of SimplePIR tput, except for first run: %f MB/s\n", stddev(tputs))
		fmt.Printf("Avg SimplePIR time query generation nanoseconds: %f s\n", avg(gener))
		fmt.Printf("Avg SimplePIR time query processing nanoseconds: %f s\n", avg(processing))

		if (stddev(offline_cs) != 0) || (stddev(online_cs) != 0) {
			fmt.Printf("%f %f SHOULD NOT HAPPEN\n", stddev(offline_cs), stddev(online_cs))
			//panic("Should not happen!")
		}
		writer.Write([]string{strconv.FormatUint(N, 10),
			strconv.FormatUint(d, 10),
			strconv.FormatFloat(avg(tputs), 'f', 4, 64),
			strconv.FormatFloat(stddev(tputs), 'f', 4, 64),
			strconv.FormatFloat(avg(offline_cs), 'f', 4, 64),
			strconv.FormatFloat(avg(online_cs), 'f', 4, 64)})
	}
}

// Benchmark SimplePIR performance with batches of increasing size.
func BenchmarkSimplePirBatchLarge(b *testing.B) {
	f, err := os.Create("simple-cpu-batch.out")
	if err != nil {
		panic("Error creating file")
	}

	flog, err := os.Create("simple-batch.log")
	if err != nil {
		panic("Error creating log file")
	}
	defer flog.Close()

	N := uint64(1 << 33)
	d := uint64(1)

	log_N, _ := strconv.Atoi(os.Getenv("LOG_N"))
	D, _ := strconv.Atoi(os.Getenv("D"))
	if log_N != 0 {
		N = uint64(1 << log_N)
	}
	if D != 0 {
		d = uint64(D)
	}

	pir := SimplePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	i := uint64(0) // index to query
	if i >= p.l*p.m {
		panic("Index out of dimensions")
	}

	DB := MakeRandomDB(N, d, &p)

	writer := csv.NewWriter(flog)
	defer writer.Flush()

	records := []string{"Batch_sz", "Good_tput", "Good_std_dev", "Num_successful_queries", "Tput"}
	writer.Write(records)

	for trial := 0; trial <= 10; trial += 1 {
		batch_sz := (1 << trial)
		var query []uint64
		for j := 0; j < batch_sz; j++ {
			query = append(query, i)
		}
		var tputs []float64
		for iter := 0; iter < 5; iter++ {
			tput, _, _, _, _, _ := RunFakePIR(&pir, DB, p, query, f, false)
			tputs = append(tputs, tput)
		}

		expected_num_empty_buckets := math.Pow(float64(batch_sz-1)/float64(batch_sz), float64(batch_sz)) * float64(batch_sz)
		expected_num_successful_queries := float64(batch_sz) - expected_num_empty_buckets
		good_tput := avg(tputs) / float64(batch_sz) * expected_num_successful_queries
		dev := stddev(tputs) / float64(batch_sz) * expected_num_successful_queries

		writer.Write([]string{strconv.Itoa(batch_sz),
			strconv.FormatFloat(good_tput, 'f', 4, 64),
			strconv.FormatFloat(dev, 'f', 4, 64),
			strconv.FormatFloat(expected_num_successful_queries, 'f', 4, 64),
			strconv.FormatFloat(avg(tputs), 'f', 4, 64)})
	}
}

// Benchmark DoublePIR performance with batches of increasing size.
func BenchmarkDoublePirBatchLarge(b *testing.B) {
	f, err := os.Create("double-cpu-batch.out")
	if err != nil {
		panic("Error creating file")
	}

	flog, err := os.Create("double-batch.log")
	if err != nil {
		panic("Error creating log file")
	}
	defer flog.Close()

	N := uint64(1 << 33)
	d := uint64(1)

	log_N, _ := strconv.Atoi(os.Getenv("LOG_N"))
	D, _ := strconv.Atoi(os.Getenv("D"))
	if log_N != 0 {
		N = uint64(1 << log_N)
	}
	if D != 0 {
		d = uint64(D)
	}

	pir := DoublePIR{}
	p := pir.PickParams(N, d, SEC_PARAM, LOGQ)

	i := uint64(0) // index to query
	if i >= p.l*p.m {
		panic("Index out of dimensions")
	}

	DB := MakeRandomDB(N, d, &p)

	writer := csv.NewWriter(flog)
	defer writer.Flush()

	records := []string{"Batch_sz", "Good_tput", "Good_std_dev", "Num_successful_queries", "Tput"}
	writer.Write(records)

	for trial := 0; trial <= 10; trial += 1 {
		batch_sz := (1 << trial)
		var query []uint64
		for j := 0; j < batch_sz; j++ {
			query = append(query, i)
		}
		var tputs []float64
		for iter := 0; iter < 5; iter++ {
			tput, _, _, _, _, _ := RunFakePIR(&pir, DB, p, query, f, false)
			tputs = append(tputs, tput)
		}
		expected_num_empty_buckets := math.Pow(float64(batch_sz-1)/float64(batch_sz), float64(batch_sz)) * float64(batch_sz)
		expected_num_successful_queries := float64(batch_sz) - expected_num_empty_buckets
		good_tput := avg(tputs) / float64(batch_sz) * expected_num_successful_queries
		dev := stddev(tputs) / float64(batch_sz) * expected_num_successful_queries

		writer.Write([]string{strconv.Itoa(batch_sz),
			strconv.FormatFloat(good_tput, 'f', 4, 64),
			strconv.FormatFloat(dev, 'f', 4, 64),
			strconv.FormatFloat(expected_num_successful_queries, 'f', 4, 64),
			strconv.FormatFloat(avg(tputs), 'f', 4, 64)})
	}
}
