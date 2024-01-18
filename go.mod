module github.com/hll-truco/experiments

go 1.18

require github.com/truquito/truco v0.1.0

require (
	github.com/filevich/canvas v0.0.0 // indirect
	github.com/filevich/combinatronics v0.0.0-20220316214652-26aa6db09482
	github.com/filevich/truco-cfr v0.0.0
)

replace github.com/truquito/truco => ../minitruco
// replace github.com/truquito/truco => ../truco

replace github.com/filevich/truco-cfr => ../truco-cfr
