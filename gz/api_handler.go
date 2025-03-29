package gz

import (
	"context"
	"github.com/gookit/goutil/strutil"
	"github.com/kongmsr/oneid-core/utilo"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strings"
)

// ConvertHeaderMD convert all r.Header to metadata
func ConvertHeaderMD(r *http.Request) context.Context {
	md := make(map[string][]string)
	for key, values := range r.Header {
		if !strutil.HasPrefix(strings.ToUpper(key), strings.ToUpper("Oneid-")) {
			continue
		}
		lowerKey := http.CanonicalHeaderKey(key)
		md[lowerKey] = values
	}

	return metadata.NewOutgoingContext(r.Context(), md)
}

// GetHeaderFromMD get header value from metadata
func GetHeaderFromMD(md metadata.MD, headerName string) *string {
	if vals, ok := md[strings.ToLower(headerName)]; !ok {
		return nil
	} else {
		return utilo.ToPrt(vals[0])
	}
}
