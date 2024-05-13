module github.com/hll-truco/hll-truco

go 1.21.5

require (
	github.com/axiomhq/hyperloglog v0.0.0-20240124082744-24bca3a5b39b
	github.com/clarkduvall/hyperloglog v0.0.0-20171127014514-a0107a5d8004
	github.com/truquito/gotruco v0.1.0
)

require github.com/dgryski/go-metro v0.0.0-20180109044635-280f6062b5bc // indirect

require (
	github.com/filevich/canvas v0.0.0 // indirect
	github.com/filevich/combinatronics v0.0.0-20220316214652-26aa6db09482
	github.com/filevich/truco-ai v0.0.0
)

replace github.com/truquito/gotruco => ../minitruco

// replace github.com/truquito/gotruco => ../gotruco

replace github.com/filevich/truco-ai => ../truco-ai

replace github.com/axiomhq/hyperloglog => ../axiom-hyperloglog

replace github.com/clarkduvall/hyperloglog => ../clarkduvall-hyperloglog
