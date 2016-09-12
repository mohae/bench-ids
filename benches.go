package main

import (
	"testing"

	dockeruuid "github.com/docker/distribution/uuid"
	googleuuid "github.com/google/uuid"
	hashicorpgouuid "github.com/hashicorp/go-uuid"
	kasworldid "github.com/kasworld/idgen"
	mitchellhuuid "github.com/mitchellh/packer/common/uuid"
	"github.com/mohae/benchutil"
	"github.com/mohae/snoflinga"
	"github.com/nats-io/nuid"
	nu7hatchuuid "github.com/nu7hatch/gouuid"
	rogpeppefastuuid "github.com/rogpeppe/fastuuid"
	"github.com/rs/xid"
	satorigouuid "github.com/satori/go.uuid"
	sdminggosnow "github.com/sdming/gosnow"
	"github.com/sony/sonyflake"
	sumoryid "github.com/sumory/idgen"
)

func rsxID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xid.New()
	}
}

func BenchRSXID() benchutil.Bench {
	bench := benchutil.NewBench("github.com/rs/xid")
	bench.Group = "pseudo-snowflake"
	bench.SubGroup = "96 bits"
	bench.Desc = "base32 hex encoded"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(rsxID))
	return bench
}

func kasworldIDGenUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kasworldid.NewUUID()
	}
}

func BenchKasworldIDGenUUID() benchutil.Bench {
	bench := benchutil.NewBench("github.com/kasworld/idgen")
	bench.Group = "pseudo-uuid"
	bench.SubGroup = "128-bit"
	bench.Desc = "crypto/rand"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(kasworldIDGenUUID))
	return bench
}

func sumoryIDGen(b *testing.B) {
	err, w := sumoryid.NewIdWorker(0)
	if err != nil {
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.NextId()
	}
}

func BenchSumoryIDGen() benchutil.Bench {
	bench := benchutil.NewBench("github.com/sumory/idgen")
	bench.Group = "snowflake"
	bench.SubGroup = "64-bit"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(sumoryIDGen))
	return bench
}

var uuid = [16]byte{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}
var name = []byte("name")

func nu7hatchGoUUIDv3(b *testing.B) {
	id := nu7hatchuuid.UUID(uuid)
	for i := 0; i < b.N; i++ {
		nu7hatchuuid.NewV3(&id, name)
	}
}

func BenchNu7hatchGoUUIDv3() benchutil.Bench {
	bench := benchutil.NewBench("github.com/nu7hatch/gouuid")
	bench.Group = "uuid v3"
	bench.SubGroup = "128-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(nu7hatchGoUUIDv3))
	return bench
}

func nu7hatchGoUUIDv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nu7hatchuuid.NewV4()
	}
}

func BenchNu7hatchGoUUIDv4() benchutil.Bench {
	bench := benchutil.NewBench("github.com/nu7hatch/gouuid")
	bench.Group = "uuid v4"
	bench.SubGroup = "128-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(nu7hatchGoUUIDv4))
	return bench
}

func nu7hatchGoUUIDv5(b *testing.B) {
	id := nu7hatchuuid.UUID(uuid)
	for i := 0; i < b.N; i++ {
		nu7hatchuuid.NewV5(&id, name)
	}
}

func BenchNu7hatchGoUUIDv5() benchutil.Bench {
	bench := benchutil.NewBench("github.com/nu7hatch/gouuid")
	bench.Group = "uuid v5"
	bench.SubGroup = "128-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(nu7hatchGoUUIDv5))
	return bench
}

func dockerUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dockeruuid.Generate()
	}
}

func BenchDockerUUID() benchutil.Bench {
	bench := benchutil.NewBench("github.com/docker/distribution/uuid")
	bench.Group = "uuid v4"
	bench.SubGroup = "128-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(dockerUUID))
	return bench
}

func MitchellhUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mitchellhuuid.TimeOrderedUUID()
	}
}

func BenchMitchellhUUID() benchutil.Bench {
	bench := benchutil.NewBench("github.com/mitchellh/packer/common/uuid")
	bench.Group = "psuedo-uuid"
	bench.SubGroup = "128-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(MitchellhUUID))
	return bench
}

func SatoriGoUUIDv1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		satorigouuid.NewV1()
	}
}

func BenchSatoriGoUUIDv1() benchutil.Bench {
	bench := benchutil.NewBench("github.com/satori/go.uuid")
	bench.Group = "uuid v1"
	bench.SubGroup = "128-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(SatoriGoUUIDv1))
	return bench
}

func SatoriGoUUIDv2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		satorigouuid.NewV2(0)
	}
}

func BenchSatoriGoUUIDv2() benchutil.Bench {
	bench := benchutil.NewBench("github.com/satori/go.uuid")
	bench.Group = "uuid v2"
	bench.SubGroup = "128-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(SatoriGoUUIDv2))
	return bench
}

func SatoriGoUUIDv3(b *testing.B) {
	s := string(name)
	for i := 0; i < b.N; i++ {
		satorigouuid.NewV3(uuid, s)
	}
}

func BenchSatoriGoUUIDv3() benchutil.Bench {
	bench := benchutil.NewBench("github.com/satori/go.uuid")
	bench.Group = "uuid v3"
	bench.SubGroup = "128-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(SatoriGoUUIDv3))
	return bench
}

func SatoriGoUUIDv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		satorigouuid.NewV4()
	}
}

func BenchSatoriGoUUIDv4() benchutil.Bench {
	bench := benchutil.NewBench("github.com/satori/go.uuid")
	bench.Group = "uuid v4"
	bench.SubGroup = "128-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(SatoriGoUUIDv4))
	return bench
}

func SatoriGoUUIDv5(b *testing.B) {
	s := string(name)
	for i := 0; i < b.N; i++ {
		satorigouuid.NewV5(uuid, s)
	}
}

func BenchSatoriGoUUIDv5() benchutil.Bench {
	bench := benchutil.NewBench("github.com/satori/go.uuid")
	bench.Group = "uuid v5"
	bench.SubGroup = "128-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(SatoriGoUUIDv5))
	return bench
}

func RogPeppeFastUUID(b *testing.B) {
	g, err := rogpeppefastuuid.NewGenerator()
	if err != nil {
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Next()
	}
}

func BenchRogPeppeFastUUID() benchutil.Bench {
	bench := benchutil.NewBench("github.com/rogpeppe/fastuuid")
	bench.Group = "psuedo-uuid"
	bench.SubGroup = "192-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(RogPeppeFastUUID))
	return bench
}

func hashicorpGoUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hashicorpgouuid.GenerateUUID()
	}
}

func BenchHashicorpGoUUID() benchutil.Bench {
	bench := benchutil.NewBench("github.com/hashicorp/go-uuid")
	bench.Group = "psuedo-uuid"
	bench.SubGroup = "128-bit"
	bench.Desc = "formatted as a uuid string"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(hashicorpGoUUID))
	return bench
}

func sonySonyFlake(b *testing.B) {
	var st sonyflake.Settings
	st.MachineID = func() (uint16, error) {
		return 42, nil
	}
	s := sonyflake.NewSonyflake(st)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.NextID()
	}
}

func BenchSonySonyFlake() benchutil.Bench {
	bench := benchutil.NewBench("github.com/sony/sonyflake")
	bench.Group = "psuedo-snowflake"
	bench.SubGroup = "63-bit"
	bench.Desc = "thread-safe"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(sonySonyFlake))
	return bench
}

func sdmingGosnow(b *testing.B) {
	s, err := sdminggosnow.NewSnowFlake(42)
	if err != nil {
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Next()
	}
}

func BenchSdmingGoSnow() benchutil.Bench {
	bench := benchutil.NewBench("github.com/0x6e6562/gosnow")
	bench.Group = "snowflake"
	bench.SubGroup = "64-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(sdmingGosnow))
	return bench
}

func googleUUIDv1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		googleuuid.NewUUID()
	}
}

func BenchGoogleUUIDv1() benchutil.Bench {
	bench := benchutil.NewBench("github.com/google/uuid")
	bench.Group = "uuid v1"
	bench.SubGroup = "128-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(googleUUIDv1))
	return bench
}

func googleUUIDv2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		googleuuid.NewDCEPerson()
	}
}

func BenchGoogleUUIDv2() benchutil.Bench {
	bench := benchutil.NewBench("github.com/google/uuid")
	bench.Group = "uuid v2"
	bench.SubGroup = "128-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(googleUUIDv2))
	return bench
}

func googleUUIDv3(b *testing.B) {
	bt := []byte(name)
	for i := 0; i < b.N; i++ {
		googleuuid.NewMD5(uuid, bt)
	}
}

func BenchGoogleUUIDv3() benchutil.Bench {
	bench := benchutil.NewBench("github.com/google/uuid")
	bench.Group = "uuid v3"
	bench.SubGroup = "128-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(googleUUIDv3))
	return bench
}

func googleUUIDv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		googleuuid.NewRandom()
	}
}

func BenchGoogleUUIDv4() benchutil.Bench {
	bench := benchutil.NewBench("github.com/google/uuid")
	bench.Group = "uuid v4"
	bench.SubGroup = "128-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(googleUUIDv4))
	return bench
}

func googleUUIDv5(b *testing.B) {
	bt := []byte(name)
	for i := 0; i < b.N; i++ {
		googleuuid.NewSHA1(uuid, bt)
	}
}

func BenchGoogleUUIDv5() benchutil.Bench {
	bench := benchutil.NewBench("github.com/google/uuid")
	bench.Group = "uuid v5"
	bench.SubGroup = "128-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(googleUUIDv5))
	return bench
}

func mohaeSnoflinga(b *testing.B) {
	gen := snoflinga.New([]byte("test"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gen.Snowflake()
	}
}

func BenchMohaeSnoflinga() benchutil.Bench {
	bench := benchutil.NewBench("github.com/mohae/snoflinga")
	bench.Group = "psuedo-snowflake"
	bench.SubGroup = "128-bit"
	bench.Desc = "thread-safe"
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(mohaeSnoflinga))
	return bench
}

func natsioNUID(b *testing.B) {
	n := nuid.New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n.Next()
	}
}

func BenchNatsIONUID() benchutil.Bench {
	bench := benchutil.NewBench("github.com/nats-io/nuid")
	bench.Group = "psuedo-snowflake"
	bench.SubGroup = "176-bit"
	bench.Desc = ""
	bench.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(natsioNUID))
	return bench
}
