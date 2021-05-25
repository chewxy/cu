package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"sort"
	"strings"

	"os"
	"path"

	"github.com/gorgonia/bindgen"
	"modernc.org/cc"
)

var pkgloc string
var apiFile string
var ctxFile string
var hdrfile string
var pkghdr string
var model = bindgen.Model()

func init() {
	gopath := os.Getenv("GOPATH")
	pkgloc = path.Join(gopath, "src/gorgonia.org/cu/dnn")
	apiFile = path.Join(pkgloc, "api.go")
	ctxFile = path.Join(pkgloc, "ctx_api.go")
	hdrfile = "cudnn.h"

	pkghdr = `package cudnn

/* Generated by gencudnn. DO NOT EDIT */

// #include <cudnn.h>
import "C"
`

}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

// yes I know goimports can be imported, but I'm lazy
func goimports(filename string) error {
	cmd := exec.Command("goimports", "-w", filename)
	return cmd.Run()
}

func main() {
	var pkg *PkgState
	pkg = parsePkg(false)
	// Step 0: run parse.py to get more sanity about inputs and outputs
	// Step 1: Explore
	// explore(hdrfile, functions, enums, otherTypes)
	// explore(hdrfile, otherTypes)
	// explore(hdrfile, functions)

	// Step 2: generate mappings for this package, then edit them manually
	// 	Specifically, the `ignored` map is edited - things that will be manually written are not removed from the list
	//	Some enum map names may also be changed
	//defaultPipeline := func(buf io.WriteCloser, t *cc.TranslationUnit) {
	// bindgen.GenNameMap(buf, t, "fnNameMap", processNameBasic, functions, true)
	// bindgen.GenNameMap(buf, t, "enumMappings", processNameBasic, enums, true)
	// generateAlphaBeta(buf, t)
	// generateCRUD(buf, t, "create")
	// generateCRUD(buf, t, "set")
	// generateCRUD(buf, t, "destroy")
	//generateCRUD(buf, t, "methods")
	//}
	//generateMappings(true, defaultPipeline)

	// Step 3: generate enums, then edit the file in the dnn package.
	//generateEnums()
	//generateEnumStrings()
	//generateStubs(false, pkg) // true/false indicates debug mode

	// Step 4: manual fix for inconsistent names (Spatial Transforms)

	// step 5:
	generateFunctions(pkg)

	// report things that aren't done yet
	pkg = parsePkg(true)
	reportPotentialNils(pkg)
	reportUnconvertedFns(pkg, hdrfile, functions)
	reportUnconvertedTypes(pkg, hdrfile, otherTypes, enums)

}

func explore(file string, things ...bindgen.FilterFunc) {
	t, err := bindgen.Parse(model, file)
	handleErr(err)
	bindgen.Explore(t, things...)
}

// find potential nils
func reportPotentialNils(pkg *PkgState) {
	nils := pkg.checkNils()
	if len(nils) > 0 {
		fmt.Printf("## Potential Nils ##\nThese functions have a `*T` return value, but a possible null exception error might happen\n\n")
		for _, n := range nils {
			fmt.Printf("* `%v`\n", n)
		}
	}
	fmt.Println()
}

// find unused/unconverted functions
func reportUnconvertedFns(pkg *PkgState, file string, things ...bindgen.FilterFunc) {
	used := pkg.usedCFn()
	t, err := bindgen.Parse(bindgen.Model(), file)
	handleErr(err)

	var allFuncs = make(map[string]struct{})
	for _, thing := range things {
		decls, err := bindgen.Get(t, thing)
		handleErr(err)
		for _, decl := range decls {
			name := bindgen.NameOf(decl)
			allFuncs[name] = struct{}{}
		}

	}
	keys := make([]string, 0, len(allFuncs))
	for k := range allFuncs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Printf("## Unconverted C Functions ##\n\n")
	for _, k := range keys {
		if _, ok := used[k]; !ok {
			fmt.Printf("* `%v`\n", k)
		}
	}
	fmt.Println()
}

func reportUnconvertedTypes(pkg *PkgState, file string, things ...bindgen.FilterFunc) {
	used := pkg.usedCTypes()
	t, err := bindgen.Parse(bindgen.Model(), file)
	handleErr(err)

	var allTypes = make(map[string]struct{})
	for _, thing := range things {
		decls, err := bindgen.Get(t, thing)
		handleErr(err)
		for _, decl := range decls {
			name := bindgen.NameOf(decl)
			allTypes[name] = struct{}{}
		}
	}
	keys := make([]string, 0, len(allTypes))
	for k := range allTypes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Printf("## Unconverted/Unused C Types ##\n\n")
	for _, k := range keys {
		if _, ok := used[k]; !ok {
			fmt.Printf("* `%v`\n", k)
		}
	}
	fmt.Println()
}

func generateEnums() {
	fullpath := path.Join(pkgloc, "generated_enums.go")
	buf, err := os.OpenFile(fullpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	handleErr(err)
	fmt.Fprintln(buf, pkghdr)

	t, err := bindgen.Parse(model, hdrfile)
	handleErr(err)

	decls, err := bindgen.Get(t, enums)
	handleErr(err)

	for _, d := range decls {
		e := d.(*bindgen.Enum)
		if isIgnored(e.Name) {
			continue
		}
		fmt.Fprintf(buf, "type %v int\nconst (\n", enumMappings[e.Name])

		var names []string
		for _, a := range e.Type.EnumeratorList() {
			cname := string(a.DefTok.S())
			names = append(names, cname)
		}

		lcp := bindgen.LongestCommonPrefix(names...)

		for _, a := range e.Type.EnumeratorList() {
			cname := string(a.DefTok.S())
			enumName := processEnumName(lcp, cname)
			fmt.Fprintf(buf, "%v %v = C.%v\n", enumName, enumMappings[e.Name], cname)
		}
		fmt.Fprintf(buf, ")\n// C returns the C representation of %v\n", enumMappings[e.Name])
		fmt.Fprintf(buf, "func (e %v) C() C.%v { return C.%v(e) }\n", enumMappings[e.Name], e.Name, e.Name)
	}
	buf.Close()
	if err := goimports(fullpath); err != nil {
		log.Printf("Failed to Goimports %q: %v", fullpath, err)
	}
}

func generateEnumStrings() {
	fullpath := path.Join(pkgloc, "generated_enums_strings.go")
	buf, err := os.OpenFile(fullpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	handleErr(err)
	fmt.Fprintln(buf, pkghdr)

	t, err := bindgen.Parse(model, hdrfile)
	handleErr(err)

	decls, err := bindgen.Get(t, enums)
	handleErr(err)

	for _, d := range decls {
		e := d.(*bindgen.Enum)
		if isIgnored(e.Name) {
			continue
		}
		fmt.Fprintf(buf, "var _%vNames = map[%v]string{\n", enumMappings[e.Name], enumMappings[e.Name])

		var names []string
		for _, a := range e.Type.EnumeratorList() {
			cname := string(a.DefTok.S())
			names = append(names, cname)
		}

		lcp := bindgen.LongestCommonPrefix(names...)

		for _, a := range e.Type.EnumeratorList() {
			cname := string(a.DefTok.S())
			enumName := processEnumName(lcp, cname)
			fmt.Fprintf(buf, "%v : %q,\n", enumName, enumName)
		}
		fmt.Fprintf(buf, "}\n func (e %v) String() string {return _%vNames[e]}\n", enumMappings[e.Name], enumMappings[e.Name])
	}
	buf.Close()
	if err := goimports(fullpath); err != nil {
		log.Printf("Failed to Goimports %q: %v", fullpath, err)
	}
}

// generateStubs creates most of the stubs
func generateStubs(debugMode bool, state *PkgState) {
	t, err := bindgen.Parse(model, hdrfile)
	handleErr(err)
	var buf io.WriteCloser
	var fullpath string
	if debugMode {
		filename := "generated_FOO.go"
		fullpath = path.Join(pkgloc, filename)
		buf, err = os.OpenFile(fullpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		handleErr(err)
		fmt.Fprintln(buf, pkghdr)
	}

	decls := state.TypeDecls()

	var todoCount int
outer:
	for k, vs := range setFns {
		if isIgnored(k) {
			continue
		}
		var hasTODO bool
		gotype, ok := ctypes2GoTypes[k]
		if !ok {
			log.Printf("Cannot generate for %q", k)
			continue
		}
		if alreadyProcessedType(gotype, decls) {
			log.Printf("Already processed %v", gotype)
			continue
		}

		// if we're not in debugging mode, then we should write out to different files per type generated
		// this makes it easier to work on all the TODOs
		if !debugMode {
			filename := fmt.Sprintf("generated_%v.go", strings.ToLower(gotype))
			fullpath = path.Join(pkgloc, filename)
			buf, err = os.OpenFile(fullpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			handleErr(err)
			fmt.Fprintln(buf, pkghdr)
		}

		for _, v := range vs {
			if isIgnored(v) {
				log.Printf("Skipped generating for %q", k)
				continue outer
			}
		}

		// get the creation function to "guess" what should be in the struct
		filter := func(decl *cc.Declarator) bool {
			if decl.Type.Kind() != cc.Function {
				return false
			}
			n := bindgen.NameOf(decl)
			for _, v := range vs {
				if n == v {
					return true
				}
			}
			return false
		}
		decls, err := bindgen.Get(t, filter)
		handleErr(err)

		cs := decls[0].(*bindgen.CSignature)
		sig := GoSignature{}
		var create, destroy string
		if creates, ok := creations[k]; ok {
			create = creates[0]
		}
		if destroys, ok := destructions[k]; ok {
			destroy = destroys[0]
		}

		if create == "" || destroy == "" {
			log.Printf("Skipped %v - No Create/Destroy", k)
			continue
		}

		body := Con{
			Ctype:   k,
			GoType:  gotype,
			Create:  create,
			Destroy: destroy,
			Set:     vs,
		}
		sig.Receiver = Param{Name: "DUMMY", Type: "*" + gotype}
		sig.RetVals = append(sig.RetVals, Param{Name: "DUMMY"})

		if _, err = csig2gosig(cs, &sig); err != nil {
			body.TODO = err.Error()
			log.Print(body.TODO)
			hasTODO = true
		}
		sig.RetVals[0] = sig.Receiver
		sig.RetVals[0].Name = "retVal"
		sig.Receiver = Param{} // Param is set to empty

		for _, p := range sig.Params {
			body.Params = append(body.Params, p.Name)
			body.ParamType = append(body.ParamType, p.Type)
		}
		sig.Doc = fmt.Sprintf("New%v creates a new %v.", gotype, gotype)
		sig.Name = fmt.Sprintf("New%v", gotype)
		constructStructTemplate.Execute(buf, body)

		fmt.Fprintf(buf, "\n%v{ \n", sig)
		if len(vs) > 1 {
			hasTODO = true
			constructionTODOTemplate.Execute(buf, body)
		} else {
			constructionTemplate.Execute(buf, body)
		}
		fmt.Fprintf(buf, "}\n")

		// getters
		retValPos := getRetVal(cs)
		if len(vs) > 1 {
			fmt.Fprintf(buf, "// TODO: Getters for %v\n", gotype)
			goto generateDestructor
		}
		for i, p := range cs.Parameters() {
			if _, ok := retValPos[i]; ok {
				continue
			}
			getterSig := GoSignature{}
			typName := goNameOf(p.Type())

			// receiver - we don;t log - adds to the noise
			if typName == gotype {
				continue
			}

			if typName == "" {
				hasTODO = true
				fmt.Fprintf(buf, "//TODO: %q: Parameter %d Skipped %q of %v - unmapped type\n", cs.Name, i, p.Name(), p.Type())
				continue
			}
			getterSig.Doc = fmt.Sprintf("%v returns the internal %v.", strings.Title(p.Name()), p.Name())
			getterSig.Receiver.Name = strings.ToLower(string(gotype[0]))
			getterSig.Receiver.Type = reqPtr(gotype)
			getterSig.Name = strings.Title(p.Name())
			getterSig.RetVals = []Param{{Type: typName}}
			fmt.Fprintf(buf, "\n%v{ return %v.%v }\n", getterSig, getterSig.Receiver.Name, p.Name())
		}

	generateDestructor:
		// destructor
		destructor := GoSignature{}
		destructor.Name = "destroy" + gotype
		destructor.Params = []Param{
			{Name: "obj", Type: gotype, IsPtr: true},
		}
		fmt.Fprintf(buf, "\n%v{", destructor)
		destructTemplate.Execute(buf, body)
		fmt.Fprintf(buf, "}\n")

		if !debugMode {
			buf.Close()
			log.Printf("Written to %v. TODO: %v", fullpath, hasTODO)
			if err := goimports(fullpath); err != nil {
				log.Printf("Failed to Goimports %q: %v", fullpath, err)
			}
		}
		if hasTODO {
			todoCount++
		}
	}
	log.Printf("%d/%d TODOs", todoCount, len(setFns))

	buf.Close()
	if err := goimports(fullpath); err != nil {
		log.Printf("Failed to Goimports %q: %v", fullpath, err)
	}
}

func generateFunctions(pkg *PkgState) {
	t, err := bindgen.Parse(model, hdrfile)
	handleErr(err)
	filename := "generated_API.go"
	fullpath := path.Join(pkgloc, filename)
	buf, err := os.OpenFile(fullpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	fmt.Fprintln(buf, pkghdr)
	decls, err := bindgen.Get(t, functions)
	handleErr(err)

	for rec, fns := range methods {
		var todoCount, dblchk int
		for _, decl := range decls {
			csig := decl.(*bindgen.CSignature)
			name := csig.Name
			if !inList(name, fns) {
				continue
			}
			if _, ok := ignored[name]; ok {
				continue
			}
			sig := GoSignature{}
			sig.Receiver.Name = strings.ToLower(depointerize(goNameOfStr(rec))[0:2])
			sig.Receiver.Type = reqPtr(goNameOfStr(rec))
			sig.Name = fnNameMap[name]

			_, err := csig2gosig(csig, &sig)

			fmt.Fprintf(buf, "%v { \n", sig)
			if err != nil {
				fmt.Fprintf(buf, "// DOUBLECHECK: %v\n", err)
				dblchk++
			}

			ab := sig.AlphaBetas()
			if len(ab) > 0 {
				check := sig.FirstTensor()
				data := AlphaBeta{
					Params:      ab,
					Check:       check,
					LSO:         len(ab) - 1,
					MultiReturn: len(sig.RetVals) > 1,
				}
				alphaTemplate.Execute(buf, data)
			}

			// update tracking matrix
			// the goal is to find out which parameter has not been coverted yet
			cparams := csig.Parameters()
			track := make([]bool, len(cparams))
			converted := make([]bool, len(cparams))
			primOut := make([]bool, len(cparams))
			enumOut := make([]bool, len(cparams))

			// receiver
			var receiverParam string
			for i, p := range cparams {
				if goNameOf(p.Type()) == sig.Receiver.Type {
					track[i] = true
					receiverParam = p.Name()
					break // only ONE receiver
				}
			}

			// alpha/betas
			for _, a := range ab {
				for i, p := range cparams {
					if safeParamName(p.Name()) == a {
						track[i] = true
						converted[i] = true
					}
				}
			}

			// retVals
			// otherParams := make(map[int]string)
			for _, p := range sig.RetVals {
				for i, cp := range cparams {
					cpName := safeParamName(cp.Name())
					if cpName != p.Name {
						continue
					}
					if isOutputPtrOfPrim(csig.Name, cp) {
						// create conversion
						fmt.Fprintf(buf, "var %vC C.%v\n", cpName, ctype2gotype2ctype(cp.Type()))
						track[i] = true
						converted[i] = true
						primOut[i] = true
						continue
					}

					if isEnumOutput(csig.Name, cp) {
						// log.Printf("%v:: %v %v is enum output", csig.Name, cp.Name(), nameOfType(cp.Type()))
						fmt.Fprintf(buf, "var %vC C.%v\n", cpName, nameOfType(cp.Type()))
						converted[i] = true
						track[i] = true
						enumOut[i] = true
						continue
					}
				}
			}

			// other params
			for _, p := range sig.Params {
				for i, cp := range cparams {
					cpName := safeParamName(cp.Name())
					if cpName != p.Name {
						continue
					}
					track[i] = true
				}
			}

			var hasTodo bool
			for i, t := range track {
				if !t {
					fmt.Fprintf(buf, "// TODO: %v %v\n", cparams[i].Name(), nameOfType(cparams[i].Type()))
					hasTodo = true
				}
			}
			if hasTodo {
				todoCount++
			}

			// make the call
			callParams := make([]Param, len(cparams))
			for i, p := range cparams {
				pname := p.Name()
				callParams[i] = cParam2GoParam(p)
				if pname == receiverParam {
					callParams[i].Name = sig.Receiver.Name
					callParams[i].Type = sig.Receiver.Type
				}
				if converted[i] {
					callParams[i].Name += "C"
					callParams[i].Type = ""
				}

				switch {
				case primOut[i]:
					callParams[i].IsPtr = true
				case enumOut[i]:
					callParams[i].IsPtr = true
				}
			}
			data := Call{
				Params:      callParams,
				CFuncName:   csig.Name,
				MultiReturn: len(sig.RetVals) > 1,
			}
			callTemplate.Execute(buf, data)

			// reverse the calls
			for _, p := range cparams {
				if isOutputPtrOfPrim(csig.Name, p) {
					goType := goNameOfStr(depointerize(nameOfType(p.Type())))
					pname := safeParamName(p.Name())
					fmt.Fprintf(buf, "%v = %v(%vC)\n", pname, goType, pname)
					continue
				}
				if isEnumOutput(csig.Name, p) {
					goType := goNameOfStr(depointerize(nameOfType(p.Type())))
					pname := safeParamName(p.Name())
					fmt.Fprintf(buf, "%v = %v(%vC)\n", pname, goType, pname)
					continue
				}
			}

			if len(sig.RetVals) > 1 {
				fmt.Fprintf(buf, "return\n")
			}

			fmt.Fprintf(buf, "}\n")
		}

		log.Printf("Receiver : %v. Functions: %d. TODOs: %d. Double Checks: %d", rec, len(fns), todoCount, dblchk)
	}
	buf.Close()
	if err := goimports(fullpath); err != nil {
		log.Printf("Failed to Goimports %q: %v", fullpath, err)
	}
}
