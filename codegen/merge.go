package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/grokify/gotilla/path/filepathutil"
	"github.com/grokify/swaggman/openapi3"
	"github.com/grokify/swaggman/swagger2"
	"github.com/jessevdk/go-flags"
)

/*
type options struct {
	Version int `short:"v" long:"version" description:"OAS Version" required:"true"`
}


func MergeOAS2(dir, outfile string) error {
	spec, err := swagger2.MergeDirectory(dir)
	if err != nil {
		return errors.Wrap(err, "E_MERGE_FAILED")
	}

	err = ioutilmore.WriteFileJSON(outfile, spec, 0644, "", "  ")
	if err != nil {
		return errors.Wrap(err, "E_WRITE_FAILED")
	}
	fmt.Printf("WROTE [%v]\n", outfile)
	return nil
}

func MergeOAS3(dir, outfile string) error {
	spec, err := openapi3.MergeDirectory(dir)
	if err != nil {
		return errors.Wrap(err, "E_MERGE_FAILED")
	}

	bytes, err := spec.MarshalJSON()
	if err != nil {
		return errors.Wrap(err, "E_JSON_ENCODING_FAILED")
	}

	err = ioutil.WriteFile(outfile, bytes, 0644)
	if err != nil {
		return errors.Wrap(err, "E_WRITE_FAILED")
	}
	fmt.Printf("WROTE [%v]\n", outfile)
	return nil
}
*/

type options struct {
	Directory string `short:"d" long:"directory" description:"OAS Directory" required:"true"`
	Version   int    `short:"v" long:"version" description:"OAS Version" required:"true"`
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	version := 3
	if opts.Version == 2 {
		version = opts.Version
	}

	dir := opts.Directory

	dir = filepathutil.TrimRight(dir)

	_, leaf := filepath.Split(dir)

	outfile := leaf + ".json"
	/*
		outfile := fmt.Sprintf("openapi-spec_v%d.0.0.json", version)
		dir := fmt.Sprintf("partial-specs_v%d.0.0", version)*/

	switch version {
	case 2:
		//err = MergeOAS2(dir, outfile)
		err = swagger2.WriteFileDirMerge(outfile, dir, 0644)
	case 3:
		//err = MergeOAS3(dir, outfile)
		err = openapi3.WriteFileDirMerge(outfile, dir, 0644)
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")

}
