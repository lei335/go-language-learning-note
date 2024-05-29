// process a big and complicated json file
// for example: ./transcript.json, 241MB size, copied from https://github.com/ethereum/kzg-ceremony/blob/main/transcript.json
package file

import (
	"embed"
	"encoding/json"
	"log"
)

// ScalarsPerBlob is the number of serialized scalars in a blob.
const ScalarsPerBlob = 32768

// JSONResult is a struct used for serializing the transcripts from/to JSON format.
type JSONResult struct {
	Transcripts []Transcript `json:"transcripts"`
}

type Transcript struct {
	NumG1Powers int `json:"numG1Powers"`
	NumG2Powers int `json:"numG2Powers"`
	Powers      PowersOfTau
}

type PowersOfTau struct {
	G1Powers []G1CompressedHexStr `json:"G1Powers"`
	G2Powers []G2CompressedHexStr `json:"G2Powers"`
}

// G1CompressedHexStr is a hex-string (with the 0x prefix) of a compressed G1 point.
type G1CompressedHexStr = string

// G2CompressedHexStr is a hex-string (with the 0x prefix) of a compressed G2 point.
type G2CompressedHexStr = string

// notice: '//go:embed' can't have space
//
//go:embed transcript.json
var content embed.FS

func ReadJsonFile() {
	config, err := content.ReadFile("transcript.json")
	if err != nil {
		log.Fatal(err)
	}

	params := new(JSONResult)
	if err = json.Unmarshal(config, params); err != nil {
		log.Fatal(err)
	}
}
