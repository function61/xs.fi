package main

import (
	"context"
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/function61/gokit/app/aws/lambdautils"
	"github.com/function61/gokit/app/cli"
	"github.com/function61/gokit/app/dynversion"
	. "github.com/function61/gokit/builtin"
	"github.com/function61/gokit/net/http/httputils"
	"github.com/function61/gokit/os/osutil"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

//go:embed index.html logo-4.png
var staticFilesXSfi embed.FS

func main() {
	if lambdautils.InLambda() {
		handler := newServerHandler()
		lambda.Start(lambdautils.NewLambdaHttpHandlerAdapter(handler))
		return
	}

	app := &cobra.Command{
		Use:     os.Args[0],
		Short:   "xs.fi",
		Version: dynversion.Version,
	}

	app.AddCommand(&cobra.Command{
		Use: "run",
		// Short: "Reticulates splines",
		Args: cobra.NoArgs,
		Run: cli.RunnerNoArgs(func(ctx context.Context, _ *log.Logger) error {
			return logic(ctx)
		}),
	})

	osutil.ExitIfError(app.Execute())
}

func logic(ctx context.Context) error {
	srv := &http.Server{
		Addr:              ":" + FirstNonEmpty(os.Getenv("PORT"), "80"),
		Handler:           newServerHandler(),
		ReadHeaderTimeout: httputils.DefaultReadHeaderTimeout,
	}

	return httputils.CancelableServer(ctx, srv, srv.ListenAndServe)
}

func newServerHandler() http.Handler {
	routes := mux.NewRouter()

	routes.PathPrefix("/9/{id}").HandlerFunc(httputils.WrapWithErrorHandling(func(w http.ResponseWriter, r *http.Request) error {
		id := mux.Vars(r)["id"]
		http.Redirect(w, r, "https://joonas.fi/assets/view?id="+id, http.StatusFound)
		return nil
	}))

	routes.PathPrefix("/").Handler(http.FileServer(http.FS(staticFilesXSfi)))

	return routes
}
