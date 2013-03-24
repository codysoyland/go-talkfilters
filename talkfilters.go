package talkfilters

/*
#cgo CFLAGS: -I/usr/local/include/talkfilters
#cgo LDFLAGS: -L/usr/local/lib
#cgo LDFLAGS: -ltalkfilters

#include <talkfilters.h>
#include <stdlib.h>
#include <string.h>

int runfilter(const char *name, const char *input, char *output) {
    return gtf_filter_lookup(name)->filter(input, output, strlen(output));
}
*/
import "C"
import "unsafe"

type filter func(string) string
type filters map[string]filter

var talkfilters = filters {
    "austro": Austro,
    "b1ff": B1ff,
    "brooklyn": Brooklyn,
    "chef": Chef,
    "cockney": Cockney,
    "drawl": Drawl,
    "dubya": Dubya,
    "fudd": Fudd,
    "funetak": Funetak,
    "jethro": Jethro,
    "jive": Jive,
    "kraut": Kraut,
    "pansy": Pansy,
    "pirate": Pirate,
    "postmodern": Postmodern,
    "redneck": Redneck,
    "valspeak": Valspeak,
    "warez": Warez,
}
func ListFilters() filters {
    return talkfilters;
}
func RunFilter(filter string, phrase string) string {
    input := C.CString(phrase)
    // TODO: The next line is a fast hack. Clean up and make this generate *C.char of empty spaces
    output := C.CString(phrase+phrase+phrase)
    defer C.free(unsafe.Pointer(output))
    C.runfilter(C.CString(filter), input, output)
    C.free(unsafe.Pointer(input))
    return C.GoString(output)
}
func GetFilter(name string) filter {
    return talkfilters[name];
}
func Austro(phrase string) string {
    return RunFilter("austro", phrase)
}
func B1ff(phrase string) string {
    return RunFilter("b1ff", phrase)
}
func Brooklyn(phrase string) string {
    return RunFilter("brooklyn", phrase)
}
func Chef(phrase string) string {
    return RunFilter("chef", phrase)
}
func Cockney(phrase string) string {
    return RunFilter("cockney", phrase)
}
func Drawl(phrase string) string {
    return RunFilter("drawl", phrase)
}
func Dubya(phrase string) string {
    return RunFilter("dubya", phrase)
}
func Fudd(phrase string) string {
    return RunFilter("fudd", phrase)
}
func Funetak(phrase string) string {
    return RunFilter("funetak", phrase)
}
func Jethro(phrase string) string {
    return RunFilter("jethro", phrase)
}
func Jive(phrase string) string {
    return RunFilter("jive", phrase)
}
func Kraut(phrase string) string {
    return RunFilter("kraut", phrase)
}
func Pansy(phrase string) string {
    return RunFilter("pansy", phrase)
}
func Pirate(phrase string) string {
    return RunFilter("pirate", phrase)
}
func Postmodern(phrase string) string {
    return RunFilter("postmodern", phrase)
}
func Redneck(phrase string) string {
    return RunFilter("redneck", phrase)
}
func Valspeak(phrase string) string {
    return RunFilter("valspeak", phrase)
}
func Warez(phrase string) string {
    return RunFilter("warez", phrase)
}
