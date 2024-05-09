module github.com/hll-truco/hll-truco

go 1.18

require (
	github.com/axiomhq/hyperloglog v0.0.0-20240124082744-24bca3a5b39b
	github.com/truquito/truco v0.1.0
)

require github.com/dgryski/go-metro v0.0.0-20180109044635-280f6062b5bc // indirect

require (
	github.com/filevich/canvas v0.0.0 // indirect
	github.com/filevich/combinatronics v0.0.0-20220316214652-26aa6db09482
	github.com/filevich/truco-cfr v0.0.0
)

replace github.com/truquito/truco => ../minitruco

// replace github.com/truquito/truco => ../truco

replace github.com/filevich/truco-cfr => ../truco-cfr
