package main

import (
	"context"
	"embed"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/function61/gokit/app/aws/lambdautils"
	"github.com/function61/gokit/app/cli"
	. "github.com/function61/gokit/builtin"
	"github.com/function61/gokit/net/http/httputils"
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
		Short: "xs.fi",
	}

	app.AddCommand(&cobra.Command{
		Use: "run",
		// Short: "Reticulates splines",
		Args: cobra.NoArgs,
		Run: cli.WrapRun(func(ctx context.Context, _ []string) error {
			return logic(ctx)
		}),
	})

	cli.Execute(app)
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
	routes := http.NewServeMux()

	routes.HandleFunc("/9/{id...}", httputils.WrapWithErrorHandling(func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")
		http.Redirect(w, r, "https://joonas.fi/assets/view?id="+id, http.StatusFound)
		return nil
	}))

	routes.HandleFunc("/a0/lubebar6/{serialNo}", httputils.WrapWithErrorHandling(func(w http.ResponseWriter, r *http.Request) error {
		// not used yet.
		// serialNo := r.PathValue("serialNo")
		http.Redirect(w, r, "https://www.asiakastieto.fi/yritykset/fi/pirkanmaan-metallitekniikka-oy/25144545/yleiskuva", http.StatusFound)
		return nil
	}))

	routes.Handle("/", http.FileServer(http.FS(staticFilesXSfi)))

	return routes
}
